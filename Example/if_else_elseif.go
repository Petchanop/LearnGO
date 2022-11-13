/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   if_else_elseif.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/12 16:35:06 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/12 19:22:13 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("if true print")
	}

	if false {
		fmt.Println("if false print")
	}

	if !true {
		fmt.Println("! with true")
	}

	if !false {
		fmt.Println("! with false")
	}
	//initialize statement
	fmt.Println("initialize statement")
	x := 42
	if x == 42 {
		fmt.Println("x is equal 42")
	} else if x == 43 {
		fmt.Println("x is equal 43")
	} else if x == 44 {
		fmt.Println("x is equal 44")
	} else {
		fmt.Println("x is not equal 42, 43")
	}

}
