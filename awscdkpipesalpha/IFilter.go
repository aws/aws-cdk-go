package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The collection of event patterns used to filter events.
// Experimental.
type IFilter interface {
	// Filters for the source.
	// Experimental.
	Filters() *[]IFilterPattern
	// Experimental.
	SetFilters(f *[]IFilterPattern)
}

// The jsii proxy for IFilter
type jsiiProxy_IFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilter) Filters() *[]IFilterPattern {
	var returns *[]IFilterPattern
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilter)SetFilters(val *[]IFilterPattern) {
	if err := j.validateSetFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filters",
		val,
	)
}

