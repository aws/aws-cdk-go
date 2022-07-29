package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Logs::Destination`.
//
// The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestination := awscdk.Aws_logs.NewCfnDestination(this, jsii.String("MyCfnDestination"), &cfnDestinationProps{
//   	destinationName: jsii.String("destinationName"),
//   	destinationPolicy: jsii.String("destinationPolicy"),
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//   })
//
type CfnDestination interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the CloudWatch Logs destination, such as `arn:aws:logs:us-west-1:123456789012:destination:MyDestination` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The name of the destination.
	DestinationName() *string
	SetDestinationName(val *string)
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	DestinationPolicy() *string
	SetDestinationPolicy(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The Amazon Resource Name (ARN) of the physical target where the log events are delivered (for example, a Kinesis stream).
	TargetArn() *string
	SetTargetArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDestination
type jsiiProxy_CfnDestination struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDestination) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) DestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) DestinationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::Destination`.
func NewCfnDestination(scope awscdk.Construct, id *string, props *CfnDestinationProps) CfnDestination {
	_init_.Initialize()

	j := jsiiProxy_CfnDestination{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::Destination`.
func NewCfnDestination_Override(c CfnDestination, scope awscdk.Construct, id *string, props *CfnDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDestination) SetDestinationName(val *string) {
	_jsii_.Set(
		j,
		"destinationName",
		val,
	)
}

func (j *jsiiProxy_CfnDestination) SetDestinationPolicy(val *string) {
	_jsii_.Set(
		j,
		"destinationPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnDestination) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDestination) SetTargetArn(val *string) {
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDestination_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnDestination",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDestination) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDestination) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDestination) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDestination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDestination) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDestination) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDestination) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDestination) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDestination) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDestination) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDestination) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestinationProps := &cfnDestinationProps{
//   	destinationName: jsii.String("destinationName"),
//   	destinationPolicy: jsii.String("destinationPolicy"),
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnDestinationProps struct {
	// The name of the destination.
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	DestinationPolicy *string `field:"required" json:"destinationPolicy" yaml:"destinationPolicy"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the physical target where the log events are delivered (for example, a Kinesis stream).
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
}

// A CloudFormation `AWS::Logs::LogGroup`.
//
// The `AWS::Logs::LogGroup` resource specifies a log group. A log group defines common properties for log streams, such as their retention and access control rules. Each log stream must belong to one log group.
//
// You can create up to 1,000,000 log groups per Region per account. You must use the following guidelines when naming a log group:
//
// - Log group names must be unique within a Region for an AWS account.
// - Log group names can be between 1 and 512 characters long.
// - Log group names consist of the following characters: a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.' (period).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogGroup := awscdk.Aws_logs.NewCfnLogGroup(this, jsii.String("MyCfnLogGroup"), &cfnLogGroupProps{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logGroupName: jsii.String("logGroupName"),
//   	retentionInDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnLogGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the log group, such as `arn:aws:logs:us-west-1:123456789012:log-group:/mystack-testgroup-12ABC1AB12A1:*`.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the AWS KMS key to use when encrypting log data.
	//
	// To associate an AWS KMS key with the log group, specify the ARN of that KMS key here. If you do so, ingested data is encrypted using this key. This association is stored as long as the data encrypted with the KMS key is still within CloudWatch Logs . This enables CloudWatch Logs to decrypt this data whenever it is requested.
	//
	// If you attempt to associate a KMS key with the log group but the KMS key doesn't exist or is deactivated, you will receive an `InvalidParameterException` error.
	//
	// Log group data is always encrypted in CloudWatch Logs . If you omit this key, the encryption does not use AWS KMS . For more information, see [Encrypt log data in CloudWatch Logs using AWS Key Management Service](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/encrypt-log-data-kms.html)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The name of the log group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	LogGroupName() *string
	SetLogGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The number of days to retain the log events in the specified log group.
	//
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, and 3653.
	//
	// To set a log group to never have log events expire, use [DeleteRetentionPolicy](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html) .
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to the log group.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLogGroup
type jsiiProxy_CfnLogGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLogGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::LogGroup`.
func NewCfnLogGroup(scope awscdk.Construct, id *string, props *CfnLogGroupProps) CfnLogGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnLogGroup{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnLogGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::LogGroup`.
func NewCfnLogGroup_Override(c CfnLogGroup, scope awscdk.Construct, id *string, props *CfnLogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnLogGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLogGroup) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnLogGroup) SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnLogGroup) SetRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnLogGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLogGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLogGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnLogGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLogGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLogGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLogGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLogGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLogGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLogGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLogGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLogGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLogGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLogGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLogGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLogGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLogGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogGroupProps := &cfnLogGroupProps{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logGroupName: jsii.String("logGroupName"),
//   	retentionInDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLogGroupProps struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key to use when encrypting log data.
	//
	// To associate an AWS KMS key with the log group, specify the ARN of that KMS key here. If you do so, ingested data is encrypted using this key. This association is stored as long as the data encrypted with the KMS key is still within CloudWatch Logs . This enables CloudWatch Logs to decrypt this data whenever it is requested.
	//
	// If you attempt to associate a KMS key with the log group but the KMS key doesn't exist or is deactivated, you will receive an `InvalidParameterException` error.
	//
	// Log group data is always encrypted in CloudWatch Logs . If you omit this key, the encryption does not use AWS KMS . For more information, see [Encrypt log data in CloudWatch Logs using AWS Key Management Service](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/encrypt-log-data-kms.html)
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the log group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The number of days to retain the log events in the specified log group.
	//
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, and 3653.
	//
	// To set a log group to never have log events expire, use [DeleteRetentionPolicy](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html) .
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// An array of key-value pairs to apply to the log group.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Logs::LogStream`.
//
// The `AWS::Logs::LogStream` resource specifies an Amazon CloudWatch Logs log stream in a specific log group. A log stream represents the sequence of events coming from an application instance or resource that you are monitoring.
//
// There is no limit on the number of log streams that you can create for a log group.
//
// You must use the following guidelines when naming a log stream:
//
// - Log stream names must be unique within the log group.
// - Log stream names can be between 1 and 512 characters long.
// - The ':' (colon) and '*' (asterisk) characters are not allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogStream := awscdk.Aws_logs.NewCfnLogStream(this, jsii.String("MyCfnLogStream"), &cfnLogStreamProps{
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   })
//
type CfnLogStream interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The name of the log group where the log stream is created.
	LogGroupName() *string
	SetLogGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the log stream.
	//
	// The name must be unique within the log group.
	LogStreamName() *string
	SetLogStreamName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLogStream
type jsiiProxy_CfnLogStream struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLogStream) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogStream) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::LogStream`.
func NewCfnLogStream(scope awscdk.Construct, id *string, props *CfnLogStreamProps) CfnLogStream {
	_init_.Initialize()

	j := jsiiProxy_CfnLogStream{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnLogStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::LogStream`.
func NewCfnLogStream_Override(c CfnLogStream, scope awscdk.Construct, id *string, props *CfnLogStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnLogStream",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLogStream) SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnLogStream) SetLogStreamName(val *string) {
	_jsii_.Set(
		j,
		"logStreamName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnLogStream_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogStream",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLogStream_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogStream",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLogStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnLogStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogStream_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnLogStream",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogStream) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLogStream) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogStream) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLogStream) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLogStream) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLogStream) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLogStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLogStream) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLogStream) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLogStream) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLogStream) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLogStream) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLogStream) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLogStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogStream) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLogStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogStreamProps := &cfnLogStreamProps{
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   }
//
type CfnLogStreamProps struct {
	// The name of the log group where the log stream is created.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log stream.
	//
	// The name must be unique within the log group.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

// A CloudFormation `AWS::Logs::MetricFilter`.
//
// The `AWS::Logs::MetricFilter` resource specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics. If you have multiple metric filters that are associated with a log group, all the filters are applied to the log streams in that group.
//
// The maximum number of metric filters that can be associated with a log group is 100.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricFilter := awscdk.Aws_logs.NewCfnMetricFilter(this, jsii.String("MyCfnMetricFilter"), &cfnMetricFilterProps{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   	metricTransformations: []interface{}{
//   		&metricTransformationProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricValue: jsii.String("metricValue"),
//
//   			// the properties below are optional
//   			defaultValue: jsii.Number(123),
//   			dimensions: []interface{}{
//   				&dimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	filterName: jsii.String("filterName"),
//   })
//
type CfnMetricFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::Logs::MetricFilter.FilterName`.
	FilterName() *string
	SetFilterName(val *string)
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern() *string
	SetFilterPattern(val *string)
	// The name of an existing log group that you want to associate with this metric filter.
	LogGroupName() *string
	SetLogGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The metric transformations.
	MetricTransformations() interface{}
	SetMetricTransformations(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMetricFilter
type jsiiProxy_CfnMetricFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMetricFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) FilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) FilterPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) MetricTransformations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricTransformations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::MetricFilter`.
func NewCfnMetricFilter(scope awscdk.Construct, id *string, props *CfnMetricFilterProps) CfnMetricFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnMetricFilter{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnMetricFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::MetricFilter`.
func NewCfnMetricFilter_Override(c CfnMetricFilter, scope awscdk.Construct, id *string, props *CfnMetricFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnMetricFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMetricFilter) SetFilterName(val *string) {
	_jsii_.Set(
		j,
		"filterName",
		val,
	)
}

func (j *jsiiProxy_CfnMetricFilter) SetFilterPattern(val *string) {
	_jsii_.Set(
		j,
		"filterPattern",
		val,
	)
}

func (j *jsiiProxy_CfnMetricFilter) SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnMetricFilter) SetMetricTransformations(val interface{}) {
	_jsii_.Set(
		j,
		"metricTransformations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnMetricFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnMetricFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMetricFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnMetricFilter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMetricFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnMetricFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnMetricFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricFilter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMetricFilter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMetricFilter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMetricFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMetricFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMetricFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMetricFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMetricFilter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMetricFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMetricFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMetricFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMetricFilter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMetricFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMetricFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMetricFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionProperty := &dimensionProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnMetricFilter_DimensionProperty struct {
	// `CfnMetricFilter.DimensionProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnMetricFilter.DimensionProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// `MetricTransformation` is a property of the `AWS::Logs::MetricFilter` resource that describes how to transform log streams into a CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricTransformationProperty := &metricTransformationProperty{
//   	metricName: jsii.String("metricName"),
//   	metricNamespace: jsii.String("metricNamespace"),
//   	metricValue: jsii.String("metricValue"),
//
//   	// the properties below are optional
//   	defaultValue: jsii.Number(123),
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   }
//
type CfnMetricFilter_MetricTransformationProperty struct {
	// The name of the CloudWatch metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A custom namespace to contain your metric in CloudWatch.
	//
	// Use namespaces to group together metrics that are similar. For more information, see [Namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace) .
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value that is published to the CloudWatch metric.
	//
	// For example, if you're counting the occurrences of a particular term like `Error` , specify 1 for the metric value. If you're counting the number of bytes transferred, reference the value that is in the log event by using $ followed by the name of the field that you specified in the filter pattern, such as `$.size` .
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// (Optional) The value to emit when a filter pattern does not match a log event.
	//
	// This value can be null.
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnMetricFilter.MetricTransformationProperty.Dimensions`.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// `CfnMetricFilter.MetricTransformationProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

// Properties for defining a `CfnMetricFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricFilterProps := &cfnMetricFilterProps{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   	metricTransformations: []interface{}{
//   		&metricTransformationProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricValue: jsii.String("metricValue"),
//
//   			// the properties below are optional
//   			defaultValue: jsii.Number(123),
//   			dimensions: []interface{}{
//   				&dimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	filterName: jsii.String("filterName"),
//   }
//
type CfnMetricFilterProps struct {
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of an existing log group that you want to associate with this metric filter.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The metric transformations.
	MetricTransformations interface{} `field:"required" json:"metricTransformations" yaml:"metricTransformations"`
	// `AWS::Logs::MetricFilter.FilterName`.
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

// A CloudFormation `AWS::Logs::QueryDefinition`.
//
// Creates a query definition for CloudWatch Logs Insights. For more information, see [Analyzing Log Data with CloudWatch Logs Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueryDefinition := awscdk.Aws_logs.NewCfnQueryDefinition(this, jsii.String("MyCfnQueryDefinition"), &cfnQueryDefinitionProps{
//   	name: jsii.String("name"),
//   	queryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	logGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   })
//
type CfnQueryDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the query definition.
	AttrQueryDefinitionId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Use this parameter if you want the query to query only certain log groups.
	LogGroupNames() *[]*string
	SetLogGroupNames(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// A name for the query definition.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The query string to use for this query definition.
	//
	// For more information, see [CloudWatch Logs Insights Query Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html) .
	QueryString() *string
	SetQueryString(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnQueryDefinition
type jsiiProxy_CfnQueryDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnQueryDefinition) AttrQueryDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrQueryDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) LogGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::QueryDefinition`.
func NewCfnQueryDefinition(scope awscdk.Construct, id *string, props *CfnQueryDefinitionProps) CfnQueryDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnQueryDefinition{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnQueryDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::QueryDefinition`.
func NewCfnQueryDefinition_Override(c CfnQueryDefinition, scope awscdk.Construct, id *string, props *CfnQueryDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnQueryDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnQueryDefinition) SetLogGroupNames(val *[]*string) {
	_jsii_.Set(
		j,
		"logGroupNames",
		val,
	)
}

func (j *jsiiProxy_CfnQueryDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnQueryDefinition) SetQueryString(val *string) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnQueryDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnQueryDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnQueryDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnQueryDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnQueryDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnQueryDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueryDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnQueryDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnQueryDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnQueryDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnQueryDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueryDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnQueryDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueryDefinitionProps := &cfnQueryDefinitionProps{
//   	name: jsii.String("name"),
//   	queryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	logGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   }
//
type CfnQueryDefinitionProps struct {
	// A name for the query definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query string to use for this query definition.
	//
	// For more information, see [CloudWatch Logs Insights Query Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html) .
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Use this parameter if you want the query to query only certain log groups.
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

// A CloudFormation `AWS::Logs::ResourcePolicy`.
//
// Creates or updates a resource policy that allows other AWS services to put log events to this account. An account can have up to 10 resource policies per AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourcePolicy := awscdk.Aws_logs.NewCfnResourcePolicy(this, jsii.String("MyCfnResourcePolicy"), &cfnResourcePolicyProps{
//   	policyDocument: jsii.String("policyDocument"),
//   	policyName: jsii.String("policyName"),
//   })
//
type CfnResourcePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The details of the policy.
	//
	// It must be formatted in JSON, and you must use backslashes to escape characters that need to be escaped in JSON strings, such as double quote marks.
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	// The name of the resource policy.
	PolicyName() *string
	SetPolicyName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResourcePolicy
type jsiiProxy_CfnResourcePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourcePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::ResourcePolicy`.
func NewCfnResourcePolicy(scope awscdk.Construct, id *string, props *CfnResourcePolicyProps) CfnResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResourcePolicy{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::ResourcePolicy`.
func NewCfnResourcePolicy_Override(c CfnResourcePolicy, scope awscdk.Construct, id *string, props *CfnResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetPolicyDocument(val *string) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnResourcePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnResourcePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResourcePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnResourcePolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnResourcePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourcePolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourcePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourcePolicyProps := &cfnResourcePolicyProps{
//   	policyDocument: jsii.String("policyDocument"),
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnResourcePolicyProps struct {
	// The details of the policy.
	//
	// It must be formatted in JSON, and you must use backslashes to escape characters that need to be escaped in JSON strings, such as double quote marks.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the resource policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

// A CloudFormation `AWS::Logs::SubscriptionFilter`.
//
// The `AWS::Logs::SubscriptionFilter` resource specifies a subscription filter and associates it with the specified log group. Subscription filters allow you to subscribe to a real-time stream of log events and have them delivered to a specific destination. Currently, the supported destinations are:
//
// - An Amazon Kinesis data stream belonging to the same account as the subscription filter, for same-account delivery.
// - A logical destination that belongs to a different account, for cross-account delivery.
// - An Amazon Kinesis Firehose delivery stream that belongs to the same account as the subscription filter, for same-account delivery.
// - An AWS Lambda function that belongs to the same account as the subscription filter, for same-account delivery.
//
// There can as many as two subscription filters associated with a log group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionFilter := awscdk.Aws_logs.NewCfnSubscriptionFilter(this, jsii.String("MyCfnSubscriptionFilter"), &cfnSubscriptionFilterProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnSubscriptionFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn() *string
	SetDestinationArn(val *string)
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern() *string
	SetFilterPattern(val *string)
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	LogGroupName() *string
	SetLogGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSubscriptionFilter
type jsiiProxy_CfnSubscriptionFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSubscriptionFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) FilterPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Logs::SubscriptionFilter`.
func NewCfnSubscriptionFilter(scope awscdk.Construct, id *string, props *CfnSubscriptionFilterProps) CfnSubscriptionFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnSubscriptionFilter{}

	_jsii_.Create(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Logs::SubscriptionFilter`.
func NewCfnSubscriptionFilter_Override(c CfnSubscriptionFilter, scope awscdk.Construct, id *string, props *CfnSubscriptionFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSubscriptionFilter) SetDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationArn",
		val,
	)
}

func (j *jsiiProxy_CfnSubscriptionFilter) SetFilterPattern(val *string) {
	_jsii_.Set(
		j,
		"filterPattern",
		val,
	)
}

func (j *jsiiProxy_CfnSubscriptionFilter) SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnSubscriptionFilter) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnSubscriptionFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSubscriptionFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSubscriptionFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_logs.CfnSubscriptionFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSubscriptionFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSubscriptionFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionFilterProps := &cfnSubscriptionFilterProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnSubscriptionFilterProps struct {
	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnRestriction := &columnRestriction{
//   	comparison: jsii.String("comparison"),
//
//   	// the properties below are optional
//   	numberValue: jsii.Number(123),
//   	stringValue: jsii.String("stringValue"),
//   }
//
// Experimental.
type ColumnRestriction struct {
	// Comparison operator to use.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Number value to compare to.
	//
	// Exactly one of 'stringValue' and 'numberValue' must be set.
	// Experimental.
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// String value to compare to.
	//
	// Exactly one of 'stringValue' and 'numberValue' must be set.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// A new CloudWatch Logs Destination for use in cross-account scenarios.
//
// CrossAccountDestinations are used to subscribe a Kinesis stream in a
// different account to a CloudWatch Subscription.
//
// Consumers will hardly ever need to use this class. Instead, directly
// subscribe a Kinesis stream using the integration class in the
// `@aws-cdk/aws-logs-destinations` package; if necessary, a
// `CrossAccountDestination` will be created automatically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   crossAccountDestination := awscdk.Aws_logs.NewCrossAccountDestination(this, jsii.String("MyCrossAccountDestination"), &crossAccountDestinationProps{
//   	role: role,
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	destinationName: jsii.String("destinationName"),
//   })
//
// Experimental.
type CrossAccountDestination interface {
	awscdk.Resource
	ILogSubscriptionDestination
	// The ARN of this CrossAccountDestination object.
	// Experimental.
	DestinationArn() *string
	// The name of this CrossAccountDestination object.
	// Experimental.
	DestinationName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Policy object of this CrossAccountDestination object.
	// Experimental.
	PolicyDocument() awsiam.PolicyDocument
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Experimental.
	AddToPolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	// Experimental.
	Bind(_scope awscdk.Construct, _sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for CrossAccountDestination
type jsiiProxy_CrossAccountDestination struct {
	internal.Type__awscdkResource
	jsiiProxy_ILogSubscriptionDestination
}

func (j *jsiiProxy_CrossAccountDestination) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) DestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCrossAccountDestination(scope constructs.Construct, id *string, props *CrossAccountDestinationProps) CrossAccountDestination {
	_init_.Initialize()

	j := jsiiProxy_CrossAccountDestination{}

	_jsii_.Create(
		"monocdk.aws_logs.CrossAccountDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCrossAccountDestination_Override(c CrossAccountDestination, scope constructs.Construct, id *string, props *CrossAccountDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.CrossAccountDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CrossAccountDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CrossAccountDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CrossAccountDestination_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.CrossAccountDestination",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) AddToPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		c,
		"addToPolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CrossAccountDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CrossAccountDestination) Bind(_scope awscdk.Construct, _sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig {
	var returns *LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CrossAccountDestination) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CrossAccountDestination) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CrossAccountDestination) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CrossAccountDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a CrossAccountDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   crossAccountDestinationProps := &crossAccountDestinationProps{
//   	role: role,
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	destinationName: jsii.String("destinationName"),
//   }
//
// Experimental.
type CrossAccountDestinationProps struct {
	// The role to assume that grants permissions to write to 'target'.
	//
	// The role must be assumable by 'logs.{REGION}.amazonaws.com'.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// The log destination target's ARN.
	// Experimental.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The name of the log destination.
	// Experimental.
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
}

// A collection of static methods to generate appropriate ILogPatterns.
//
// Example:
//   // Search for lines that contain both "ERROR" and "MainThread"
//   pattern1 := logs.filterPattern.allTerms(jsii.String("ERROR"), jsii.String("MainThread"))
//
//   // Search for lines that either contain both "ERROR" and "MainThread", or
//   // both "WARN" and "Deadlock".
//   pattern2 := logs.filterPattern.anyTermGroup([]*string{
//   	jsii.String("ERROR"),
//   	jsii.String("MainThread"),
//   }, []*string{
//   	jsii.String("WARN"),
//   	jsii.String("Deadlock"),
//   })
//
// Experimental.
type FilterPattern interface {
}

// The jsii proxy struct for FilterPattern
type jsiiProxy_FilterPattern struct {
	_ byte // padding
}

// Experimental.
func NewFilterPattern() FilterPattern {
	_init_.Initialize()

	j := jsiiProxy_FilterPattern{}

	_jsii_.Create(
		"monocdk.aws_logs.FilterPattern",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFilterPattern_Override(f FilterPattern) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.FilterPattern",
		nil, // no parameters
		f,
	)
}

// A JSON log pattern that matches if all given JSON log patterns match.
// Experimental.
func FilterPattern_All(patterns ...JsonPattern) JsonPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"all",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches all events.
// Experimental.
func FilterPattern_AllEvents() IFilterPattern {
	_init_.Initialize()

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"allEvents",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A log pattern that matches if all the strings given appear in the event.
// Experimental.
func FilterPattern_AllTerms(terms ...*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range terms {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"allTerms",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if any of the given JSON log patterns match.
// Experimental.
func FilterPattern_Any(patterns ...JsonPattern) JsonPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"any",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches if any of the strings given appear in the event.
// Experimental.
func FilterPattern_AnyTerm(terms ...*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range terms {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"anyTerm",
		args,
		&returns,
	)

	return returns
}

// A log pattern that matches if any of the given term groups matches the event.
//
// A term group matches an event if all the terms in it appear in the event string.
// Experimental.
func FilterPattern_AnyTermGroup(termGroups ...*[]*string) IFilterPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range termGroups {
		args = append(args, a)
	}

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"anyTermGroup",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field exists and equals the boolean value.
// Experimental.
func FilterPattern_BooleanValue(jsonField *string, value *bool) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"booleanValue",
		[]interface{}{jsonField, value},
		&returns,
	)

	return returns
}

// A JSON log patter that matches if the field exists.
//
// This is a readable convenience wrapper over 'field = *'.
// Experimental.
func FilterPattern_Exists(jsonField *string) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"exists",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field exists and has the special value 'null'.
// Experimental.
func FilterPattern_IsNull(jsonField *string) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"isNull",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// Use the given string as log pattern.
//
// See https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
// for information on writing log patterns.
// Experimental.
func FilterPattern_Literal(logPatternString *string) IFilterPattern {
	_init_.Initialize()

	var returns IFilterPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"literal",
		[]interface{}{logPatternString},
		&returns,
	)

	return returns
}

// A JSON log pattern that matches if the field does not exist.
// Experimental.
func FilterPattern_NotExists(jsonField *string) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"notExists",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// A JSON log pattern that compares numerical values.
//
// This pattern only matches if the event is a JSON event, and the indicated field inside
// compares with the value in the indicated way.
//
// Use '$' to indicate the root of the JSON structure. The comparison operator can only
// compare equality or inequality. The '*' wildcard may appear in the value may at the
// start or at the end.
//
// For more information, see:
//
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
// Experimental.
func FilterPattern_NumberValue(jsonField *string, comparison *string, value *float64) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"numberValue",
		[]interface{}{jsonField, comparison, value},
		&returns,
	)

	return returns
}

// A space delimited log pattern matcher.
//
// The log event is divided into space-delimited columns (optionally
// enclosed by "" or [] to capture spaces into column values), and names
// are given to each column.
//
// '...' may be specified once to match any number of columns.
//
// Afterwards, conditions may be added to individual columns.
// Experimental.
func FilterPattern_SpaceDelimited(columns ...*string) SpaceDelimitedTextPattern {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range columns {
		args = append(args, a)
	}

	var returns SpaceDelimitedTextPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"spaceDelimited",
		args,
		&returns,
	)

	return returns
}

// A JSON log pattern that compares string values.
//
// This pattern only matches if the event is a JSON event, and the indicated field inside
// compares with the string value.
//
// Use '$' to indicate the root of the JSON structure. The comparison operator can only
// compare equality or inequality. The '*' wildcard may appear in the value may at the
// start or at the end.
//
// For more information, see:
//
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
// Experimental.
func FilterPattern_StringValue(jsonField *string, comparison *string, value *string) JsonPattern {
	_init_.Initialize()

	var returns JsonPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.FilterPattern",
		"stringValue",
		[]interface{}{jsonField, comparison, value},
		&returns,
	)

	return returns
}

// Interface for objects that can render themselves to log patterns.
// Experimental.
type IFilterPattern interface {
	// Experimental.
	LogPatternString() *string
}

// The jsii proxy for IFilterPattern
type jsiiProxy_IFilterPattern struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilterPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}

// Experimental.
type ILogGroup interface {
	awsiam.IResourceWithPolicy
	// Create a new Metric Filter on this Log Group.
	// Experimental.
	AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter
	// Create a new Log Stream for this Log Group.
	// Experimental.
	AddStream(id *string, props *StreamOptions) LogStream
	// Create a new Subscription Filter on this Log Group.
	// Experimental.
	AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter
	// Extract a metric from structured log events in the LogGroup.
	//
	// Creates a MetricFilter on this LogGroup that will extract the value
	// of the indicated JSON field in all records where it occurs.
	//
	// The metric will be available in CloudWatch Metrics under the
	// indicated namespace and name.
	//
	// Returns: A Metric object representing the extracted metric.
	// Experimental.
	ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric
	// Give the indicated permissions on this log group and all streams.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Give permissions to write to create and write to streams in this log group.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Public method to get the physical name of this log group.
	// Experimental.
	LogGroupPhysicalName() *string
	// The ARN of this log group, with ':*' appended.
	// Experimental.
	LogGroupArn() *string
	// The name of this log group.
	// Experimental.
	LogGroupName() *string
}

// The jsii proxy for ILogGroup
type jsiiProxy_ILogGroup struct {
	internal.Type__awsiamIResourceWithPolicy
}

func (i *jsiiProxy_ILogGroup) AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter {
	var returns MetricFilter

	_jsii_.Invoke(
		i,
		"addMetricFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) AddStream(id *string, props *StreamOptions) LogStream {
	var returns LogStream

	_jsii_.Invoke(
		i,
		"addStream",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter {
	var returns SubscriptionFilter

	_jsii_.Invoke(
		i,
		"addSubscriptionFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"extractMetric",
		[]interface{}{jsonField, metricNamespace, metricName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) LogGroupPhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"logGroupPhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILogGroup) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogGroup) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

// Experimental.
type ILogStream interface {
	awscdk.IResource
	// The name of this log stream.
	// Experimental.
	LogStreamName() *string
}

// The jsii proxy for ILogStream
type jsiiProxy_ILogStream struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILogStream) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

// Interface for classes that can be the destination of a log Subscription.
// Experimental.
type ILogSubscriptionDestination interface {
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	// Experimental.
	Bind(scope awscdk.Construct, sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig
}

// The jsii proxy for ILogSubscriptionDestination
type jsiiProxy_ILogSubscriptionDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogSubscriptionDestination) Bind(scope awscdk.Construct, sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig {
	var returns *LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, sourceLogGroup},
		&returns,
	)

	return returns
}

// Base class for patterns that only match JSON log events.
//
// Example:
//   // Search for all events where the component field is equal to
//   // "HttpServer" and either error is true or the latency is higher
//   // than 1000.
//   pattern := logs.filterPattern.all(logs.filterPattern.stringValue(jsii.String("$.component"), jsii.String("="), jsii.String("HttpServer")), logs.filterPattern.any(logs.filterPattern.booleanValue(jsii.String("$.error"), jsii.Boolean(true)), logs.filterPattern.numberValue(jsii.String("$.latency"), jsii.String(">"), jsii.Number(1000))))
//
// Experimental.
type JsonPattern interface {
	IFilterPattern
	// Experimental.
	JsonPatternString() *string
	// Experimental.
	LogPatternString() *string
}

// The jsii proxy struct for JsonPattern
type jsiiProxy_JsonPattern struct {
	jsiiProxy_IFilterPattern
}

func (j *jsiiProxy_JsonPattern) JsonPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPatternString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}


// Experimental.
func NewJsonPattern_Override(j JsonPattern, jsonPatternString *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.JsonPattern",
		[]interface{}{jsonPatternString},
		j,
	)
}

// Define a CloudWatch Log Group.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
//   logBucket := s3.NewBucket(this, jsii.String("S3 Bucket"))
//
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
//   	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
//   	jobDriver: &jobDriver{
//   		sparkSubmitJobDriver: &sparkSubmitJobDriver{
//   			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
//   		},
//   	},
//   	monitoring: &monitoring{
//   		logGroup: logGroup,
//   		logBucket: logBucket,
//   	},
//   })
//
// Experimental.
type LogGroup interface {
	awscdk.Resource
	ILogGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this log group.
	// Experimental.
	LogGroupArn() *string
	// The name of this log group.
	// Experimental.
	LogGroupName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Create a new Metric Filter on this Log Group.
	// Experimental.
	AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter
	// Create a new Log Stream for this Log Group.
	// Experimental.
	AddStream(id *string, props *StreamOptions) LogStream
	// Create a new Subscription Filter on this Log Group.
	// Experimental.
	AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter
	// Adds a statement to the resource policy associated with this log group.
	//
	// A resource policy will be automatically created upon the first call to `addToResourcePolicy`.
	//
	// Any ARN Principals inside of the statement will be converted into AWS Account ID strings
	// because CloudWatch Logs Resource Policies do not accept ARN principals.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Extract a metric from structured log events in the LogGroup.
	//
	// Creates a MetricFilter on this LogGroup that will extract the value
	// of the indicated JSON field in all records where it occurs.
	//
	// The metric will be available in CloudWatch Metrics under the
	// indicated namespace and name.
	//
	// Returns: A Metric object representing the extracted metric.
	// Experimental.
	ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Give the indicated permissions on this log group and all streams.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Give permissions to create and write to streams in this log group.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Public method to get the physical name of this log group.
	//
	// Returns: Physical name of log group.
	// Experimental.
	LogGroupPhysicalName() *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for LogGroup
type jsiiProxy_LogGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ILogGroup
}

func (j *jsiiProxy_LogGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogGroup(scope constructs.Construct, id *string, props *LogGroupProps) LogGroup {
	_init_.Initialize()

	j := jsiiProxy_LogGroup{}

	_jsii_.Create(
		"monocdk.aws_logs.LogGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogGroup_Override(l LogGroup, scope constructs.Construct, id *string, props *LogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.LogGroup",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing LogGroup given its ARN.
// Experimental.
func LogGroup_FromLogGroupArn(scope constructs.Construct, id *string, logGroupArn *string) ILogGroup {
	_init_.Initialize()

	var returns ILogGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogGroup",
		"fromLogGroupArn",
		[]interface{}{scope, id, logGroupArn},
		&returns,
	)

	return returns
}

// Import an existing LogGroup given its name.
// Experimental.
func LogGroup_FromLogGroupName(scope constructs.Construct, id *string, logGroupName *string) ILogGroup {
	_init_.Initialize()

	var returns ILogGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogGroup",
		"fromLogGroupName",
		[]interface{}{scope, id, logGroupName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LogGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LogGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter {
	var returns MetricFilter

	_jsii_.Invoke(
		l,
		"addMetricFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddStream(id *string, props *StreamOptions) LogStream {
	var returns LogStream

	_jsii_.Invoke(
		l,
		"addStream",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter {
	var returns SubscriptionFilter

	_jsii_.Invoke(
		l,
		"addSubscriptionFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		l,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LogGroup) ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"extractMetric",
		[]interface{}{jsonField, metricNamespace, metricName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) LogGroupPhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"logGroupPhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a LogGroup.
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
//   	encryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
//   	encryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: kmsKey,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchLogGroup: logGroup,
//   			cloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			s3Bucket: execBucket,
//   			s3EncryptionEnabled: jsii.Boolean(true),
//   			s3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		logging: ecs.executeCommandLogging_OVERRIDE,
//   	},
//   })
//
// Experimental.
type LogGroupProps struct {
	// The KMS Key to encrypt the log group with.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Name of the log group.
	// Experimental.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Determine the removal policy of this log group.
	//
	// Normally you want to retain the log group so you can diagnose issues
	// from logs even after a deployment that no longer includes the log group.
	// In that case, use the normal date-based retention policy to age out your
	// logs.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// How long, in days, the log contents will be retained.
	//
	// To retain all logs, set this value to RetentionDays.INFINITE.
	// Experimental.
	Retention RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

// Creates a custom resource to control the retention policy of a CloudWatch Logs log group.
//
// The log group is created if it doesn't already exist. The policy
// is removed when `retentionDays` is `undefined` or equal to `Infinity`.
// Log group can be created in the region that is different from stack region by
// specifying `logGroupRegion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var role role
//
//   logRetention := awscdk.Aws_logs.NewLogRetention(this, jsii.String("MyLogRetention"), &logRetentionProps{
//   	logGroupName: jsii.String("logGroupName"),
//   	retention: awscdk.*Aws_logs.retentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	logGroupRegion: jsii.String("logGroupRegion"),
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: duration,
//   		maxRetries: jsii.Number(123),
//   	},
//   	role: role,
//   })
//
// Experimental.
type LogRetention interface {
	awscdk.Construct
	// The ARN of the LogGroup.
	// Experimental.
	LogGroupArn() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for LogRetention
type jsiiProxy_LogRetention struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_LogRetention) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogRetention) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogRetention(scope constructs.Construct, id *string, props *LogRetentionProps) LogRetention {
	_init_.Initialize()

	j := jsiiProxy_LogRetention{}

	_jsii_.Create(
		"monocdk.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogRetention_Override(l LogRetention, scope constructs.Construct, id *string, props *LogRetentionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.LogRetention",
		[]interface{}{scope, id, props},
		l,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func LogRetention_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogRetention",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogRetention) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogRetention) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogRetention) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogRetention) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogRetention) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a LogRetention.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var role role
//
//   logRetentionProps := &logRetentionProps{
//   	logGroupName: jsii.String("logGroupName"),
//   	retention: awscdk.Aws_logs.retentionDays_ONE_DAY,
//
//   	// the properties below are optional
//   	logGroupRegion: jsii.String("logGroupRegion"),
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: duration,
//   		maxRetries: jsii.Number(123),
//   	},
//   	role: role,
//   }
//
// Experimental.
type LogRetentionProps struct {
	// The log group name.
	// Experimental.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The number of days log events are kept in CloudWatch Logs.
	// Experimental.
	Retention RetentionDays `field:"required" json:"retention" yaml:"retention"`
	// The region where the log group should be created.
	// Experimental.
	LogGroupRegion *string `field:"optional" json:"logGroupRegion" yaml:"logGroupRegion"`
	// Retry options for all AWS API calls.
	// Experimental.
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// Retry options for all AWS API calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   logRetentionRetryOptions := &logRetentionRetryOptions{
//   	base: duration,
//   	maxRetries: jsii.Number(123),
//   }
//
// Experimental.
type LogRetentionRetryOptions struct {
	// The base duration to use in the exponential backoff for operation retries.
	// Experimental.
	Base awscdk.Duration `field:"optional" json:"base" yaml:"base"`
	// The maximum amount of retries.
	// Experimental.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

// Define a Log Stream in a Log Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup logGroup
//
//   logStream := awscdk.Aws_logs.NewLogStream(this, jsii.String("MyLogStream"), &logStreamProps{
//   	logGroup: logGroup,
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   })
//
// Experimental.
type LogStream interface {
	awscdk.Resource
	ILogStream
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The name of this log stream.
	// Experimental.
	LogStreamName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for LogStream
type jsiiProxy_LogStream struct {
	internal.Type__awscdkResource
	jsiiProxy_ILogStream
}

func (j *jsiiProxy_LogStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogStream) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogStream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogStream(scope constructs.Construct, id *string, props *LogStreamProps) LogStream {
	_init_.Initialize()

	j := jsiiProxy_LogStream{}

	_jsii_.Create(
		"monocdk.aws_logs.LogStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogStream_Override(l LogStream, scope constructs.Construct, id *string, props *LogStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.LogStream",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing LogGroup.
// Experimental.
func LogStream_FromLogStreamName(scope constructs.Construct, id *string, logStreamName *string) ILogStream {
	_init_.Initialize()

	var returns ILogStream

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogStream",
		"fromLogStreamName",
		[]interface{}{scope, id, logStreamName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LogStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LogStream_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.LogStream",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LogStream) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogStream) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogStream) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogStream) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LogStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogStream) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a LogStream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup logGroup
//
//   logStreamProps := &logStreamProps{
//   	logGroup: logGroup,
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   }
//
// Experimental.
type LogStreamProps struct {
	// The log group to create a log stream for.
	// Experimental.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
	// The name of the log stream to create.
	//
	// The name must be unique within the log group.
	// Experimental.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// Determine what happens when the log stream resource is removed from the app.
	//
	// Normally you want to retain the log stream so you can diagnose issues from
	// logs even after a deployment that no longer includes the log stream.
	//
	// The date-based retention policy of your log group will age out the logs
	// after a certain time.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

// Properties returned by a Subscription destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logSubscriptionDestinationConfig := &logSubscriptionDestinationConfig{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	role: role,
//   }
//
// Experimental.
type LogSubscriptionDestinationConfig struct {
	// The ARN of the subscription's destination.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The role to assume to write log events to the destination.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// A filter that extracts information from CloudWatch Logs and emits to CloudWatch Metrics.
//
// Example:
//   awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &metricFilterProps{
//   	logGroup: logGroup,
//   	metricNamespace: jsii.String("MyApp"),
//   	metricName: jsii.String("Latency"),
//   	filterPattern: awscdk.FilterPattern.exists(jsii.String("$.latency")),
//   	metricValue: jsii.String("$.latency"),
//   })
//
// Experimental.
type MetricFilter interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Return the given named metric for this Metric Filter.
	// Experimental.
	Metric(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for MetricFilter
type jsiiProxy_MetricFilter struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_MetricFilter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricFilter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetricFilter(scope constructs.Construct, id *string, props *MetricFilterProps) MetricFilter {
	_init_.Initialize()

	j := jsiiProxy_MetricFilter{}

	_jsii_.Create(
		"monocdk.aws_logs.MetricFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetricFilter_Override(m MetricFilter, scope constructs.Construct, id *string, props *MetricFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.MetricFilter",
		[]interface{}{scope, id, props},
		m,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func MetricFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.MetricFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func MetricFilter_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.MetricFilter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_MetricFilter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) Metric(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metric",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetricFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_MetricFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetricFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_MetricFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a MetricFilter created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterPattern iFilterPattern
//
//   metricFilterOptions := &metricFilterOptions{
//   	filterPattern: filterPattern,
//   	metricName: jsii.String("metricName"),
//   	metricNamespace: jsii.String("metricNamespace"),
//
//   	// the properties below are optional
//   	defaultValue: jsii.Number(123),
//   	metricValue: jsii.String("metricValue"),
//   }
//
// Experimental.
type MetricFilterOptions struct {
	// Pattern to search for log events.
	// Experimental.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of the metric to emit.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to emit.
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value to emit if the pattern does not match a particular event.
	// Experimental.
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The value to emit for the metric.
	//
	// Can either be a literal number (typically "1"), or the name of a field in the structure
	// to take the value from the matched event. If you are using a field value, the field
	// value must have been matched using the pattern.
	//
	// If you want to specify a field from a matched JSON structure, use '$.fieldName',
	// and make sure the field is in the pattern (if only as '$.fieldName = *').
	//
	// If you want to specify a field from a matched space-delimited structure,
	// use '$fieldName'.
	// Experimental.
	MetricValue *string `field:"optional" json:"metricValue" yaml:"metricValue"`
}

// Properties for a MetricFilter.
//
// Example:
//   awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &metricFilterProps{
//   	logGroup: logGroup,
//   	metricNamespace: jsii.String("MyApp"),
//   	metricName: jsii.String("Latency"),
//   	filterPattern: awscdk.FilterPattern.exists(jsii.String("$.latency")),
//   	metricValue: jsii.String("$.latency"),
//   })
//
// Experimental.
type MetricFilterProps struct {
	// Pattern to search for log events.
	// Experimental.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of the metric to emit.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to emit.
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value to emit if the pattern does not match a particular event.
	// Experimental.
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The value to emit for the metric.
	//
	// Can either be a literal number (typically "1"), or the name of a field in the structure
	// to take the value from the matched event. If you are using a field value, the field
	// value must have been matched using the pattern.
	//
	// If you want to specify a field from a matched JSON structure, use '$.fieldName',
	// and make sure the field is in the pattern (if only as '$.fieldName = *').
	//
	// If you want to specify a field from a matched space-delimited structure,
	// use '$fieldName'.
	// Experimental.
	MetricValue *string `field:"optional" json:"metricValue" yaml:"metricValue"`
	// The log group to create the filter on.
	// Experimental.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

// Define a query definition for CloudWatch Logs Insights.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryDefinition interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ID of the query definition.
	// Experimental.
	QueryDefinitionId() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for QueryDefinition
type jsiiProxy_QueryDefinition struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_QueryDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryDefinition) QueryDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewQueryDefinition(scope constructs.Construct, id *string, props *QueryDefinitionProps) QueryDefinition {
	_init_.Initialize()

	j := jsiiProxy_QueryDefinition{}

	_jsii_.Create(
		"monocdk.aws_logs.QueryDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewQueryDefinition_Override(q QueryDefinition, scope constructs.Construct, id *string, props *QueryDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.QueryDefinition",
		[]interface{}{scope, id, props},
		q,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func QueryDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.QueryDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func QueryDefinition_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.QueryDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		q,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (q *jsiiProxy_QueryDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		q,
		"onPrepare",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		q,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (q *jsiiProxy_QueryDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) Prepare() {
	_jsii_.InvokeVoid(
		q,
		"prepare",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QueryDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		q,
		"synthesize",
		[]interface{}{session},
	)
}

func (q *jsiiProxy_QueryDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueryDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a QueryDefinition.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryDefinitionProps struct {
	// Name of the query definition.
	// Experimental.
	QueryDefinitionName *string `field:"required" json:"queryDefinitionName" yaml:"queryDefinitionName"`
	// The query string to use for this query definition.
	// Experimental.
	QueryString QueryString `field:"required" json:"queryString" yaml:"queryString"`
	// Specify certain log groups for the query definition.
	// Experimental.
	LogGroups *[]ILogGroup `field:"optional" json:"logGroups" yaml:"logGroups"`
}

// Define a QueryString.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryString interface {
	// String representation of this QueryString.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueryString
type jsiiProxy_QueryString struct {
	_ byte // padding
}

// Experimental.
func NewQueryString(props *QueryStringProps) QueryString {
	_init_.Initialize()

	j := jsiiProxy_QueryString{}

	_jsii_.Create(
		"monocdk.aws_logs.QueryString",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewQueryString_Override(q QueryString, props *QueryStringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.QueryString",
		[]interface{}{props},
		q,
	)
}

func (q *jsiiProxy_QueryString) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a QueryString.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryStringProps struct {
	// Specifies which fields to display in the query results.
	// Experimental.
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Retrieves the specified fields from log events for display.
	// Experimental.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Filters the results of a query that's based on one or more conditions.
	// Experimental.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Specifies the number of log events returned by the query.
	// Experimental.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Extracts data from a log field and creates one or more ephemeral fields that you can process further in the query.
	// Experimental.
	Parse *string `field:"optional" json:"parse" yaml:"parse"`
	// Sorts the retrieved log events.
	// Experimental.
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Uses log field values to calculate aggregate statistics.
	// Experimental.
	Stats *string `field:"optional" json:"stats" yaml:"stats"`
}

// Resource Policy for CloudWatch Log Groups.
//
// Policies define the operations that are allowed on this resource.
//
// You almost never need to define this construct directly.
//
// All AWS resources that support resource policies have a method called
// `addToResourcePolicy()`, which will automatically create a new resource
// policy if one doesn't exist yet, otherwise it will add to the existing
// policy.
//
// Prefer to use `addToResourcePolicy()` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyStatement policyStatement
//
//   resourcePolicy := awscdk.Aws_logs.NewResourcePolicy(this, jsii.String("MyResourcePolicy"), &resourcePolicyProps{
//   	policyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	resourcePolicyName: jsii.String("resourcePolicyName"),
//   })
//
// Experimental.
type ResourcePolicy interface {
	awscdk.Resource
	// The IAM policy document for this resource policy.
	// Experimental.
	Document() awsiam.PolicyDocument
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ResourcePolicy
type jsiiProxy_ResourcePolicy struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ResourcePolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewResourcePolicy(scope constructs.Construct, id *string, props *ResourcePolicyProps) ResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_ResourcePolicy{}

	_jsii_.Create(
		"monocdk.aws_logs.ResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewResourcePolicy_Override(r ResourcePolicy, scope constructs.Construct, id *string, props *ResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.ResourcePolicy",
		[]interface{}{scope, id, props},
		r,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.ResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ResourcePolicy_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.ResourcePolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ResourcePolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_ResourcePolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_ResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to define Cloudwatch log group resource policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyStatement policyStatement
//
//   resourcePolicyProps := &resourcePolicyProps{
//   	policyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	resourcePolicyName: jsii.String("resourcePolicyName"),
//   }
//
// Experimental.
type ResourcePolicyProps struct {
	// Initial statements to add to the resource policy.
	// Experimental.
	PolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// Name of the log group resource policy.
	// Experimental.
	ResourcePolicyName *string `field:"optional" json:"resourcePolicyName" yaml:"resourcePolicyName"`
}

// How long, in days, the log contents will be retained.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var myLogsPublishingRole role
//   var vpc vpc
//
//
//   // Exporting logs from a cluster
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
//   		version: rds.auroraEngineVersion_VER_1_17_9(),
//   	}),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("error"),
//   		jsii.String("general"),
//   		jsii.String("slowquery"),
//   		jsii.String("audit"),
//   	},
//   	 // Export all available MySQL-based logs
//   	cloudwatchLogsRetention: logs.retentionDays_THREE_MONTHS,
//   	 // Optional - default is to never expire logs
//   	cloudwatchLogsRetentionRole: myLogsPublishingRole,
//   })
//
//   // Exporting logs from an instance
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	vpc: vpc,
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("postgresql"),
//   	},
//   })
//
// Experimental.
type RetentionDays string

const (
	// 1 day.
	// Experimental.
	RetentionDays_ONE_DAY RetentionDays = "ONE_DAY"
	// 3 days.
	// Experimental.
	RetentionDays_THREE_DAYS RetentionDays = "THREE_DAYS"
	// 5 days.
	// Experimental.
	RetentionDays_FIVE_DAYS RetentionDays = "FIVE_DAYS"
	// 1 week.
	// Experimental.
	RetentionDays_ONE_WEEK RetentionDays = "ONE_WEEK"
	// 2 weeks.
	// Experimental.
	RetentionDays_TWO_WEEKS RetentionDays = "TWO_WEEKS"
	// 1 month.
	// Experimental.
	RetentionDays_ONE_MONTH RetentionDays = "ONE_MONTH"
	// 2 months.
	// Experimental.
	RetentionDays_TWO_MONTHS RetentionDays = "TWO_MONTHS"
	// 3 months.
	// Experimental.
	RetentionDays_THREE_MONTHS RetentionDays = "THREE_MONTHS"
	// 4 months.
	// Experimental.
	RetentionDays_FOUR_MONTHS RetentionDays = "FOUR_MONTHS"
	// 5 months.
	// Experimental.
	RetentionDays_FIVE_MONTHS RetentionDays = "FIVE_MONTHS"
	// 6 months.
	// Experimental.
	RetentionDays_SIX_MONTHS RetentionDays = "SIX_MONTHS"
	// 1 year.
	// Experimental.
	RetentionDays_ONE_YEAR RetentionDays = "ONE_YEAR"
	// 13 months.
	// Experimental.
	RetentionDays_THIRTEEN_MONTHS RetentionDays = "THIRTEEN_MONTHS"
	// 18 months.
	// Experimental.
	RetentionDays_EIGHTEEN_MONTHS RetentionDays = "EIGHTEEN_MONTHS"
	// 2 years.
	// Experimental.
	RetentionDays_TWO_YEARS RetentionDays = "TWO_YEARS"
	// 5 years.
	// Experimental.
	RetentionDays_FIVE_YEARS RetentionDays = "FIVE_YEARS"
	// 6 years.
	// Experimental.
	RetentionDays_SIX_YEARS RetentionDays = "SIX_YEARS"
	// 7 years.
	// Experimental.
	RetentionDays_SEVEN_YEARS RetentionDays = "SEVEN_YEARS"
	// 8 years.
	// Experimental.
	RetentionDays_EIGHT_YEARS RetentionDays = "EIGHT_YEARS"
	// 9 years.
	// Experimental.
	RetentionDays_NINE_YEARS RetentionDays = "NINE_YEARS"
	// 10 years.
	// Experimental.
	RetentionDays_TEN_YEARS RetentionDays = "TEN_YEARS"
	// Retain logs forever.
	// Experimental.
	RetentionDays_INFINITE RetentionDays = "INFINITE"
)

// Space delimited text pattern.
//
// Example:
//   // Search for all events where the component is "HttpServer" and the
//   // result code is not equal to 200.
//   pattern := logs.filterPattern.spaceDelimited(jsii.String("time"), jsii.String("component"), jsii.String("..."), jsii.String("result_code"), jsii.String("latency")).whereString(jsii.String("component"), jsii.String("="), jsii.String("HttpServer")).whereNumber(jsii.String("result_code"), jsii.String("!="), jsii.Number(200))
//
// Experimental.
type SpaceDelimitedTextPattern interface {
	IFilterPattern
	// Experimental.
	LogPatternString() *string
	// Restrict where the pattern applies.
	// Experimental.
	WhereNumber(columnName *string, comparison *string, value *float64) SpaceDelimitedTextPattern
	// Restrict where the pattern applies.
	// Experimental.
	WhereString(columnName *string, comparison *string, value *string) SpaceDelimitedTextPattern
}

// The jsii proxy struct for SpaceDelimitedTextPattern
type jsiiProxy_SpaceDelimitedTextPattern struct {
	jsiiProxy_IFilterPattern
}

func (j *jsiiProxy_SpaceDelimitedTextPattern) LogPatternString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logPatternString",
		&returns,
	)
	return returns
}


// Experimental.
func NewSpaceDelimitedTextPattern(columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) SpaceDelimitedTextPattern {
	_init_.Initialize()

	j := jsiiProxy_SpaceDelimitedTextPattern{}

	_jsii_.Create(
		"monocdk.aws_logs.SpaceDelimitedTextPattern",
		[]interface{}{columns, restrictions},
		&j,
	)

	return &j
}

// Experimental.
func NewSpaceDelimitedTextPattern_Override(s SpaceDelimitedTextPattern, columns *[]*string, restrictions *map[string]*[]*ColumnRestriction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.SpaceDelimitedTextPattern",
		[]interface{}{columns, restrictions},
		s,
	)
}

// Construct a new instance of a space delimited text pattern.
//
// Since this class must be public, we can't rely on the user only creating it through
// the `LogPattern.spaceDelimited()` factory function. We must therefore validate the
// argument in the constructor. Since we're returning a copy on every mutation, and we
// don't want to re-validate the same things on every construction, we provide a limited
// set of mutator functions and only validate the new data every time.
// Experimental.
func SpaceDelimitedTextPattern_Construct(columns *[]*string) SpaceDelimitedTextPattern {
	_init_.Initialize()

	var returns SpaceDelimitedTextPattern

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.SpaceDelimitedTextPattern",
		"construct",
		[]interface{}{columns},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) WhereNumber(columnName *string, comparison *string, value *float64) SpaceDelimitedTextPattern {
	var returns SpaceDelimitedTextPattern

	_jsii_.Invoke(
		s,
		"whereNumber",
		[]interface{}{columnName, comparison, value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpaceDelimitedTextPattern) WhereString(columnName *string, comparison *string, value *string) SpaceDelimitedTextPattern {
	var returns SpaceDelimitedTextPattern

	_jsii_.Invoke(
		s,
		"whereString",
		[]interface{}{columnName, comparison, value},
		&returns,
	)

	return returns
}

// Properties for a new LogStream created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamOptions := &streamOptions{
//   	logStreamName: jsii.String("logStreamName"),
//   }
//
// Experimental.
type StreamOptions struct {
	// The name of the log stream to create.
	//
	// The name must be unique within the log group.
	// Experimental.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

// A new Subscription on a CloudWatch log group.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &subscriptionFilterProps{
//   	logGroup: logGroup,
//   	destination: destinations.NewLambdaDestination(fn),
//   	filterPattern: logs.filterPattern.allTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   })
//
// Experimental.
type SubscriptionFilter interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for SubscriptionFilter
type jsiiProxy_SubscriptionFilter struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_SubscriptionFilter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionFilter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSubscriptionFilter(scope constructs.Construct, id *string, props *SubscriptionFilterProps) SubscriptionFilter {
	_init_.Initialize()

	j := jsiiProxy_SubscriptionFilter{}

	_jsii_.Create(
		"monocdk.aws_logs.SubscriptionFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSubscriptionFilter_Override(s SubscriptionFilter, scope constructs.Construct, id *string, props *SubscriptionFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.SubscriptionFilter",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func SubscriptionFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.SubscriptionFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SubscriptionFilter_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_logs.SubscriptionFilter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SubscriptionFilter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SubscriptionFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SubscriptionFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SubscriptionFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a new SubscriptionFilter created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterPattern iFilterPattern
//   var logSubscriptionDestination iLogSubscriptionDestination
//
//   subscriptionFilterOptions := &subscriptionFilterOptions{
//   	destination: logSubscriptionDestination,
//   	filterPattern: filterPattern,
//   }
//
// Experimental.
type SubscriptionFilterOptions struct {
	// The destination to send the filtered events to.
	//
	// For example, a Kinesis stream or a Lambda function.
	// Experimental.
	Destination ILogSubscriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Log events matching this pattern will be sent to the destination.
	// Experimental.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
}

// Properties for a SubscriptionFilter.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &subscriptionFilterProps{
//   	logGroup: logGroup,
//   	destination: destinations.NewLambdaDestination(fn),
//   	filterPattern: logs.filterPattern.allTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   })
//
// Experimental.
type SubscriptionFilterProps struct {
	// The destination to send the filtered events to.
	//
	// For example, a Kinesis stream or a Lambda function.
	// Experimental.
	Destination ILogSubscriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Log events matching this pattern will be sent to the destination.
	// Experimental.
	FilterPattern IFilterPattern `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to create the subscription on.
	// Experimental.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

