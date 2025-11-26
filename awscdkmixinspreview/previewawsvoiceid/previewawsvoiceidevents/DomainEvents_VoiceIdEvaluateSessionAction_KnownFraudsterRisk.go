package previewawsvoiceidevents


// Type definition for KnownFraudsterRisk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   knownFraudsterRisk := &KnownFraudsterRisk{
//   	GeneratedFraudsterId: []*string{
//   		jsii.String("generatedFraudsterId"),
//   	},
//   	RiskScore: []*string{
//   		jsii.String("riskScore"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_KnownFraudsterRisk struct {
	// generatedFraudsterId property.
	//
	// Specify an array of string values to match this event if the actual value of generatedFraudsterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GeneratedFraudsterId *[]*string `field:"optional" json:"generatedFraudsterId" yaml:"generatedFraudsterId"`
	// riskScore property.
	//
	// Specify an array of string values to match this event if the actual value of riskScore is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RiskScore *[]*string `field:"optional" json:"riskScore" yaml:"riskScore"`
}

