package main

func p33() int {
	dp, np := 1, 1
	for i := 1; i < 10; i++ {
		for den := 1; den < i; den++ {
			for nom := 1; nom < den; nom++ {
				if (nom*10+i)*den == nom*(i*10+den) {
					dp *= den
					np *= nom
				}
			}
		}
	}
	return dp / Gcd(np, dp)
}
