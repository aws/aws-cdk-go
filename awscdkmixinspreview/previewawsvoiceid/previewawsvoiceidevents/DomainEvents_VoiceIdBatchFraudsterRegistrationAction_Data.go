package previewawsvoiceidevents


// Type definition for Data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   data := &Data{
//   	DataAccessRoleArn: []*string{
//   		jsii.String("dataAccessRoleArn"),
//   	},
//   	InputDataConfig: &InputDataConfig{
//   		S3Uri: []*string{
//   			jsii.String("s3Uri"),
//   		},
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		KmsKeyId: []*string{
//   			jsii.String("kmsKeyId"),
//   		},
//   		S3Uri: []*string{
//   			jsii.String("s3Uri"),
//   		},
//   	},
//   	RegistrationConfig: &RegistrationConfig{
//   		DuplicateRegistrationAction: []*string{
//   			jsii.String("duplicateRegistrationAction"),
//   		},
//   		FraudsterSimilarityThreshold: []*string{
//   			jsii.String("fraudsterSimilarityThreshold"),
//   		},
//   		WatchlistIds: []*string{
//   			jsii.String("watchlistIds"),
//   		},
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchFraudsterRegistrationAction_Data struct {
	// dataAccessRoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of dataAccessRoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataAccessRoleArn *[]*string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// inputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of inputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_InputDataConfig `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// outputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of outputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// registrationConfig property.
	//
	// Specify an array of string values to match this event if the actual value of registrationConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegistrationConfig *DomainEvents_VoiceIdBatchFraudsterRegistrationAction_RegistrationConfig `field:"optional" json:"registrationConfig" yaml:"registrationConfig"`
}

