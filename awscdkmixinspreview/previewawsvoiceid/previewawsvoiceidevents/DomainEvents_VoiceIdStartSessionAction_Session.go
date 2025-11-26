package previewawsvoiceidevents


// Type definition for Session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   session := &Session{
//   	AuthenticationAudioProgress: &AuthenticationAudioProgress{
//   		AudioAggregationEndedAt: []*string{
//   			jsii.String("audioAggregationEndedAt"),
//   		},
//   		AudioAggregationStartedAt: []*string{
//   			jsii.String("audioAggregationStartedAt"),
//   		},
//   	},
//   	AuthenticationConfiguration: &AuthenticationConfiguration{
//   		AcceptanceThreshold: []*string{
//   			jsii.String("acceptanceThreshold"),
//   		},
//   	},
//   	EnrollmentAudioProgress: &EnrollmentAudioProgress{
//   		AudioAggregationEndedAt: []*string{
//   			jsii.String("audioAggregationEndedAt"),
//   		},
//   		AudioAggregationStartedAt: []*string{
//   			jsii.String("audioAggregationStartedAt"),
//   		},
//   		AudioAggregationStatus: []*string{
//   			jsii.String("audioAggregationStatus"),
//   		},
//   	},
//   	FraudDetectionAudioProgress: &AuthenticationAudioProgress{
//   		AudioAggregationEndedAt: []*string{
//   			jsii.String("audioAggregationEndedAt"),
//   		},
//   		AudioAggregationStartedAt: []*string{
//   			jsii.String("audioAggregationStartedAt"),
//   		},
//   	},
//   	FraudDetectionConfiguration: &FraudDetectionConfiguration{
//   		RiskThreshold: []*string{
//   			jsii.String("riskThreshold"),
//   		},
//   		WatchlistId: []*string{
//   			jsii.String("watchlistId"),
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
//   	StreamingConfiguration: &StreamingConfiguration{
//   		AuthenticationMinimumSpeechInSeconds: []*string{
//   			jsii.String("authenticationMinimumSpeechInSeconds"),
//   		},
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdStartSessionAction_Session struct {
	// authenticationAudioProgress property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationAudioProgress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationAudioProgress *DomainEvents_VoiceIdStartSessionAction_AuthenticationAudioProgress `field:"optional" json:"authenticationAudioProgress" yaml:"authenticationAudioProgress"`
	// authenticationConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationConfiguration *DomainEvents_VoiceIdStartSessionAction_AuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// enrollmentAudioProgress property.
	//
	// Specify an array of string values to match this event if the actual value of enrollmentAudioProgress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnrollmentAudioProgress *DomainEvents_VoiceIdStartSessionAction_EnrollmentAudioProgress `field:"optional" json:"enrollmentAudioProgress" yaml:"enrollmentAudioProgress"`
	// fraudDetectionAudioProgress property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionAudioProgress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionAudioProgress *DomainEvents_VoiceIdStartSessionAction_AuthenticationAudioProgress `field:"optional" json:"fraudDetectionAudioProgress" yaml:"fraudDetectionAudioProgress"`
	// fraudDetectionConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionConfiguration *DomainEvents_VoiceIdStartSessionAction_FraudDetectionConfiguration `field:"optional" json:"fraudDetectionConfiguration" yaml:"fraudDetectionConfiguration"`
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
	// streamingConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of streamingConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamingConfiguration *DomainEvents_VoiceIdStartSessionAction_StreamingConfiguration `field:"optional" json:"streamingConfiguration" yaml:"streamingConfiguration"`
}

