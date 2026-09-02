package interfacesawsiot


// A reference to a Stream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamReference := &StreamReference{
//   	StreamArn: jsii.String("streamArn"),
//   }
//
type StreamReference struct {
	// The Arn of the Stream resource.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

