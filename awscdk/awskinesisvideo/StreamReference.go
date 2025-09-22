package awskinesisvideo


// A reference to a Stream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamReference := &StreamReference{
//   	StreamArn: jsii.String("streamArn"),
//   	StreamName: jsii.String("streamName"),
//   }
//
type StreamReference struct {
	// The ARN of the Stream resource.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// The Name of the Stream resource.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

