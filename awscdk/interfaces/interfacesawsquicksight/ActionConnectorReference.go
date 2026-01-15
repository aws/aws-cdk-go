package interfacesawsquicksight


// A reference to a ActionConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionConnectorReference := &ActionConnectorReference{
//   	ActionConnectorArn: jsii.String("actionConnectorArn"),
//   	ActionConnectorId: jsii.String("actionConnectorId"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   }
//
type ActionConnectorReference struct {
	// The ARN of the ActionConnector resource.
	ActionConnectorArn *string `field:"required" json:"actionConnectorArn" yaml:"actionConnectorArn"`
	// The ActionConnectorId of the ActionConnector resource.
	ActionConnectorId *string `field:"required" json:"actionConnectorId" yaml:"actionConnectorId"`
	// The AwsAccountId of the ActionConnector resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
}

