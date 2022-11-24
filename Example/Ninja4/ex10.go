/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex10.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 15:22:44 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 15:25:31 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
	//delete data in map
	s_map := map[string][]string{
		"Bond_James":      {"Shaken, not stirred", "Martinis", "Women"},
		"Moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being Evil", "Ice cream", "Sunsets"},
	}

	for i := range s_map {
		delete(s_map, i)
	}
	fmt.Println(len(s_map))
}
