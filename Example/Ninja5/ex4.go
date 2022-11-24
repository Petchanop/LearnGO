/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex4.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 20:07:35 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 20:11:45 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import (
	"fmt"
)

func	main(){
	ano := struct{
		name 	string
		last	string
		age		int
	}{
		name: "Araiva",
		last: "PorMueng",
		age: 23,
	}

	fmt.Println(ano)
}