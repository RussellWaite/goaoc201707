package main

import (
	"bufio"
	"errors"
	"learngo/goaoc201707/discs"
	"os"
)

func SolvePart1(path string) (string, error) {
	data, err := readInput(path)
	if err != nil {
		return "", err
	}
	mappy := discs.ParseStrings(data)
	return findOrphan(mappy)
}

func SolvePart2(path string, rootNode string) (int, error) {
	data, err := readInput(path)
	if err != nil {
		return 0, err
	}
	mappy := discs.ParseStrings(data)
	return findImbalance(mappy, rootNode)
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func findOrphan(data map[string]*discs.Disc) (string, error) {
	for _, v := range data {
		if v.Parent == nil {
			return v.Name, nil
		}
	}
	return "", errors.New("failed to find a parentless disc")
}

func findImbalance(data map[string]*discs.Disc, rootNode string) (int, error) {
	result, err := recurseToFindImbalance(&data, rootNode)
	if err != nil {
		return 0, err
	}

	discrepancy, err := result.Discrepancy()
	if err != nil {
		return 0, err
	}

	return discrepancy, nil
}

func recurseToFindImbalance(data *map[string]*discs.Disc, element string) (*discs.Disc, error) {
	disc, exists := (*data)[element]
	if !exists {
		return nil, errors.New("The key is not in the map" + element)
	}

	for _, v := range disc.Kids {
		child, _ := recurseToFindImbalance(data, v.Name)
		if child != nil && !child.IsBalanced() {
			return child, nil
		}
	}
	if !disc.IsBalanced() {
		return disc, nil
	}

	return nil, errors.New("ran out of options")
}
