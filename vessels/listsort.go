package vessels

import "cmp"

/**
* Sort the list
 */
func SortList[T cmp.Ordered](list *List[T]) {
	if list.Len() > 0 {
		sortNodes(list.Begin(), list.Len())
	}
}

/** Merge two sorted lists of nodes separated by a pivot (mid)
 * list1 is [first, mid)
 * list2 is [mid, last)
 * list2 is merged into list1
 * returns the new first (last is also the new last)
 */
func mergeOrderedNodes[T cmp.Ordered](
	first *ListNode[T], mid *ListNode[T], last *ListNode[T],
) *ListNode[T] {
	// determine which node will be the new first
	newFirst := first
	if mid.value < first.value {
		newFirst = mid
	}

	// Step across the already ordered elements of list1 while inserting any runs
	// of list2 where they belong.
	for first != mid && mid != last {
		if mid.value < first.value {
			// determine the list2 run of values and splice them in
			run := mid
			next := run.next
			for next != last {
				if !(next.value < first.value) {
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

/**
* Simple recursive merge sort. Avoids walking the list by recursing by half
* down to single nodes and merging them back up.
* Returns: (newFirst, newLast]
 */
func sortNodes[T cmp.Ordered](first *ListNode[T], size int) (*ListNode[T], *ListNode[T]) {
	switch size {
	case 0:
		return first, first
	case 1:
		return first, first.next
	default:
		break
	}

	newFirst, mid := sortNodes(first, size/2)
	mid, newLast := sortNodes(mid, size-(size/2))
	newFirst = mergeOrderedNodes(newFirst, mid, newLast)
	return newFirst, newLast
}
