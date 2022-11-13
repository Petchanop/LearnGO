/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex9.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/13 14:26:17 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/13 14:29:16 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	sport := "football"
	switch sport {
	case "basketball":
		fmt.Println("Coby Brian")
	case "Boxing":
		fmt.Println("Roy Jones")
	case "football":
		fmt.Println("Messi")
	default:
		fmt.Println(sport)
	}
}
