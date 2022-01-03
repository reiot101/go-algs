# skiplist

https://en.wikipedia.org/wiki/Skip_list

Skiplist is a data struct that supports the following operations:
- Insert: O(logn) avg, O(n) worst case
- Search: O(logn) avg, O(n) worst case
- Delete: O(logn) avg, O(n) worst case

Space complexity: O(n) avg, O(nlogn) worst case

Skiplist has multiple levels, each level is a sorted (singly) linked list with jump links from current level to the next level of the elements with the same value.

The last level has every elements, each level up contains a subset of the elements in the below level.

```
# Fast forward, like playing video.
 1  _  _  _  _  _  _  _  _ $     8x

 1  _  _  _  _  6  _  _  _ $     4x

 1  _  3  _  _  6  _  8  _ $     2x

 1  2  3  4  _  6  7  8  9 $     1x
```

### Search(T): Start from the top-left head
- find the greatest node N in current layer that has value <= T
- if N.val == T: return True else go to the next layer
- repeat steps 1-2 until current nodes becomes nil.

```
# Search(7) = True
 1  _  _  _  _  _  _  _  _ $
 |
 1  -  -  -  -  6  _  _  _ $
                |
 1  _  3  _  _  6  _  8  _ $ 
                |
 1  2  3  4  _  6  7  8  9 $ 

 1 6 6 7
```

```
# Search(5) = False
 1  _  _  _  _  _  _  _  _ $
 |
 1  _  _  _  _  6  _  _  _ $
 |
 1  -  3  _  _  6  _  8  _ $
       |
 1  2  3  4  _  6  7  8  9 $

 1 1 3 4
```

### Delete(T): Start from the top-left head
- find the greatest node N in current layer that has value <= T
- if N.next.val == T: N.next.next go to the next layer
- repeat steps 1-3 until current nodes becomes nil.

```
# Delete(6) = True
 1  _  _  _  _  _  _  _  _ $
 |
 1  _  _  _  _ [6] _  _  _ $  1 -> $
 |
 1  -  3  _  _ [6] _  8  _ $  3 -> 8
       |
 1  2  3  4  _ [6] 7  8  9 $  4 -> 7
```

### Insert(T): Start from the top-left head
- For each level i, find and collect the largest node N (i) st, N(i).val < T
- Go to the next level
- Repeat 1-2 until to the bottom
- From the bottom to top, insert the element right after N(i) with probability of 1/2 (d-i)

```
level   1: 1/2 (d-1)
...
level d-2: 0.25
level d-1: 0.5
level   d: 1
```
- Connect new elements between layers

Why 1/2?
Each level has half size of the net level,
total space is n+n/2+n/4+... = 2n = O(n)

```
 1  _  _  _  _  _  _  _  _ $     8x
 |
 1  _  _  _  _  6  _  _  _ $     4x
 |              |
 1  _  3  _  _  6  _  8  _ $     2x
 |     |        |     |
 1  2  3  4  _  6  7  8  9 $     1x
```


```
 # insert(5)

 1  _  _  _  _  _  _  _  _ $              1  _  _  _  _  _  _  _  _ $   p=0.125
 |                                        |
 1  _  _  _  _  6  _  _  _ $              1  _  _  _  5  6  _  _  _ $   p=0.25
 |              |                 =>      |           |  |   
 1  _  3  _  _  6  _  8  _ $              1  _  3  _  5  6  _  8  _ $   p=0.5
 |     |        |     |                   |     |     |  |     |
 1  2  3  4  _  6  7  8  9 $              1  2  3  4  5  6  7  8  9 $   p=1

  Collect largest nodes < T                 Insert right after collected 
                                              nodes with prob 1/2 (d-i)
```