/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex2.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 16:11:52 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 16:48:40 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja6 Ex2")
	num_array := []int{1, 2, 3, 4, 5}
	sum_foo := foo(num_array...)
	sum_bar := bar(num_array)
	fmt.Println(sum_foo)
	fmt.Println(sum_bar)
}

func foo(n ...int) int {
	var sum int
	for _, num := range n {
		sum += num
	}
	return sum
}

func bar(n []int) int {
	var sum int
	for _, num := range n {
		sum += num
	}
	return sum
}
