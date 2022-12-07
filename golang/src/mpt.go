package main

//导入 eth-merkle-patricia-tree
import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
)

func storePolicy() {
	//创建一个空的trie
	trie, err := trie.New(common.Hash{}, trie.NewDatabase(nil))
	if err != nil {
		panic(err)
	}

	//将数据存入trie
	trie.Update([]byte("key"), []byte("value"))
	//将trie的根节点存入数据库
	trie.Commit(nil)
}
