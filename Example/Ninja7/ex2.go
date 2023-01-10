/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex2.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/10 13:58:18 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/10 14:23:04 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Narit"
	(*p).last = "Wattana"
	p.age = 28
}

func main() {
	fmt.Println("Ninja 7 ex2")
	p1 := person{
		"Saifa", "Meethong", 18,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
