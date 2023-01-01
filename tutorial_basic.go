/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tutorial_basic.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 22:29:09 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 23:30:04 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

const (
	amii  = "Soontaree"
	petch = "Panupong"
	Nii   = iota
	Noii  = iota
)

type Person struct {
	first_name, last_name string
	age                   int
}

type Police struct {
	Person
	weapon string
}

func main() {
	var i int = 42
	const k int = 50
	j := 42
	fmt.Println(i, k, j)

	fmt.Println(amii, petch)
	fmt.Println(Nii, Noii)

	//for (initialize :  condition : iterate)
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	j = 9
	for j < 20 {
		fmt.Println(j)
		j++
	}

	p1 := Person{
		first_name: "panupong",
		last_name:  "mikada",
		age:        23,
	}

	police := Police{
		Person: Person{
			first_name: "Panupong",
			last_name:  "mikada",
			age:        23,
		},
		weapon: "Magnum",
	}
	fmt.Println(p1)
	fmt.Println(police)
}
