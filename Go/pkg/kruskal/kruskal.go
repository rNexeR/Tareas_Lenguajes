package kruskal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type tableRow struct {
	From   string
	To     string
	Weight int
	Flag   int
}

type Edge struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Weight int    `json:"weight"`
}

type ByEdge []Edge

type Graph struct {
	Edges []Edge `json:"edges"`
}

func (s ByEdge) Len() int {
	return len(s)
}

func (s ByEdge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByEdge) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

/*
func main() {
	fmt.Println("Working")
	res := Graph{}
	data := `{"edges" :
				[
					{"from" : "a", "to" : "b", "weight" : 6},
					{"from" : "a", "to" : "g", "weight" : 8},
					{"from" : "a", "to" : "d", "weight" : 10},
					{"from" : "d", "to" : "e", "weight" : 6},
					{"from" : "e", "to" : "b", "weight" : 15},
					{"from" : "b", "to" : "h", "weight" : 13},
					{"from" : "b", "to" : "c", "weight" : 11},
					{"from" : "c", "to" : "h", "weight" : 3},
					{"from" : "g", "to" : "h", "weight" : 5},
					{"from" : "g", "to" : "i", "weight" : 5},
					{"from" : "e", "to" : "f", "weight" : 2},
					{"from" : "f", "to" : "g", "weight" : 4},
					{"from" : "f", "to" : "i", "weight" : 6},
					{"from" : "h", "to" : "i", "weight" : 7}
				]}`
	json.Unmarshal([]byte(data), &res)
	fmt.Println("graph", res)
	Kruskal(res)
}

//*/

func Kruskal(graph Graph) Graph {
	nodes := getNodes(graph)
	//fmt.Println("nodes: ", nodes)
	ret := doKruskal(graph, len(nodes))
	str, _ := json.Marshal(ret)
	fmt.Println(string(str))
	return ret
}

func getNodes(graph Graph) []string {
	length := len(graph.Edges)
	nodes := []string{}
	for i := 0; i < length; i++ {
		nodes = append(nodes, graph.Edges[i].From)
		nodes = append(nodes, graph.Edges[i].To)
	}
	return removeDuplicates(nodes)
}

func removeDuplicates(src []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for x := range src {
		if !encountered[src[x]] {
			result = append(result, src[x])
			encountered[src[x]] = true
		}
	}
	return result
}

func pressAnyKey() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func doKruskal(graph Graph, nNodes int) Graph {
	ret := Graph{}
	kruskalTable := []tableRow{}
	sort.Sort(ByEdge(graph.Edges))
	//printGraph(graph)
	counter := 0
	iterator := 0
	flag := 0
	for iterator < len(graph.Edges) {
		fmt.Println("-->", graph.Edges[iterator])
		cflag := nodeIsConnected(kruskalTable, graph.Edges[iterator].From, graph.Edges[iterator].To)
		if cflag < 0 {
			kruskalTable = append(kruskalTable, tableRow{graph.Edges[iterator].From, graph.Edges[iterator].To, graph.Edges[iterator].Weight, flag})
			ret.Edges = append(ret.Edges, graph.Edges[iterator])
			flag++
			counter++
		} else if notCycle(kruskalTable, graph.Edges[iterator].From, graph.Edges[iterator].To) {
			kruskalTable = updateFlag(kruskalTable, graph.Edges[iterator].From, cflag)
			kruskalTable = updateFlag(kruskalTable, graph.Edges[iterator].To, cflag)
			kruskalTable = append(kruskalTable, tableRow{graph.Edges[iterator].From, graph.Edges[iterator].To, graph.Edges[iterator].Weight, cflag})
			ret.Edges = append(ret.Edges, graph.Edges[iterator])
			counter++
			//fmt.Println("\t", kruskalTable)
		}
		iterator++
	}
	//fmt.Println(kruskalTable)
	//pressAnyKey()
	//printGraph(ret)

	return ret
}

func PrintGraph(graph Graph) {
	fmt.Println("Printing Graph")
	length := len(graph.Edges)
	for i := 0; i < length; i++ {
		fmt.Println("edges: ", graph.Edges[i])
	}
}

func nodeIsConnected(table []tableRow, node1, node2 string) int {
	length := len(table)
	flags := []int{}
	for i := 0; i < length; i++ {
		row := table[i]
		if row.From == node1 || row.To == node1 || row.From == node2 || row.To == node2 {
			flags = append(flags, row.Flag)
		}
	}
	if len(flags) > 0 {
		sort.Ints(flags)
		return flags[0]
	}
	return -1
}

func notCycle(table []tableRow, node1, node2 string) bool {
	incidences := []tableRow{}
	for i := 0; i < len(table); i++ {
		row := table[i]
		if row.From == node1 || row.To == node1 || row.From == node2 || row.To == node2 {
			incidences = append(incidences, row)
		}
	}

	//checking if incidences has 2 in the same flag
	flagN1 := -1
	flagN2 := -1
	for i := 0; i < len(incidences); i++ {
		row := incidences[i]
		if row.From == node1 || row.To == node1 {
			flagN1 = row.Flag
		}
		if row.From == node2 || row.To == node2 {
			flagN2 = row.Flag
		}
	}
	return flagN1 != flagN2
}

func updateFlag(table []tableRow, node string, flag int) []tableRow {
	for i := 0; i < len(table); i++ {
		if table[i].From == node || table[i].To == node {
			table[i].Flag = flag
			fmt.Println(table[i], " Flag updated to", flag)
		}
	}
	return table
}
