package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The collection of event patterns used to filter events.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//
//
//   sourceFilter := pipes.NewFilter([]IFilterPattern{
//   	pipes.FilterPattern_FromObject(map[string]interface{}{
//   		"body": map[string][]*string{
//   			// only forward events with customerType B2B or B2C
//   			"customerType": []*string{
//   				jsii.String("B2B"),
//   				jsii.String("B2C"),
//   			},
//   		},
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   	Filter: sourceFilter,
//   })
//
// Experimental.
type Filter interface {
	IFilter
	// Filters for the source.
	// Experimental.
	Filters() *[]IFilterPattern
	// Experimental.
	SetFilters(val *[]IFilterPattern)
}

// The jsii proxy struct for Filter
type jsiiProxy_Filter struct {
	jsiiProxy_IFilter
}

func (j *jsiiProxy_Filter) Filters() *[]IFilterPattern {
	var returns *[]IFilterPattern
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}


// Experimental.
func NewFilter(filter *[]IFilterPattern) Filter {
	_init_.Initialize()

	if err := validateNewFilterParameters(filter); err != nil {
		panic(err)
	}
	j := jsiiProxy_Filter{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.Filter",
		[]interface{}{filter},
		&j,
	)

	return &j
}

// Experimental.
func NewFilter_Override(f Filter, filter *[]IFilterPattern) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.Filter",
		[]interface{}{filter},
		f,
	)
}

func (j *jsiiProxy_Filter)SetFilters(val *[]IFilterPattern) {
	if err := j.validateSetFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filters",
		val,
	)
}

