package awsiotsitewise


// The task execution configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskConfigurationProperty := &TaskConfigurationProperty{
//   	ContainerTaskConfiguration: &ContainerTaskConfigurationProperty{
//   		EcrUri: jsii.String("ecrUri"),
//   		ProcessingType: jsii.String("processingType"),
//   		ProcessingUnit: jsii.String("processingUnit"),
//   		TaskExecutionRole: jsii.String("taskExecutionRole"),
//
//   		// the properties below are optional
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-taskconfiguration.html
//
type CfnTask_TaskConfigurationProperty struct {
	// Configuration for running a custom container image on managed compute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-taskconfiguration.html#cfn-iotsitewise-task-taskconfiguration-containertaskconfiguration
	//
	ContainerTaskConfiguration interface{} `field:"required" json:"containerTaskConfiguration" yaml:"containerTaskConfiguration"`
}

