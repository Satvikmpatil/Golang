# Go Programming - 50 Fundamental Questions


> Easy but Important questions that build strong programming fundamentals.


---


## Section 1: Basics & Operators (Q1-Q5)


### Q1: Swap Two Numbers Using Temporary Variable
**Why Important:** Foundation of variable manipulation and memory understanding.


**Objective:** Read two integers and swap them using a third (temporary) variable.


| Input | Output |
|-------|--------|
| a = 5, b = 10 | a = 10, b = 5 |


---


### Q2: Swap Two Numbers WITHOUT Temporary Variable
**Why Important:** Shows arithmetic trick, asked frequently in interviews.


**Objective:** Swap two numbers using only arithmetic operations (+ and -).


| Input | Output |
|-------|--------|
| a = 5, b = 10 | a = 10, b = 5 |


---


### Q3: Check Even or Odd Using Modulus
**Why Important:** Most basic conditional logic, foundation for divisibility problems.


**Objective:** Check if a number is even or odd using `%` operator.


| Input | Output |
|-------|--------|
| 7 | Odd |
| 12 | Even |


---


### Q4: Find Maximum of Three Numbers
**Why Important:** Foundation of comparison logic, used everywhere.


**Objective:** Find the largest among three numbers using if-else.


| Input | Output |
|-------|--------|
| 10, 25, 15 | 25 |


---


### Q5: Check Positive, Negative, or Zero using switch
**Why Important:** Basic conditional branching practice (using switch).


**Objective:** Classify a number as positive, negative, or zero.


| Input | Output |
|-------|--------|
| -5 | Negative |
| 0 | Zero |
| 10 | Positive |


---


## Section 2: Loops - The Foundation (Q6-Q10)


### Q6: Print Numbers 1 to N
**Why Important:** Most basic loop pattern, must be automatic.


**Objective:** Print all numbers from 1 to N.


| Input | Output |
|-------|--------|
| N = 5 | 1 2 3 4 5 |


---


### Q7: Sum of First N Natural Numbers
**Why Important:** Accumulator pattern - used in 90% of problems.


**Objective:** Calculate 1 + 2 + 3 + ... + N using a loop.


| Input | Output |
|-------|--------|
| N = 5 | 15 |


---


### Q8: Factorial of a Number
**Why Important:** Classic loop problem, foundation for recursion later.


**Objective:** Calculate N! = N × (N-1) × ... × 1


| Input | Output |
|-------|--------|
| 5 | 120 |
| 0 | 1 |


---


### Q9: Print Multiplication Table
**Why Important:** Practical loop application everyone understands.


**Objective:** Print multiplication table of N from 1 to 10.


| Input | Output |
|-------|--------|
| N = 5 | 5×1=5, 5×2=10, ... 5×10=50 |


---


### Q10: Count Digits in a Number
**Why Important:** Teaches digit extraction pattern used in many problems.


**Objective:** Count how many digits are in a number.


| Input | Output |
|-------|--------|
| 12345 | 5 |
| 7 | 1 |


---


## Section 3: Number Problems (Q11-Q15)


### Q11: Reverse a Number
**Why Important:** Core pattern for digit manipulation, asked in every interview.


**Objective:** Reverse the digits of a number.


| Input | Output |
|-------|--------|
| 1234 | 4321 |
| 100 | 1 |


---


### Q12: Check Palindrome Number
**Why Important:** Combines reversal + comparison, very common question.


**Objective:** Check if a number reads same forwards and backwards.


| Input | Output |
|-------|--------|
| 121 | Palindrome |
| 123 | Not Palindrome |


---


### Q13: Check Prime Number
**Why Important:** Fundamental algorithm, teaches optimization thinking.


**Objective:** Check if a number is prime (divisible only by 1 and itself).


| Input | Output |
|-------|--------|
| 7 | Prime |
| 12 | Not Prime |


---


### Q14: Print Fibonacci Series (First N Terms)
**Why Important:** Classic sequence problem, foundation for DP thinking.


**Objective:** Print: 0, 1, 1, 2, 3, 5, 8, 13...


| Input | Output |
|-------|--------|
| N = 7 | 0 1 1 2 3 5 8 |


---


### Q15: Find GCD of Two Numbers (Euclidean Algorithm)
**Why Important:** Classic algorithm, used in many math problems.


**Objective:** Find Greatest Common Divisor of two numbers.


| Input | Output |
|-------|--------|
| 48, 18 | 6 |


---


## Section 4: Pattern Printing (Q16-Q17)


### Q16: Right Triangle Star Pattern
**Why Important:** Teaches nested loops, most asked pattern type.


**Objective:** Print right triangle of stars.


| Input | Output |
|-------|--------|
| N = 4 | `*` <br> `**` <br> `***` <br> `****` |


---


### Q17: Number Pyramid Pattern
**Why Important:** Combines counting with nested loops.


**Objective:** Print number pattern.


| Input | Output |
|-------|--------|
| N = 4 | `1` <br> `12` <br> `123` <br> `1234` |


---


## Section 5: Slices - Most Important Topic (Q18-Q24)


### Q18: Find Largest Element in Slice
**Why Important:** Most basic slice traversal, building block for everything.


**Objective:** Find the maximum element in a slice.


| Input | Output |
|-------|--------|
| [3, 7, 2, 9, 1] | 9 |


---


### Q19: Find Second Largest Element
**Why Important:** Slightly tricky, tests edge case thinking.


**Objective:** Find the second largest element (no sorting allowed).


| Input | Output |
|-------|--------|
| [3, 7, 2, 9, 1] | 7 |


---


### Q20: Reverse a Slice
**Why Important:** Two-pointer technique foundation, asked everywhere.


**Objective:** Reverse slice in-place using two pointers.


| Input | Output |
|-------|--------|
| [1, 2, 3, 4, 5] | [5, 4, 3, 2, 1] |


---


### Q21: Linear Search
**Why Important:** Most basic searching algorithm.


**Objective:** Find if an element exists in slice, return its index.


| Input | Output |
|-------|--------|
| arr=[10,20,30], key=20 | Found at index 1 |
| arr=[10,20,30], key=25 | Not Found |


---


### Q22: Check if Slice is Sorted
**Why Important:** Simple but important check, used as helper in many problems.


**Objective:** Check if slice is sorted in ascending order.


| Input | Output |
|-------|--------|
| [1, 2, 3, 4] | Sorted |
| [1, 3, 2, 4] | Not Sorted |


---


### Q23: Find Duplicate Elements in Slice
**Why Important:** Classic problem, teaches frequency counting with maps.


**Objective:** Find all elements that appear more than once.


| Input | Output |
|-------|--------|
| [1, 2, 3, 2, 4, 3] | [2, 3] |


---


### Q24: Move All Zeros to End
**Why Important:** In-place slice manipulation, common interview question.


**Objective:** Move all 0s to end while maintaining order of other elements.


| Input | Output |
|-------|--------|
| [0, 1, 0, 3, 12] | [1, 3, 12, 0, 0] |


---


## Section 6: Strings - Second Most Important (Q25-Q30)


### Q25: Reverse a String
**Why Important:** Most basic string manipulation.


**Objective:** Reverse a string without using built-in reverse.


| Input | Output |
|-------|--------|
| "hello" | "olleh" |


---


### Q26: Check String Palindrome
**Why Important:** Very commonly asked, clean logic required.


**Objective:** Check if string reads same forwards and backwards.


| Input | Output |
|-------|--------|
| "madam" | Palindrome |
| "hello" | Not Palindrome |


---


### Q27: Count Vowels and Consonants
**Why Important:** Character classification practice.


**Objective:** Count vowels (a,e,i,o,u) and consonants in a string.


| Input | Output |
|-------|--------|
| "Hello World" | Vowels: 3, Consonants: 7 |


---


### Q28: Check if Two Strings are Anagrams
**Why Important:** Teaches frequency counting with maps, very common question.


**Objective:** Check if two strings have same characters with same frequency.


| Input | Output |
|-------|--------|
| "listen", "silent" | Anagrams |
| "hello", "world" | Not Anagrams |


---


### Q29: Find First Non-Repeating Character
**Why Important:** Classic map problem, asked in Amazon/Google.


**Objective:** Find first character that appears only once.


| Input | Output |
|-------|--------|
| "swiss" | 'w' |
| "aabbcc" | None |


---


### Q30: Remove All Spaces from String
**Why Important:** String cleaning, practical skill.


**Objective:** Remove all whitespace characters.


| Input | Output |
|-------|--------|
| "Hello World" | "HelloWorld" |


---


## Section 7: Sorting Algorithms (Q31-Q32)


### Q31: Bubble Sort
**Why Important:** Simplest sorting algorithm, must know for interviews.


**Objective:** Sort slice by repeatedly swapping adjacent elements.


| Input | Output |
|-------|--------|
| [5, 3, 8, 4, 2] | [2, 3, 4, 5, 8] |


---


### Q32: Selection Sort
**Why Important:** Easy to understand, teaches "find minimum" pattern.


**Objective:** Sort by repeatedly finding minimum and placing at beginning.


| Input | Output |
|-------|--------|
| [64, 25, 12, 22] | [12, 22, 25, 64] |


---


## Section 8: Recursion - Think Different (Q33-Q36)


### Q33: Factorial Using Recursion
**Why Important:** First recursion problem everyone should solve.


**Objective:** Calculate N! using recursion.


| Input | Output |
|-------|--------|
| 5 | 120 |


---


### Q34: Sum of N Numbers Using Recursion
**Why Important:** Reinforces recursion pattern.


**Objective:** Calculate 1 + 2 + ... + N recursively.


| Input | Output |
|-------|--------|
| 5 | 15 |


---


### Q35: Fibonacci Using Recursion
**Why Important:** Classic recursion, shows overlapping subproblems.


**Objective:** Find Nth Fibonacci number recursively.


| Input | Output |
|-------|--------|
| 6 | 8 (0,1,1,2,3,5,8) |


---


### Q36: Print 1 to N Using Recursion (No Loop)
**Why Important:** Shows recursion can replace loops.


**Objective:** Print 1, 2, 3, ... N without using any loop.


| Input | Output |
|-------|--------|
| 5 | 1 2 3 4 5 |


---


## Section 9: Bit Manipulation (Q37-Q38)


### Q37: Check Even/Odd Using Bitwise AND
**Why Important:** Shows how bits work, faster than modulus.


**Objective:** Use `&` operator to check even/odd.


| Input | Output |
|-------|--------|
| 5 | Odd |
| 8 | Even |


---


### Q38: Count Number of Set Bits (1s)
**Why Important:** Fundamental bit manipulation, asked frequently.


**Objective:** Count how many 1s are in binary representation.


| Input | Output |
|-------|--------|
| 13 (binary: 1101) | 3 |
| 7 (binary: 111) | 3 |


---


## Section 10: 2D Slices (Q39-Q40)


### Q39: Print a 2D Matrix
**Why Important:** Foundation for all matrix problems.


**Objective:** Create and print a 2D slice (matrix).


| Input | Output |
|-------|--------|
| 2x3 matrix: `1 2 3` and `4 5 6` | `1 2 3` <br> `4 5 6` |


---


### Q40: Sum of Matrix Elements
**Why Important:** Basic matrix traversal.


**Objective:** Find sum of all elements in a matrix.


| Input | Output |
|-------|--------|
| [[1,2],[3,4]] | 10 |


---


## Section 11: Structs & Methods (Q41-Q42)


### Q41: Create a Struct with Constructor Function
**Why Important:** Foundation of Go's struct-based design.


**Objective:** Create a `Student` struct with name and age, use a constructor function to initialize, and display details.


| Input | Output |
|-------|--------|
| name="Rahul", age=20 | "Rahul - 20" |


---


### Q42: Demonstrate Struct Embedding (Composition)
**Why Important:** Go uses composition over inheritance.


**Objective:** Create `Animal` struct with `Eat()` method. Create `Dog` struct that embeds Animal and adds `Bark()` method.


| Method Call | Output |
|-------------|--------|
| dog.Eat() | "Eating..." |
| dog.Bark() | "Barking..." |


---


## Section 12: Linked List (Q43-Q44)


### Q43: Create and Traverse a Linked List
**Why Important:** Foundation of all linked list problems.


**Objective:** Create a linked list with nodes 1→2→3 and print all elements.


| Input | Output |
|-------|--------|
| Create nodes: 1, 2, 3 | 1 → 2 → 3 |


---


### Q44: Reverse a Linked List
**Why Important:** MOST ASKED linked list question in interviews.


**Objective:** Reverse a linked list iteratively.


| Input | Output |
|-------|--------|
| 1→2→3→4 | 4→3→2→1 |


---


## Section 13: Stack (Q45)


### Q45: Check Balanced Parentheses
**Why Important:** Classic stack problem, asked everywhere.


**Objective:** Check if parentheses are balanced: `(){}[]`


| Input | Output |
|-------|--------|
| "{[()]}" | Balanced |
| "{[(])}" | Not Balanced |
| "(((" | Not Balanced |


---


## Section 14: Queue (Q46)


### Q46: Implement Queue Operations
**Why Important:** Understand FIFO concept clearly.


**Objective:** Implement enqueue (add), dequeue (remove), and peek operations using slices.


| Operation | Queue State | Returns |
|-----------|-------------|---------|
| add(10) | [10] | - |
| add(20) | [10, 20] | - |
| add(30) | [10, 20, 30] | - |
| remove() | [20, 30] | 10 |
| peek() | [20, 30] | 20 |


---


## Section 15: Binary Tree (Q47)


### Q47: Inorder Traversal of Binary Tree
**Why Important:** Most fundamental tree operation, asked in every interview.


**Objective:** Print tree nodes in Inorder (Left → Root → Right).


```
   1
  / \
 2   3
```


| Traversal | Output |
|-----------|--------|
| Inorder (L-Root-R) | 2, 1, 3 |
| Preorder (Root-L-R) | 1, 2, 3 |
| Postorder (L-R-Root) | 2, 3, 1 |


---


## Section 16: Binary Search Tree (Q48)


### Q48: Search in BST
**Why Important:** Shows BST property usage, O(log n) search.


**Objective:** Search for a value in BST.


```
   5
  / \
 3   7
/ \
2   4
```


| Search | Output |
|--------|--------|
| 4 | Found |
| 6 | Not Found |


---


## Section 17: Maps (Q49)


### Q49: Two Sum Problem Using Map
**Why Important:** MOST ASKED interview question (Leetcode #1).


**Objective:** Find two numbers in slice that add up to target. Return their indices.


| Input | Output |
|-------|--------|
| arr=[2,7,11,15], target=9 | [0, 1] (indices of 2 and 7) |
| arr=[3,2,4], target=6 | [1, 2] (indices of 2 and 4) |


---


## Section 18: Backtracking (Q50)


### Q50: Print All Subsets of a Slice
**Why Important:** Foundation of backtracking, pattern used in many problems.


**Objective:** Print all possible subsets (power set).


| Input | Output |
|-------|--------|
| [1, 2] | [], [1], [2], [1,2] |
| [1, 2, 3] | [], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3] |


---


## Quick Reference: Topic → Question Mapping


| Topic | Questions |
|-------|-----------|
| Basics & Operators | Q1-Q5 |
| Loops | Q6-Q10 |
| Number Problems | Q11-Q15 |
| Patterns | Q16-Q17 |
| Slices (1D) | Q18-Q24 |
| Strings | Q25-Q30 |
| Sorting | Q31-Q32 |
| Recursion | Q33-Q36 |
| Bit Manipulation | Q37-Q38 |
| 2D Slices | Q39-Q40 |
| Structs & Methods | Q41-Q42 |
| Linked List | Q43-Q44 |
| Stack | Q45 |
| Queue | Q46 |
| Binary Tree | Q47 |
| BST | Q48 |
| Maps | Q49 |
| Backtracking | Q50 |


---


## Study Order Recommendation


```
Week 1: Q1-Q17  (Basics, Loops, Numbers, Patterns)
Week 2: Q18-Q32 (Slices, Strings, Sorting)
Week 3: Q33-Q42 (Recursion, Bits, 2D Slices, Structs)
Week 4: Q43-Q50 (Data Structures - LinkedList to Backtracking)
```


---


**All 50 questions build STRONG FUNDAMENTALS and are frequently asked in interviews!**
