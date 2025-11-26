package previewawsvoiceidevents


// Type definition for RiskDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   riskDetails := &RiskDetails{
//   	KnownFraudsterRisk: &KnownFraudsterRisk{
//   		GeneratedFraudsterId: []*string{
//   			jsii.String("generatedFraudsterId"),
//   		},
//   		RiskScore: []*string{
//   			jsii.String("riskScore"),
//   		},
//   	},
//   	VoiceSpoofingRisk: &VoiceSpoofingRisk{
//   		RiskScore: []*string{
//   			jsii.String("riskScore"),
//   		},
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_RiskDetails struct {
	// knownFraudsterRisk property.
	//
	// Specify an array of string values to match this event if the actual value of knownFraudsterRisk is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KnownFraudsterRisk *DomainEvents_VoiceIdEvaluateSessionAction_KnownFraudsterRisk `field:"optional" json:"knownFraudsterRisk" yaml:"knownFraudsterRisk"`
	// voiceSpoofingRisk property.
	//
	// Specify an array of string values to match this event if the actual value of voiceSpoofingRisk is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VoiceSpoofingRisk *DomainEvents_VoiceIdEvaluateSessionAction_VoiceSpoofingRisk `field:"optional" json:"voiceSpoofingRisk" yaml:"voiceSpoofingRisk"`
}

