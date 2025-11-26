package previewawsvoiceidevents


// Type definition for Session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   session := &Session{
//   	AuthenticationConfiguration: &AuthenticationConfiguration{
//   		AcceptanceThreshold: []*string{
//   			jsii.String("acceptanceThreshold"),
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
//   }
//
// Experimental.
type DomainEvents_VoiceIdUpdateSessionAction_Session struct {
	// authenticationConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationConfiguration *DomainEvents_VoiceIdUpdateSessionAction_AuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// fraudDetectionConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionConfiguration *DomainEvents_VoiceIdUpdateSessionAction_FraudDetectionConfiguration `field:"optional" json:"fraudDetectionConfiguration" yaml:"fraudDetectionConfiguration"`
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
}

