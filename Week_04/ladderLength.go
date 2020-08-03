package main

import "fmt"

func main() {

	fmt.Println(ladderLength("hit", "dot", []string{"hot","dot","dog","lot","log","cog"}))

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if  inList(endWord, wordList) == false || beginWord ==""|| endWord==""|| len(wordList) == 0{
		return 0
	}

	n := len(beginWord)
	wordmap := make(map[string] []string)
	for _, word := range wordList{
		for i:=0; i< n; i++{
			tempS := word[:i] + "*" + word[i+1:]
			wordmap[tempS] = append(wordmap[tempS], word)
		}
	}
	// queue
	queueBegin := [][]interface{}{{beginWord, 1}}
	queueEnd := [][]interface{}{{endWord, 1}}

	// map	make sure don't repeat
	visitedBegin := map[string] int{beginWord:1}
	visitedEnd := map[string] int{endWord:1}
	//var ans int
	for len(queueBegin) > 0 && len(queueEnd) >0 {

		ans := visitWordNode(&queueBegin, n, visitedBegin, visitedEnd, wordmap)
		if ans > 0 {
			return ans
		}
		ans = visitWordNode(&queueEnd, n, visitedEnd,visitedBegin, wordmap)
		if ans > 0 {
			return ans
		}
	}
	return 0
}

func inList(word string, wordList []string )  bool{
	for _,v:=range wordList{

		if v == word{
			return true
		}
	}
	return false
}

func visitWordNode(queue *[][]interface{}, n int, visited, visitedOther map[string] int, wordmap map[string] []string)  int{
	currentWord, level := (*queue)[0][0].(string), (*queue)[0][1].(int)
	*queue = (*queue)[1:]

	for i:=0 ; i < n; i++ {

		intermediateWord := currentWord[:i] + "*" + currentWord[i+1:]
		for _, word := range wordmap[intermediateWord]{
			if visitedOther[word] > 0 {
				return level + visitedOther[word]
			}
			if visited[word] <= 0 {
				visited[word] = level + 1
				*queue=  append(*queue, []interface{}{word, level+1})
			}
		}
	}
	return 0
}