/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Array_Slice.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/13 14:53:14 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/13 19:28:57 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	var ar [5]int // zero base index
	fmt.Println(ar)
	// x := type{values} // composite literal
	x := []int{4, 10, 5, 23, 115, 6}
	fmt.Println(x)
	//Slice allows you to group VALUES of the same TYPE
	for i, v := range x {
		fmt.Println("value : ", v)
		fmt.Println("index : ", i)
	}
	//slice a slice
	fmt.Println(x[1:])
	fmt.Println(x[:4])
	fmt.Println(x[2:5])
	fmt.Println(x[2:4])

	//append slices
	y := []int{8, 11, 9}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
	x = append(x, 90, 94, 23)
	fmt.Println(x)
	z := []int{1, 32, 59, 20, 43}
	x = append(x[:4], z[3:]...)
	fmt.Println(x)
	//delete data in slices

	//keyword make
	a := make([]int, 10, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	//if append when full capacity slice grow to 2 times capacity
	//multidimension slice
	jb := []string{"Mobile", "Palmmy", "Mhew", "Sound"}
	sb := []string{"Muay", "fah", "Perth", "Tatar"}

	mm := [][]string{jb, sb}
	fmt.Println(mm)
}
