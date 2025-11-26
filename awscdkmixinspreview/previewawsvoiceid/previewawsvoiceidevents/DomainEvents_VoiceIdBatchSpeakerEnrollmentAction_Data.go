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
//   	EnrollmentConfig: &EnrollmentConfig{
//   		ExistingEnrollmentAction: []*string{
//   			jsii.String("existingEnrollmentAction"),
//   		},
//   		FraudDetectionConfig: &FraudDetectionConfig{
//   			FraudDetectionAction: []*string{
//   				jsii.String("fraudDetectionAction"),
//   			},
//   			FraudDetectionThreshold: []*string{
//   				jsii.String("fraudDetectionThreshold"),
//   			},
//   			WatchlistIds: []*string{
//   				jsii.String("watchlistIds"),
//   			},
//   		},
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
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_Data struct {
	// dataAccessRoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of dataAccessRoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataAccessRoleArn *[]*string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// enrollmentConfig property.
	//
	// Specify an array of string values to match this event if the actual value of enrollmentConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnrollmentConfig *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_EnrollmentConfig `field:"optional" json:"enrollmentConfig" yaml:"enrollmentConfig"`
	// inputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of inputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_InputDataConfig `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// outputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of outputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
}

