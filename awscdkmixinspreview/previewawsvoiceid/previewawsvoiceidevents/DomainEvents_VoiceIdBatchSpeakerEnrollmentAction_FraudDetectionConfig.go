package previewawsvoiceidevents


// Type definition for FraudDetectionConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fraudDetectionConfig := &FraudDetectionConfig{
//   	FraudDetectionAction: []*string{
//   		jsii.String("fraudDetectionAction"),
//   	},
//   	FraudDetectionThreshold: []*string{
//   		jsii.String("fraudDetectionThreshold"),
//   	},
//   	WatchlistIds: []*string{
//   		jsii.String("watchlistIds"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_FraudDetectionConfig struct {
	// fraudDetectionAction property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionAction *[]*string `field:"optional" json:"fraudDetectionAction" yaml:"fraudDetectionAction"`
	// fraudDetectionThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionThreshold *[]*string `field:"optional" json:"fraudDetectionThreshold" yaml:"fraudDetectionThreshold"`
	// watchlistIds property.
	//
	// Specify an array of string values to match this event if the actual value of watchlistIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WatchlistIds *[]*string `field:"optional" json:"watchlistIds" yaml:"watchlistIds"`
}

