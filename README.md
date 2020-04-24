# 30-Day-LeetCoding-Challenge
30-Day LeetCoding Challenge

## Week 1: April 1st–April 7th
###  [Single Number](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-singleNumber.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-singleNumber_test.go)
```
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
Example 1:
Input: [2,2,1]
Output: 1
Example 2:
Input: [4,1,2,1,2]
Output: 4
```

###  [Happy Number](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-happyNumber.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-happyNumber_test.go)
```
Write an algorithm to determine if a number n is "happy".
A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
Return True if n is a happy number, and False if not.
Example:
Input: 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```

###  [Maximum Subarray](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-maximumSubarray.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-maximumSubarray_test.go)
```
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
Example:
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

```

###  [Move Zeroes](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-moveZeroes.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-moveZeroes_test.go)
```
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
   Hide Hint #1
In-place means we should not be allocating any space for extra array. But we are allowed to modify the existing array. However, as a first step, try coming up with a solution that makes use of additional space. For this problem as well, first apply the idea discussed using an additional array and the in-place solution will pop up eventually.
   Hide Hint #2
A two-pointer approach could be helpful here. The idea would be to have one pointer for iterating the array and another pointer that just works on the non-zero elements of the array.
```

###  [Best Time to Buy and Sell Stock II](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-bestTimetoBuyandSellStockII.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-bestTimetoBuyandSellStockII_test.go)
```
Say you have an array prices for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
Example 1:
Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:
Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
Constraints:
1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4
```

###  [Group Anagrams](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-groupAnagrams.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-groupAnagrams_test.go)
```
Given an array of strings, group anagrams together.
Example:
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:
All inputs will be in lowercase.
The order of your output does not matter.
```

###  [Counting Elements](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-countingElements.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week1-countingElements_test.go)
```
Given an integer array arr, count element x such that x + 1 is also in arr.
If there're duplicates in arr, count them seperately.
Example 1:
Input: arr = [1,2,3]
Output: 2
Explanation: 1 and 2 are counted cause 2 and 3 are in arr.
Example 2:
Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Explanation: No numbers are counted, cause there's no 2, 4, 6, or 8 in arr.
Example 3:
Input: arr = [1,3,2,3,5,0]
Output: 3
Explanation: 0, 1 and 2 are counted cause 1, 2 and 3 are in arr.
Example 4:
Input: arr = [1,1,2,2]
Output: 2
Explanation: Two 1s are counted cause 2 is in arr.
Constraints:
1 <= arr.length <= 1000
0 <= arr[i] <= 1000
```


## Week 2: Week 2: April 8th–April 14th
###  [Middle of the Linked List](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-middleoftheLinkedList.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-middleoftheLinkedList_test.go)
```
Given a non-empty, singly linked list with head node head, return a middle node of linked list.
If there are two middle nodes, return the second middle node.
Example 1:
Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
Example 2:
Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.
Note:
The number of nodes in the given list will be between 1 and 100.
```

###  [Backspace String Compare](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-backspaceStringCompare.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-backspaceStringCompare_test.go)
```
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
Note that after backspacing an empty text, the text will continue empty.
Example 1:
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:
1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:
Can you solve it in O(N) time and O(1) space?
```

###  [Min Stack](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-minStack.go)
```
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
Output
[null,null,null,null,-3,null,0,-2]
Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
Constraints:
Methods pop, top and getMin operations will always be called on non-empty stacks.
   Hide Hint #1  
Consider each node in the stack having a minimum value. (Credits to @aakarshmadhavan)
```

###  [Diameter of Binary Tree](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-diameterofBinaryTree.go)
```
Diameter of Binary Tree
Solution
Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
Note: The length of path between two nodes is represented by the number of edges between them.
```

###  [Last Stone Weight](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-lastStoneWeight.go)
```
We have a collection of stones, each stone has a positive integer weight.
Each turn, we choose the two heaviest stones and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:
If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)
Example 1:
Input: [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
Note:
1 <= stones.length <= 30
1 <= stones[i] <= 1000
   Hide Hint #1
Simulate the process. We can do it with a heap, or by sorting some list of stones every time we take a turn.
```

###  [Contiguous Array](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-contiguousArray.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week2-contiguousArray_test.go)
```
Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
Example 2:
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
Note: The length of the given binary array will not exceed 50,000.
```

## Week 3: Week 3: April 15th–April 21st
###  [Leftmost Column with at Least a One](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week3-leftmostColumnwithatLeastaOne.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week3-leftmostColumnwithatLeastaOne_test.go)
```
(This problem is an interactive problem.)
A binary matrix means that all elements are 0 or 1. For each individual row of the matrix, this row is sorted in non-decreasing order.
Given a row-sorted binary matrix binaryMatrix, return leftmost column index(0-indexed) with at least a 1 in it. If such index doesn't exist, return -1.
You can't access the Binary Matrix directly.  You may only access the matrix using a BinaryMatrix interface:
BinaryMatrix.get(x, y) returns the element of the matrix at index (x, y) (0-indexed).
BinaryMatrix.dimensions() returns a list of 2 elements [n, m], which means the matrix is n * m.
Submissions making more than 1000 calls to BinaryMatrix.get will be judged Wrong Answer.  Also, any solutions that attempt to circumvent the judge will result in disqualification.
For custom testing purposes you're given the binary matrix mat as input in the following four examples. You will not have access the binary matrix directly.
Example 1:
Input: mat = [[0,0],[1,1]]
Output: 0
Example 2:
Input: mat = [[0,0],[0,1]]
Output: 1
Example 3:
Input: mat = [[0,0],[0,0]]
Output: -1
Example 4:
Input: mat = [[0,0,0,1],[0,0,1,1],[0,1,1,1]]
Output: 1
Constraints:
1 <= mat.length, mat[i].length <= 100
mat[i][j] is either 0 or 1.
mat[i] is sorted in a non-decreasing way.
   Hide Hint #1
1. (Binary Search) For each row do a binary search to find the leftmost one on that row and update the answer.
   Hide Hint #2
2. (Optimal Approach) Imagine there is a pointer p(x, y) starting from top right corner. p can only move left or down. If the value at p is 0, move down. If the value at p is 1, move left. Try to figure out the correctness and time complexity of this algorithm.

```

## Week 4: April 22nd–April 28th
###  [Subarray Sum Equals K](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-subarraySumEqualsK.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-subarraySumEqualsK_test.go)
```
Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
   Hide Hint #1  
Will Brute force work here? Try to optimize it.
   Hide Hint #2  
Can we optimize it by using some extra space?
   Hide Hint #3  
What about storing sum frequencies in a hash table? Will it be useful?
   Hide Hint #4  
sum(i,j)=sum(0,j)-sum(0,i), where sum(i,j) represents the sum of all the elements from index i to j-1. Can we use this property to optimize it.
https://leetcode.com/problems/subarray-sum-equals-k/solution/
```

###  [Subarray Sum Equals K](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-bitwiseANDofNumbersRange.go) [(unit test case)](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-bitwiseANDofNumbersRange_test.go)
```
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
Example 1:
Input: [5,7]
Output: 4
Example 2:
Input: [0,1]
Output: 0
```