package previewawssagemakerevents


// Type definition for ObjectiveStatusCounters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectiveStatusCounters := &ObjectiveStatusCounters{
//   	Failed: []*string{
//   		jsii.String("failed"),
//   	},
//   	Pending: []*string{
//   		jsii.String("pending"),
//   	},
//   	Succeeded: []*string{
//   		jsii.String("succeeded"),
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_ObjectiveStatusCounters struct {
	// Failed property.
	//
	// Specify an array of string values to match this event if the actual value of Failed is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Failed *[]*string `field:"optional" json:"failed" yaml:"failed"`
	// Pending property.
	//
	// Specify an array of string values to match this event if the actual value of Pending is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Pending *[]*string `field:"optional" json:"pending" yaml:"pending"`
	// Succeeded property.
	//
	// Specify an array of string values to match this event if the actual value of Succeeded is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Succeeded *[]*string `field:"optional" json:"succeeded" yaml:"succeeded"`
}

