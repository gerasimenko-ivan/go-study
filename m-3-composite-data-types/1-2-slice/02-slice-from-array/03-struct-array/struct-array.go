package main

import "fmt"

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	arr := []Person{{name: "a"}, {name: "b"}, {name: "c"}, {name: "d"}}
	sla := arr[1:3]
	fmt.Println(sla) // [{B  } {C  }]

	sla[0].name = "BB"
	fmt.Println(arr) // [{a  } {BB  } {c  } {d  }] - AFFECTED!
	fmt.Println(sla) // [{BB  } {c  }]

	sla = append(sla, Person{name: "DD"})
	fmt.Println(arr) // [{a  } {BB  } {c  } {DD  }] - AFFECTED!
	fmt.Println(sla) // [{BB  } {c  } {DD  }]

	sla = append(sla[:0], sla[1:]...)
	fmt.Println(arr) // [{a  } {c  } {DD  } {DD  }] - AFFECTED!
	fmt.Println(sla) // [{c  } {DD  }]
}
