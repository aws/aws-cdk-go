package previewawsvoiceidevents


// Type definition for FraudDetectionResult.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fraudDetectionResult := &FraudDetectionResult{
//   	AudioAggregationEndedAt: []*string{
//   		jsii.String("audioAggregationEndedAt"),
//   	},
//   	AudioAggregationStartedAt: []*string{
//   		jsii.String("audioAggregationStartedAt"),
//   	},
//   	Configuration: &ConfigurationFraud{
//   		RiskThreshold: []*string{
//   			jsii.String("riskThreshold"),
//   		},
//   		WatchlistId: []*string{
//   			jsii.String("watchlistId"),
//   		},
//   	},
//   	Decision: []*string{
//   		jsii.String("decision"),
//   	},
//   	FraudDetectionResultId: []*string{
//   		jsii.String("fraudDetectionResultId"),
//   	},
//   	Reasons: []*string{
//   		jsii.String("reasons"),
//   	},
//   	RiskDetails: &RiskDetails{
//   		KnownFraudsterRisk: &KnownFraudsterRisk{
//   			GeneratedFraudsterId: []*string{
//   				jsii.String("generatedFraudsterId"),
//   			},
//   			RiskScore: []*string{
//   				jsii.String("riskScore"),
//   			},
//   		},
//   		VoiceSpoofingRisk: &VoiceSpoofingRisk{
//   			RiskScore: []*string{
//   				jsii.String("riskScore"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_FraudDetectionResult struct {
	// audioAggregationEndedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationEndedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationEndedAt *[]*string `field:"optional" json:"audioAggregationEndedAt" yaml:"audioAggregationEndedAt"`
	// audioAggregationStartedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationStartedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationStartedAt *[]*string `field:"optional" json:"audioAggregationStartedAt" yaml:"audioAggregationStartedAt"`
	// configuration property.
	//
	// Specify an array of string values to match this event if the actual value of configuration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Configuration *DomainEvents_VoiceIdEvaluateSessionAction_ConfigurationFraud `field:"optional" json:"configuration" yaml:"configuration"`
	// decision property.
	//
	// Specify an array of string values to match this event if the actual value of decision is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Decision *[]*string `field:"optional" json:"decision" yaml:"decision"`
	// fraudDetectionResultId property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionResultId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionResultId *[]*string `field:"optional" json:"fraudDetectionResultId" yaml:"fraudDetectionResultId"`
	// reasons property.
	//
	// Specify an array of string values to match this event if the actual value of reasons is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reasons *[]*string `field:"optional" json:"reasons" yaml:"reasons"`
	// riskDetails property.
	//
	// Specify an array of string values to match this event if the actual value of riskDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RiskDetails *DomainEvents_VoiceIdEvaluateSessionAction_RiskDetails `field:"optional" json:"riskDetails" yaml:"riskDetails"`
}

