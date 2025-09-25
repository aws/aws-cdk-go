package awsec2


// A reference to a NetworkInsightsPath resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInsightsPathReference := &NetworkInsightsPathReference{
//   	NetworkInsightsPathArn: jsii.String("networkInsightsPathArn"),
//   	NetworkInsightsPathId: jsii.String("networkInsightsPathId"),
//   }
//
type NetworkInsightsPathReference struct {
	// The ARN of the NetworkInsightsPath resource.
	NetworkInsightsPathArn *string `field:"required" json:"networkInsightsPathArn" yaml:"networkInsightsPathArn"`
	// The NetworkInsightsPathId of the NetworkInsightsPath resource.
	NetworkInsightsPathId *string `field:"required" json:"networkInsightsPathId" yaml:"networkInsightsPathId"`
}

