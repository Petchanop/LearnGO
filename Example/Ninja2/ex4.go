/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex4.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/10 18:32:41 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/10 18:40:25 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Shift bits exercise")
	num := 128
	fmt.Printf("%d  %b %x\n", num, num, num)
	a := num << 1
	b := a << 1
	c := b << 1
	fmt.Printf("%d  %b %x\n", a, a, a)
	fmt.Printf("%d  %b %x\n", b, b, b)
	fmt.Printf("%d %b %x\n", c, c, c)
}
