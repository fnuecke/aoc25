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

func Run2(input []Action) int {
	count := 0
	acc := 50
	for _, action := range input {
		delta := int(action.direction) * action.distance
		if delta >= 100 {
			count += delta / 100
			delta = delta % 100
		}
		if delta <= -100 {
			count += (-delta) / 100
			delta = -((-delta) % 100)
		}

		if delta == 0 {
			continue
		}

		acc += delta
		if acc == 0 {
			count += 1
		} else if acc < 0 {
			if acc != delta {
				count += 1
			}
			acc += 100
		} else if acc > 99 {
			acc -= 100
			count += 1
		}
	}
	return count
}
