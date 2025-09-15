package awsec2


// A reference to a NetworkInsightsAccessScope resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInsightsAccessScopeReference := &NetworkInsightsAccessScopeReference{
//   	NetworkInsightsAccessScopeArn: jsii.String("networkInsightsAccessScopeArn"),
//   	NetworkInsightsAccessScopeId: jsii.String("networkInsightsAccessScopeId"),
//   }
//
type NetworkInsightsAccessScopeReference struct {
	// The ARN of the NetworkInsightsAccessScope resource.
	NetworkInsightsAccessScopeArn *string `field:"required" json:"networkInsightsAccessScopeArn" yaml:"networkInsightsAccessScopeArn"`
	// The NetworkInsightsAccessScopeId of the NetworkInsightsAccessScope resource.
	NetworkInsightsAccessScopeId *string `field:"required" json:"networkInsightsAccessScopeId" yaml:"networkInsightsAccessScopeId"`
}

