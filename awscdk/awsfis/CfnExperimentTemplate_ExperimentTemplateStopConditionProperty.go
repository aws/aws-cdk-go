package awsfis


// Specifies a stop condition for an experiment template.
//
// For more information, see [Stop conditions](https://docs.aws.amazon.com/fis/latest/userguide/stop-conditions.html) in the *AWS Fault Injection Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateStopConditionProperty := &ExperimentTemplateStopConditionProperty{
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html
//
type CfnExperimentTemplate_ExperimentTemplateStopConditionProperty struct {
	// The source for the stop condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html#cfn-fis-experimenttemplate-experimenttemplatestopcondition-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The Amazon Resource Name (ARN) of the CloudWatch alarm, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.html#cfn-fis-experimenttemplate-experimenttemplatestopcondition-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

