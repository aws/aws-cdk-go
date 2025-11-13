package interfacesawsgameliftstreams


// A reference to a StreamGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamGroupReference := &StreamGroupReference{
//   	StreamGroupArn: jsii.String("streamGroupArn"),
//   }
//
type StreamGroupReference struct {
	// The Arn of the StreamGroup resource.
	StreamGroupArn *string `field:"required" json:"streamGroupArn" yaml:"streamGroupArn"`
}

