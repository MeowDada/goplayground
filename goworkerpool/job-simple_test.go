package workerpool

import (
	"testing"
)

func simpleAdd(args ...interface{}) (interface{}, error) {
	ret := 0
	for _, num := range args {
		ret += num.(int)
	}
	return ret, nil
}

func TestNewSimpleJob(t *testing.T) {
	job := NewSimpleJob(simpleAdd, 1, 2, 3, 4)
	result := job.Execute()

	want := 1+2+3+4
	get := result.result

	if want != get {
		t.Errorf("expect %d, but get %d", want, get)
	}
}