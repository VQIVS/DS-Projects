### 2-SAT

#### proof of correctness

**implication graph and sccs:**
1. **construct the implication graph:**
   each clause $(a \lor b)$ can be converted to two implications: $(\neg a \rightarrow b)$ and $(\neg b \rightarrow a)$. this means that if $a$ is false, then $b$ must be true, and if $b$ is false, then $a$ must be true. these two implications can be represented as directed edges in the implication graph.

2. **find strongly connected components (sccs):**
   use kosaraju’s or tarjan’s algorithm to find all sccs in the graph. sccs are subgraphs where every pair of vertices is mutually reachable.

3. **check satisfiability:**
   if $x_i$ and $\neg x_i$ are in the same scc, the formula is unsatisfiable because it implies that $x_i$ and $\neg x_i$ must both be true simultaneously, which is impossible. otherwise, assign values based on the topological order of the sccs. this means that if $\neg x_i$ comes before $x_i$ in the topological order, $x_i$ must be true, and vice versa.

**example:**
consider the instance $(x_1 \lor x_2) \land (\neg x_1 \lor x_3) \land (\neg x_2 \lor \neg x_3)$:
- construct the graph: 
  - from $(x_1 \lor x_2)$, we get $(\neg x_1 \rightarrow x_2)$ and $(\neg x_2 \rightarrow x_1)$.
  - from $(\neg x_1 \lor x_3)$, we get $(x_1 \rightarrow x_3)$ and $(\neg x_3 \rightarrow \neg x_1)$.
  - from $(\neg x_2 \lor \neg x_3)$, we get $(x_2 \rightarrow \neg x_3)$ and $(x_3 \rightarrow \neg x_2)$.
- identify sccs: use kosaraju’s or tarjan’s algorithm to find sccs.
- assign values: if $x_1$ and $\neg x_1$ are not in the same scc, assign values based on the topological order.

### fibonacci heap

#### proof of amortized time complexity

**potential function:**
\[
\Phi(H) = t(H) + 2m(H)
\]
where $t(H)$ is the number of trees and $m(H)$ is the number of marked nodes. this potential function is used to calculate the amortized costs.

**insert operation:**
- actual cost: $O(1)$
- change in potential ($\Delta \Phi$): a new tree is added, so $\Delta \Phi = 1$
- amortized cost: $O(1) + 1 = O(1)$

**extract-min operation:**
- actual cost: $O(\log n + k)$ where $k$ is the number of children of the root that need to be added to the main heap.
- change in potential ($\Delta \Phi$): the number of trees decreases and some nodes may be unmarked, so $\Delta \Phi = - (k - 1)$
- amortized cost: $O(\log n + k) - (k - 1) = O(\log n)$

**decrease-key operation:**
- actual cost: $O(1)$
- change in potential ($\Delta \Phi$): a node might be marked or added to the root list, so $\Delta \Phi = 1$
- amortized cost: $O(1) + 1 = O(1)$

**delete operation:**
- amortized cost: $O(\log n)$ similar to extract-min, because deletion involves extracting the minimum and then decreasing the key.

### hashing

#### proof of expected time complexity

**load factor and collision probability:**
\[
\alpha = \frac{n}{m}
\]
where $n$ is the number of elements and $m$ is the number of slots in the hash table.

**open addressing:**
- expected number of probes: $\frac{1}{1 - \alpha}$ for load factor $\alpha < 1$. as $\alpha$ approaches 1, the probability of collision increases.

**separate chaining:**
- expected chain length: $\alpha$. on average, each linked list has a length of $\alpha$.
- expected time complexity: $O(1 + \alpha)$. the operations insert, search, and delete are expected to take $O(1 + \alpha)$ time on average.

**universal hashing:**
- expected number of collisions: $O(1)$. universal hashing chooses hash functions randomly to minimize collisions.

**conclusion:**
- expected time complexity: $O(1)$ with a constant load factor $\alpha$.

### minimum spanning tree (mst)

#### proof of correctness and time complexity

**kruskal’s algorithm:**

**correctness:**
1. sort all edges in non-decreasing order of their weights.
2. pick the smallest edge. check if it forms a cycle with the spanning tree formed so far using union-find.
3. if no cycle is formed, include this edge. repeat until there are $v-1$ edges in the spanning tree.

**time complexity:**
- sorting edges: $O(e \log e)$.
- union-find operations: $O(e \log v)$ using path compression and union by rank.

**prim’s algorithm:**

**correctness:**
1. start with a vertex and greedily grow the spanning tree by adding the minimum weight edge from the tree to a vertex not in the tree.
2. repeat until all vertices are included.

**time complexity:**
- using a priority queue (binary heap): $O(e \log v)$

### segment tree

#### proof of construction and query operation complexities

**construction:**
- time complexity: $O(n \log n)$. for each level of the tree, which has $O(\log n)$ levels, we have $O(n)$ nodes.

**query operation:**
- time complexity: $O(\log n)$. each query starts at the root and traverses down to the leaves.

**update operation:**
- time complexity: $O(\log n)$. each update requires traversal from the root to a leaf.

**proof:**
1. the height of the segment tree is $\log n$.
2. each query/update operation traverses $\log n$ nodes.

**conclusion:**
- segment tree construction: $O(n \log n)$
- range query: $O(\log n)$
- point update: $O(\log n)$

### skip list

#### proof of average-case time complexities

**expected height:**
- each level in a skip list contains approximately half of the nodes from the level below it.
- the expected height $h$ of the skip list is $O(\log n)$.

**search operation:**
- expected time complexity: $O(\log n)$.
- each level reduces the search space by half, similar to binary search.

**insert/delete operation:**
- expected time complexity: $O(\log n)$.
- each insertion/deletion operation involves searching for the position and updating pointers at multiple levels.

**conclusion:**
- expected height: $O(\log n)$.
- search: $O(\log n)$.
- insert/delete: $O(\log n)$.

### traveling salesman problem (tsp)

#### proof of approximation algorithms

**nearest neighbor heuristic:**
- start from a vertex, repeatedly visit the nearest unvisited vertex until all vertices are visited.
- approximation ratio: can be $O(\log n)$ in the worst case.

**christofides' algorithm:**
- construct mst, find a minimum-weight perfect matching in the induced subgraph of odd-degree vertices, and form an eulerian circuit.
- approximation ratio: $\frac{3}{2}$.

### trie

#### proof of space and time complexities

**insert operation:**
- time complexity: $O(L)$ where $L$ is the length of the string.

**search operation:**
- time complexity: $O(L)$ where $L$ is the length of the string.

**space complexity:**
- space complexity: $O(N \cdot L)$ where $N$ is the number of strings and $L$ is the average length.

**proof:**
1. each character of a string is stored in a node.
2. each node has at most 26 children (for lowercase english letters).

### conclusion

these detailed proofs provide mathematical backing for the time and space complexities of various data structures and algorithms, ensuring a clear understanding of their performance characteristics.
