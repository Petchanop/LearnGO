/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex5.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/14 13:31:23 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/14 13:34:17 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sl = append(sl[:3], sl[6:]...)
	fmt.Println(sl)
}