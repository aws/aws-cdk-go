package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter criteria for Lambda event filtering.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table Table
//
//   var fn Function
//
//   fn.AddEventSource(awscdk.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
//   			"dynamodb": map[string]map[string]map[string]interface{}{
//   				"NewImage": map[string]map[string]interface{}{
//   					"id": map[string]interface{}{
//   						"BOOL": lambda.FilterRule_isEqual(jsii.Boolean(true)),
//   					},
//   				},
//   			},
//   		}),
//   	},
//   }))
//
type FilterCriteria interface {
}

// The jsii proxy struct for FilterCriteria
type jsiiProxy_FilterCriteria struct {
	_ byte // padding
}

func NewFilterCriteria() FilterCriteria {
	_init_.Initialize()

	j := jsiiProxy_FilterCriteria{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FilterCriteria",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFilterCriteria_Override(f FilterCriteria) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FilterCriteria",
		nil, // no parameters
		f,
	)
}

// Filter for event source.
func FilterCriteria_Filter(filter *map[string]interface{}) *map[string]interface{} {
	_init_.Initialize()

	if err := validateFilterCriteria_FilterParameters(filter); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterCriteria",
		"filter",
		[]interface{}{filter},
		&returns,
	)

	return returns
}

