package previewawsvoiceidevents


// Type definition for VoiceSpoofingRisk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceSpoofingRisk := &VoiceSpoofingRisk{
//   	RiskScore: []*string{
//   		jsii.String("riskScore"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_VoiceSpoofingRisk struct {
	// riskScore property.
	//
	// Specify an array of string values to match this event if the actual value of riskScore is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RiskScore *[]*string `field:"optional" json:"riskScore" yaml:"riskScore"`
}

