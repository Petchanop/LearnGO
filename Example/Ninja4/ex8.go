/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex8.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 15:09:53 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 15:21:15 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	s_map := map[string][]string{
		"Bond_James":      {"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being Evil", "Ice cream", "Sunsets"},
	}
	for i := range s_map {
		fmt.Println(s_map[i])
	}
	for i, val := range s_map {
		fmt.Println(i)
		for j, data := range val {
			fmt.Println("index ", j, " ", data)
		}
	}
}
