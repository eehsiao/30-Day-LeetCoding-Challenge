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

## Week 2: Week 2: April 8th–April 14th
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