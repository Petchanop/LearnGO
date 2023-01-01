/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   closure.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/01 22:43:29 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/01 23:08:25 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*one scope enhance other scope*/
func main() {
	var x int
	x++
	fmt.Println(x)
	{
		y := 1
		fmt.Println(y)
		x += y
	}
	fmt.Println(x)
	y := incrementor()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println("Hello world")
}

func incrementor() func() int {
	var x int
	return func() int {
		x += 1
		return x
	}
}
