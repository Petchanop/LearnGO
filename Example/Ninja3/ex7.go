/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex7.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/13 14:19:48 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/13 14:22:06 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	test := "test"
	if test == "no" {
		fmt.Println("no")
	} else if test == "test" {
		fmt.Println("test")
	} else {
		fmt.Println("else")
	}
}
