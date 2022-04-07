package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Synthetics::Canary`.
//
// Creates or updates a canary. Canaries are scripts that monitor your endpoints and APIs from the outside-in. Canaries help you check the availability and latency of your web services and troubleshoot anomalies by investigating load time data, screenshots of the UI, logs, and metrics. You can set up a canary to run continuously or just once.
//
// To create canaries, you must have the `CloudWatchSyntheticsFullAccess` policy. If you are creating a new IAM role for the canary, you also need the the `iam:CreateRole` , `iam:CreatePolicy` and `iam:AttachRolePolicy` permissions. For more information, see [Necessary Roles and Permissions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles) .
//
// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the Amazon Resource Name (ARN) for the canary, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   cfnCanary := synthetics.NewCfnCanary(this, jsii.String("MyCfnCanary"), &cfnCanaryProps{
//   	artifactS3Location: jsii.String("artifactS3Location"),
//   	code: &codeProperty{
//   		handler: jsii.String("handler"),
//
//   		// the properties below are optional
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		script: jsii.String("script"),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	name: jsii.String("name"),
//   	runtimeVersion: jsii.String("runtimeVersion"),
//   	schedule: &scheduleProperty{
//   		expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		durationInSeconds: jsii.String("durationInSeconds"),
//   	},
//   	startCanaryAfterCreation: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	artifactConfig: &artifactConfigProperty{
//   		s3Encryption: &s3EncryptionProperty{
//   			encryptionMode: jsii.String("encryptionMode"),
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	failureRetentionPeriod: jsii.Number(123),
//   	runConfig: &runConfigProperty{
//   		activeTracing: jsii.Boolean(false),
//   		environmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		memoryInMb: jsii.Number(123),
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	successRetentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visualReference: &visualReferenceProperty{
//   		baseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   		// the properties below are optional
//   		baseScreenshots: []interface{}{
//   			&baseScreenshotProperty{
//   				screenshotName: jsii.String("screenshotName"),
//
//   				// the properties below are optional
//   				ignoreCoordinates: []*string{
//   					jsii.String("ignoreCoordinates"),
//   				},
//   			},
//   		},
//   	},
//   	vpcConfig: &vPCConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		vpcId: jsii.String("vpcId"),
//   	},
//   })
//
type CfnCanary interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig() interface{}
	SetArtifactConfig(val interface{})
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of this canary.
	//
	// Artifacts include the log file, screenshots, and HAR files. Specify the full location path, including `s3://` at the beginning of the path.
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	// The ID of the canary.
	AttrId() *string
	// The state of the canary.
	//
	// For example, `RUNNING` .
	AttrState() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Use this structure to input your script code for the canary.
	//
	// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
	Code() interface{}
	SetCode(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN of the IAM role to be used to run the canary.
	//
	// This role must already exist, and must include `lambda.amazonaws.com` as a principal in the trust policy. The role must also have the following permissions:
	//
	// - `s3:PutObject`
	// - `s3:GetBucketLocation`
	// - `s3:ListAllMyBuckets`
	// - `cloudwatch:PutMetricData`
	// - `logs:CreateLogGroup`
	// - `logs:CreateLogStream`
	// - `logs:PutLogEvents`.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The number of days to retain data about failed runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name for this canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A structure that contains input information for a canary run.
	//
	// If you omit this structure, the frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.
	RunConfig() interface{}
	SetRunConfig(val interface{})
	// Specifies the runtime version to use for the canary.
	//
	// For more information about runtime versions, see [Canary Runtime Versions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) .
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	// A structure that contains information about how often the canary is to run, and when these runs are to stop.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specify TRUE to have the canary start making runs immediately after it is created.
	//
	// A canary that you create using CloudFormation can't be used to monitor the CloudFormation stack that creates the canary or to roll back that stack if there is a failure.
	StartCanaryAfterCreation() interface{}
	SetStartCanaryAfterCreation(val interface{})
	// The number of days to retain data about successful runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	// The list of key-value pairs that are associated with the canary.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// If this canary performs visual monitoring by comparing screenshots, this structure contains the ID of the canary run to use as the baseline for screenshots, and the coordinates of any parts of the screen to ignore during the visual monitoring comparison.
	VisualReference() interface{}
	SetVisualReference(val interface{})
	// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
	//
	// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCanary
type jsiiProxy_CfnCanary struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCanary) ArtifactConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RunConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) StartCanaryAfterCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryAfterCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) VisualReference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary(scope constructs.Construct, id *string, props *CfnCanaryProps) CfnCanary {
	_init_.Initialize()

	j := jsiiProxy_CfnCanary{}

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary_Override(c CfnCanary, scope constructs.Construct, id *string, props *CfnCanaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCanary) SetArtifactConfig(val interface{}) {
	_jsii_.Set(
		j,
		"artifactConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetArtifactS3Location(val *string) {
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetFailureRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRunConfig(val interface{}) {
	_jsii_.Set(
		j,
		"runConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetStartCanaryAfterCreation(val interface{}) {
	_jsii_.Set(
		j,
		"startCanaryAfterCreation",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSuccessRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetVisualReference(val interface{}) {
	_jsii_.Set(
		j,
		"visualReference",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetVpcConfig(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCanary_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCanary_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCanary_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCanary) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCanary) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCanary) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCanary) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCanary) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCanary) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCanary) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCanary) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCanary) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCanary) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCanary) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCanary) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCanary) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCanary) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   artifactConfigProperty := &artifactConfigProperty{
//   	s3Encryption: &s3EncryptionProperty{
//   		encryptionMode: jsii.String("encryptionMode"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnCanary_ArtifactConfigProperty struct {
	// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
	//
	// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
	S3Encryption interface{} `json:"s3Encryption" yaml:"s3Encryption"`
}

// A structure representing a screenshot that is used as a baseline during visual monitoring comparisons made by the canary.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   baseScreenshotProperty := &baseScreenshotProperty{
//   	screenshotName: jsii.String("screenshotName"),
//
//   	// the properties below are optional
//   	ignoreCoordinates: []*string{
//   		jsii.String("ignoreCoordinates"),
//   	},
//   }
//
type CfnCanary_BaseScreenshotProperty struct {
	// The name of the screenshot.
	//
	// This is generated the first time the canary is run after the `UpdateCanary` operation that specified for this canary to perform visual monitoring.
	ScreenshotName *string `json:"screenshotName" yaml:"screenshotName"`
	// Coordinates that define the part of a screen to ignore during screenshot comparisons.
	//
	// To obtain the coordinates to use here, use the CloudWatch Logs console to draw the boundaries on the screen. For more information, see {LINK}
	IgnoreCoordinates *[]*string `json:"ignoreCoordinates" yaml:"ignoreCoordinates"`
}

// Use this structure to input your script code for the canary.
//
// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   codeProperty := &codeProperty{
//   	handler: jsii.String("handler"),
//
//   	// the properties below are optional
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	script: jsii.String("script"),
//   }
//
type CfnCanary_CodeProperty struct {
	// The entry point to use for the source code when running the canary.
	//
	// For canaries that use the `syn-python-selenium-1.0` runtime or a `syn-nodejs.puppeteer` runtime earlier than `syn-nodejs.puppeteer-3.4` , the handler must be specified as `*fileName* .handler` . For `syn-python-selenium-1.1` , `syn-nodejs.puppeteer-3.4` , and later runtimes, the handler can be specified as `*fileName* . *functionName*` , or you can specify a folder where canary scripts reside as `*folder* / *fileName* . *functionName*` .
	Handler *string `json:"handler" yaml:"handler"`
	// If your canary script is located in S3, specify the bucket name here.
	//
	// The bucket must already exist.
	S3Bucket *string `json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of your script.
	//
	// For more information, see [Working with Amazon S3 Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html) .
	S3Key *string `json:"s3Key" yaml:"s3Key"`
	// The S3 version ID of your script.
	S3ObjectVersion *string `json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// If you input your canary script directly into the canary instead of referring to an S3 location, the value of this parameter is the script in plain text.
	//
	// It can be up to 5 MB.
	Script *string `json:"script" yaml:"script"`
}

// A structure that contains input information for a canary run.
//
// This structure is required.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   runConfigProperty := &runConfigProperty{
//   	activeTracing: jsii.Boolean(false),
//   	environmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	memoryInMb: jsii.Number(123),
//   	timeoutInSeconds: jsii.Number(123),
//   }
//
type CfnCanary_RunConfigProperty struct {
	// Specifies whether this canary is to use active AWS X-Ray tracing when it runs.
	//
	// Active tracing enables this canary run to be displayed in the ServiceLens and X-Ray service maps even if the canary does not hit an endpoint that has X-Ray tracing enabled. Using X-Ray tracing incurs charges. For more information, see [Canaries and X-Ray tracing](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_tracing.html) .
	//
	// You can enable active tracing only for canaries that use version `syn-nodejs-2.0` or later for their canary runtime.
	ActiveTracing interface{} `json:"activeTracing" yaml:"activeTracing"`
	// Specifies the keys and values to use for any environment variables used in the canary script.
	//
	// Use the following format:
	//
	// { "key1" : "value1", "key2" : "value2", ...}
	//
	// Keys must start with a letter and be at least two characters. The total size of your environment variables cannot exceed 4 KB. You can't specify any Lambda reserved environment variables as the keys for your environment variables. For more information about reserved keys, see [Runtime environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime) .
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// The maximum amount of memory that the canary can use while running.
	//
	// This value must be a multiple of 64. The range is 960 to 3008.
	MemoryInMb *float64 `json:"memoryInMb" yaml:"memoryInMb"`
	// How long the canary is allowed to run before it must stop.
	//
	// You can't set this time to be longer than the frequency of the runs of this canary.
	//
	// If you omit this field, the frequency of the canary is used as this value, up to a maximum of 900 seconds.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   s3EncryptionProperty := &s3EncryptionProperty{
//   	encryptionMode: jsii.String("encryptionMode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnCanary_S3EncryptionProperty struct {
	// The encryption method to use for artifacts created by this canary.
	//
	// Specify `SSE_S3` to use server-side encryption (SSE) with an Amazon S3-managed key. Specify `SSE-KMS` to use server-side encryption with a customer-managed AWS KMS key.
	//
	// If you omit this parameter, an AWS -managed AWS KMS key is used.
	EncryptionMode *string `json:"encryptionMode" yaml:"encryptionMode"`
	// The ARN of the customer-managed AWS KMS key to use, if you specify `SSE-KMS` for `EncryptionMode`.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// This structure specifies how often a canary is to make runs and the date and time when it should stop making runs.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   scheduleProperty := &scheduleProperty{
//   	expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	durationInSeconds: jsii.String("durationInSeconds"),
//   }
//
type CfnCanary_ScheduleProperty struct {
	// A `rate` expression or a `cron` expression that defines how often the canary is to run.
	//
	// For a rate expression, The syntax is `rate( *number unit* )` . *unit* can be `minute` , `minutes` , or `hour` .
	//
	// For example, `rate(1 minute)` runs the canary once a minute, `rate(10 minutes)` runs it once every 10 minutes, and `rate(1 hour)` runs it once every hour. You can specify a frequency between `rate(1 minute)` and `rate(1 hour)` .
	//
	// Specifying `rate(0 minute)` or `rate(0 hour)` is a special value that causes the canary to run only once when it is started.
	//
	// Use `cron( *expression* )` to specify a cron expression. You can't schedule a canary to wait for more than a year before running. For information about the syntax for cron expressions, see [Scheduling canary runs using cron](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html) .
	Expression *string `json:"expression" yaml:"expression"`
	// How long, in seconds, for the canary to continue making regular runs according to the schedule in the `Expression` value.
	//
	// If you specify 0, the canary continues making runs until you stop it. If you omit this field, the default of 0 is used.
	DurationInSeconds *string `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
//
// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   vPCConfigProperty := &vPCConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnCanary_VPCConfigProperty struct {
	// The IDs of the security groups for this canary.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets where this canary is to run.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC where this canary is to run.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary.
//
// If you omit this parameter, no changes are made to any baseline screenshots that the canary might be using already.
//
// Visual monitoring is supported only on canaries running the *syn-puppeteer-node-3.2* runtime or later. For more information, see [Visual monitoring](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html) and [Visual monitoring blueprint](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html)
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   visualReferenceProperty := &visualReferenceProperty{
//   	baseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   	// the properties below are optional
//   	baseScreenshots: []interface{}{
//   		&baseScreenshotProperty{
//   			screenshotName: jsii.String("screenshotName"),
//
//   			// the properties below are optional
//   			ignoreCoordinates: []*string{
//   				jsii.String("ignoreCoordinates"),
//   			},
//   		},
//   	},
//   }
//
type CfnCanary_VisualReferenceProperty struct {
	// Specifies which canary run to use the screenshots from as the baseline for future visual monitoring with this canary.
	//
	// Valid values are `nextrun` to use the screenshots from the next run after this update is made, `lastrun` to use the screenshots from the most recent run before this update was made, or the value of `Id` in the [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) from any past run of this canary.
	BaseCanaryRunId *string `json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// An array of screenshots that are used as the baseline for comparisons during visual monitoring.
	BaseScreenshots interface{} `json:"baseScreenshots" yaml:"baseScreenshots"`
}

// Properties for defining a `CfnCanary`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import synthetics "github.com/aws/aws-cdk-go/awscdk/aws_synthetics"
//   cfnCanaryProps := &cfnCanaryProps{
//   	artifactS3Location: jsii.String("artifactS3Location"),
//   	code: &codeProperty{
//   		handler: jsii.String("handler"),
//
//   		// the properties below are optional
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		script: jsii.String("script"),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	name: jsii.String("name"),
//   	runtimeVersion: jsii.String("runtimeVersion"),
//   	schedule: &scheduleProperty{
//   		expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		durationInSeconds: jsii.String("durationInSeconds"),
//   	},
//   	startCanaryAfterCreation: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	artifactConfig: &artifactConfigProperty{
//   		s3Encryption: &s3EncryptionProperty{
//   			encryptionMode: jsii.String("encryptionMode"),
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	failureRetentionPeriod: jsii.Number(123),
//   	runConfig: &runConfigProperty{
//   		activeTracing: jsii.Boolean(false),
//   		environmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		memoryInMb: jsii.Number(123),
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	successRetentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visualReference: &visualReferenceProperty{
//   		baseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   		// the properties below are optional
//   		baseScreenshots: []interface{}{
//   			&baseScreenshotProperty{
//   				screenshotName: jsii.String("screenshotName"),
//
//   				// the properties below are optional
//   				ignoreCoordinates: []*string{
//   					jsii.String("ignoreCoordinates"),
//   				},
//   			},
//   		},
//   	},
//   	vpcConfig: &vPCConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnCanaryProps struct {
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of this canary.
	//
	// Artifacts include the log file, screenshots, and HAR files. Specify the full location path, including `s3://` at the beginning of the path.
	ArtifactS3Location *string `json:"artifactS3Location" yaml:"artifactS3Location"`
	// Use this structure to input your script code for the canary.
	//
	// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
	Code interface{} `json:"code" yaml:"code"`
	// The ARN of the IAM role to be used to run the canary.
	//
	// This role must already exist, and must include `lambda.amazonaws.com` as a principal in the trust policy. The role must also have the following permissions:
	//
	// - `s3:PutObject`
	// - `s3:GetBucketLocation`
	// - `s3:ListAllMyBuckets`
	// - `cloudwatch:PutMetricData`
	// - `logs:CreateLogGroup`
	// - `logs:CreateLogStream`
	// - `logs:PutLogEvents`.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name for this canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
	Name *string `json:"name" yaml:"name"`
	// Specifies the runtime version to use for the canary.
	//
	// For more information about runtime versions, see [Canary Runtime Versions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) .
	RuntimeVersion *string `json:"runtimeVersion" yaml:"runtimeVersion"`
	// A structure that contains information about how often the canary is to run, and when these runs are to stop.
	Schedule interface{} `json:"schedule" yaml:"schedule"`
	// Specify TRUE to have the canary start making runs immediately after it is created.
	//
	// A canary that you create using CloudFormation can't be used to monitor the CloudFormation stack that creates the canary or to roll back that stack if there is a failure.
	StartCanaryAfterCreation interface{} `json:"startCanaryAfterCreation" yaml:"startCanaryAfterCreation"`
	// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig interface{} `json:"artifactConfig" yaml:"artifactConfig"`
	// The number of days to retain data about failed runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod *float64 `json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// A structure that contains input information for a canary run.
	//
	// If you omit this structure, the frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.
	RunConfig interface{} `json:"runConfig" yaml:"runConfig"`
	// The number of days to retain data about successful runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod *float64 `json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// The list of key-value pairs that are associated with the canary.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// If this canary performs visual monitoring by comparing screenshots, this structure contains the ID of the canary run to use as the baseline for screenshots, and the coordinates of any parts of the screen to ignore during the visual monitoring comparison.
	VisualReference interface{} `json:"visualReference" yaml:"visualReference"`
	// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
	//
	// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
	VpcConfig interface{} `json:"vpcConfig" yaml:"vpcConfig"`
}

