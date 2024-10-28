package lib

// Make sure the function name is uppercase, otherwise it will be private!!!
func IsPrimeNumber(number uint64) (isPrime bool) {
	// check if the number is prime or not

	if number < 2 {
		return false
	}
	for i := uint64(2); i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
