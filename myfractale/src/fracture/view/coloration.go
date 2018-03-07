package view

type Coloration func(float64) interface{}

var ascii []rune= []rune{'.', ',', '-', '~', ':', '*', '(', '\\', '!', '|', '?', '/', ')', '£', '$', '@', '%', '#'}

func ASCII(val float64) interface{} {
	intval := int(val)
	var c rune

	if intval < 0.0 {
		c = ' '
	} else {
		intval %= 2 * len(ascii)
		if intval/len(ascii) == 0 {
			c = ascii[intval]
		} else {
			c = ascii[len(ascii)-(intval%len(ascii)) - 1]
		}
	}

	return c
}

func BlackAndWhite(val float64) interface{} {
  intval := int(val) % 510
  var res uint8

  if intval / 255 == 0 {
    res = uint8(intval)
  } else {
    res = uint8(255 - (intval - 255))
  }

  return res
}
