package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awsmsk/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var logGroup logGroup
//
//   brokerLogging := &brokerLogging{
//   	cloudwatchLogGroup: logGroup,
//   	firehoseDeliveryStreamName: jsii.String("firehoseDeliveryStreamName"),
//   	s3: &s3LoggingConfiguration{
//   		bucket: bucket,
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
// Experimental.
type BrokerLogging struct {
	// The CloudWatch Logs group that is the destination for broker logs.
	// Experimental.
	CloudwatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	// Experimental.
	FirehoseDeliveryStreamName *string `field:"optional" json:"firehoseDeliveryStreamName" yaml:"firehoseDeliveryStreamName"`
	// Details of the Amazon S3 destination for broker logs.
	// Experimental.
	S3 *S3LoggingConfiguration `field:"optional" json:"s3" yaml:"s3"`
}

// A CloudFormation `AWS::MSK::BatchScramSecret`.
//
// Represents a secret stored in the Amazon Secrets Manager that can be used to authenticate with a cluster using a user name and a password.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBatchScramSecret := awscdk.Aws_msk.NewCfnBatchScramSecret(this, jsii.String("MyCfnBatchScramSecret"), &cfnBatchScramSecretProps{
//   	clusterArn: jsii.String("clusterArn"),
//
//   	// the properties below are optional
//   	secretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   })
//
type CfnBatchScramSecret interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn() *string
	SetClusterArn(val *string)
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list of Amazon Secrets Manager secret ARNs.
	SecretArnList() *[]*string
	SetSecretArnList(val *[]*string)
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

// The jsii proxy struct for CfnBatchScramSecret
type jsiiProxy_CfnBatchScramSecret struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBatchScramSecret) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) SecretArnList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretArnList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBatchScramSecret) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecret(scope awscdk.Construct, id *string, props *CfnBatchScramSecretProps) CfnBatchScramSecret {
	_init_.Initialize()

	j := jsiiProxy_CfnBatchScramSecret{}

	_jsii_.Create(
		"monocdk.aws_msk.CfnBatchScramSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::BatchScramSecret`.
func NewCfnBatchScramSecret_Override(c CfnBatchScramSecret, scope awscdk.Construct, id *string, props *CfnBatchScramSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.CfnBatchScramSecret",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBatchScramSecret) SetClusterArn(val *string) {
	_jsii_.Set(
		j,
		"clusterArn",
		val,
	)
}

func (j *jsiiProxy_CfnBatchScramSecret) SetSecretArnList(val *[]*string) {
	_jsii_.Set(
		j,
		"secretArnList",
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
func CfnBatchScramSecret_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnBatchScramSecret",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBatchScramSecret_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnBatchScramSecret",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBatchScramSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnBatchScramSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBatchScramSecret_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_msk.CfnBatchScramSecret",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBatchScramSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBatchScramSecret) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnBatchScramSecret`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBatchScramSecretProps := &cfnBatchScramSecretProps{
//   	clusterArn: jsii.String("clusterArn"),
//
//   	// the properties below are optional
//   	secretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }
//
type CfnBatchScramSecretProps struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// A list of Amazon Secrets Manager secret ARNs.
	SecretArnList *[]*string `field:"optional" json:"secretArnList" yaml:"secretArnList"`
}

// A CloudFormation `AWS::MSK::Cluster`.
//
// The `AWS::MSK::Cluster` resource creates an Amazon MSK cluster . For more information, see [What Is Amazon MSK?](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html) in the *Amazon MSK Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_msk.NewCfnCluster(this, jsii.String("MyCfnCluster"), &cfnClusterProps{
//   	brokerNodeGroupInfo: &brokerNodeGroupInfoProperty{
//   		clientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		brokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		connectivityInfo: &connectivityInfoProperty{
//   			publicAccess: &publicAccessProperty{
//   				type: jsii.String("type"),
//   			},
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		storageInfo: &storageInfoProperty{
//   			ebsStorageInfo: &eBSStorageInfoProperty{
//   				provisionedThroughput: &provisionedThroughputProperty{
//   					enabled: jsii.Boolean(false),
//   					volumeThroughput: jsii.Number(123),
//   				},
//   				volumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	kafkaVersion: jsii.String("kafkaVersion"),
//   	numberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	clientAuthentication: &clientAuthenticationProperty{
//   		sasl: &saslProperty{
//   			iam: &iamProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			scram: &scramProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		tls: &tlsProperty{
//   			certificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticated: &unauthenticatedProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	configurationInfo: &configurationInfoProperty{
//   		arn: jsii.String("arn"),
//   		revision: jsii.Number(123),
//   	},
//   	currentVersion: jsii.String("currentVersion"),
//   	encryptionInfo: &encryptionInfoProperty{
//   		encryptionAtRest: &encryptionAtRestProperty{
//   			dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		encryptionInTransit: &encryptionInTransitProperty{
//   			clientBroker: jsii.String("clientBroker"),
//   			inCluster: jsii.Boolean(false),
//   		},
//   	},
//   	enhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	loggingInfo: &loggingInfoProperty{
//   		brokerLogs: &brokerLogsProperty{
//   			cloudWatchLogs: &cloudWatchLogsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				logGroup: jsii.String("logGroup"),
//   			},
//   			firehose: &firehoseProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				deliveryStream: jsii.String("deliveryStream"),
//   			},
//   			s3: &s3Property{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				bucket: jsii.String("bucket"),
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	openMonitoring: &openMonitoringProperty{
//   		prometheus: &prometheusProperty{
//   			jmxExporter: &jmxExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   			nodeExporter: &nodeExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// The setup to be used for brokers in the cluster.
	//
	// AWS CloudFormation may replace the cluster when you update certain `BrokerNodeGroupInfo` properties. To understand the update behavior for your use case, you should review the child properties for [`BrokerNodeGroupInfo`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#aws-properties-msk-cluster-brokernodegroupinfo-properties) .
	BrokerNodeGroupInfo() interface{}
	SetBrokerNodeGroupInfo(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Includes information related to client authentication.
	ClientAuthentication() interface{}
	SetClientAuthentication(val interface{})
	// The name of the cluster.
	ClusterName() *string
	SetClusterName(val *string)
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo() interface{}
	SetConfigurationInfo(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The version of the cluster that you want to update.
	CurrentVersion() *string
	SetCurrentVersion(val *string)
	// Includes all encryption-related information.
	EncryptionInfo() interface{}
	SetEncryptionInfo(val interface{})
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring() *string
	SetEnhancedMonitoring(val *string)
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion() *string
	SetKafkaVersion(val *string)
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
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
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes() *float64
	SetNumberOfBrokerNodes(val *float64)
	// The settings for open monitoring.
	OpenMonitoring() interface{}
	SetOpenMonitoring(val interface{})
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
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) BrokerNodeGroupInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClientAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ConfigurationInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EncryptionInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LoggingInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OpenMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster(scope awscdk.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"monocdk.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope awscdk.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetBrokerNodeGroupInfo(val interface{}) {
	_jsii_.Set(
		j,
		"brokerNodeGroupInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClientAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"clientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfigurationInfo(val interface{}) {
	_jsii_.Set(
		j,
		"configurationInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetCurrentVersion(val *string) {
	_jsii_.Set(
		j,
		"currentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEncryptionInfo(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEnhancedMonitoring(val *string) {
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKafkaVersion(val *string) {
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLoggingInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetNumberOfBrokerNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetOpenMonitoring(val interface{}) {
	_jsii_.Set(
		j,
		"openMonitoring",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_msk.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This configuration specifies the details of these destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerLogsProperty := &brokerLogsProperty{
//   	cloudWatchLogs: &cloudWatchLogsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		logGroup: jsii.String("logGroup"),
//   	},
//   	firehose: &firehoseProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		deliveryStream: jsii.String("deliveryStream"),
//   	},
//   	s3: &s3Property{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bucket: jsii.String("bucket"),
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnCluster_BrokerLogsProperty struct {
	// Details of the CloudWatch Logs destination for broker logs.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Details of the Amazon MSK destination for broker logs.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

// The setup to be used for brokers in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerNodeGroupInfoProperty := &brokerNodeGroupInfoProperty{
//   	clientSubnets: []*string{
//   		jsii.String("clientSubnets"),
//   	},
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	brokerAzDistribution: jsii.String("brokerAzDistribution"),
//   	connectivityInfo: &connectivityInfoProperty{
//   		publicAccess: &publicAccessProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	storageInfo: &storageInfoProperty{
//   		ebsStorageInfo: &eBSStorageInfoProperty{
//   			provisionedThroughput: &provisionedThroughputProperty{
//   				enabled: jsii.Boolean(false),
//   				volumeThroughput: jsii.Number(123),
//   			},
//   			volumeSize: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// The list of subnets to connect to in the client virtual private cloud (VPC).
	//
	// Amazon creates elastic network interfaces inside these subnets. Client applications use elastic network interfaces to produce and consume data.
	//
	// Specify exactly two subnets if you are using the US West (N. California) Region. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	//
	// Client subnets can't occupy the Availability Zone with ID `use1-az3` .
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, and kafka.m5.24xlarge.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// This parameter is currently not in use.
	BrokerAzDistribution *string `field:"optional" json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// Information about the cluster's connectivity setting.
	ConnectivityInfo interface{} `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo interface{} `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

// Includes information related to client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientAuthenticationProperty := &clientAuthenticationProperty{
//   	sasl: &saslProperty{
//   		iam: &iamProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		scram: &scramProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	tls: &tlsProperty{
//   		certificateAuthorityArnList: []*string{
//   			jsii.String("certificateAuthorityArnList"),
//   		},
//   		enabled: jsii.Boolean(false),
//   	},
//   	unauthenticated: &unauthenticatedProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_ClientAuthenticationProperty struct {
	// Details for ClientAuthentication using SASL.
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Details for client authentication using TLS.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Details for ClientAuthentication using no authentication.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

// Details of the CloudWatch Logs destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsProperty := &cloudWatchLogsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logGroup: jsii.String("logGroup"),
//   }
//
type CfnCluster_CloudWatchLogsProperty struct {
	// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudWatch Logs group that is the destination for broker logs.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

// Specifies the Amazon MSK configuration to use for the brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationInfoProperty := &configurationInfoProperty{
//   	arn: jsii.String("arn"),
//   	revision: jsii.Number(123),
//   }
//
type CfnCluster_ConfigurationInfoProperty struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, `arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1` .
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

// Specifies whether the cluster's brokers are publicly accessible.
//
// By default, they are not.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectivityInfoProperty := &connectivityInfoProperty{
//   	publicAccess: &publicAccessProperty{
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnCluster_ConnectivityInfoProperty struct {
	// Specifies whether the cluster's brokers are accessible from the internet.
	//
	// Public access is off by default.
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

// Contains information about the EBS storage volumes attached to brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSStorageInfoProperty := &eBSStorageInfoProperty{
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		enabled: jsii.Boolean(false),
//   		volumeThroughput: jsii.Number(123),
//   	},
//   	volumeSize: jsii.Number(123),
//   }
//
type CfnCluster_EBSStorageInfoProperty struct {
	// Specifies whether provisioned throughput is turned on and the volume throughput target.
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

// The data volume encryption details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestProperty := &encryptionAtRestProperty{
//   	dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   }
//
type CfnCluster_EncryptionAtRestProperty struct {
	// The ARN of the Amazon KMS key for encrypting data at rest.
	//
	// If you don't specify a KMS key, MSK creates one for you and uses it on your behalf.
	DataVolumeKmsKeyId *string `field:"required" json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

// The settings for encrypting data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInTransitProperty := &encryptionInTransitProperty{
//   	clientBroker: jsii.String("clientBroker"),
//   	inCluster: jsii.Boolean(false),
//   }
//
type CfnCluster_EncryptionInTransitProperty struct {
	// Indicates the encryption setting for data in transit between clients and brokers. The following are the possible values.
	//
	// - `TLS` means that client-broker communication is enabled with TLS only.
	// - `TLS_PLAINTEXT` means that client-broker communication is enabled for both TLS-encrypted, as well as plain text data.
	// - `PLAINTEXT` means that client-broker communication is enabled in plain text only.
	//
	// The default value is `TLS` .
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted.
	//
	// When set to false, the communication happens in plain text. The default value is true.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}

// Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInfoProperty := &encryptionInfoProperty{
//   	encryptionAtRest: &encryptionAtRestProperty{
//   		dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   	},
//   	encryptionInTransit: &encryptionInTransitProperty{
//   		clientBroker: jsii.String("clientBroker"),
//   		inCluster: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_EncryptionInfoProperty struct {
	// The data-volume encryption details.
	EncryptionAtRest interface{} `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// The details for encryption in transit.
	EncryptionInTransit interface{} `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &firehoseProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	deliveryStream: jsii.String("deliveryStream"),
//   }
//
type CfnCluster_FirehoseProperty struct {
	// Specifies whether broker logs get sent to the specified Kinesis Data Firehose delivery stream.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

// Details for IAM access control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamProperty := &iamProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_IamProperty struct {
	// Whether IAM access control is enabled.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Indicates whether you want to enable or disable the JMX Exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jmxExporterProperty := &jmxExporterProperty{
//   	enabledInBroker: jsii.Boolean(false),
//   }
//
type CfnCluster_JmxExporterProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	EnabledInBroker interface{} `field:"required" json:"enabledInBroker" yaml:"enabledInBroker"`
}

// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingInfoProperty := &loggingInfoProperty{
//   	brokerLogs: &brokerLogsProperty{
//   		cloudWatchLogs: &cloudWatchLogsProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			logGroup: jsii.String("logGroup"),
//   		},
//   		firehose: &firehoseProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			deliveryStream: jsii.String("deliveryStream"),
//   		},
//   		s3: &s3Property{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			bucket: jsii.String("bucket"),
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
type CfnCluster_LoggingInfoProperty struct {
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	BrokerLogs interface{} `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

// Indicates whether you want to enable or disable the Node Exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeExporterProperty := &nodeExporterProperty{
//   	enabledInBroker: jsii.Boolean(false),
//   }
//
type CfnCluster_NodeExporterProperty struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	EnabledInBroker interface{} `field:"required" json:"enabledInBroker" yaml:"enabledInBroker"`
}

// JMX and Node monitoring for the MSK cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openMonitoringProperty := &openMonitoringProperty{
//   	prometheus: &prometheusProperty{
//   		jmxExporter: &jmxExporterProperty{
//   			enabledInBroker: jsii.Boolean(false),
//   		},
//   		nodeExporter: &nodeExporterProperty{
//   			enabledInBroker: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCluster_OpenMonitoringProperty struct {
	// Prometheus exporter settings.
	Prometheus interface{} `field:"required" json:"prometheus" yaml:"prometheus"`
}

// Prometheus settings for open monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prometheusProperty := &prometheusProperty{
//   	jmxExporter: &jmxExporterProperty{
//   		enabledInBroker: jsii.Boolean(false),
//   	},
//   	nodeExporter: &nodeExporterProperty{
//   		enabledInBroker: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_PrometheusProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter interface{} `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter interface{} `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

// Specifies whether provisioned throughput is turned on and the volume throughput target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	enabled: jsii.Boolean(false),
//   	volumeThroughput: jsii.Number(123),
//   }
//
type CfnCluster_ProvisionedThroughputProperty struct {
	// Specifies whether provisioned throughput is turned on for the cluster.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The provisioned throughput rate in MiB per second.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

// Specifies whether the cluster's brokers are accessible from the internet.
//
// Public access is off by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessProperty := &publicAccessProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnCluster_PublicAccessProperty struct {
	// Set to `DISABLED` to turn off public access or to `SERVICE_PROVIDED_EIPS` to turn it on.
	//
	// Public access if off by default.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The details of the Amazon S3 destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bucket: jsii.String("bucket"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnCluster_S3Property struct {
	// Specifies whether broker logs get sent to the specified Amazon S3 destination.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for broker logs.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Details for client authentication using SASL.
//
// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   saslProperty := &saslProperty{
//   	iam: &iamProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	scram: &scramProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_SaslProperty struct {
	// Details for IAM access control.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Details for SASL/SCRAM client authentication.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

// Details for SASL/SCRAM client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scramProperty := &scramProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_ScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Contains information about storage volumes attached to MSK broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageInfoProperty := &storageInfoProperty{
//   	ebsStorageInfo: &eBSStorageInfoProperty{
//   		provisionedThroughput: &provisionedThroughputProperty{
//   			enabled: jsii.Boolean(false),
//   			volumeThroughput: jsii.Number(123),
//   		},
//   		volumeSize: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_StorageInfoProperty struct {
	// EBS volume information.
	EbsStorageInfo interface{} `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

// Details for client authentication using TLS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsProperty := &tlsProperty{
//   	certificateAuthorityArnList: []*string{
//   		jsii.String("certificateAuthorityArnList"),
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_TlsProperty struct {
	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList *[]*string `field:"optional" json:"certificateAuthorityArnList" yaml:"certificateAuthorityArnList"`
	// TLS authentication is enabled or not.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Details for allowing no client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unauthenticatedProperty := &unauthenticatedProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_UnauthenticatedProperty struct {
	// Unauthenticated is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	brokerNodeGroupInfo: &brokerNodeGroupInfoProperty{
//   		clientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		brokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		connectivityInfo: &connectivityInfoProperty{
//   			publicAccess: &publicAccessProperty{
//   				type: jsii.String("type"),
//   			},
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		storageInfo: &storageInfoProperty{
//   			ebsStorageInfo: &eBSStorageInfoProperty{
//   				provisionedThroughput: &provisionedThroughputProperty{
//   					enabled: jsii.Boolean(false),
//   					volumeThroughput: jsii.Number(123),
//   				},
//   				volumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	kafkaVersion: jsii.String("kafkaVersion"),
//   	numberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	clientAuthentication: &clientAuthenticationProperty{
//   		sasl: &saslProperty{
//   			iam: &iamProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			scram: &scramProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		tls: &tlsProperty{
//   			certificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticated: &unauthenticatedProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	configurationInfo: &configurationInfoProperty{
//   		arn: jsii.String("arn"),
//   		revision: jsii.Number(123),
//   	},
//   	currentVersion: jsii.String("currentVersion"),
//   	encryptionInfo: &encryptionInfoProperty{
//   		encryptionAtRest: &encryptionAtRestProperty{
//   			dataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		encryptionInTransit: &encryptionInTransitProperty{
//   			clientBroker: jsii.String("clientBroker"),
//   			inCluster: jsii.Boolean(false),
//   		},
//   	},
//   	enhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	loggingInfo: &loggingInfoProperty{
//   		brokerLogs: &brokerLogsProperty{
//   			cloudWatchLogs: &cloudWatchLogsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				logGroup: jsii.String("logGroup"),
//   			},
//   			firehose: &firehoseProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				deliveryStream: jsii.String("deliveryStream"),
//   			},
//   			s3: &s3Property{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				bucket: jsii.String("bucket"),
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	openMonitoring: &openMonitoringProperty{
//   		prometheus: &prometheusProperty{
//   			jmxExporter: &jmxExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   			nodeExporter: &nodeExporterProperty{
//   				enabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnClusterProps struct {
	// The setup to be used for brokers in the cluster.
	//
	// AWS CloudFormation may replace the cluster when you update certain `BrokerNodeGroupInfo` properties. To understand the update behavior for your use case, you should review the child properties for [`BrokerNodeGroupInfo`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#aws-properties-msk-cluster-brokernodegroupinfo-properties) .
	BrokerNodeGroupInfo interface{} `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Includes information related to client authentication.
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo interface{} `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// The version of the cluster that you want to update.
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Includes all encryption-related information.
	EncryptionInfo interface{} `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The settings for open monitoring.
	OpenMonitoring interface{} `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::MSK::Configuration`.
//
// Creates a new MSK configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfiguration := awscdk.Aws_msk.NewCfnConfiguration(this, jsii.String("MyCfnConfiguration"), &cfnConfigurationProps{
//   	name: jsii.String("name"),
//   	serverProperties: jsii.String("serverProperties"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kafkaVersionsList: []*string{
//   		jsii.String("kafkaVersionsList"),
//   	},
//   })
//
type CfnConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the configuration.
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
	// The description of the configuration.
	Description() *string
	SetDescription(val *string)
	// A list of the versions of Apache Kafka with which you can use this MSK configuration.
	//
	// You can use this configuration for an MSK cluster only if the Apache Kafka version specified for the cluster appears in this list.
	KafkaVersionsList() *[]*string
	SetKafkaVersionsList(val *[]*string)
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
	// The name of the configuration.
	//
	// Configuration names are strings that match the regex "^[0-9A-Za-z][0-9A-Za-z-]{0,}$".
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Contents of the server.properties file. When using the API, you must ensure that the contents of the file are base64 encoded. When using the console, the SDK, or the CLI, the contents of server.properties can be in plaintext.
	ServerProperties() *string
	SetServerProperties(val *string)
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

// The jsii proxy struct for CfnConfiguration
type jsiiProxy_CfnConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) KafkaVersionsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kafkaVersionsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) ServerProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::Configuration`.
func NewCfnConfiguration(scope awscdk.Construct, id *string, props *CfnConfigurationProps) CfnConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnConfiguration{}

	_jsii_.Create(
		"monocdk.aws_msk.CfnConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::Configuration`.
func NewCfnConfiguration_Override(c CfnConfiguration, scope awscdk.Construct, id *string, props *CfnConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.CfnConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetKafkaVersionsList(val *[]*string) {
	_jsii_.Set(
		j,
		"kafkaVersionsList",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnConfiguration) SetServerProperties(val *string) {
	_jsii_.Set(
		j,
		"serverProperties",
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
func CfnConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_msk.CfnConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &cfnConfigurationProps{
//   	name: jsii.String("name"),
//   	serverProperties: jsii.String("serverProperties"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kafkaVersionsList: []*string{
//   		jsii.String("kafkaVersionsList"),
//   	},
//   }
//
type CfnConfigurationProps struct {
	// The name of the configuration.
	//
	// Configuration names are strings that match the regex "^[0-9A-Za-z][0-9A-Za-z-]{0,}$".
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contents of the server.properties file. When using the API, you must ensure that the contents of the file are base64 encoded. When using the console, the SDK, or the CLI, the contents of server.properties can be in plaintext.
	ServerProperties *string `field:"required" json:"serverProperties" yaml:"serverProperties"`
	// The description of the configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of the versions of Apache Kafka with which you can use this MSK configuration.
	//
	// You can use this configuration for an MSK cluster only if the Apache Kafka version specified for the cluster appears in this list.
	KafkaVersionsList *[]*string `field:"optional" json:"kafkaVersionsList" yaml:"kafkaVersionsList"`
}

// Configuration properties for client authentication.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClientAuthentication interface {
	// - properties for SASL authentication.
	// Experimental.
	SaslProps() *SaslAuthProps
	// - properties for TLS authentication.
	// Experimental.
	TlsProps() *TlsAuthProps
}

// The jsii proxy struct for ClientAuthentication
type jsiiProxy_ClientAuthentication struct {
	_ byte // padding
}

func (j *jsiiProxy_ClientAuthentication) SaslProps() *SaslAuthProps {
	var returns *SaslAuthProps
	_jsii_.Get(
		j,
		"saslProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClientAuthentication) TlsProps() *TlsAuthProps {
	var returns *TlsAuthProps
	_jsii_.Get(
		j,
		"tlsProps",
		&returns,
	)
	return returns
}


// SASL authentication.
// Experimental.
func ClientAuthentication_Sasl(props *SaslAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.ClientAuthentication",
		"sasl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// TLS authentication.
// Experimental.
func ClientAuthentication_Tls(props *TlsAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.ClientAuthentication",
		"tls",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Indicates the encryption setting for data in transit between clients and brokers.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClientBrokerEncryption string

const (
	// TLS means that client-broker communication is enabled with TLS only.
	// Experimental.
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	// TLS_PLAINTEXT means that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// Experimental.
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	// PLAINTEXT means that client-broker communication is enabled in plaintext only.
	// Experimental.
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)

// Create a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type Cluster interface {
	awscdk.Resource
	ICluster
	// Get the list of brokers that a client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more hostname:port pairs.
	// Experimental.
	BootstrapBrokers() *string
	// Get the list of brokers that a SASL/SCRAM authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more dns name (or IP) and SASL SCRAM port pairs.
	// Experimental.
	BootstrapBrokersSaslScram() *string
	// Get the list of brokers that a TLS authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more DNS names (or IP) and TLS port pairs.
	// Experimental.
	BootstrapBrokersTls() *string
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
	// Manages connections for the cluster.
	// Experimental.
	Connections() awsec2.Connections
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
	// Key used to encrypt SASL/SCRAM users.
	// Experimental.
	SaslScramAuthenticationKey() awskms.IKey
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Get the ZooKeeper Connection string.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to the Apache ZooKeeper cluster.
	// Experimental.
	ZookeeperConnectionString() *string
	// Get the ZooKeeper Connection string for a TLS enabled cluster.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to zookeeper cluster on TLS port.
	// Experimental.
	ZookeeperConnectionStringTls() *string
	// A list of usersnames to register with the cluster.
	//
	// The password will automatically be generated using Secrets
	// Manager and the { username, password } JSON object stored in Secrets Manager as `AmazonMSK_username`.
	//
	// Must be using the SASL/SCRAM authentication mechanism.
	// Experimental.
	AddUser(usernames ...*string)
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

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SaslScramAuthenticationKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"saslScramAuthenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionStringTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionStringTls",
		&returns,
	)
	return returns
}


// Experimental.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"monocdk.aws_msk.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Reference an existing cluster, defined outside of the CDK code, by name.
// Experimental.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Cluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddUser(usernames ...*string) {
	args := []interface{}{}
	for _, a := range usernames {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUser",
		args,
	)
}

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Cluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon MSK configuration to use for the cluster.
//
// Note: There is currently no Cloudformation Resource to create a Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigurationInfo := &clusterConfigurationInfo{
//   	arn: jsii.String("arn"),
//   	revision: jsii.Number(123),
//   }
//
// Experimental.
type ClusterConfigurationInfo struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

// The level of monitoring for the MSK cluster.
// See: https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html#metrics-details
//
// Experimental.
type ClusterMonitoringLevel string

const (
	// Default metrics are the essential metrics to monitor.
	// Experimental.
	ClusterMonitoringLevel_DEFAULT ClusterMonitoringLevel = "DEFAULT"
	// Per Broker metrics give you metrics at the broker level.
	// Experimental.
	ClusterMonitoringLevel_PER_BROKER ClusterMonitoringLevel = "PER_BROKER"
	// Per Topic Per Broker metrics help you understand volume at the topic level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_BROKER ClusterMonitoringLevel = "PER_TOPIC_PER_BROKER"
	// Per Topic Per Partition metrics help you understand consumer group lag at the topic partition level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION ClusterMonitoringLevel = "PER_TOPIC_PER_PARTITION"
)

// Properties for a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClusterProps struct {
	// The physical name of the cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	// Experimental.
	KafkaVersion KafkaVersion `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Defines the virtual networking environment for this cluster.
	//
	// Must have at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Configuration properties for client authentication.
	//
	// MSK supports using private TLS certificates or SASL/SCRAM to authenticate the identity of clients.
	// Experimental.
	ClientAuthentication ClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	// Experimental.
	ConfigurationInfo *ClusterConfigurationInfo `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// Information about storage volumes attached to MSK broker nodes.
	// Experimental.
	EbsStorageInfo *EbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
	// Config details for encryption in transit.
	// Experimental.
	EncryptionInTransit *EncryptionInTransitConfig `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// The EC2 instance type that you want Amazon MSK to use when it creates your brokers.
	// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html#broker-instance-types
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Configure your MSK cluster to send broker logs to different destination types.
	// Experimental.
	Logging *BrokerLogging `field:"optional" json:"logging" yaml:"logging"`
	// Cluster monitoring configuration.
	// Experimental.
	Monitoring *MonitoringConfiguration `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Number of Apache Kafka brokers deployed in each Availability Zone.
	// Experimental.
	NumberOfBrokerNodes *float64 `field:"optional" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// What to do when this resource is deleted from a stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The AWS security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the nodes within the VPC.
	//
	// Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	// The subnets that you specify must be in distinct Availability Zones.
	// Client subnets can't be in Availability Zone us-east-1e.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// EBS volume information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   ebsStorageInfo := &ebsStorageInfo{
//   	encryptionKey: key,
//   	volumeSize: jsii.Number(123),
//   }
//
// Experimental.
type EbsStorageInfo struct {
	// The AWS KMS key for encrypting data at rest.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

// The settings for encrypting data in transit.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-encryption.html#msk-encryption-in-transit
//
// Experimental.
type EncryptionInTransitConfig struct {
	// Indicates the encryption setting for data in transit between clients and brokers.
	// Experimental.
	ClientBroker ClientBrokerEncryption `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Indicates that data communication among the broker nodes of the cluster is encrypted.
	// Experimental.
	EnableInCluster *bool `field:"optional" json:"enableInCluster" yaml:"enableInCluster"`
}

// Represents a MSK Cluster.
// Experimental.
type ICluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Kafka cluster version.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type KafkaVersion interface {
	// cluster version number.
	// Experimental.
	Version() *string
}

// The jsii proxy struct for KafkaVersion
type jsiiProxy_KafkaVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
// Experimental.
func KafkaVersion_Of(version *string) KafkaVersion {
	_init_.Initialize()

	var returns KafkaVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.KafkaVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KafkaVersion_V1_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V1_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_2_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_3_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_4_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_4_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_5_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_5_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_3() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_3",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_8_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_8_1",
		&returns,
	)
	return returns
}

// Monitoring Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringConfiguration := &monitoringConfiguration{
//   	clusterMonitoringLevel: awscdk.Aws_msk.clusterMonitoringLevel_DEFAULT,
//   	enablePrometheusJmxExporter: jsii.Boolean(false),
//   	enablePrometheusNodeExporter: jsii.Boolean(false),
//   }
//
// Experimental.
type MonitoringConfiguration struct {
	// Specifies the level of monitoring for the MSK cluster.
	// Experimental.
	ClusterMonitoringLevel ClusterMonitoringLevel `field:"optional" json:"clusterMonitoringLevel" yaml:"clusterMonitoringLevel"`
	// Indicates whether you want to enable or disable the JMX Exporter.
	// Experimental.
	EnablePrometheusJmxExporter *bool `field:"optional" json:"enablePrometheusJmxExporter" yaml:"enablePrometheusJmxExporter"`
	// Indicates whether you want to enable or disable the Prometheus Node Exporter.
	//
	// You can use the Prometheus Node Exporter to get CPU and disk metrics for the broker nodes.
	// Experimental.
	EnablePrometheusNodeExporter *bool `field:"optional" json:"enablePrometheusNodeExporter" yaml:"enablePrometheusNodeExporter"`
}

// Details of the Amazon S3 destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3LoggingConfiguration := &s3LoggingConfiguration{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type S3LoggingConfiguration struct {
	// The S3 bucket that is the destination for broker logs.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// SASL authentication properties.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type SaslAuthProps struct {
	// Enable IAM access control.
	// Experimental.
	Iam *bool `field:"optional" json:"iam" yaml:"iam"`
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Enable SASL/SCRAM authentication.
	// Experimental.
	Scram *bool `field:"optional" json:"scram" yaml:"scram"`
}

// TLS authentication properties.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.tls(&tlsAuthProps{
//   		certificateAuthorities: []iCertificateAuthority{
//   			acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `field:"optional" json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

