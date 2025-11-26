package previewawsgluemixins


// Properties for CfnWorkflowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var defaultRunProperties interface{}
//   var tags interface{}
//
//   cfnWorkflowMixinProps := &CfnWorkflowMixinProps{
//   	DefaultRunProperties: defaultRunProperties,
//   	Description: jsii.String("description"),
//   	MaxConcurrentRuns: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html
//
type CfnWorkflowMixinProps struct {
	// A collection of properties to be used as part of each execution of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-defaultrunproperties
	//
	DefaultRunProperties interface{} `field:"optional" json:"defaultRunProperties" yaml:"defaultRunProperties"`
	// A description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs.
	//
	// If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-maxconcurrentruns
	//
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The name of the workflow representing the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to use with this workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

