package utils

import (
	"errors"
)

var (
	_l int    = 3
	_m int    = 0
	_x string = "X"
)

func IsMutant(dna []string) (mutant bool, err error) {

	for _, value := range dna {
		if len(dna) != len(value) {
			return false, errors.New("array does not meet size nxn")
		}
	}

	for i := 0; i < 4; i++ {
		order(dna, i)
	}

	if _m > 1 {
		_m = 0
		return true, nil
	} else {
		_m = 0
		return false, nil
	}
}

func order(dna []string, flag int) {
	dnsNew := []string{}
	switch flag {
	case 1:
		for x, _ := range dna {
			_cm := ""
			for _, c := range dna {
				v := string(c[x])
				_cm += v
			}
			dnsNew = append(dnsNew, _cm)
		}
	case 2:
		for a, _ := range dna {
			if len(dna)-3 > a {
				_cmu := ""
				_cmd := ""
				for i, c := range dna {
					if len(dna) > i+a {
						_cmu += string(c[i+a])
					} else {
						_cmu += _x
					}
					if -1 < i-a {
						_cmd += string(c[i-a])
					} else {
						_cmd += _x
					}
				}
				if a != 0 {
					dnsNew = append(dnsNew, _cmd)
				}
				dnsNew = append(dnsNew, _cmu)
			}
		}
	case 3:
		for a, _ := range dna {
			if len(dna)-3 > a {
				_cmu := ""
				_cmd := ""
				for i, c := range dna {
					if len(dna) > i+a {
						_cmu += string(c[len(c)-(i+1+a)])
					} else {
						_cmu += _x
					}
					if -1 < i-a {
						_cmd += string(c[len(c)-(i+1-a)])
					} else {
						_cmd += _x
					}
				}
				if a != 0 {
					dnsNew = append(dnsNew, _cmd)
				}

				dnsNew = append(dnsNew, _cmu)
			}
		}
	default:
		dnsNew = dna
	}
	executed(dnsNew)
}

func executed(dna []string) {
	for _, b := range dna {
		_cm := ""
		c := len(b)
		for i := 0; i < c; i++ {
			if i != 0 {
				ls := string(b[i-1])
				v := string(b[i])
				if v == ls {
					if len(_cm) == _l {
						_m++
						_cm = ""
					} else {
						if len(_cm) == 0 {
							_cm += v
						}
						_cm += v
					}
				} else {
					_cm = ""
				}
			}
		}
	}
}
