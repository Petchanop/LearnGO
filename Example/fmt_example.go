/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fmt_example.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/07 17:41:24 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/10 17:53:20 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

var y int = 42

const (
	_ = iota
	d = 1 << (iota * 10)
	e = 1 << (iota * 10)
	f = 1 << (iota * 10)
	a = 43
	b = 43.123
	c = "test const"
)

type NiiNoi int

var test NiiNoi

func main() {
	test = 50
	fmt.Println(test)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Printf("%T %T %T\n", d, e, f)
	fmt.Printf("%T\n", test)
	fmt.Printf("%b\n", test)
	fmt.Printf("%x\n", test)
	fmt.Printf("%#x\n", test)
	fmt.Printf("%b\t%x\t%#x\n", test, test, test)
}
