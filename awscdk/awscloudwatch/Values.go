package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class for providing values for use with {@link VariableInputType.SELECT} and {@link VariableInputType.RADIO} dashboard variables.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   dashboard := cw.NewDashboard(this, jsii.String("Dash"), &DashboardProps{
//   	DefaultInterval: awscdk.Duration_Days(jsii.Number(7)),
//   	Variables: []iVariable{
//   		cw.NewDashboardVariable(&DashboardVariableOptions{
//   			Id: jsii.String("region"),
//   			Type: cw.VariableType_PROPERTY,
//   			Label: jsii.String("Region"),
//   			InputType: cw.VariableInputType_RADIO,
//   			Value: jsii.String("region"),
//   			Values: cw.Values_FromValues(&VariableValue{
//   				Label: jsii.String("IAD"),
//   				Value: jsii.String("us-east-1"),
//   			}, &VariableValue{
//   				Label: jsii.String("DUB"),
//   				Value: jsii.String("us-west-2"),
//   			}),
//   			DefaultValue: cw.DefaultValue_Value(jsii.String("us-east-1")),
//   			Visible: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type Values interface {
	ToJson() interface{}
}

// The jsii proxy struct for Values
type jsiiProxy_Values struct {
	_ byte // padding
}

func NewValues_Override(v Values) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.Values",
		nil, // no parameters
		v,
	)
}

// Create values from a search expression.
func Values_FromSearch(expression *string, populateFrom *string) Values {
	_init_.Initialize()

	if err := validateValues_FromSearchParameters(expression, populateFrom); err != nil {
		panic(err)
	}
	var returns Values

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Values",
		"fromSearch",
		[]interface{}{expression, populateFrom},
		&returns,
	)

	return returns
}

// Create values from the components of search expression.
func Values_FromSearchComponents(components *SearchComponents) Values {
	_init_.Initialize()

	if err := validateValues_FromSearchComponentsParameters(components); err != nil {
		panic(err)
	}
	var returns Values

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Values",
		"fromSearchComponents",
		[]interface{}{components},
		&returns,
	)

	return returns
}

// Create values from an array of possible variable values.
func Values_FromValues(values ...*VariableValue) Values {
	_init_.Initialize()

	if err := validateValues_FromValuesParameters(&values); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns Values

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.Values",
		"fromValues",
		args,
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Values) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

