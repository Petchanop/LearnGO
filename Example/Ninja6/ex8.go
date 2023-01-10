/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex8.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/06 16:45:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/06 16:59:31 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja 6 ex 8")
	plus := func(n int, i int) int {
		return n + i
	}
	test := plus(2, 3)
	fmt.Println(test)
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
