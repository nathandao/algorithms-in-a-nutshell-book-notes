# Sorting algorithms

## Stable sort

Is an algorithm that keeps the original odering of 2 elements if they are equal.

## Types of Sorting algorithms 

### Tranposition Sorting: 

Done by comparing and shifting elements inside the provided list. Includes:

- Insertion sort: Insert smaller element before its larger element. Best: O(n), Worst: O(n^2)
- Selection sort: Select largest and place at the end. O(n^2)
- Heap sort: Construct tree structure to quickly figure out the largest value in n/2 steps. Repeat for nth times. O(nlog(n))

### Partition-Based Sorting:

Divide the list into partitions (smaller chunks) in order to reduce processing step required to finish sorting. Example:

- Quick sort: Create 2 partition with all smaller values on left & larger values on right of the pivot. Use recurring func to keep doing this for each of the 2 parts. Possible to use Insertion sort once partition size is reduced to a certain value. Best Avg: O(nlog(n)), Worst: O(n^2)

### Bucket Sort - Sorting without Comparisons

Use hash function to categorize values into buckets. Once done, just join the buckets together to get the final sorted list.

- Hash function: Convert value into an index number corresponds to a bucket.
- Number of buckets: If too large, will affect O(n), since bucket sort is generally O(n + k).
- Implemented correctly: O(n) for all cases

## Usage suggestions

Only a few items: Insertion Sort
Items arn mostly sorted: Insertion Sort
Concerned about worst-case scenarios: Heap Sort
Interested in a goot average-case behaviour: Quick Sort 
Items are drawn from a unitform dense universe: Bucket Sort
Desire to write as little code as possible: Insertion Sort 
Require stable sort: Merge Sort
