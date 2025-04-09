package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Dashboard Variable.
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
type DashboardVariable interface {
	IVariable
	// Return the variable JSON for use in the dashboard.
	ToJson() interface{}
}

// The jsii proxy struct for DashboardVariable
type jsiiProxy_DashboardVariable struct {
	jsiiProxy_IVariable
}

func NewDashboardVariable(options *DashboardVariableOptions) DashboardVariable {
	_init_.Initialize()

	if err := validateNewDashboardVariableParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardVariable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.DashboardVariable",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewDashboardVariable_Override(d DashboardVariable, options *DashboardVariableOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.DashboardVariable",
		[]interface{}{options},
		d,
	)
}

func (d *jsiiProxy_DashboardVariable) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

