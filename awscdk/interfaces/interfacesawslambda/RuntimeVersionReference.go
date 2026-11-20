package interfacesawslambda


// A reference to a RuntimeVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeVersionReference := &RuntimeVersionReference{
//   	RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   }
//
type RuntimeVersionReference struct {
	// The Arn of the RuntimeVersion resource.
	RuntimeVersionArn *string `field:"required" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}

