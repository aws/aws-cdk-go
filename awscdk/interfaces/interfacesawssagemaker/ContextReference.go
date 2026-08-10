package interfacesawssagemaker


// A reference to a Context resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextReference := &ContextReference{
//   	ContextArn: jsii.String("contextArn"),
//   }
//
type ContextReference struct {
	// The Arn of the Context resource.
	ContextArn *string `field:"required" json:"contextArn" yaml:"contextArn"`
}

