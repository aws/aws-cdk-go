package awslambda


// A reference to a Url resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   urlReference := &UrlReference{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
type UrlReference struct {
	// The FunctionArn of the Url resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

