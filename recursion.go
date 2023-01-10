/* ************************************************************************** */
/*                                                                   */
/*                                                        :::      ::::::::   */
/*   recursion.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 15:30:15 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 15:30:15 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Recursion")
	num := 11
	fmt.Println(Factorial(num))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
