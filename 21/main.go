package main

import (
	//"stevee2112/aoc-2020/types"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"regexp"
	"strings"
	//"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	ingredientFoodCounter := map[string]int{}
	ingredientIndex := map[string]map[string]bool{}
	allergenIndex := map[string][]map[string]bool{}

	for scanner.Scan() {
		inputStr := scanner.Text()

		r := regexp.MustCompile(`(^[^\(]+)\(contains ([^)]+)\)`)
		match := r.FindAllSubmatch([]byte(inputStr), -1)

		ingredientsStr := strings.TrimSpace(string(match[0][1]))
		allergensStr := strings.TrimSpace(string(match[0][2]))

		ingredients := strings.Split(ingredientsStr, " ")
		allergens := strings.Split(allergensStr, ", ")

		for _,i := range ingredients {
			ingredientFoodCounter[i]++

			if _,ok := ingredientIndex[i]; !ok {
				ingredientIndex[i] = map[string]bool{}
			}

			for _,a := range allergens {
				ingredientIndex[i][a] = true
			}
		}

		for _,a := range allergens {
			if _,ok := allergenIndex[a]; !ok {
				allergenIndex[a] = []map[string]bool{}
			}

			iIndex := map[string]bool{}
			for _,i := range ingredients {
				iIndex[i] = true
			}

			allergenIndex[a] = append(allergenIndex[a], iIndex)
		}
	}

	for i,allergens := range ingredientIndex {
		for a := range allergens {
			possible := possibleAllergen(i, a, allergenIndex[a])

			if !possible {
				delete(ingredientIndex[i], a)
			}
		}
	}

	// Part 1
	sum := 0
	for i, possibleAllergens := range ingredientIndex {
		fmt.Println(i, possibleAllergens)
		if len(possibleAllergens) == 0 {
			sum += ingredientFoodCounter[i]
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", 0)
}

func possibleAllergen(ingredient string, allergen string, foods []map[string]bool) bool {

	for _, ingredients := range foods {
		if _, ok := ingredients[ingredient]; !ok {
			return false
		}
	}

	return true;
}
