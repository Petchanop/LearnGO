/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex6.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 13:35:28 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 13:52:28 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	state_name := make([]string, 50, 50)
	state_name = []string{"Alabama"}
	state_name = append(state_name, "Alaska")
	state_name = append(state_name, "Arizona")
	state_name = append(state_name, "Arkansas")
	state_name = append(state_name, "California")
	state_name = append(state_name, "Colorado")
	state_name = append(state_name, "Connecticut")
	state_name = append(state_name, "Delaware")
	state_name = append(state_name, "Florida")
	state_name = append(state_name, "Georgia")
	state_name = append(state_name, "Hawaii")
	state_name = append(state_name, "Idaho")
	state_name = append(state_name, "Illinois")
	state_name = append(state_name, "Indiana")
	state_name = append(state_name, "Iowa")
	state_name = append(state_name, "Kansas")
	state_name = append(state_name, "Kentucky")
	for i := 0; i < len(state_name); i++ {
		fmt.Println(i, state_name[i])
	}
	fmt.Println(len(state_name), " len\t", cap(state_name), " cap")
}
