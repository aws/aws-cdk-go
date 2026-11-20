package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter criteria for Lambda event filtering.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//   var table Table
//
//
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_or(jsii.String("INSERT"), jsii.String("MODIFY")),
//   			"dynamodb": map[string]map[string]map[string][]map[string]*string{
//   				"NewImage": map[string]map[string][]map[string]*string{
//   					"status": map[string][]map[string]*string{
//   						"S": lambda.FilterRule_beginsWith(jsii.String("approved")),
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

