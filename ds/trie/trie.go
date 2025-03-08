package trie

type Node struct {
	children map[rune]*Node
	isWord   bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
			isWord:   false,
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &Node{
				children: make(map[rune]*Node),
				isWord:   false,
			}
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}

func (t *Trie) GetAllWord() []string {
	result := []string{}
	t.collectWords(t.root, "", &result)
	return result
}

func (t *Trie) Delete(word string) bool {
	return t.deleteAux(t.root, word, 0)
}

func (t *Trie) deleteAux(node *Node, word string, index int) bool {
	if index == len(word) {
		if !node.isWord {
			return false
		}
		node.isWord = false
		return len(node.children) == 0
	}

	ch := rune(word[index])
	child, ok := node.children[ch]
	if !ok {
		return false
	}

	deleteChild := t.deleteAux(child, word, index+1)
	if !deleteChild {
		return false
	}

	delete(node.children, ch)
	return len(node.children) == 0 && !node.isWord
}

func (t *Trie) searchPrefix(prefix string) *Node {
	curr := t.root
	for _, ch := range prefix {
		if _, ok := curr.children[ch]; !ok {
			return nil
		}
		curr = curr.children[ch]
	}

	return curr
}

func (t *Trie) collectWords(node *Node, currentString string, result *[]string) {
	if node.isWord {
		*result = append(*result, currentString)
	}

	for ch, child := range node.children {
		t.collectWords(child, currentString+string(ch), result)
	}
}
