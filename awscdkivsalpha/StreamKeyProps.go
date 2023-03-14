// The CDK Construct Library for AWS::IVS
package awscdkivsalpha


// Properties for creating a new Stream Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ivs_alpha "github.com/aws/aws-cdk-go/awscdkivsalpha"
//
//   var channel channel
//
//   streamKeyProps := &StreamKeyProps{
//   	Channel: channel,
//   }
//
// Experimental.
type StreamKeyProps struct {
	// Channel ARN for the stream.
	// Experimental.
	Channel IChannel `field:"required" json:"channel" yaml:"channel"`
}

