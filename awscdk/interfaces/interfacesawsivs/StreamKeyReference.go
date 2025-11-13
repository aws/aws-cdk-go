package interfacesawsivs


// A reference to a StreamKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamKeyReference := &StreamKeyReference{
//   	StreamKeyArn: jsii.String("streamKeyArn"),
//   }
//
type StreamKeyReference struct {
	// The Arn of the StreamKey resource.
	StreamKeyArn *string `field:"required" json:"streamKeyArn" yaml:"streamKeyArn"`
}

