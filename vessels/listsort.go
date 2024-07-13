package vessels

import "cmp"

/**
 * Sort the list according to cmp.Less
 */
func SortList[T cmp.Ordered](list *List[T]) {
	if list.Len() > 0 {
		sortNodes(list.Begin(), list.Len(), cmp.Less)
	}
}

/**
 * Sort the list according to comp comparison function
 */
func SortListFunc[T cmp.Ordered](list *List[T], comp func(a, b T) bool) {
	if list.Len() > 0 {
		sortNodes(list.Begin(), list.Len(), comp)
	}
}

/**
* Simple recursive merge sort. Avoids walking the list by recursing by half
* down to single nodes and merging them back up.
* Returns: (newFirst, newLast]
 */
func sortNodes[T cmp.Ordered](first *ListNode[T], size int,
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

/** Merge two sorted lists of nodes separated by a pivot (mid)
 * list1 is [first, mid)
 * list2 is [mid, last)
 * list2 is merged into list1
 * comp is a comparison function
 * returns the new first (last is also the new last)
 */
func mergeOrderedNodes[T cmp.Ordered](
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
