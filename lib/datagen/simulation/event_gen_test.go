package simulation

import (
	"github.com/chaseWilliams/family-map/lib/datagen/family"
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
	"sort"
	"testing"
)

/*
generateTestInitPopulation creates a population in its init phase - that is,
just one generation, everyone was just born. generation 1 will have a random
number of people, from [min, max]
*/
func generateTestInitPopulation(min int, max int) (pop family.Population) {
	numPeople := rand.Intn(min+1) + max - min
	pop = family.CreatePopulation(numPeople)
	return
}

func basicMetrics(t *testing.T, data []float64) {
	sort.Float64s(data)
	t.Logf("mean = %v", stat.Mean(data, nil))
	t.Logf("std-dev = %v", math.Sqrt(stat.Variance(data, nil)))
	t.Logf(
		"10%% percentile: %v, 90%% percentile: %v",
		stat.Quantile(0.1, stat.Empirical, data, nil),
		stat.Quantile(0.9, stat.Empirical, data, nil),
	)
}

func TestDeath(t *testing.T) {
	iterations := 100
	means := make([]float64, iterations)
	stddev := make([]float64, iterations)
	tenPercentile := make([]float64, iterations)
	ninetyPercentile := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		year := 0
		pop := generateTestInitPopulation(30, 50)
		data := make([]float64, len(pop[0]))
		for !pop[0].AllDead() {
			for j := 0; j < len(pop[0]); j++ {
				deathCheck(pop[0][j], &pop, year)
			}
			year++
		}

		for i, person := range pop[0] {
			data[i] = float64(person.DeathYear())
		}
		sort.Float64s(data)
		means[i] = stat.Mean(data, nil)
		stddev[i] = math.Sqrt(stat.Variance(data, nil))
		tenPercentile[i] = stat.Quantile(0.1, stat.Empirical, data, nil)
		ninetyPercentile[i] = stat.Quantile(0.9, stat.Empirical, data, nil)
	}
	t.Logf("Average metrics for %v iterations. Population size ranges from 20 - 50", iterations)
	t.Logf("mean = %v", stat.Mean(means, nil))
	t.Logf("std-dev = %v", stat.Mean(stddev, nil))
	t.Logf(
		"10%% percentile: %v, 90%% percentile: %v",
		stat.Mean(tenPercentile, nil),
		stat.Mean(ninetyPercentile, nil),
	)
}

// TODO
// SAMPLE TIMELINE
func TestMarriageAndDivorce(t *testing.T) {
	/*
		metrics:
			average percent of homosexual marriages
			number of people married x times
			average age for x marriage (1st marriage...)
			average length of x marriage
			average time person was single
			sample person's marriage timeline (marriages + divorces)
	*/
	iterations := 5
	avgPercentHomosexual := make([]float64, iterations)
	// index 0: data for people married 0 times, etc.
	avgPercentPeopleMarriedXTimes := make([][]float64, 0)
	avgAgeForXMarriage := make([][]float64, 0)
	avgLengthOfXMarriage := make([][]float64, 0)
	avgTimeSingleBeforeMarriageX := make([][]float64, 0)

	for iter := 0; iter < iterations; iter++ {
		pop := generateTestInitPopulation(70, 200)
		year := 0
		for !pop[0].AllDead() {
			for j := 0; j < len(pop[0]); j++ {
				deathCheck(pop[0][j], &pop, year)
				marriageCheck(pop[0][j], &pop, year)
				divorceCheck(pop[0][j], &pop, year)
			}
			year++
		}
		// these are all running totals
		numHomosexualMarriages := 0.0
		numPeopleMarriedXTimes := make([]float64, 0)
		ageForXMarriage := make([][]float64, 0)
		lengthOfXMarriage := make([][]float64, 0)
		timeSingleBeforeMarriageX := make([][]float64, 0)

		for _, person := range pop[0] {
			if !person.IsStraight() {
				numHomosexualMarriages++
			}
			marriageCount := len(person.Spouses())

			numPeopleMarriedXTimes = expandIfNeeded(numPeopleMarriedXTimes, marriageCount + 1)
			ageForXMarriage = expand2DIfNeeded(ageForXMarriage, marriageCount)
			lengthOfXMarriage = expand2DIfNeeded(lengthOfXMarriage, marriageCount)
			timeSingleBeforeMarriageX = expand2DIfNeeded(timeSingleBeforeMarriageX, marriageCount)

			numPeopleMarriedXTimes[marriageCount]++

			for i, year := range person.MarriageYears() {
				ageForXMarriage[i] = append(ageForXMarriage[i], float64(person.Age(year)))
			}

			divorceYears := person.DivorceYears()
			for i, year := range person.MarriageYears() {
				length := 0.0
				if i >= len(divorceYears) {
					length = float64(person.DeathYear() - year)
				} else {
					length = float64(divorceYears[i] - year)
				}
				lengthOfXMarriage[i] = append(lengthOfXMarriage[i], length)
			}

			for i, year := range person.MarriageYears() {
				if i == 0 {
					timeSingleBeforeMarriageX[i] = append(timeSingleBeforeMarriageX[i], float64(person.Age(year) - 18)) // single since eligible to be married
				} else {
					timeSingleBeforeMarriageX[i] = append(timeSingleBeforeMarriageX[i], float64(year - divorceYears[i-1]))
				}
			}
		}

		populationSize := float64(len(pop[0]))
		avgPercentHomosexual[iter] = numHomosexualMarriages / populationSize
		avgPercentPeopleMarriedXTimes = expand2DIfNeeded(
			avgPercentPeopleMarriedXTimes, 
			len(numPeopleMarriedXTimes), 
		)
		avgAgeForXMarriage = expand2DIfNeeded(
			avgAgeForXMarriage, 
			len(ageForXMarriage), 
		)
		avgLengthOfXMarriage = expand2DIfNeeded(
			avgLengthOfXMarriage, 
			len(lengthOfXMarriage), 
		)
		avgTimeSingleBeforeMarriageX = expand2DIfNeeded(
			avgTimeSingleBeforeMarriageX, 
			len(timeSingleBeforeMarriageX), 
		)
		for i, num := range numPeopleMarriedXTimes {
			avgPercentPeopleMarriedXTimes[i] = append(avgPercentPeopleMarriedXTimes[i], num / populationSize)
		}
		for i, nums := range ageForXMarriage {
			avg := stat.Mean(nums, nil)
			if !math.IsNaN(avg) {
				avgAgeForXMarriage[i] = append(avgAgeForXMarriage[i], avg)
			}
		}
		for i, nums := range lengthOfXMarriage {
			avgLengthOfXMarriage[i] = append(avgLengthOfXMarriage[i], stat.Mean(nums, nil))
		} 
		for i, nums := range timeSingleBeforeMarriageX {
			avgTimeSingleBeforeMarriageX[i] = append(avgTimeSingleBeforeMarriageX[i], stat.Mean(nums, nil))
		}
	}

	t.Logf("Metrics for %v iterations", iterations)
	t.Logf(
		"Average percent of homosexual marriages: %.2f%%",
		stat.Mean(avgPercentHomosexual, nil) * 100,
	)
	t.Log("Average percent of people married x times:")
	for i, nums := range avgPercentPeopleMarriedXTimes {
		t.Logf(
			"%d\t%.2f%%",
			i,
			stat.Mean(nums, nil) * 100,
		)
	}
	t.Log("Average age of people at x marriage:")
	for i, nums := range avgAgeForXMarriage {
		t.Logf(
			"%d\t%.2f +/- %.2f years old",
			i + 1,
			stat.Mean(nums, nil),
			math.Sqrt(stat.Variance(nums, nil)),
		)
	}
	t.Log("Average length of x marriage:")
	for i, nums := range avgLengthOfXMarriage {
		t.Logf(
			"%d\t%.2f +/- %.2f years",
			i + 1,
			stat.Mean(nums, nil),
			math.Sqrt(stat.Variance(nums, nil)),
		)
	}
	t.Log("Average time single before marriage x")
	for i, nums := range avgTimeSingleBeforeMarriageX {
		t.Logf(
			"%d\t%.2f +/- %.2f years",
			i + 1,
			stat.Mean(nums, nil),
			math.Sqrt(stat.Variance(nums, nil)),
		)
	}
}

func TestBabies(t *testing.T) {
	/*
	metrics:
		- average num of children per marriage
		- sample timeline of having children
	*/
	iterations := 5
	avgNumChildren := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		year := 0
		pop := generateTestInitPopulation(30, 50)
		data := make([]float64, len(pop[0]))
		for !pop[0].AllDead() {
			for j := 0; j < len(pop[0]); j++ {
				deathCheck(pop[0][j], &pop, year)
				marriageCheck(pop[0][j], &pop, year)
				divorceCheck(pop[0][j], &pop, year)
				babyCheck(pop[0][j], &pop, year)
			}
			year++
		}

		for i, person := range pop[0] {
			data[i] = float64(person.DeathYear())
		}
		sort.Float64s(data)
		means[i] = stat.Mean(data, nil)
		stddev[i] = math.Sqrt(stat.Variance(data, nil))
		tenPercentile[i] = stat.Quantile(0.1, stat.Empirical, data, nil)
		ninetyPercentile[i] = stat.Quantile(0.9, stat.Empirical, data, nil)
	}
	t.Logf("Average metrics for %v iterations. Population size ranges from 20 - 50", iterations)
	t.Logf("mean = %v", stat.Mean(means, nil))
	t.Logf("std-dev = %v", stat.Mean(stddev, nil))
	t.Logf(
		"10%% percentile: %v, 90%% percentile: %v",
		stat.Mean(tenPercentile, nil),
		stat.Mean(ninetyPercentile, nil),
	)
}

func TestPopulationSizeChange(t *testing.T) {

}

func expandIfNeeded(slice []float64, length int) []float64 {
	if len(slice) >= length {
		return slice
	}
	newSlice := make([]float64, length)
	copy(newSlice, slice)
	return newSlice
}

func expand2DIfNeeded(slice [][]float64, length int) [][]float64 {
	if len(slice) >= length {
		return slice
	}
	newSlice := make([][]float64, length)
	copy(newSlice, slice)
	for i := len(slice); i < length; i++ {
		newSlice[i] = make([]float64, 0)
	}
	return newSlice
}
