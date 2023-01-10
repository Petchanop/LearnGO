/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex10.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/06 17:19:01 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/06 18:03:34 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja 6 ex 10")
	x := 1
	y := decrement(x)
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(x)
}

func decrement(n int) func() int {
	x := 1
	return func() int {
		n -= x
		return n
	}
}
