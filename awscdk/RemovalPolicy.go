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
//   var myRole role
//
//   cr.NewAwsCustomResource(this, jsii.String("Customized"), &AwsCustomResourceProps{
//   	Role: myRole,
//   	 // must be assumable by the `lambda.amazonaws.com` service principal
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	 // defaults to 2 minutes
//   	LogGroup: logs.NewLogGroup(this, jsii.String("AwsCustomResourceLogs"), &LogGroupProps{
//   		Retention: logs.RetentionDays_ONE_DAY,
//   	}),
//   	FunctionName: jsii.String("my-custom-name"),
//   	 // defaults to a CloudFormation generated name
//   	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	 // defaults to `RemovalPolicy.DESTROY`
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type RemovalPolicy string

const (
	// This is the default removal policy.
	//
	// It means that when the resource is
	// removed from the app, it will be physically destroyed.
	RemovalPolicy_DESTROY RemovalPolicy = "DESTROY"
	// This uses the 'Retain' DeletionPolicy, which will cause the resource to be retained in the account, but orphaned from the stack.
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

