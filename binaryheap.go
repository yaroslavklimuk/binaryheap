package binaryheap

// BinaryHeap
// parent:    (i-1)/2
//
//	            |
//	          index
//	         /     \
//	left: 2i+1     right: 2i+2
type BinaryHeap struct {
	internalSlice []int
	swap          int
	maxTop        bool
}

func (bh *BinaryHeap) PopTop() int {
	top := bh.internalSlice[0]
	bh.internalSlice[0] = bh.internalSlice[len(bh.internalSlice)-1]
	bh.internalSlice = bh.internalSlice[:len(bh.internalSlice)-1]
	defer func() { bh.heapifyDown(0) }()
	return top
}

func (bh *BinaryHeap) heapifyDown(index int) {
	var indexToSwap int
	var stop bool

	for !stop {
		leftIndex := getLeftChildIndex(index)
		rightIndex := getRightChildIndex(index)
		currLength := len(bh.internalSlice)
		if leftIndex >= currLength {
			return
		}
		if rightIndex >= currLength {
			indexToSwap = leftIndex
		} else {
			if bh.firstBelowSecond(leftIndex, rightIndex) {
				indexToSwap = leftIndex
			} else {
				indexToSwap = rightIndex
			}
		}

		bh.swap = bh.internalSlice[indexToSwap]
		bh.internalSlice[indexToSwap] = bh.internalSlice[index]
		bh.internalSlice[index] = bh.swap
		index = indexToSwap
	}
}

func (bh *BinaryHeap) Insert(item int) {
	bh.internalSlice = append(bh.internalSlice, item)
	bh.heapifyUp(len(bh.internalSlice) - 1)
}

func (bh *BinaryHeap) heapifyUp(index int) {
	prntIndex := getParentIndex(index)
	if prntIndex <= 0 {
		return
	}
	for bh.firstBelowSecond(index, prntIndex) {
		bh.swap = bh.internalSlice[prntIndex]
		bh.internalSlice[prntIndex] = bh.internalSlice[index]
		bh.internalSlice[index] = bh.swap

		index = prntIndex
		prntIndex = getParentIndex(index)
		if prntIndex <= 0 {
			return
		}
	}
}

func (bh *BinaryHeap) firstBelowSecond(first int, second int) bool {
	secondValue := bh.internalSlice[second]
	firstValue := bh.internalSlice[first]
	if bh.maxTop {
		return secondValue < firstValue
	} else {
		return secondValue > firstValue
	}
}

func getLeftChildIndex(index int) int {
	return 2*index + 1
}

func getRightChildIndex(index int) int {
	return 2*index + 2
}

func getParentIndex(index int) int {
	var parentIndex int
	parentIndex = (index - 1) / 2
	return parentIndex
}
