/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fmt_example.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/07 17:41:24 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/07 18:09:28 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

var y int = 42

type NiiNoi int

var test NiiNoi

func main() {
	test = 50
	fmt.Println(test)
	fmt.Printf("%T\n", test)
	fmt.Printf("%b\n", test)
	fmt.Printf("%x\n", test)
	fmt.Printf("%#x\n", test)
	fmt.Printf("%b\t%x\t%#x\n", test, test, test)
}
