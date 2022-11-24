/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex7.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 13:56:13 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 15:22:08 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	s1 := [][]string{{"James", "Bonds", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hello, not James"}}
	for _, val := range s1 {
		fmt.Println(val)
		for _, rec := range val {
			fmt.Println(rec)
		}
	}
	smap := map[string]string{
		"Bonds":"Shaken, not stirred",
		"Moneypenny"
	}
}
