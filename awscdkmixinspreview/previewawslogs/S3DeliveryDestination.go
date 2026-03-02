package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an S3 delivery destination for CloudWatch Logs.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//   import logDestinations "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var origin IBucket
//
//
//   destinationAccount := "123456789012"
//   sourceAccount := "234567890123"
//   region := "us-east-1"
//
//   app := awscdk.NewApp()
//
//   destStack := awscdk.NewStack(app, jsii.String("destination-stack"), &StackProps{
//   	Env: &Environment{
//   		Account: destinationAccount,
//   		Region: jsii.String(*Region),
//   	},
//   })
//
//   // Create destination bucket
//   destBucket := s3.NewBucket(destStack, jsii.String("DeliveryBucket"))
//   logDestinations.NewS3DeliveryDestination(destStack, jsii.String("Destination"), &S3DeliveryDestinationProps{
//   	Bucket: destBucket,
//   	SourceAccountId: sourceAccount,
//   })
//
//   sourceStack := awscdk.NewStack(app, jsii.String("source-stack"), &StackProps{
//   	Env: &Environment{
//   		Account: sourceAccount,
//   		Region: jsii.String(*Region),
//   	},
//   })
//   distribution := cloudfront.NewDistribution(sourceStack, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
//   	},
//   })
//
//   destination := logs.CfnDeliveryDestination_FromDeliveryDestinationArn(sourceStack, jsii.String("Destination"), jsii.String("arn of Delivery Destination in destinationAccount"))
//
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
//
// Experimental.
type S3DeliveryDestination interface {
	awslogs.CfnDeliveryDestination
	// The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.
	// Experimental.
	AttrArn() *string
	// Tag Manager which manages the tags for this resource.
	// Experimental.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	// Experimental.
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// An IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	// Experimental.
	DeliveryDestinationPolicy() interface{}
	// Experimental.
	SetDeliveryDestinationPolicy(val interface{})
	// A reference to a DeliveryDestination resource.
	// Experimental.
	DeliveryDestinationRef() *interfacesawslogs.DeliveryDestinationReference
	// Displays whether this delivery destination is CloudWatch Logs, Amazon S3, Firehose, or X-Ray.
	// Experimental.
	DeliveryDestinationType() *string
	// Experimental.
	SetDeliveryDestinationType(val *string)
	// The ARN of the AWS destination that this delivery destination represents.
	// Experimental.
	DestinationResourceArn() *string
	// Experimental.
	SetDestinationResourceArn(val *string)
	// Experimental.
	Env() *interfaces.ResourceEnvironment
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
	// The name of this delivery destination.
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The format of the logs that are sent to this delivery destination.
	// Experimental.
	OutputFormat() *string
	// Experimental.
	SetOutputFormat(val *string)
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
	// An array of key-value pairs to apply to the delivery destination.
	// Experimental.
	Tags() *[]*awscdk.CfnTag
	// Experimental.
	SetTags(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Experimental.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	// Experimental.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	// Experimental.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Experimental.
	RemoveDependency(target awscdk.CfnResource)
	// Experimental.
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	// Experimental.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Experimental.
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for S3DeliveryDestination
type jsiiProxy_S3DeliveryDestination struct {
	internal.Type__awslogsCfnDeliveryDestination
}

func (j *jsiiProxy_S3DeliveryDestination) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) DeliveryDestinationPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryDestinationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) DeliveryDestinationRef() *interfacesawslogs.DeliveryDestinationReference {
	var returns *interfacesawslogs.DeliveryDestinationReference
	_jsii_.Get(
		j,
		"deliveryDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) DeliveryDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) DestinationResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DeliveryDestination) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3DeliveryDestination(scope constructs.Construct, id *string, props *S3DeliveryDestinationProps) S3DeliveryDestination {
	_init_.Initialize()

	if err := validateNewS3DeliveryDestinationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3DeliveryDestination{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3DeliveryDestination_Override(s S3DeliveryDestination, scope constructs.Construct, id *string, props *S3DeliveryDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetDeliveryDestinationPolicy(val interface{}) {
	if err := j.validateSetDeliveryDestinationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryDestinationPolicy",
		val,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetDeliveryDestinationType(val *string) {
	_jsii_.Set(
		j,
		"deliveryDestinationType",
		val,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetDestinationResourceArn(val *string) {
	_jsii_.Set(
		j,
		"destinationResourceArn",
		val,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetOutputFormat(val *string) {
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_S3DeliveryDestination)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Experimental.
func S3DeliveryDestination_ArnForDeliveryDestination(resource interfacesawslogs.IDeliveryDestinationRef) *string {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_ArnForDeliveryDestinationParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"arnForDeliveryDestination",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IDeliveryDestinationRef from an ARN.
// Experimental.
func S3DeliveryDestination_FromDeliveryDestinationArn(scope constructs.Construct, id *string, arn *string) interfacesawslogs.IDeliveryDestinationRef {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_FromDeliveryDestinationArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns interfacesawslogs.IDeliveryDestinationRef

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"fromDeliveryDestinationArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new IDeliveryDestinationRef from a deliveryDestinationName.
// Experimental.
func S3DeliveryDestination_FromDeliveryDestinationName(scope constructs.Construct, id *string, deliveryDestinationName *string) interfacesawslogs.IDeliveryDestinationRef {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_FromDeliveryDestinationNameParameters(scope, id, deliveryDestinationName); err != nil {
		panic(err)
	}
	var returns interfacesawslogs.IDeliveryDestinationRef

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"fromDeliveryDestinationName",
		[]interface{}{scope, id, deliveryDestinationName},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnDeliveryDestination.
// Experimental.
func S3DeliveryDestination_IsCfnDeliveryDestination(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_IsCfnDeliveryDestinationParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"isCfnDeliveryDestination",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func S3DeliveryDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
// Experimental.
func S3DeliveryDestination_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func S3DeliveryDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3DeliveryDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func S3DeliveryDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_logs.S3DeliveryDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) AddDeletionOverride(path *string) {
	if err := s.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddDependency(target awscdk.CfnResource) {
	if err := s.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependency",
		[]interface{}{target},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddDependsOn(target awscdk.CfnResource) {
	if err := s.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddMetadata(key *string, value interface{}) {
	if err := s.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddPropertyDeletionOverride(propertyPath *string) {
	if err := s.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := s.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := s.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := s.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		s,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) GetMetadata(key *string) interface{} {
	if err := s.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) Inspect(inspector awscdk.TreeInspector) {
	if err := s.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"inspect",
		[]interface{}{inspector},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		s,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		s,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) RemoveDependency(target awscdk.CfnResource) {
	if err := s.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"removeDependency",
		[]interface{}{target},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := s.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := s.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DeliveryDestination) ValidateProperties(_properties interface{}) {
	if err := s.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (s *jsiiProxy_S3DeliveryDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

