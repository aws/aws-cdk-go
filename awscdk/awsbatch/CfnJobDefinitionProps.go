package awsbatch


// Properties for defining a `CfnJobDefinition`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html
//
type CfnJobDefinitionProps struct {
	// The type of job definition.
	//
	// For more information about multi-node parallel jobs, see [Creating a multi-node parallel job definition](https://docs.aws.amazon.com/batch/latest/userguide/multi-node-job-def.html) in the *AWS Batch User Guide* .
	//
	// - If the value is `container` , then one of the following is required: `containerProperties` , `ecsProperties` , or `eksProperties` .
	// - If the value is `multinode` , then `nodeProperties` is required.
	//
	// > If the job is run on Fargate resources, then `multinode` isn't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// An object with properties specific to Amazon ECS-based jobs.
	//
	// When `containerProperties` is used in the job definition, it can't be used in addition to `eksProperties` , `ecsProperties` , or `nodeProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-containerproperties
	//
	ContainerProperties interface{} `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// An object that contains the properties for the Amazon ECS resources of a job.When `ecsProperties` is used in the job definition, it can't be used in addition to `containerProperties` , `eksProperties` , or `nodeProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-ecsproperties
	//
	EcsProperties interface{} `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// An object with properties that are specific to Amazon EKS-based jobs.
	//
	// When `eksProperties` is used in the job definition, it can't be used in addition to `containerProperties` , `ecsProperties` , or `nodeProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-eksproperties
	//
	EksProperties interface{} `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// The name of the job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-jobdefinitionname
	//
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// An object with properties that are specific to multi-node parallel jobs.
	//
	// When `nodeProperties` is used in the job definition, it can't be used in addition to `containerProperties` , `ecsProperties` , or `eksProperties` .
	//
	// > If the job runs on Fargate resources, don't specify `nodeProperties` . Use `containerProperties` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-nodeproperties
	//
	NodeProperties interface{} `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Default parameters or parameter substitution placeholders that are set in the job definition.
	//
	// Parameters are specified as a key-value pair mapping. Parameters in a `SubmitJob` request override any corresponding parameter defaults from the job definition. For more information about specifying parameters, see [Job definition parameters](https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The platform capabilities required by the job definition.
	//
	// If no value is specified, it defaults to `EC2` . Jobs run on Fargate resources specify `FARGATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-platformcapabilities
	//
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the tasks when the tasks are created. For tags with the same name, job tags are given priority over job definitions tags. If the total number of combined tags from the job and job definition is over 50, the job is moved to the `FAILED` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-propagatetags
	//
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The retry strategy to use for failed jobs that are submitted with this job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-retrystrategy
	//
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// The scheduling priority of the job definition.
	//
	// This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-schedulingpriority
	//
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The tags that are applied to the job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes, AWS Batch terminates your jobs if they aren't finished.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

