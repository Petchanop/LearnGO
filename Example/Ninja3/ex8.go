/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex8.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/13 14:22:29 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/13 14:25:54 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	x := "switch"
	switch x {
	case "not":
		fmt.Println("not")
	case "switch":
		fmt.Println(x)
	default:
		fmt.Println("default")
	}
}
