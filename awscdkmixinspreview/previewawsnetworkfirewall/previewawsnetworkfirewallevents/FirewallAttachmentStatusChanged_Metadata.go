package previewawsnetworkfirewallevents


// Type definition for Metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadata := &Metadata{
//   	StateChangeId: []*string{
//   		jsii.String("stateChangeId"),
//   	},
//   }
//
// Experimental.
type FirewallAttachmentStatusChanged_Metadata struct {
	// State Change ID property.
	//
	// Specify an array of string values to match this event if the actual value of State Change ID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateChangeId *[]*string `field:"optional" json:"stateChangeId" yaml:"stateChangeId"`
}

