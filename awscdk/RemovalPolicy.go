package awscdk


// Possible values for a resource's Removal Policy.
//
// The removal policy controls what happens to the resource if it stops being
// managed by CloudFormation. This can happen in one of three situations:
//
// - The resource is removed from the template, so CloudFormation stops managing it;
// - A change to the resource is made that requires it to be replaced, so CloudFormation stops
//   managing it;
// - The stack is deleted, so CloudFormation stops managing all resources in it.
//
// The Removal Policy applies to all above cases.
//
// Many stateful resources in the AWS Construct Library will accept a
// `removalPolicy` as a property, typically defaulting it to `RETAIN`.
//
// If the AWS Construct Library resource does not accept a `removalPolicy`
// argument, you can always configure it by using the escape hatch mechanism,
// as shown in the following example:
//
// ```ts
// declare const bucket: s3.Bucket;
//
// const cfnBucket = bucket.node.findChild('Resource') as CfnResource;
// cfnBucket.applyRemovalPolicy(RemovalPolicy.DESTROY);
// ```.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("memoryBucket"), &BucketProps{
//   	BucketName: jsii.String("test-memory"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   topic := sns.NewTopic(this, jsii.String("topic"))
//
//   // Create a custom semantic memory strategy
//   selfManagedStrategy := agentcore.MemoryStrategy_UsingSelfManaged(&SelfManagedStrategyProps{
//   	Name: jsii.String("selfManagedStrategy"),
//   	Description: jsii.String("self managed memory strategy"),
//   	HistoricalContextWindowSize: jsii.Number(5),
//   	InvocationConfiguration: &InvocationConfiguration{
//   		Topic: topic,
//   		S3Location: &Location{
//   			BucketName: bucket.BucketName,
//   			ObjectKey: jsii.String("memory/"),
//   		},
//   	},
//   	TriggerConditions: &TriggerConditions{
//   		MessageBasedTrigger: jsii.Number(1),
//   		TimeBasedTrigger: cdk.Duration_Seconds(jsii.Number(10)),
//   		TokenBasedTrigger: jsii.Number(100),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		selfManagedStrategy,
//   	},
//   })
//
type RemovalPolicy string

const (
	// When this removal policy is applied, the resource will be physically destroyed when it is removed from the stack or when the stack is deleted.
	RemovalPolicy_DESTROY RemovalPolicy = "DESTROY"
	// This uses the 'Retain' DeletionPolicy, which will cause the resource to be retained in the account, but orphaned from the stack.
	//
	// Most resources default to this removal policy.
	RemovalPolicy_RETAIN RemovalPolicy = "RETAIN"
	// This retention policy deletes the resource, but saves a snapshot of its data before deleting, so that it can be re-created later.
	//
	// Only available for some stateful resources,
	// like databases, EC2 volumes, etc.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	RemovalPolicy_SNAPSHOT RemovalPolicy = "SNAPSHOT"
	// Resource will be retained when they are requested to be deleted during a stack delete request or need to be replaced due to a stack update request.
	//
	// Resource are not retained, if the creation is rolled back.
	//
	// The result is that new, empty, and unused resources are deleted,
	// while in-use resources and their data are retained.
	//
	// This uses the 'RetainExceptOnCreate' DeletionPolicy,
	// and the 'Retain' UpdateReplacePolicy, when `applyToUpdateReplacePolicy` is set.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	RemovalPolicy_RETAIN_ON_UPDATE_OR_DELETE RemovalPolicy = "RETAIN_ON_UPDATE_OR_DELETE"
)

