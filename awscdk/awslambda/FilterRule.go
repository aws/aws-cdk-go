package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter rules for Lambda event filtering.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Stream: dynamodb.StreamViewType_NEW_IMAGE,
//   })
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
//   		}),
//   	},
//   }))
//
type FilterRule interface {
}

// The jsii proxy struct for FilterRule
type jsiiProxy_FilterRule struct {
	_ byte // padding
}

func NewFilterRule() FilterRule {
	_init_.Initialize()

	j := jsiiProxy_FilterRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FilterRule",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFilterRule_Override(f FilterRule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FilterRule",
		nil, // no parameters
		f,
	)
}

// Begins with comparison operator.
func FilterRule_BeginsWith(elem *string) *[]*map[string]*string {
	_init_.Initialize()

	if err := validateFilterRule_BeginsWithParameters(elem); err != nil {
		panic(err)
	}
	var returns *[]*map[string]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"beginsWith",
		[]interface{}{elem},
		&returns,
	)

	return returns
}

// Numeric range comparison operator.
func FilterRule_Between(first *float64, second *float64) *[]*map[string]*[]interface{} {
	_init_.Initialize()

	if err := validateFilterRule_BetweenParameters(first, second); err != nil {
		panic(err)
	}
	var returns *[]*map[string]*[]interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"between",
		[]interface{}{first, second},
		&returns,
	)

	return returns
}

// Empty comparison operator.
func FilterRule_Empty() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"empty",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Exists comparison operator.
func FilterRule_Exists() *[]*map[string]*bool {
	_init_.Initialize()

	var returns *[]*map[string]*bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"exists",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Equals comparison operator.
func FilterRule_IsEqual(item interface{}) interface{} {
	_init_.Initialize()

	if err := validateFilterRule_IsEqualParameters(item); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"isEqual",
		[]interface{}{item},
		&returns,
	)

	return returns
}

// Not equals comparison operator.
func FilterRule_NotEquals(elem *string) *[]*map[string]*[]*string {
	_init_.Initialize()

	if err := validateFilterRule_NotEqualsParameters(elem); err != nil {
		panic(err)
	}
	var returns *[]*map[string]*[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"notEquals",
		[]interface{}{elem},
		&returns,
	)

	return returns
}

// Not exists comparison operator.
func FilterRule_NotExists() *[]*map[string]*bool {
	_init_.Initialize()

	var returns *[]*map[string]*bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"notExists",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Null comparison operator.
func FilterRule_Null() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"null",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Or comparison operator.
func FilterRule_Or(elem ...*string) *[]*string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range elem {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FilterRule",
		"or",
		args,
		&returns,
	)

	return returns
}

