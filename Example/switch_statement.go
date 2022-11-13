/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   switch_statement.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/12 16:57:49 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/12 19:19:13 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	n := 4
	switch n {
	case 2:
		fmt.Println("equal 2")
	case 4:
		fmt.Println("equal 4")
		fallthrough
	default:
		fmt.Println("this is default")
	}
}

//fallthrough use for program to reach last statement
