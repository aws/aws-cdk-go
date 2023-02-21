package awspipes


// The parameters for using an AWS Batch job as a target.
//
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
	// The job definition used by this job.
	//
	// This value can be one of `name` , `name:revision` , or the Amazon Resource Name (ARN) for the job definition. If name is specified without a revision then the latest active revision is used.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// The name of the job.
	//
	// It can be up to 128 letters long. The first character must be alphanumeric, can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The array properties for the submitted job, such as the size of the array.
	//
	// The array size can be between 2 and 10,000. If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.
	ArrayProperties interface{} `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// The overrides that are sent to a container.
	ContainerOverrides interface{} `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// A list of dependencies for the job.
	//
	// A job can depend upon a maximum of 20 jobs. You can specify a `SEQUENTIAL` type dependency without specifying a job ID for array jobs so that each child array job completes sequentially, starting at index 0. You can also specify an `N_TO_N` type dependency with a job ID for array jobs. In that case, each index child of this job must wait for the corresponding index child of each dependency to complete before it can begin.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Additional parameters passed to the job that replace parameter substitution placeholders that are set in the job definition.
	//
	// Parameters are specified as a key and value pair mapping. Parameters included here override any corresponding parameter defaults from the job definition.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The retry strategy to use for failed jobs.
	//
	// When a retry strategy is specified here, it overrides the retry strategy defined in the job definition.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

