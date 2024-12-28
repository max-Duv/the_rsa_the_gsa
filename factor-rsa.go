package main

import (
	"fmt"
	"math/big"
)

func factorize(n *big.Int) (p *big.Int, q *big.Int) {
	i := big.NewInt(2)
	sqrtN := new(big.Int).Sqrt(n)

	for i.Cmp(sqrtN) <= 0 {
		if new(big.Int).Mod(n, i).Cmp(big.NewInt(1)) == 0 {
			p = new(big.Int).Set(i)
			q = new(big.Int).Div(n, i)
			return p, q
		}
		i.Add(i, big.NewInt(1))
	}
	return nil, nil
}

func derivePrivKey(p, q, e *big.Int) *big.Int {
	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, big.NewInt(1)),
		new(big.Int).Sub(q, big.NewInt(1)),
	)
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		fmt.Println("Failed to compute Modular Inverse. Set e and phi(n) to coprime values.")
		return nil
	}
	return d
}

func main() {
	modulus := "21205784530233862304670408898601978289868859779205008114018793568459885766749330937868603266318844349189672854499743502036986798132989395172980924706589031814909451810089410757627246967629865390902746619613801758191697561006002760929583381840459561046034052845254926325182106210145017588991470754231159163364995711352195558339022043194134405161017611403300000598690300333746197072853630025431061644225223367113454874956878993488881043155112373759757829767484770710303654742616725646044082574019498090490235337069026713270090328036962689158030704887374065418132632144182617742262070554891198741145781763243470157539819" //Ex modulus (replace with true RSA num)
	n := big.NewInt(2048)
	n.SetString(modulus, 16)
	e := big.NewInt(65537) // Replace w public exponent

	p, q := factorize(n)
	if p != nil && q != nil {
		fmt.Printf("Factors of %d \n Are: %d, and: %d\n", n, p, q)
	} else {
		fmt.Printf("No Factors found for %d\n", n)
	}
	d := derivePrivKey(p, q, e)
	if d != nil {
		fmt.Printf("\n Derived Private Key d: %s\n", d.String())
	} else {
		fmt.Println("No Derived Private Key")
	}
}

// Using OpenSSL; download and view certs from websites
// Run: openssl s_client -connect example.com:443 showcerts
// Save cert to a file
// Extract Public Key Info: openssl x509 -in cert.pem -pubkey -noout > pubkey.pem
// View Public Key Components: openssl rsa -pubin -in pubkey.pem -text -noout
// Look for modulus "nn" in the output.
