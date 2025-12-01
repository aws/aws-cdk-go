package previewawssagemakerevents


// Type definition for HyperParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hyperParameters := &HyperParameters{
//   	EvalMetric: []*string{
//   		jsii.String("evalMetric"),
//   	},
//   	NumRound: []*string{
//   		jsii.String("numRound"),
//   	},
//   	Objective: []*string{
//   		jsii.String("objective"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_HyperParameters struct {
	// eval_metric property.
	//
	// Specify an array of string values to match this event if the actual value of eval_metric is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EvalMetric *[]*string `field:"optional" json:"evalMetric" yaml:"evalMetric"`
	// num_round property.
	//
	// Specify an array of string values to match this event if the actual value of num_round is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumRound *[]*string `field:"optional" json:"numRound" yaml:"numRound"`
	// objective property.
	//
	// Specify an array of string values to match this event if the actual value of objective is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Objective *[]*string `field:"optional" json:"objective" yaml:"objective"`
}

