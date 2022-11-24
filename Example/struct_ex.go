/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   struct_ex.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 15:27:49 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 20:00:46 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type Mobile struct {
	weight int
	height int
	body   string
}

type artist struct {
	name      string
	equipment string
	apperence Mobile
}

func main() {
	mo := Mobile{weight: 60, height: 165, body: "chubby"}
	fmt.Println(mo)

	//Embedded struct
	Nii := artist{
		name:      "Niinoi",
		equipment: "Piano",
		apperence: Mobile{
			weight: 55,
			height: 167,
			body:   `fat`,
		},
	}
	fmt.Println(Nii)
}
