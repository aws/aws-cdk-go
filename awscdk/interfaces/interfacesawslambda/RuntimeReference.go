package interfacesawslambda


// A reference to a Runtime resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeReference := &RuntimeReference{
//   	RuntimeArn: jsii.String("runtimeArn"),
//   }
//
type RuntimeReference struct {
	// The Arn of the Runtime resource.
	RuntimeArn *string `field:"required" json:"runtimeArn" yaml:"runtimeArn"`
}

