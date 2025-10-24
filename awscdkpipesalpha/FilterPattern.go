package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Generate a filter pattern from an input.
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
type FilterPattern interface {
}

// The jsii proxy struct for FilterPattern
type jsiiProxy_FilterPattern struct {
	_ byte // padding
}

// Experimental.
func NewFilterPattern() FilterPattern {
	_init_.Initialize()

	j := jsiiProxy_FilterPattern{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.FilterPattern",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFilterPattern_Override(f FilterPattern) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.FilterPattern",
		nil, // no parameters
		f,
	)
}

// Generates a filter pattern from a JSON object.
// Experimental.
func FilterPattern_FromObject(patternObject *map[string]interface{}) IFilterPattern {
	_init_.Initialize()

	if err := validateFilterPattern_FromObjectParameters(patternObject); err != nil {
		panic(err)
	}
	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.FilterPattern",
		"fromObject",
		[]interface{}{patternObject},
		&returns,
	)

	return returns
}

