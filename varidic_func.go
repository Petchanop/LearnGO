/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   varidic_func.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/29 14:13:06 by npiya-is          #+#    #+#             */
/*   Updated: 2022/12/29 14:37:03 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	VaridicSum(2, 3, 4, 5, 6, 7)
	ex := []int{2, 3, 4, 5, 6, 7}
	ret := SumSlices(ex...)
	fmt.Println(ret)
}

func VaridicSum(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Printf("%d", v)
	}
	fmt.Println()
	fmt.Println(sum)
}

func SumSlices(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return (sum)
}
