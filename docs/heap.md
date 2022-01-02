# heap

“A heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap property” — wikipedia

Heap property (min heap): parent’s key is less than children’s key

> Heap actually stores in an array.
```
Index relationship (0-based)
Parent => children: n => 2n + 1, 2n + 2
Child => parent: n => (n-1)/2
```

> Build a heap: convert an array to heap
```
for i=(n/1)2 to 0
    siftDown(i)

Time complexity: n/2 * log(n) = O(nlogn)
This is an upper bound. Tighter bound: O(n)
```

Applications:
- Heapsort O(nlogn)
- Dijkstra’s algorithm O(|E|log|v|)
- Priority Queue
- Selection algorithm
    - Select top k elements among n
    - Sorting: O(nlogn)
    - Binary heap: O(n+klogn)