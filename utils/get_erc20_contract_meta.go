package utils

import (
	"context"
	"encoding/json"
	"eth-watcher/chain"
	"eth-watcher/consts"
	"eth-watcher/global"
	"eth-watcher/types"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetERC20ContractMeta(chainId int64, contractAddress *ethCommon.Address, client *ethclient.Client) (*types.ContractMeta, error) {
	if contractAddress == nil {
		return nil, nil
	}

	cacheTableKey := fmt.Sprintf(consts.CACHE_KEY_ERC20_META, chainId)
	cacheRowKey := contractAddress.Hex()

	var contractMeta types.ContractMeta

	// Get from cache
	exist, err := global.Redis.HExists(context.Background(), cacheTableKey, cacheRowKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to check meta existence in cache with error: %v", err)
	} else if exist {
		// Get from cache
		contractMetaBytes, err := global.Redis.HGet(context.Background(), cacheTableKey, cacheRowKey).Bytes()
		if err != nil {
			global.Logger.Errorf("Failed to get meta bytes from cache with error: %v", err)
		} else {
			err = json.Unmarshal(contractMetaBytes, &contractMeta)
			if err != nil {
				global.Logger.Errorf("Failed to unmarshal meta bytes %s into meta object with error: %v", contractMetaBytes, err)
			} else {
				// All done successfully
				return &contractMeta, nil
			}
		}
	}

	// Get from blockchain
	contractMeta.Name, contractMeta.Symbol, contractMeta.Decimals, err = chain.GetERC20Meta(contractAddress, client)
	if err != nil {
		global.Logger.Errorf("Failed to get token meta for contract %s from blockchain #%d with error: %v", contractAddress.Hex(), chainId, err)
		return nil, err
	}

	// Save into cache
	contractMetaBytes, err := json.Marshal(&contractMeta)
	if err != nil {
		global.Logger.Errorf("Failed to parse token meta %v for contract %s into cache with error: %v", contractMeta, contractAddress.Hex(), err)
	}

	global.Redis.HSet(context.Background(), cacheTableKey, cacheRowKey, contractMetaBytes)
	if err != nil {
		global.Logger.Errorf("Failed to save token meta for contract %s into cache with error: %v", contractAddress.Hex(), err)
	}

	return &contractMeta, nil

}
