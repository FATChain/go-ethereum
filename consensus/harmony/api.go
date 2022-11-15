// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package harmony

import (
	"bytes"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

// API is a user facing RPC API to allow controlling the delegate and voting
// mechanisms of the delegated-proof-of-stake
type API struct {
	chain  consensus.ChainHeaderReader
	engine *Harmony
}

// MintKeyLen -> mint trie 的key 长度
const MintKeyLen = 28

// GetValidators retrieves the list of the validators at specified block
func (api *API) GetValidators(number *rpc.BlockNumber) ([]common.Address, error) {
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	if header == nil {
		return nil, errUnknownBlock
	}

	ctxTrie, err := newEpochTrie(header.EngineInfo.EpochHash, api.engine.db)
	if err != nil {
		return nil, err
	}
	ctx := Context{}
	ctx.SetEpochTrie(ctxTrie)
	validators, err := ctx.GetValidators()
	if err != nil {
		return nil, err
	}
	return validators, nil
}

// GetConfirmedBlockNumber retrieves the latest irreversible block
func (api *API) GetConfirmedBlockNumber() (uint64, error) {
	height := api.chain.CurrentHeader().Number.Uint64() - 7
	return height, nil
}

// GetVoteList retrieves voteTrie return Delegate and its vote guy
// VoteTrie数据格式：
// key：vote-投票人地址
// value：候选人地址
// 这里需要看到开始的几个见证人自己给自己的投票--取vote树完了后，然后取当前委托者列表，取个并集
func (api *API) GetVoteList() (map[common.Address]common.Address, error) {
	candidates := map[common.Address]common.Address{}

	var header *types.Header

	header = api.chain.CurrentHeader()

	if header == nil {
		return nil, errUnknownBlock
	}

	voteTrie, err := newVoteTrie(header.EngineInfo.VoteHash, api.engine.db)
	if err != nil {
		return nil, err
	}

	iterCandidate := voteTrie.Iterator(nil)
	existCandidate := iterCandidate.Next()

	for existCandidate {
		addr := iterCandidate.Key
		candidate := iterCandidate.Value
		candidates[common.BytesToAddress(addr)] = common.BytesToAddress(candidate)

		existCandidate = iterCandidate.Next()
	}

	return candidates, nil
}

// GetCurEpoch retrieves current epoch
func (api *API) GetCurEpoch(number *rpc.BlockNumber) (uint64, error) {
	var header *types.Header

	header = api.chain.CurrentHeader()

	if header == nil {
		return 0, errUnknownBlock
	}
	return header.Time / epochInterval, nil
}

// GetMintCnt 得到epoch的所有验证节点和对应的出块数
// MintCntTrie数据格式：
// key；mintCnt-周期数（2进制）-验证人
// value：当前验证人本周期总共挖块数
func (api *API) GetMintCnt(epochID int) (map[common.Address]uint64, error) {
	mintMap := map[common.Address]uint64{}

	var header *types.Header

	header = api.chain.CurrentHeader()

	if header == nil {
		return nil, errUnknownBlock
	}

	mintTrie, err := newMintTrie(header.EngineInfo.MintCntHash, api.engine.db)
	if err != nil {
		return nil, err
	}

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(epochID))

	iterMint := mintTrie.Iterator(key)
	existMint := iterMint.Next()
	if !existMint {
		//return nil, errors.New("no candidates")
	}
	for existMint {
		iterkey := iterMint.Key
		itervalue := iterMint.Value

		if len(iterkey) != MintKeyLen {
			log.Info("existMint len is not expected", "key", MintKeyLen)
		}
		//前缀 0 - 7  epoch 8-27 validator addr
		epoch := binary.BigEndian.Uint64(iterkey[0:8])

		validator := common.BytesToAddress(iterkey[8:28])
		cnt := binary.BigEndian.Uint64(itervalue)

		log.Info("votes", "epoch", epoch, "validator", validator, "value", cnt)

		mintMap[validator] = cnt
		existMint = iterMint.Next()
	}

	return mintMap, nil
}

// GetValidatorMintCnt --- 得到epoch中validator的出块数
// MintCntTrie数据格式：
// key；mintCnt-周期数（2进制）-验证人
// value：当前验证人本周期总共挖块数
func (api *API) GetValidatorMintCnt(epochID int, addr string) (uint64, error) {
	count := uint64(0)
	var header *types.Header

	header = api.chain.CurrentHeader()

	if header == nil {
		return 0, errUnknownBlock
	}

	mintTrie, err := newMintTrie(header.EngineInfo.MintCntHash, api.engine.db)
	if err != nil {
		return 0, err
	}

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(epochID))
	validator := common.HexToAddress(addr)

	cntBytes, err := mintTrie.t.TryGet(append(key, validator.Bytes()...))
	if cntBytes != nil {
		count = binary.BigEndian.Uint64(cntBytes)
	}

	return count, err
}

// GetCandidates retrieves current candidates
func (api *API) GetCandidates() ([]common.Address, error) {
	var header *types.Header
	var candidates []common.Address

	header = api.chain.CurrentHeader()

	if header == nil {
		return nil, errUnknownBlock
	}

	candidateTrie, err := newCandidateTrie(header.EngineInfo.CandidateHash, api.engine.db)
	if err != nil {
		return nil, err
	}

	iterCandidates := candidateTrie.Iterator(nil)
	existCandidate := iterCandidates.Next()
	if !existCandidate {
		//return nil, errors.New("no candidates")
	}
	for existCandidate {
		addr1 := iterCandidates.Key
		addr2 := iterCandidates.Value
		if bytes.Equal(addr2, addr1) {

		}
		candidates = append(candidates, common.BytesToAddress(addr2))
		existCandidate = iterCandidates.Next()
	}
	return candidates, nil
}
