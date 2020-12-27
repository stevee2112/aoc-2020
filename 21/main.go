package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"regexp"
	"strings"
	"sort"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

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
		if len(possibleAllergens) == 0 {
			sum += ingredientFoodCounter[i]
		}
	}

	// Part 2
	// reduce
	for needToReduce(ingredientIndex) {
		for i, possibleAllergens := range ingredientIndex {
			if len(possibleAllergens) == 1 {
				for a, _ := range possibleAllergens {
					ingredientIndex = reduce(i, a, ingredientIndex)
				}
			}
		}
	}

	finalAllergens := map[string]string{}
	for i, possibleAllergens := range ingredientIndex {
		if len(possibleAllergens) == 1 {
			for a, _ := range possibleAllergens {
				finalAllergens[a] = i
			}
		}
	}

	dangerIngredients := []string{}
	keys := make([]string, 0, len(finalAllergens))

	for k := range finalAllergens {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		dangerIngredients = append(dangerIngredients, finalAllergens[k])
	}


	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %s\n", strings.Join(dangerIngredients, ","))
}

func needToReduce(ingredientIndex map[string]map[string]bool) bool {
	for _, possibleAllergens := range ingredientIndex {
		if len(possibleAllergens) > 1 {
			return true
		}
	}

	return false
}

func reduce(ingredient string, allergen string, ingredientIndex map[string]map[string]bool) map[string]map[string]bool {
	for i, _ := range ingredientIndex {
		if i == ingredient {
			continue
		}

		if _,ok := ingredientIndex[i][allergen]; ok {
			delete(ingredientIndex[i], allergen)
		}
	}

	return ingredientIndex
}

func possibleAllergen(ingredient string, allergen string, foods []map[string]bool) bool {

	for _, ingredients := range foods {
		if _, ok := ingredients[ingredient]; !ok {
			return false
		}
	}

	return true;
}
