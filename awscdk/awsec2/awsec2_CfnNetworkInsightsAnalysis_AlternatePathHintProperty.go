package awsec2


// Describes an potential intermediate component of a feasible path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alternatePathHintProperty := &alternatePathHintProperty{
//   	componentArn: jsii.String("componentArn"),
//   	componentId: jsii.String("componentId"),
//   }
//
type CfnNetworkInsightsAnalysis_AlternatePathHintProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// The ID of the component.
	ComponentId *string `field:"optional" json:"componentId" yaml:"componentId"`
}

