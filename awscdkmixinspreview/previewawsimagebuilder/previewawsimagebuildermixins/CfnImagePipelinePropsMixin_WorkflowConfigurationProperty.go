package previewawsimagebuildermixins


// Contains control settings and configurable inputs for a workflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnImagePipelinePropsMixin_WorkflowConfigurationProperty struct {
	// The action to take if the workflow fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-onfailure
	//
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Test workflows are defined within named runtime groups called parallel groups.
	//
	// The parallel group is the named group that contains this test workflow. Test workflows within a parallel group can run at the same time. Image Builder starts up to five test workflows in the group at the same time, and starts additional workflows as others complete, until all workflows in the group have completed. This field only applies for test workflows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-parallelgroup
	//
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// Contains parameter values for each of the parameters that the workflow document defined for the workflow resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the workflow resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html#cfn-imagebuilder-imagepipeline-workflowconfiguration-workflowarn
	//
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
}

