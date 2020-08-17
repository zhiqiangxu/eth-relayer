/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
*/
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/polynetwork/eth_relayer/log"
)

const (
	ETH_MONITOR_INTERVAL = 3 * time.Second
	ONT_MONITOR_INTERVAL = 3 * time.Second
	KEY_UNLOCK_TIME      = 30 * time.Second

	ETH_USEFUL_BLOCK_NUM     = 3
	ETH_PROOF_USERFUL_BLOCK  = 12
	ONT_USEFUL_BLOCK_NUM     = 1
	ETH_CHAIN_ID             = 2
	DEFAULT_CONFIG_FILE_NAME = "./config.json"
	Version                  = "1.0"

	DEFAULT_LOG_LEVEL = log.InfoLog
)

//type ETH struct {
//	Chain             string // eth or etc
//	ChainId           uint64
//	RpcAddress        string
//	ConfirmedBlockNum uint
//	//Tokens            []*Token
//}

type ServiceConfig struct {
	MultiChainConfig *MultiChainConfig
	ETHConfig        *ETHConfig
	BoltDbPath       string
	RoutineNum int64
}

type MultiChainConfig struct {
	RestURL                 string
	EntranceContractAddress string
	WalletFile              string
	WalletPwd               string
}

type ETHConfig struct {
	RestURL             string
	ECCMContractAddress string
	ECCDContractAddress string
	CapitalOwnersPath   string
	CapitalPassword     string
	BlockConfig         uint64
}

type ONTConfig struct {
	RestURL string
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Errorf("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func NewServiceConfig(configFilePath string) *ServiceConfig {
	fileContent, err := ReadFile(configFilePath)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	servConfig := &ServiceConfig{}
	err = json.Unmarshal(fileContent, servConfig)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}

	return servConfig
}