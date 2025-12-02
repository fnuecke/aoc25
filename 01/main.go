package main

func Run(input []Action) int {
	count := 0
	acc := 50
	for _, action := range input {
		acc += int(action.direction) * action.distance
		for acc < 0 {
			acc += 100
		}
		for acc > 99 {
			acc -= 100
		}
		if acc == 0 {
			count += 1
		}

	}
	return count
}
