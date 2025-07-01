package awscloudwatch


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
type VariableInputType string

const (
	// Freeform text input box.
	VariableInputType_INPUT VariableInputType = "INPUT"
	// A dropdown of pre-defined values, or values filled in from a metric search query.
	VariableInputType_RADIO VariableInputType = "RADIO"
	// A set of pre-defined radio buttons, which can also be defined from a metric search query.
	VariableInputType_SELECT VariableInputType = "SELECT"
)

