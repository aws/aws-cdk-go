package awscloudwatch


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
type VariableValue struct {
	// Value of the selected item.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Optional label for the selected item.
	// Default: - the variable's value.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
}

