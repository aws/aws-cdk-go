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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	DeletionProtection: jsii.Boolean(true),
//   	// only the replica in us-east-1 will be deleted during stack deletion
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			DeletionProtection: jsii.Boolean(false),
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			DeletionProtection: jsii.Boolean(true),
//   		},
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

