package previewawssagemakerevents


// Type definition for SageMakerTrainingJobStateChangeItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTrainingJobStateChangeItem1 := &SageMakerTrainingJobStateChangeItem1{
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StatusMessage: []*string{
//   		jsii.String("statusMessage"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeItem1 struct {
	// StartTime property.
	//
	// Specify an array of string values to match this event if the actual value of StartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// Status property.
	//
	// Specify an array of string values to match this event if the actual value of Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// StatusMessage property.
	//
	// Specify an array of string values to match this event if the actual value of StatusMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusMessage *[]*string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}

