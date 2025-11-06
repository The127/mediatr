package internal

import "reflect"

func TypeOf[T any]() reflect.Type {
	var tp *T
	return reflect.TypeOf(tp).Elem()
}

func Zero[T any]() T {
	var tp T
	return tp
}
