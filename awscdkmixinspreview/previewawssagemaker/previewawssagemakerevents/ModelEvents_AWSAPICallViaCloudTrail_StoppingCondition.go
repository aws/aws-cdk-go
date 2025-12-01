package previewawssagemakerevents


// Type definition for StoppingCondition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stoppingCondition := &StoppingCondition{
//   	MaxRuntimeInSeconds: []*string{
//   		jsii.String("maxRuntimeInSeconds"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_StoppingCondition struct {
	// maxRuntimeInSeconds property.
	//
	// Specify an array of string values to match this event if the actual value of maxRuntimeInSeconds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxRuntimeInSeconds *[]*string `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

