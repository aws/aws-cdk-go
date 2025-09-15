package awsmediapackagev2


// A reference to a OriginEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originEndpointReference := &OriginEndpointReference{
//   	OriginEndpointArn: jsii.String("originEndpointArn"),
//   }
//
type OriginEndpointReference struct {
	// The Arn of the OriginEndpoint resource.
	OriginEndpointArn *string `field:"required" json:"originEndpointArn" yaml:"originEndpointArn"`
}

