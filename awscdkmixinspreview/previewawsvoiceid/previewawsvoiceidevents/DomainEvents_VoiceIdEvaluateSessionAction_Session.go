package previewawsvoiceidevents


// Type definition for Session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   session := &Session{
//   	AuthenticationResult: &AuthenticationResult{
//   		AudioAggregationEndedAt: []*string{
//   			jsii.String("audioAggregationEndedAt"),
//   		},
//   		AudioAggregationStartedAt: []*string{
//   			jsii.String("audioAggregationStartedAt"),
//   		},
//   		AuthenticationResultId: []*string{
//   			jsii.String("authenticationResultId"),
//   		},
//   		Configuration: &ConfigurationAuthentication{
//   			AcceptanceThreshold: []*string{
//   				jsii.String("acceptanceThreshold"),
//   			},
//   		},
//   		Decision: []*string{
//   			jsii.String("decision"),
//   		},
//   		Score: []*string{
//   			jsii.String("score"),
//   		},
//   	},
//   	FraudDetectionResult: &FraudDetectionResult{
//   		AudioAggregationEndedAt: []*string{
//   			jsii.String("audioAggregationEndedAt"),
//   		},
//   		AudioAggregationStartedAt: []*string{
//   			jsii.String("audioAggregationStartedAt"),
//   		},
//   		Configuration: &ConfigurationFraud{
//   			RiskThreshold: []*string{
//   				jsii.String("riskThreshold"),
//   			},
//   			WatchlistId: []*string{
//   				jsii.String("watchlistId"),
//   			},
//   		},
//   		Decision: []*string{
//   			jsii.String("decision"),
//   		},
//   		FraudDetectionResultId: []*string{
//   			jsii.String("fraudDetectionResultId"),
//   		},
//   		Reasons: []*string{
//   			jsii.String("reasons"),
//   		},
//   		RiskDetails: &RiskDetails{
//   			KnownFraudsterRisk: &KnownFraudsterRisk{
//   				GeneratedFraudsterId: []*string{
//   					jsii.String("generatedFraudsterId"),
//   				},
//   				RiskScore: []*string{
//   					jsii.String("riskScore"),
//   				},
//   			},
//   			VoiceSpoofingRisk: &VoiceSpoofingRisk{
//   				RiskScore: []*string{
//   					jsii.String("riskScore"),
//   				},
//   			},
//   		},
//   	},
//   	GeneratedSpeakerId: []*string{
//   		jsii.String("generatedSpeakerId"),
//   	},
//   	SessionId: []*string{
//   		jsii.String("sessionId"),
//   	},
//   	SessionName: []*string{
//   		jsii.String("sessionName"),
//   	},
//   	StreamingStatus: []*string{
//   		jsii.String("streamingStatus"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_Session struct {
	// authenticationResult property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationResult is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationResult *DomainEvents_VoiceIdEvaluateSessionAction_AuthenticationResult `field:"optional" json:"authenticationResult" yaml:"authenticationResult"`
	// fraudDetectionResult property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionResult is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionResult *DomainEvents_VoiceIdEvaluateSessionAction_FraudDetectionResult `field:"optional" json:"fraudDetectionResult" yaml:"fraudDetectionResult"`
	// generatedSpeakerId property.
	//
	// Specify an array of string values to match this event if the actual value of generatedSpeakerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GeneratedSpeakerId *[]*string `field:"optional" json:"generatedSpeakerId" yaml:"generatedSpeakerId"`
	// sessionId property.
	//
	// Specify an array of string values to match this event if the actual value of sessionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SessionId *[]*string `field:"optional" json:"sessionId" yaml:"sessionId"`
	// sessionName property.
	//
	// Specify an array of string values to match this event if the actual value of sessionName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SessionName *[]*string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// streamingStatus property.
	//
	// Specify an array of string values to match this event if the actual value of streamingStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamingStatus *[]*string `field:"optional" json:"streamingStatus" yaml:"streamingStatus"`
}

