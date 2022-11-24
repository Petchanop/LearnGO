/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex2.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 19:08:55 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 19:30:10 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import (
	"fmt"
)

type Person struct {
	first_name	string
	last_name	string
	favourite	[]string
}

func	main(){
	p1 := Person{
		first_name: "Prinda",
		last_name: "Ruengdech",
		favourite: []string{"Milk", "Chocolate"},
	}

	p2 := Person{
		first_name: "Patcharamon",
		last_name: "Namwaykor",
		favourite: []string{"Kanom", "Chesse"},
	}

	People := map[string]Person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	// People["Ruengdech"] = p1
	// People["Namwaykor"] = p2



	fmt.Println(People["Ruengdech"])
	fmt.Println(People["Namwaykor"])
	for _, val := range People {
		fmt.Println(val)
	}


}