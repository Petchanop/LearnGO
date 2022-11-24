/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex1.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 19:00:54 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 19:08:39 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import (
	"fmt"
)

type Person struct {
	first_name	string
	last_name	string
	favourite	string
}

func main(){
	p1 := Person{
		first_name: "Petcha",
		last_name: "npiya-is",
		favourite: "Somtum",
	}

	p2 := Person{
		first_name: "NiiNoi",
		last_name: "Rchiewli",
		favourite: "Mhala",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	P := []Person{p1, p2}
	for i, val := range P{
		fmt.Println("index ", i, val)
	}
}