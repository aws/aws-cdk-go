package previewawsvoiceidevents


// Type definition for AuthenticationConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationConfiguration := &AuthenticationConfiguration{
//   	AcceptanceThreshold: []*string{
//   		jsii.String("acceptanceThreshold"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdStartSessionAction_AuthenticationConfiguration struct {
	// acceptanceThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of acceptanceThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AcceptanceThreshold *[]*string `field:"optional" json:"acceptanceThreshold" yaml:"acceptanceThreshold"`
}

