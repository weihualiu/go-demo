package main

import "fmt"
import "bytes"
import "time"
import "runtime"

func main() {
	//var set_t HashSet
	set_t := NewHashSet()
	done := make(chan int)
	
	for i := 0; i < 10; i++ {
		go test(set_t, done, i)
	}

	for j := 0; j < 10; j++ {
		<-done
	}
	set_t.Clear()
	time.Sleep(10*time.Second)
	runtime.GC()
	time.Sleep(10*time.Second)
	
	fmt.Println(set_t.String())
}

func test(t *HashSet, done chan int, prefix int) {
	for i := 0; i < 10; i++ {
		t.Add("test" + string(i) + string(prefix))
	}
	//time.Sleep(10*time.Second)
	done <- 1
}


type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	fmt.Println("new hashset pointer")
	return &HashSet{m: make(map[interface{}]bool)}
}

//add
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true	
	}
	return false
}

//delete
func (set *HashSet) Delete(e interface{}) {
	delete(set.m, e)
}

//clear
func (set *HashSet) Clear() {
	//直接将map指向新地址对象
	set.m = make(map[interface{}]bool)
}

//contains
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//len
func (set *HashSet) Len() int {
	return len(set.m)
}

//same
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//elements
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		}else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot

}

//string
func (set *HashSet) String() string {
	// buffer 
	var buf bytes.Buffer
	buf.WriteString("set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

