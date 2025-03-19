package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Default value for use in {@link DashboardVariableOptions}.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   dashboard := cw.NewDashboard(this, jsii.String("Dash"), &DashboardProps{
//   	DefaultInterval: awscdk.Duration_Days(jsii.Number(7)),
//   	Variables: []iVariable{
//   		cw.NewDashboardVariable(&DashboardVariableOptions{
//   			Id: jsii.String("functionName"),
//   			Type: cw.VariableType_PATTERN,
//   			Label: jsii.String("Function"),
//   			InputType: cw.VariableInputType_RADIO,
//   			Value: jsii.String("originalFuncNameInDashboard"),
//   			// equivalent to cw.Values.fromSearch('{AWS/Lambda,FunctionName} MetricName=\"Duration\"', 'FunctionName')
//   			Values: cw.Values_FromSearchComponents(&SearchComponents{
//   				Namespace: jsii.String("AWS/Lambda"),
//   				Dimensions: []*string{
//   					jsii.String("FunctionName"),
//   				},
//   				MetricName: jsii.String("Duration"),
//   				PopulateFrom: jsii.String("FunctionName"),
//   			}),
//   			DefaultValue: cw.DefaultValue_FIRST(),
//   			Visible: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type DefaultValue interface {
	Val() interface{}
}

// The jsii proxy struct for DefaultValue
type jsiiProxy_DefaultValue struct {
	_ byte // padding
}

func (j *jsiiProxy_DefaultValue) Val() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"val",
		&returns,
	)
	return returns
}


// Create a default value.
func DefaultValue_Value(value interface{}) DefaultValue {
	_init_.Initialize()

	if err := validateDefaultValue_ValueParameters(value); err != nil {
		panic(err)
	}
	var returns DefaultValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.DefaultValue",
		"value",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DefaultValue_FIRST() DefaultValue {
	_init_.Initialize()
	var returns DefaultValue
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.DefaultValue",
		"FIRST",
		&returns,
	)
	return returns
}

