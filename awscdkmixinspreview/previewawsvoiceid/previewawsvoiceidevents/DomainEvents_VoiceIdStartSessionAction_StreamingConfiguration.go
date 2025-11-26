package previewawsvoiceidevents


// Type definition for StreamingConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamingConfiguration := &StreamingConfiguration{
//   	AuthenticationMinimumSpeechInSeconds: []*string{
//   		jsii.String("authenticationMinimumSpeechInSeconds"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdStartSessionAction_StreamingConfiguration struct {
	// authenticationMinimumSpeechInSeconds property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationMinimumSpeechInSeconds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationMinimumSpeechInSeconds *[]*string `field:"optional" json:"authenticationMinimumSpeechInSeconds" yaml:"authenticationMinimumSpeechInSeconds"`
}

