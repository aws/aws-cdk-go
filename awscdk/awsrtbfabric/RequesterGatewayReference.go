package awsrtbfabric


// A reference to a RequesterGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requesterGatewayReference := &RequesterGatewayReference{
//   	RequesterGatewayArn: jsii.String("requesterGatewayArn"),
//   }
//
type RequesterGatewayReference struct {
	// The Arn of the RequesterGateway resource.
	RequesterGatewayArn *string `field:"required" json:"requesterGatewayArn" yaml:"requesterGatewayArn"`
}

