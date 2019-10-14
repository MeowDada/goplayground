package job

import (
	"fmt"
	"math/rand"
	"strconv"
)

func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandomNumberToStricleng(100))
	}
	return jobs
}

func RandomNumberToString(maxDuration int) string {
	num := rand.Intn(maxDuration)
	return strconv.FormatInt(int64(num), 10)
}

func DoWork(numString string, id int) {
	fmt.Printf("Worker<%d>: %s\n", id, numString)
}