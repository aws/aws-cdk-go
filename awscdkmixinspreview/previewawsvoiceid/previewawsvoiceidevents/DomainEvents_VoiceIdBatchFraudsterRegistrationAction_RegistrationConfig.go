package previewawsvoiceidevents


// Type definition for RegistrationConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registrationConfig := &RegistrationConfig{
//   	DuplicateRegistrationAction: []*string{
//   		jsii.String("duplicateRegistrationAction"),
//   	},
//   	FraudsterSimilarityThreshold: []*string{
//   		jsii.String("fraudsterSimilarityThreshold"),
//   	},
//   	WatchlistIds: []*string{
//   		jsii.String("watchlistIds"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchFraudsterRegistrationAction_RegistrationConfig struct {
	// duplicateRegistrationAction property.
	//
	// Specify an array of string values to match this event if the actual value of duplicateRegistrationAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DuplicateRegistrationAction *[]*string `field:"optional" json:"duplicateRegistrationAction" yaml:"duplicateRegistrationAction"`
	// fraudsterSimilarityThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of fraudsterSimilarityThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FraudsterSimilarityThreshold *[]*string `field:"optional" json:"fraudsterSimilarityThreshold" yaml:"fraudsterSimilarityThreshold"`
	// watchlistIds property.
	//
	// Specify an array of string values to match this event if the actual value of watchlistIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WatchlistIds *[]*string `field:"optional" json:"watchlistIds" yaml:"watchlistIds"`
}

