# The mathematics of algorithms

## Technical terms

Before digging into different performance classifications of algorithms, it's useful to first understand these terms:

- Problem instance: A problem instance is one particular dataset input to the algorithm. Dataset's structure depends on how we map each datapoint to the problem's features. For example, for a problem of definining housing price, we can define each data point has a set of numeric values of (area, year of construction). And the dataset is a collection of houses, which is the collection of (area, year of construction) points.
- Size of a problem instance: The number of data points in a dataset (or in a problem instance). In the housing example, this is the number of pairs of values (area, year of construction) in each dataset.
- Performance: A mathematical representation of the relationship of a algorithm's performance (processing time) against its instance size. In another word, how much processing time increases if the instance size increases by n times. 
- Rate of growth: How longer does the processing time take (denoted by a function of f(n)) if instance size increases by n times.
- Upper bound: A function that represents the worst case performance. Denoted by big O.
- Lower bound: A function that represents the best case performance. Denoted by Ω.

## Why not use absolute processing time (seconds, milliseconds, minutes, ...) as a metric for algorithm performance:

This is because absolute time indicator depends on hardware. A faster CPU with more cores obviously performs the same algorithm better than a slower CPU, and hardware performance is constantly changing. This means overtime, as processing power becomes faster, what takes 10s to process previously might take only 10ms today - which makes absolute processing time less of a concern.

A more appropriate performance metric should only concern about how processing time is affected as the problem instance size increases (also called rate of growth - which by nature, is independant from hardware performance) - meaning when instance size increases n times, how much longer does the algorithm take to process the dataset.

## Instance cases

- Best case: The best case scenario when the algorithm exhibits best runtime behaviour.
- Average case: The expected behaviour when executing the algorithm on random problem instances.
- Worst case: Algorithm exhibits worst runtime behaviour.

By knowing the performance at each case, we can better pick what algorithm to use to yeld the best overall performance.

## Performance family:

O( f(n) ), where f(n) is a function of n, means how much work does the algorithm's implementation has to do in respect to size n of the dataset.

For example, considering a set of data [n.1, n.2, n.3, ..., n.n] (with n number of elements) . If the algorithm needs to compare each element with the rest of n - 1 other elements, this means the number of comparisons needed to be done will be (n - 1) * n (since we need to do the n - 1 comparisons every element in the datase). From this, the big O performance family is O(n ^ 2).

Performance family in order, from best to worst.

- Constant: O(1)
- Logarithmic: O(log(n))
- Sublinear: O(n^d) for d < 1
- Linear: O(n)
- Linearithmic: O(nlog(n))
- Quadratic: O(n^2)
- Exponential: O(2^n)
  
