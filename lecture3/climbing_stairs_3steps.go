package lecture3

/*
Problem:
	Climbing Stairs (3 steps)
	You are climbing a staircase. It takes n steps to reach to the top.
	Each time you can either climb 1, 2 or 3 steps.
	In how many distinct ways can you climb to the top?
Framework for Solving DP Problems:
	1. Define the objective function
		f(i) is the number of distinct ways to reach the i-th stair.
	2. Identify base cases
		f(0) = 1
		f(1) = 1
		f(2) = 2
	3. Write down a recurrence relation for the optimized objective function
		f(n) = f(n-1) + f(n-2) + f(n-3)
	4. What's the order of execution?
		bottom-up
	5. Where to look for the answer?
		f(n)
*/

// Time complexity: O(n)
// Space complexity: O(n)

func climbStairs3Steps(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
