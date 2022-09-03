package reflection

type Dict map[interface{}]interface{}
type Coll []interface{}

type PersonKey struct{}
type Person struct{}

func PutPerson(p Person, d Dict) {
	d[PersonKey{}] = p
}

func GetPerson(d Dict) Person {
	switch val := d[Person{}].(type) {
	case Person:
		return val
	default:
		panic("Shit, types doesn't match")
	}
}

func PutToMap(item ...interface{}) map[string]interface{} {
	if len(item)%2 != 0 {
		panic("Wrong list of arguments")
	}

	ret := make(map[string]interface{})

	for i := 0; i < len(item); i += 2 {
		key := item[i].(string)
		ret[key] = item[i+1]

	}
	return ret
}
