/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex3.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 12:26:57 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 12:41:37 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	s2l := make([][]int, 4)
	num := 42
	for i := 0; i < cap(s2l); i++ {
		s2l[i] = make([]int, 5)
		if i == 2 {
			num = 44
		} else if i == 3 {
			num = 43
		}
		for j := 0; j < cap(s2l[i]); j++ {
			s2l[i][j] = num
			num++
		}
	}
	fmt.Println(s2l)

	//or using slice in slice
}
