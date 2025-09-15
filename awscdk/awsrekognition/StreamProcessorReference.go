package awsrekognition


// A reference to a StreamProcessor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamProcessorReference := &StreamProcessorReference{
//   	StreamProcessorArn: jsii.String("streamProcessorArn"),
//   	StreamProcessorName: jsii.String("streamProcessorName"),
//   }
//
type StreamProcessorReference struct {
	// The ARN of the StreamProcessor resource.
	StreamProcessorArn *string `field:"required" json:"streamProcessorArn" yaml:"streamProcessorArn"`
	// The Name of the StreamProcessor resource.
	StreamProcessorName *string `field:"required" json:"streamProcessorName" yaml:"streamProcessorName"`
}

