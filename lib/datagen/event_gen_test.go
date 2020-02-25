package datagen

import (
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
	"sort"
	"testing"
)

/*
generateTestInitPopulation creates a population in its init phase - that is,
just one generation, everyone was just born. generation 1 will have a random
number of people, between 20-50
*/
func generateTestInitPopulation() (pop population) {
	numPeople := rand.Intn(31) + 20
	pop = make(population, 1)
	pop[0] = make(generation, numPeople)
	// create the born people, born at year 0
	for i := range pop[0] {
		pop[0][i] = randomPerson(0)
	}
	return
}

func TestDeath(t *testing.T) {
	iterations := 100
	means := make([]float64, iterations)
	stddev := make([]float64, iterations)
	tenPercentile := make([]float64, iterations)
	ninetyPercentile := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		year := 0
		pop := generateTestInitPopulation()
		data := make([]float64, len(pop[0]))
		for !pop[0].allDead() {
			for j := 0; j < len(pop[0]); j++ {
				deathCheck(&pop[0][j], &pop, year)
			}
			year++
		}
		year--
		for i, person := range pop[0] {
			data[i] = float64(person.deathYear)
		}
		sort.Float64s(data)
		means[i] = stat.Mean(data, nil)
		stddev[i] = math.Sqrt(stat.Variance(data, nil))
		tenPercentile[i] = stat.Quantile(0.1, stat.Empirical, data, nil)
		ninetyPercentile[i] = stat.Quantile(0.9, stat.Empirical, data, nil)
	}
	t.Logf("Average metrics for %v iterations. Population size ranges from 20 - 30", iterations)
	t.Logf("mean = %v", stat.Mean(means, nil))
	t.Logf("std-dev = %v", stat.Mean(stddev, nil))
	t.Logf(
		"10%% percentile: %v, 90%% percentile: %v",
		stat.Mean(tenPercentile, nil),
		stat.Mean(ninetyPercentile, nil),
	)
}
