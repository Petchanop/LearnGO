/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex3.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/11/24 19:32:59 by npiya-is          #+#    #+#             */
/*   Updated: 2022/11/24 19:47:20 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import (
	"fmt"
)

type vehicle struct {
	doors	int
	color	string
}

type truck struct {
	vehicle
	fourWheel	bool
}

type sedan struct {
	vehicle
	luxury	bool
}

func main()  {
	car1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Blue", 
		},
		fourWheel: true,
	}
	
	fmt.Println(car1)

	car2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		luxury: false,
	}

	fmt.Println(car2)

	fmt.Println(car1.vehicle)
	fmt.Println(car2.vehicle)
}