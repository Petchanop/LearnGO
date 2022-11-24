/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   anonymous_struct.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 18:23:44 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 18:27:08 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

// type person struct {
// 	first	string
// 	last	string
// 	age		int
// }

// create no name struct

func main(){
	p1 := struct {
		first	string
		last	string
		age		int
	}{
		first:	"James",
		last:	"Bond",
		age:	32,
	}

	fmt.Println(p1)
}