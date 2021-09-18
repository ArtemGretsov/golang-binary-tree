package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree `json:"left"`
	Right *Tree `json:"right"`
	Value int   `json:"value"`
}

func (t *Tree) Create(tree *Tree, maxIncludeSize int) Tree {
	if maxIncludeSize == 0 {
		return *tree
	}

	tree.Value = generateRandomInt()
	isRight := generateRandomBool()
	isLeft := generateRandomBool()

	if isRight {
		tree.Right = &Tree{}
		t.Create(tree.Right, maxIncludeSize - 1)
	}

	if isLeft {
		tree.Left = &Tree{}
		t.Create(tree.Left, maxIncludeSize - 1)
	}

	return *tree
}

func (t *Tree) Generate(maxIncludeSize int) {
	t.Create(t, maxIncludeSize)
}

func (t *Tree) Print() {
	out, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func (t *Tree) calculate(tree *Tree) int {
	sum := tree.Value

	if tree.Left != nil {
		sum += t.calculate(tree.Left)
	}

	if tree.Right != nil {
		sum += t.calculate(tree.Right)
	}

	return sum
}

func (t *Tree) Calculate() int {
	return t.calculate(t)
}

func generateRandomBool() bool {
	randSource := rand.New(rand.NewSource(time.Now().Unix()))
	min := 0
	max := 1
	randValue := randSource.Intn(max - min + 1) + min
	return randValue != 0
}

func generateRandomInt() int {
	randSource := rand.New(rand.NewSource(time.Now().Unix()))
	min := 0
	max := 10
	return randSource.Intn(max - min + 1) + min
}

func main() {
	tree := Tree{}
	tree.Generate(10)
	tree.Print()
	sum := tree.Calculate()
	fmt.Println(sum)
}
