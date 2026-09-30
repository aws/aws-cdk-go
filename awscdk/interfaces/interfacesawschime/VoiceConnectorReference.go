package interfacesawschime


// A reference to a VoiceConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   voiceConnectorReference := &VoiceConnectorReference{
//   	VoiceConnectorArn: jsii.String("voiceConnectorArn"),
//   }
//
type VoiceConnectorReference struct {
	// The VoiceConnectorArn of the VoiceConnector resource.
	VoiceConnectorArn *string `field:"required" json:"voiceConnectorArn" yaml:"voiceConnectorArn"`
}

