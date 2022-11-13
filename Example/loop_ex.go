/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   loop_ex.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/12 14:52:45 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/12 16:24:10 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	//for init; condition ;post {}
	// for statement with for clause
	fmt.Println("for statement with for clause")
	fmt.Println("for init; condition ; post {}")
	for i := 0; i <= 10; i++ {
		fmt.Printf("i is increment by 1 equal : %d\n", i)
	}

	// nested loop
	fmt.Println("nested for loop")
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Printf("nested loop i outer loop is %d X by j inner loop is %d equal %d\n", i, j, i*j)
		}
	}
	// for statement with single condition
	fmt.Println("for statement with single condition")
	x := 3
	for x < 15 {
		fmt.Printf("x is %d less than 15\n", x)
		fmt.Println("x multiply by 2")
		x *= 2
	}
	// for while
	fmt.Println("for while with break keyword")
	x = 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	// for with break and continue keywords
	fmt.Println("for with break and continue keywords")
	x = 0
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	//for loop printing ascii
	fmt.Println("for loop printing ascii")
	for x := 33; x <= 122; x++ {
		fmt.Printf("%v\t%#U\n", x, x)
	}
}
