package awsevents


// A reference to a Endpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointReference := &EndpointReference{
//   	EndpointArn: jsii.String("endpointArn"),
//   	EndpointName: jsii.String("endpointName"),
//   }
//
type EndpointReference struct {
	// The ARN of the Endpoint resource.
	EndpointArn *string `field:"required" json:"endpointArn" yaml:"endpointArn"`
	// The Name of the Endpoint resource.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
}

