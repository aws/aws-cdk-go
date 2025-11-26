package previewawsguarddutyevents


// Type definition for DnsRequestAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dnsRequestAction := &DnsRequestAction{
//   	Blocked: []*string{
//   		jsii.String("blocked"),
//   	},
//   	Domain: []*string{
//   		jsii.String("domain"),
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_DnsRequestAction struct {
	// blocked property.
	//
	// Specify an array of string values to match this event if the actual value of blocked is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Blocked *[]*string `field:"optional" json:"blocked" yaml:"blocked"`
	// domain property.
	//
	// Specify an array of string values to match this event if the actual value of domain is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Domain *[]*string `field:"optional" json:"domain" yaml:"domain"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
}

