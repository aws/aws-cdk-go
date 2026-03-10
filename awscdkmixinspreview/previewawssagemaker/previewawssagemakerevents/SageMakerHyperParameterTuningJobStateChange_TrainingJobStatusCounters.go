package previewawssagemakerevents


// Type definition for TrainingJobStatusCounters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trainingJobStatusCounters := &TrainingJobStatusCounters{
//   	Completed: []*string{
//   		jsii.String("completed"),
//   	},
//   	InProgress: []*string{
//   		jsii.String("inProgress"),
//   	},
//   	NonRetryableError: []*string{
//   		jsii.String("nonRetryableError"),
//   	},
//   	RetryableError: []*string{
//   		jsii.String("retryableError"),
//   	},
//   	Stopped: []*string{
//   		jsii.String("stopped"),
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_TrainingJobStatusCounters struct {
	// Completed property.
	//
	// Specify an array of string values to match this event if the actual value of Completed is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Completed *[]*string `field:"optional" json:"completed" yaml:"completed"`
	// InProgress property.
	//
	// Specify an array of string values to match this event if the actual value of InProgress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InProgress *[]*string `field:"optional" json:"inProgress" yaml:"inProgress"`
	// NonRetryableError property.
	//
	// Specify an array of string values to match this event if the actual value of NonRetryableError is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NonRetryableError *[]*string `field:"optional" json:"nonRetryableError" yaml:"nonRetryableError"`
	// RetryableError property.
	//
	// Specify an array of string values to match this event if the actual value of RetryableError is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetryableError *[]*string `field:"optional" json:"retryableError" yaml:"retryableError"`
	// Stopped property.
	//
	// Specify an array of string values to match this event if the actual value of Stopped is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Stopped *[]*string `field:"optional" json:"stopped" yaml:"stopped"`
}

