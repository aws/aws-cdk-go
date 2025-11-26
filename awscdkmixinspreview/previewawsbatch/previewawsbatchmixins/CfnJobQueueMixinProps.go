package previewawsbatchmixins


// Properties for CfnJobQueuePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnJobQueueMixinProps := &CfnJobQueueMixinProps{
//   	ComputeEnvironmentOrder: []interface{}{
//   		&ComputeEnvironmentOrderProperty{
//   			ComputeEnvironment: jsii.String("computeEnvironment"),
//   			Order: jsii.Number(123),
//   		},
//   	},
//   	JobQueueName: jsii.String("jobQueueName"),
//   	JobQueueType: jsii.String("jobQueueType"),
//   	JobStateTimeLimitActions: []interface{}{
//   		&JobStateTimeLimitActionProperty{
//   			Action: jsii.String("action"),
//   			MaxTimeSeconds: jsii.Number(123),
//   			Reason: jsii.String("reason"),
//   			State: jsii.String("state"),
//   		},
//   	},
//   	Priority: jsii.Number(123),
//   	SchedulingPolicyArn: jsii.String("schedulingPolicyArn"),
//   	ServiceEnvironmentOrder: []interface{}{
//   		&ServiceEnvironmentOrderProperty{
//   			Order: jsii.Number(123),
//   			ServiceEnvironment: jsii.String("serviceEnvironment"),
//   		},
//   	},
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html
//
type CfnJobQueueMixinProps struct {
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to determine which compute environment runs a specific job. Compute environments must be in the `VALID` state before you can associate them with a job queue. You can associate up to three compute environments with a job queue. All of the compute environments must be either EC2 ( `EC2` or `SPOT` ) or Fargate ( `FARGATE` or `FARGATE_SPOT` ); EC2 and Fargate compute environments can't be mixed.
	//
	// > All compute environments that are associated with a job queue must share the same architecture. AWS Batch doesn't support mixing compute environment architecture types in a single job queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-computeenvironmentorder
	//
	ComputeEnvironmentOrder interface{} `field:"optional" json:"computeEnvironmentOrder" yaml:"computeEnvironmentOrder"`
	// The name of the job queue.
	//
	// It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobqueuename
	//
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// The type of job queue.
	//
	// For service jobs that run on SageMaker AI , this value is `SAGEMAKER_TRAINING` . For regular container jobs, this value is `EKS` , `ECS` , or `ECS_FARGATE` depending on the compute environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobqueuetype
	//
	JobQueueType *string `field:"optional" json:"jobQueueType" yaml:"jobQueueType"`
	// The set of actions that AWS Batch perform on jobs that remain at the head of the job queue in the specified state longer than specified times.
	//
	// AWS Batch will perform each action after `maxTimeSeconds` has passed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobstatetimelimitactions
	//
	JobStateTimeLimitActions interface{} `field:"optional" json:"jobStateTimeLimitActions" yaml:"jobStateTimeLimitActions"`
	// The priority of the job queue.
	//
	// Job queues with a higher priority (or a higher integer value for the `priority` parameter) are evaluated first when associated with the same compute environment. Priority is determined in descending order. For example, a job queue with a priority value of `10` is given scheduling preference over a job queue with a priority value of `1` . All of the compute environments must be either EC2 ( `EC2` or `SPOT` ) or Fargate ( `FARGATE` or `FARGATE_SPOT` ); EC2 and Fargate compute environments can't be mixed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the scheduling policy.
	//
	// The format is `aws: *Partition* :batch: *Region* : *Account* :scheduling-policy/ *Name*` . For example, `aws:aws:batch:us-west-2:123456789012:scheduling-policy/MySchedulingPolicy` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-schedulingpolicyarn
	//
	SchedulingPolicyArn *string `field:"optional" json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
	// The order of the service environment associated with the job queue.
	//
	// Job queues with a higher priority are evaluated first when associated with the same service environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-serviceenvironmentorder
	//
	ServiceEnvironmentOrder interface{} `field:"optional" json:"serviceEnvironmentOrder" yaml:"serviceEnvironmentOrder"`
	// The state of the job queue.
	//
	// If the job queue state is `ENABLED` , it is able to accept jobs. If the job queue state is `DISABLED` , new jobs can't be added to the queue, but jobs already in the queue can finish.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags that are applied to the job queue.
	//
	// For more information, see [Tagging your AWS Batch resources](https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html) in *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

