package awsgroundstation


// Provides information about how AWS Ground Station should track the satellite through the sky during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackingConfigProperty := &trackingConfigProperty{
//   	autotrack: jsii.String("autotrack"),
//   }
//
type CfnConfig_TrackingConfigProperty struct {
	// Specifies whether or not to use autotrack.
	//
	// `REMOVED` specifies that program track should only be used during the contact. `PREFERRED` specifies that autotracking is preferred during the contact but fallback to program track if the signal is lost. `REQUIRED` specifies that autotracking is required during the contact and not to use program track if the signal is lost.
	Autotrack *string `field:"optional" json:"autotrack" yaml:"autotrack"`
}

