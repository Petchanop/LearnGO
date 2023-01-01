/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   callback.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/01 17:22:55 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/01 22:24:03 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

/*- passing func as a argument
  - functional programming concept*/
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(ii...)
	fmt.Println(s)
	s2 := even(sum, ii...)
	fmt.Println("sum of even number from 1 to 10", s2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
