/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex2.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/10 18:02:00 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/10 18:28:17 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	a := 42
	if a == 42 {
		a = 43
	}
	if a != 42 {
		a = 42
	}
	if a <= 50 {
		fmt.Printf("%t\n", a <= 50)
	}
	if a >= 50 {
		fmt.Printf("%t\n", a >= 50)
	}
	fmt.Printf("finish a equal %d\n", a)
}
