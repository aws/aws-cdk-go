package awscognito


// The compromised credentials actions type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compromisedCredentialsActionsTypeProperty := &compromisedCredentialsActionsTypeProperty{
//   	eventAction: jsii.String("eventAction"),
//   }
//
type CfnUserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsTypeProperty struct {
	// The event action.
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
}

