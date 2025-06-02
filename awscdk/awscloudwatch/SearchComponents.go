package awscloudwatch


// Search components for use with {@link Values.fromSearchComponents}.
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
type SearchComponents struct {
	// The list of dimensions to be used in the search expression.
	Dimensions *[]*string `field:"required" json:"dimensions" yaml:"dimensions"`
	// The metric name to be used in the search expression.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace to be used in the search expression.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The dimension name, that the search expression retrieves, whose values will be used to populate the values to choose from.
	PopulateFrom *string `field:"required" json:"populateFrom" yaml:"populateFrom"`
}

