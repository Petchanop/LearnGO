/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tutorial_interfaces.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/31 16:10:56 by npiya-is          #+#    #+#             */
/*   Updated: 2022/12/31 16:39:51 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am secretAgent type", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am person type ", p.first, p.last)
}

type human interface {
	speak()
} // if empty interface every types can put in.

func greetings(h human) {
	fmt.Printf("Hello ")
	h.speak()
	switch h.(type) {
	case person:
		fmt.Println("Nice to meet you.")
	case secretAgent:
		fmt.Println("This is secret missions.")
	}
}

func main() {
	opal := person{
		first: "Thanaporn",
		last:  "Sirirakwongsa",
	}
	James := secretAgent{
		person{
			first: "James",
			last:  "Bond",
		},
		true,
	}
	greetings(opal)
	greetings(James)
}
