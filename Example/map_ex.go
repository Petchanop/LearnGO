/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   map_ex.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/13 19:30:11 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 12:03:59 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Mobile": 1,
		"Mhew":   2,
		"Moss":   3,
	}
	fmt.Println(m)
	fmt.Println(m["Mobile"])

	//read key and value from map
	v, ok := m["Mhew"]
	fmt.Println(v)
	fmt.Println(ok)

	//addd elements to map
	m["May"] = 4
	for value, _ := range m {
		fmt.Println(value)
	}
	//delete member in map
	delete(m, "Mobile")
	fmt.Println(m)
}
