package awsfis


// Specifies a filter used for the target resource input in an experiment template.
//
// For more information, see [Resource filters](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateTargetFilterProperty := &experimentTemplateTargetFilterProperty{
//   	path: jsii.String("path"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateTargetFilterProperty struct {
	// The attribute path for the filter.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The attribute values for the filter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

