/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   embedded_struct.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 17:53:35 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 18:02:41 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type person struct {
	first	string
	last	string
	age		int
}

type secretAgent struct {
	person
	ltk	bool
}

func main(){
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	//doesn't have to call sa1.person you can directly call its properties innner type
	fmt.Println(sa1.first, sa1.last, sa1.age)

}