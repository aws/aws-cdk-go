package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTaskPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTaskMixinProps := &CfnTaskMixinProps{
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskConfiguration: &TaskConfigurationProperty{
//   		ContainerTaskConfiguration: &ContainerTaskConfigurationProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			EcrUri: jsii.String("ecrUri"),
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			ProcessingType: jsii.String("processingType"),
//   			ProcessingUnit: jsii.String("processingUnit"),
//   			TaskExecutionRole: jsii.String("taskExecutionRole"),
//   			TimeoutSeconds: jsii.Number(123),
//   		},
//   	},
//   	TaskName: jsii.String("taskName"),
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html
//
type CfnTaskMixinProps struct {
	// A description of the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html#cfn-iotsitewise-task-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html#cfn-iotsitewise-task-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The task execution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html#cfn-iotsitewise-task-taskconfiguration
	//
	TaskConfiguration interface{} `field:"optional" json:"taskConfiguration" yaml:"taskConfiguration"`
	// The name of the task.
	//
	// Must be unique within the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html#cfn-iotsitewise-task-taskname
	//
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-task.html#cfn-iotsitewise-task-workspacename
	//
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

