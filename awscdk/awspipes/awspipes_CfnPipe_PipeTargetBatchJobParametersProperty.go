package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetBatchJobParametersProperty := &pipeTargetBatchJobParametersProperty{
//   	jobDefinition: jsii.String("jobDefinition"),
//   	jobName: jsii.String("jobName"),
//
//   	// the properties below are optional
//   	arrayProperties: &batchArrayPropertiesProperty{
//   		size: jsii.Number(123),
//   	},
//   	containerOverrides: &batchContainerOverridesProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&batchEnvironmentVariableProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		resourceRequirements: []interface{}{
//   			&batchResourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	dependsOn: []interface{}{
//   		&batchJobDependencyProperty{
//   			jobId: jsii.String("jobId"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	retryStrategy: &batchRetryStrategyProperty{
//   		attempts: jsii.Number(123),
//   	},
//   }
//
type CfnPipe_PipeTargetBatchJobParametersProperty struct {
	// `CfnPipe.PipeTargetBatchJobParametersProperty.JobDefinition`.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.JobName`.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.ArrayProperties`.
	ArrayProperties interface{} `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.ContainerOverrides`.
	ContainerOverrides interface{} `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.DependsOn`.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.Parameters`.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// `CfnPipe.PipeTargetBatchJobParametersProperty.RetryStrategy`.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

