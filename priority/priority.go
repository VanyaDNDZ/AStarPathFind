package priority

type PriorityItem interface {
	GetPriority() int
	SetIndex(value int)
	Eq(interface{}) bool
}


type PriorityQueue []*interface{}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return (*pq[i]).(PriorityItem).GetPriority() > (*pq[j]).(PriorityItem).GetPriority()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	(*pq[i]).(PriorityItem).SetIndex(j)
	(*pq[j]).(PriorityItem).SetIndex(i)
}

func (pq *PriorityQueue) Get(i int) interface{} {
	return *(*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x
	(item).(PriorityItem).SetIndex(n)
	*pq = append(*pq, &item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	(*item).(PriorityItem).SetIndex(-1) // for safety
	*pq = old[0 : n-1]
	return *item
}

func (pq *PriorityQueue) Has(element PriorityItem) bool {
	result := false
	for i := 0; i < pq.Len(); i++ {
		if element.Eq(pq.Get(i)) {
			result = true
		}
	}
	return result
}
