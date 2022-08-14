package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block
}

/*
	Singleton Pattern
   - 변수의 instance를 직접 공유하지 않음
   - 변수의 instance를 대신해서 드러내주는 function 생성
   - 다른 패키지에서 blockchain이 어떻게 드러날 지를 제어할 수 있음
*/
var b *blockchain

/*
	Sync Package
	 - 동기적으로 처리해야 하는 부분을 처리하도록 도와줌
	 - sync.Once -> Do func -> 단 한번만 호출되도록 도와주는 함수
*/
var once sync.Once

// 복사된 블록이 아닌 블록자체에 hash값 추가
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

// 마지막 블록 hash 값 리턴
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

/*
	블록생성
	getLastHash() -> prevHash return
	calculateHash() -> newBlock.hash 계산 및 추가
*/
func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
