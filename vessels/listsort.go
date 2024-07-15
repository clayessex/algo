package vessels

import "cmp"

// Sort the list according to cmp.Less
// Requires that T be cmp.Ordered, which is not a requirment of the underlying list
func SortList[T cmp.Ordered](list *List[T]) {
	if list.Len() > 0 {
		sortNodes(list.Begin(), list.Len(), cmp.Less)
	}
}

// Sort the list according to comp comparison function
// comp is a comparison function that returns true if a is ordered before b
func SortListFunc[T any](list *List[T], comp func(a, b T) bool) {
	if list.Len() > 0 {
		sortNodes(list.Begin(), list.Len(), comp)
	}
}

// Simple recursive merge sort. Avoids walking the list by recursing by half
// down to single nodes and merging them back up.
// comp is a comparison function that returns true if a is ordered before b
// Returns: (newFirst, newLast]
func sortNodes[T any](first *ListNode[T], size int,
	comp func(a, b T) bool,
) (*ListNode[T], *ListNode[T]) {
	switch size {
	case 0:
		return first, first
	case 1:
		return first, first.next
	default:
		break
	}

	newFirst, mid := sortNodes(first, size/2, comp)
	mid, newLast := sortNodes(mid, size-(size/2), comp)
	newFirst = mergeOrderedNodes(newFirst, mid, newLast, comp)
	return newFirst, newLast
}

// Merge two sorted lists of nodes separated by a pivot (mid)
// list1 is [first, mid)
// list2 is [mid, last)
// list2 is merged into list1
// comp is a comparison function that returns true if a is ordered before b
// returns the new first (last is also the new last)
func mergeOrderedNodes[T any](
	first, mid, last *ListNode[T],
	comp func(a, b T) bool,
) *ListNode[T] {
	// determine which node will be the new first
	newFirst := first
	if comp(mid.value, first.value) {
		newFirst = mid
	}

	// Step across the already ordered elements of list1 while inserting any runs
	// of list2 where they belong.
	for first != mid && mid != last {
		if comp(mid.value, first.value) {
			// determine the list2 run of values and splice them in
			run := mid
			next := run.next
			for next != last {
				if !comp(next.value, first.value) {
					break
				}
				next = next.next
			}
			splice(first, run, next)
			mid = next
		} else {
			// advance the insertion point across list1
			first = first.next
		}
	}

	return newFirst
}

// Alternative non-recursive merge sort
func SortListAlt[T cmp.Ordered](list *List[T]) {
	SortListFuncAlt(list, cmp.Less)
}

// Non-recursive merge sort
func SortListFuncAlt[T any](list *List[T], comp func(a, b T) bool) {
	if list.Len() <= 1 {
		return
	}

	buckets := make([]*List[T], 64)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = NewList[T]()
	}
	topBucket := 0
	hold := NewList[T]()

	for list.Len() > 0 {
		splice(hold.Begin(), list.Begin(), list.Begin().Next()) // take 1
		list.len--
		hold.len++

		i := 0
		for i != topBucket && buckets[i].len != 0 {
			ListMergeFunc(buckets[i], hold, comp)
			hold.Swap(buckets[i])
			i++
		}
		hold.Swap(buckets[i])
		if i == topBucket {
			topBucket++
		}
	}

	for i := 1; i < topBucket; i++ {
		ListMergeFunc(buckets[i], buckets[i-1], comp)
	}

	buckets[topBucket-1].Swap(list)
}
