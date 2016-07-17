# Searching

## Sequential Search

Also called linear search - go through all elements in sequence until a match is found.

In order to declare a no match, the whole list has to be processed. Due to this nature, sequetial search is most basic and slowest way to search for an item in a list.

## Binary Search

Only works on ordered list. Divide list into 2 parts and narrow down the comparison by half each time by comparing search value to the mid value of the remaining list. If search_val > mid, narrow down to the bottom half, otherwise, narrow down to the top half. Performance: O(log(n))

The only let down is that the list has to be ordered in order to perform binary search. As such, binary search is not appropriate for always-changing dataset.

## Hash-based Search

Use hash table with hased-keys pointing to different values, which can be accessed immediately through memory. Hash keys are generated from a hash function.

This is a genreral workflow of a hash-based searh:

```
# 1. Index dataset values into a hash table
for each val of dataset:
    hashKey = hashFunction(val)
    hashTable[hashKey] = val
    
# 2. Search for value
search_value = user_enter
hashKey = hashFunction(search_value)
if hashTable[hashKey] contains search_value:
    return true # match found
else:
    return false # match not found
```

The key for hash-based search performance is the hash key-value table, which is stored in memory or RAM and allows instant access of the value through key. This is also one weakness of hash-based search: since it relies on memory storage of the hash table, if manage poorly, the hash table size can eat up all of your hardware's memory. 

### Some key concepts of hash-based search:

- Managing colisions: colisoin happens when hash function returns the same hashed key for different values. 2 ways of handling colisions: 1. Point the hash key to a list instead of an array. 2. Move down to the next empty hash-key value in the hash table and add the value there.

- Choosing size of hash table: Usually a prime number, or 2^k - 1 - although that value is not always prime. This is to maximize the chance of even distribution of data accross hash keys.
