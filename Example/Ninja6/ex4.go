/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex4.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/01/05 17:56:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/01/05 18:04:47 by npiya-is         ###   ########.fr       */
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

func (p person) speak() {
	fmt.Println("Hello my name is", p.first, ".I 'm", p.age, "years old.")
}

func main() {
	fmt.Println("Ninja6 ex4")
	p1 := person{
		first: "Nam",
		last:  "eknipitphong",
		age:   36,
	}
	p1.speak()
}
