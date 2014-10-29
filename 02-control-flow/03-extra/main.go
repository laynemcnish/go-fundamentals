//TODO type assertion swtich

import "reflect"

func whatAmI(v interface{}) {
	switch t := relect.TypeOf(v); t {
	case t.Kind() == reflect.String:
		return "String"

	default:
		return "Unknown"
	}

}
