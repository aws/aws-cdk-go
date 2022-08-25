package awsfis


// Specifies a stop condition for an experiment template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateStopConditionProperty := &experimentTemplateStopConditionProperty{
//   	source: jsii.String("source"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnExperimentTemplate_ExperimentTemplateStopConditionProperty struct {
	// The source for the stop condition.
	//
	// Specify `aws:cloudwatch:alarm` if the stop condition is defined by a CloudWatch alarm. Specify `none` if there is no stop condition.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The Amazon Resource Name (ARN) of the CloudWatch alarm.
	//
	// This is required if the source is a CloudWatch alarm.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

