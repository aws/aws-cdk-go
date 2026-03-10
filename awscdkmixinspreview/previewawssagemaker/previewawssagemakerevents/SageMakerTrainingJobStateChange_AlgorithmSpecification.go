package previewawssagemakerevents


// Type definition for AlgorithmSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   algorithmSpecification := &AlgorithmSpecification{
//   	TrainingImage: []*string{
//   		jsii.String("trainingImage"),
//   	},
//   	TrainingInputMode: []*string{
//   		jsii.String("trainingInputMode"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_AlgorithmSpecification struct {
	// TrainingImage property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingImage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingImage *[]*string `field:"optional" json:"trainingImage" yaml:"trainingImage"`
	// TrainingInputMode property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingInputMode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingInputMode *[]*string `field:"optional" json:"trainingInputMode" yaml:"trainingInputMode"`
}

