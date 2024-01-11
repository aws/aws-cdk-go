package awsfis


// Describes a filter used for the target resources in an experiment template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateTargetFilterProperty := &ExperimentTemplateTargetFilterProperty{
//   	Path: jsii.String("path"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html
//
type CfnExperimentTemplate_ExperimentTemplateTargetFilterProperty struct {
	// The attribute path for the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html#cfn-fis-experimenttemplate-experimenttemplatetargetfilter-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// The attribute values for the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetargetfilter.html#cfn-fis-experimenttemplate-experimenttemplatetargetfilter-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

