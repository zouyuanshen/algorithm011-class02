### 字典树 Trie
- 基本性质
    + 结点本身不存完整单词
    + 从根结点到某一结点，路径经过的字符连接起来，为改结点对应的字符串
    + 每个结点的所有子结点路径代表的字符都不同
- 核心思想
    + 空间换时间
    + 利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的
