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
type VariableType string

const (
	// A property variable changes the values of all instances of a property in the list of widgets in the dashboard.
	VariableType_PROPERTY VariableType = "PROPERTY"
	// A pattern variable is one that changes a regex pattern across the dashboard JSON.
	VariableType_PATTERN VariableType = "PATTERN"
)

