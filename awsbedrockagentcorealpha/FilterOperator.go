package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Filter operators for online evaluation filtering.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   filterOperator := bedrock_agentcore_alpha.FilterOperator_CONTAINS()
//
// Experimental.
type FilterOperator interface {
	// The string value of the filter operator.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for FilterOperator
type jsiiProxy_FilterOperator struct {
	_ byte // padding
}

func (j *jsiiProxy_FilterOperator) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewFilterOperator(value *string) FilterOperator {
	_init_.Initialize()

	if err := validateNewFilterOperatorParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_FilterOperator{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewFilterOperator_Override(f FilterOperator, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		[]interface{}{value},
		f,
	)
}

func FilterOperator_CONTAINS() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"CONTAINS",
		&returns,
	)
	return returns
}

func FilterOperator_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_GREATER_THAN() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"GREATER_THAN",
		&returns,
	)
	return returns
}

func FilterOperator_GREATER_THAN_OR_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"GREATER_THAN_OR_EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_LESS_THAN() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"LESS_THAN",
		&returns,
	)
	return returns
}

func FilterOperator_LESS_THAN_OR_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"LESS_THAN_OR_EQUAL",
		&returns,
	)
	return returns
}

func FilterOperator_NOT_CONTAINS() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"NOT_CONTAINS",
		&returns,
	)
	return returns
}

func FilterOperator_NOT_EQUAL() FilterOperator {
	_init_.Initialize()
	var returns FilterOperator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterOperator",
		"NOT_EQUAL",
		&returns,
	)
	return returns
}

