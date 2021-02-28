package utils

import (
	"errors"
	"sync"
)

var (
	_l int    = 3
	_m int    = 0
	_x string = "X"
)

func monitorWorker(wg *sync.WaitGroup, cs chan int) {
	wg.Wait()
	close(cs)
}

func IsMutant(dna []string) (mutant bool, err error) {

	wg := &sync.WaitGroup{}
	cs := make(chan int)

	for _, value := range dna {
		if len(dna) != len(value) {
			return false, errors.New("array does not meet size nxn")
		}
	}

	wg.Add(4)
	go order(dna, 0, wg, cs)
	go order(dna, 1, wg, cs)
	go order(dna, 2, wg, cs)
	go order(dna, 3, wg, cs)
	go monitorWorker(wg, cs)

	if <-cs > 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func order(dna []string, flag int, wg *sync.WaitGroup, cs chan int) {
	defer wg.Done()
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
	executed(dnsNew, cs)
}

func executed(dna []string, cs chan int) {
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
	cs <- _m
}
