package awsconnect


// The contact hooks for the task template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contactHooksProperty := &ContactHooksProperty{
//   	WrapUpFlowArn: jsii.String("wrapUpFlowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-contacthooks.html
//
type CfnTaskTemplatePropsMixin_ContactHooksProperty struct {
	// The identifier (arn) of the flow module that runs when an agent attempts to disconnect from a task created by referencing this template.
	//
	// If no agent ever connects to the task, the flow module runs when the task is terminated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-contacthooks.html#cfn-connect-tasktemplate-contacthooks-wrapupflowarn
	//
	WrapUpFlowArn *string `field:"optional" json:"wrapUpFlowArn" yaml:"wrapUpFlowArn"`
}

