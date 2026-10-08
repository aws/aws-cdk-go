package interfacesawscomprehend


// A reference to a EntityRecognizerEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognizerEndpointReference := &EntityRecognizerEndpointReference{
//   	EntityRecognizerEndpointArn: jsii.String("entityRecognizerEndpointArn"),
//   }
//
type EntityRecognizerEndpointReference struct {
	// The Arn of the EntityRecognizerEndpoint resource.
	EntityRecognizerEndpointArn *string `field:"required" json:"entityRecognizerEndpointArn" yaml:"entityRecognizerEndpointArn"`
}

