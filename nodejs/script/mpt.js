const { level } = require('level');
const { fs } = require('fs');
const { csv } = require('csv-parser');
const { BaseTrie: Trie } = require('merkle-patricia-tree');

const trie = new Trie();
let parser = parse({columns: true}, function (err, records) {
    console.log(records);
});
//将policy.csv中的数据存入trie
async function test() {
  fs.createReadStream('../config/policy.csv').pipe(parser);
  
  for (let i = 0; i < records.length; i++) {
    await trie.put(records[i].p_type, records[i].v0);
  }

}

test();