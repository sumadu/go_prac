package main

import fmt

function main() var stack []string
{
	stack = append(stack,"cool")

	stack = append(stack, "budy")
	for len(stack) > 0 {
		n := len(stack) -1
		fmt.print(stack[n])

		stack = stack[:n]
}

}