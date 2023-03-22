package awsivs


// Properties for creating a new Stream Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var channel channel
//
//   streamKeyProps := &streamKeyProps{
//   	channel: channel,
//   }
//
// Experimental.
type StreamKeyProps struct {
	// Channel ARN for the stream.
	// Experimental.
	Channel IChannel `field:"required" json:"channel" yaml:"channel"`
}

