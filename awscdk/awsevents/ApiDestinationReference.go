package awsevents


// A reference to a ApiDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiDestinationReference := &ApiDestinationReference{
//   	ApiDestinationArn: jsii.String("apiDestinationArn"),
//   	ApiDestinationName: jsii.String("apiDestinationName"),
//   }
//
type ApiDestinationReference struct {
	// The ARN of the ApiDestination resource.
	ApiDestinationArn *string `field:"required" json:"apiDestinationArn" yaml:"apiDestinationArn"`
	// The Name of the ApiDestination resource.
	ApiDestinationName *string `field:"required" json:"apiDestinationName" yaml:"apiDestinationName"`
}

