package previewawsdeadlinemixins


// Properties for CfnQueueEnvironmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueueEnvironmentMixinProps := &CfnQueueEnvironmentMixinProps{
//   	FarmId: jsii.String("farmId"),
//   	Priority: jsii.Number(123),
//   	QueueId: jsii.String("queueId"),
//   	Template: jsii.String("template"),
//   	TemplateType: jsii.String("templateType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html
//
type CfnQueueEnvironmentMixinProps struct {
	// The identifier assigned to the farm that contains the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html#cfn-deadline-queueenvironment-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The queue environment's priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html#cfn-deadline-queueenvironment-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The unique identifier of the queue that contains the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html#cfn-deadline-queueenvironment-queueid
	//
	QueueId *string `field:"optional" json:"queueId" yaml:"queueId"`
	// A JSON or YAML template that describes the processing environment for the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html#cfn-deadline-queueenvironment-template
	//
	Template *string `field:"optional" json:"template" yaml:"template"`
	// Specifies whether the template for the queue environment is JSON or YAML.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html#cfn-deadline-queueenvironment-templatetype
	//
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}

