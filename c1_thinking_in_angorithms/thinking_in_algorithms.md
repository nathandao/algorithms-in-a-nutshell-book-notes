# Thinking in algorithms

## 1. Understand the problem

The first step of tackling a problem is obviously understanding the problem. The example used is a convex hull problem: given a set of points. Stretch a rubber band to form an outer line for all the points produces a line called convex hull. Create a angorithm to find the outter set of points that form the convex hull.

- Identify the requirements and be able to come out with a quick visual sketch of what the problem discribes. For the convex hull example, this is a set of points, out of which we can draw a boundary line my connecting all outer points. The problem is to find a mathematical way to find the points that create this line for any set of points provided.

## 2. After identifying the problem, we can now solve the problem with one of these approaches:

### Naive solution: 

Most of the time, naive is the most obvioubs solution by reasoning, without going through much mathematical abstraction. in context of the convex hull problem, this means going through every point in the set and list out all points that cannot be included inside a triangle created my other 3 points.

### More intelligent approaches:

Intelligent approaches can be further sub-divided into 5 categories, ordered in layers of implementation abstractions:

- Greedy: Identify a mathematical solution to solve the problem, one datapoint at a time. 
- Divide and conquer: Use the mathematical solution in Greedy, but slit the dataset into 2 and process them at the same time. After that, combine the result to get the final solution.
- Parallel: Like divide and conquer, but now we can have ```n``` number of subsets instead of 2. This can make use of parallel processing power of multi-core processors.
- Approximation: Sometimes, when the problem does not require an exact answer, we can trade accuracy for speed by implementing an approximation algorithm that arrives at the final answer much quicker. In this case, the approximation should also provide an error rate for the result.
- Generalization: Use an already created general algorithm, and implement that algorithm to fit the problem's domain-specific questions.
