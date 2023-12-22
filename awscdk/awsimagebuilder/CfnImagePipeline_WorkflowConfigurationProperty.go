package awsimagebuilder


// The workflow configuration of the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowConfigurationProperty := &WorkflowConfigurationProperty{
//   	OnFailure: jsii.String("onFailure"),
//   	ParallelGroup: jsii.String("parallelGroup"),
//   	Parameters: []interface{}{
//   		&WorkflowParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	WorkflowArn: jsii.String("workflowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html
//
type CfnImagePipeline_WorkflowConfigurationProperty struct {
	// Define execution decision in case of workflow failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-onfailure
	//
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The parallel group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-parallelgroup
	//
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// The parameters associated with the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-workflowarn
	//
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
}

