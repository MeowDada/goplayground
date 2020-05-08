package concurmap

import (
	"testing"
	"math/rand"
)

type item struct {
	key interface{}
	value interface{}
}

func randomKey() interface{} {
	return rand.Uint64()
}

func randomValue() interface{} {
	return rand.Uint64()
}

func generateItem() item {
	return item{
		key:   randomKey(),   
		value: randomValue(),
	}
}

func generateItems(count int) []item {
	items := make([]item, count)
	for i := 0 ; i < count ; i++ {
		items[i] = generateItem()
	}
	return items
}

func TestSimpleStore(t *testing.T) {
	cm := New()
	items := []item{
		{"hello", 1994},
		{2019, 1213},
		{"world", "beautiful"},
		{"qnap", "goodgood"},
	}

	for _, item := range items {
		cm.Store(item.key, item.value)
	}
	cm.Dump()
}

func TestSimpleLoad(t *testing.T) {
	cm := New()
	items := []item{
		{"hello", 1994},
		{2019, 1213},
		{"world", "beautiful"},
		{"qnap", "goodgood"},
	}

	for _, item := range items {
		cm.Store(item.key, item.value)
	}

	// Should return corresponding value
	for _, item := range items {
		get, _ := cm.Load(item.key)
		want := item.value
		if get != want {
			t.Errorf("expect %v, but get %v", want, get)
		}
	}

	// Should return nil object and false
	val, ok := cm.Load("???")
	if val != nil || ok != false {
		t.Errorf("expect to get nil object and false, but get %v, %v", val, ok)
	}
}

func TestSimpleDelete(t *testing.T) {
	cm := New()
	key, value := 123, 456
	cm.Store(key, value)
	cm.Delete(key)
	val, ok := cm.Load(key)
	if val != nil || ok != false {
		t.Errorf("expect to get a nil object and false, but get %v, %v", val, ok)
	}
}

func benchmarkNormalStore(m map[interface{}]interface{}, items []item) {
	for _, item := range items {
		m[item.key] = item.value
	}
}

func BenchmarkNormalStore10(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(10)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalStore100(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(100)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalStore500(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(500)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalStore1000(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(1000)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalStore2500(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(2500)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalStore10000(b *testing.B) {
	m := make(map[interface{}]interface{})
	items := generateItems(10000)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkNormalStore(m, items)
	}
}

func BenchmarkNormalLoad(b *testing.B) {

}

func BenchmakrNormalDelete(b *testing.B) {

}

func benchmarkStore(m ConcurrentMap, items []item) {
	for _, item := range items {
		m.Store(item.key, item.value)
	}
}

func BenchmarkStore10(b *testing.B) {
	cm := New()
	items := generateItems(10)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}

func BenchmarkStore100(b *testing.B) {
	cm := New()
	items := generateItems(100)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}

func BenchmarkStore500(b *testing.B) {
	cm := New()
	items := generateItems(500)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}

func BenchmarkStore1000(b *testing.B) {
	cm := New()
	items := generateItems(1000)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}

func BenchmarkStore2500(b *testing.B) {
	cm := New()
	items := generateItems(2500)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}

func BenchmarkStore10000(b *testing.B) {
	cm := New()
	items := generateItems(10000)
	for n := 0 ; n <= b.N ; n++ {
		benchmarkStore(cm, items)
	}
}


func BenchmarkLoad(b *testing.B) {

}

func BenchmarkDelete(b *testing.B) {

}