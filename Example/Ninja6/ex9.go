/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex9.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/06 16:54:22 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/06 17:30:35 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja 6 ex 9")
	foo(bar, "test", "plus")
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := []int{1, 3, 5, 7, 9}
	fmt.Println(g(x))
}

func foo(f func(string, string) string, s, st string) {
	fmt.Println(f(s, st))
}

func bar(s string, st string) string {
	return s + st
}
