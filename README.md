## Algorithms in a nutshell

Random reading notes on [Algorithms in a nutshell](https://www.amazon.com/Algorithms-Nutshell-Practical-George-Heineman-ebook/dp/B01DAWPK6S/ref=mt_kindle?_encoding=UTF8&me=), and some examples re-written in Golang. 

## Installation

The examples require GO 1.5 and above to run.

## Running the code

Since I tried to isolate each chapter in its own folder, ```GOGPATH``` will have to be changed to each chapter's root folder in order to run properly.

For example, to run the sample codes for the sort chapter:

```
cd c4_sorting_algorithms
export GOPATH=`pwd` # Set GOPATH to be the chapter's root folder
cd src/sorts
go run *.go
```
