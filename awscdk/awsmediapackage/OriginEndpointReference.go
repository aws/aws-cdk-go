package awsmediapackage


// A reference to a OriginEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originEndpointReference := &OriginEndpointReference{
//   	OriginEndpointArn: jsii.String("originEndpointArn"),
//   	OriginEndpointId: jsii.String("originEndpointId"),
//   }
//
type OriginEndpointReference struct {
	// The ARN of the OriginEndpoint resource.
	OriginEndpointArn *string `field:"required" json:"originEndpointArn" yaml:"originEndpointArn"`
	// The Id of the OriginEndpoint resource.
	OriginEndpointId *string `field:"required" json:"originEndpointId" yaml:"originEndpointId"`
}

