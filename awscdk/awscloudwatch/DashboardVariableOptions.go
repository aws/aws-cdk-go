package awscloudwatch


// Options for {@link DashboardVariable}.
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
type DashboardVariableOptions struct {
	// Unique id.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The way the variable value is selected.
	InputType VariableInputType `field:"required" json:"inputType" yaml:"inputType"`
	// Type of the variable.
	Type VariableType `field:"required" json:"type" yaml:"type"`
	// Pattern or property value to replace.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Optional default value.
	// Default: - no default value is set.
	//
	DefaultValue DefaultValue `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Optional label in the toolbar.
	// Default: - the variable's value.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Optional values (required for {@link VariableInputType.RADIO} and {@link VariableInputType.SELECT} dashboard variables).
	// Default: - no values.
	//
	Values Values `field:"optional" json:"values" yaml:"values"`
	// Whether the variable is visible.
	// Default: - true.
	//
	Visible *bool `field:"optional" json:"visible" yaml:"visible"`
}

