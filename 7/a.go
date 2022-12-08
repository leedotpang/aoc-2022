package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NodeType int

const (
	DirNode NodeType = iota
	FileNode
)

type Node struct {
	size     int
	name     string
	nType    NodeType
	children []*Node
	parent   *Node
}

func (n *Node) AddChild(c *Node) {
	n.children = append(n.children, c)
}

func (n *Node) GetFolderSize() int {
	size := 0

	for _, c := range n.children {
		if c.nType == DirNode {
			size += c.GetFolderSize()
			continue
		}
		size += c.size
	}

	return size
}

func get_folder_sizes(parent *Node) map[string]int {
	sizes := map[string]int{
		parent.name: parent.GetFolderSize(),
	}

	// Loop over files
	for _, dir := range parent.children {
		if dir.nType == FileNode {
			continue
		}
		for name, size := range get_folder_sizes(dir) {
			// adding parent name to child dir to distinguish directories with the same name
			sizes[dir.name+name] = size
		}
	}

	return sizes
}

func get_file_size(s string) int {
	size, _ := strconv.Atoi(s)
	return size
}

func build_tree(commands string) *Node {
	// Loop through file and create directory tree based on input
	root := &Node{
		name:     "/",
		children: []*Node{},
		nType:    DirNode,
	}
	pwd := root
	cmd := []string{}
	temp_node := &Node{}

	for _, dir := range strings.Split(commands, "$ cd ") {
		if len(dir) == 0 {
			continue
		}

		for i, node := range strings.Split(dir, "\n") {
			if len(node) == 0 || regexp.MustCompile(`(\$ ls|dir )`).FindString(node) != "" {
				continue
			}
			if i == 0 && node == "/" {
				pwd = root
				continue
			}
			if node == ".." {
				pwd = pwd.parent
				continue
			}

			cmd = strings.Split(node, " ")

			// $ cd = new subdir is being accessed
			if len(cmd) == 1 {
				temp_node = &Node{
					name:   cmd[0],
					parent: pwd,
					nType:  DirNode,
				}
				pwd.AddChild(temp_node)
				pwd = temp_node
				continue
			}

			pwd.AddChild(&Node{
				size:   get_file_size(cmd[0]),
				name:   cmd[1],
				parent: pwd,
				nType:  FileNode,
			})
		}
	}

	return root
}

func find_solution(content string) int {
	const MAX_SIZE int = 100000
	root := build_tree(content)
	folder_sizes := get_folder_sizes(root)

	// loop to get top three
	sum_of_folder_sizes := 0
	for _, size := range folder_sizes {
		if size > MAX_SIZE {
			continue
		}
		sum_of_folder_sizes += size
	}
	return sum_of_folder_sizes
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)
	// 	content_str := `$ cd /
	// $ ls
	// dir a
	// 14848514 b.txt
	// 8504156 c.dat
	// dir d
	// $ cd a
	// $ ls
	// dir e
	// 29116 f
	// 2557 g
	// 62596 h.lst
	// $ cd e
	// $ ls
	// 584 i
	// $ cd ..
	// $ cd ..
	// $ cd d
	// $ ls
	// 4060174 j
	// 8033020 d.log
	// 5626152 d.ext
	// 7214296 k`

	fmt.Println("Total of Dirs Under 10k: ", find_solution(content_str))
}
