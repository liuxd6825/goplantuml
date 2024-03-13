package cmd

import "github.com/jfeliu007/goplantuml/pkg/tocode/classdiagram"

type Aggregate struct {
	Class     *classdiagram.Class
	Namespace *classdiagram.Namespace
}

func NewAggregate(namespace *classdiagram.Namespace, class *classdiagram.Class) *Aggregate {
	return &Aggregate{
		Namespace: namespace,
		Class:     class,
	}
}
