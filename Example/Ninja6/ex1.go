/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex1.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 15:43:23 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 16:05:39 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja6 Ex1")
	foo_re := foo()
	bar_num, bar_str := bar()
	fmt.Println(foo_re)
	fmt.Println(bar_num)
	fmt.Println(bar_str)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	num := 43
	str := "String"
	return num, str
}
