package awsoam


// A reference to a Sink resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sinkReference := &SinkReference{
//   	SinkArn: jsii.String("sinkArn"),
//   }
//
type SinkReference struct {
	// The Arn of the Sink resource.
	SinkArn *string `field:"required" json:"sinkArn" yaml:"sinkArn"`
}

