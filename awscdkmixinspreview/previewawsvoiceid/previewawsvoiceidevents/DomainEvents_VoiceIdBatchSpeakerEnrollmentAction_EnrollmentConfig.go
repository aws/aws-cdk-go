package previewawsvoiceidevents


// Type definition for EnrollmentConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enrollmentConfig := &EnrollmentConfig{
//   	ExistingEnrollmentAction: []*string{
//   		jsii.String("existingEnrollmentAction"),
//   	},
//   	FraudDetectionConfig: &FraudDetectionConfig{
//   		FraudDetectionAction: []*string{
//   			jsii.String("fraudDetectionAction"),
//   		},
//   		FraudDetectionThreshold: []*string{
//   			jsii.String("fraudDetectionThreshold"),
//   		},
//   		WatchlistIds: []*string{
//   			jsii.String("watchlistIds"),
//   		},
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_EnrollmentConfig struct {
	// existingEnrollmentAction property.
	//
	// Specify an array of string values to match this event if the actual value of existingEnrollmentAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExistingEnrollmentAction *[]*string `field:"optional" json:"existingEnrollmentAction" yaml:"existingEnrollmentAction"`
	// fraudDetectionConfig property.
	//
	// Specify an array of string values to match this event if the actual value of fraudDetectionConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudDetectionConfig *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_FraudDetectionConfig `field:"optional" json:"fraudDetectionConfig" yaml:"fraudDetectionConfig"`
}

