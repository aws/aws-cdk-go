package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::SageMaker::App`.
//
// Creates a running app for the specified UserProfile. Supported apps are `JupyterServer` and `KernelGateway` . This operation is automatically invoked by Amazon SageMaker Studio upon access to the associated Domain, and when new kernel configurations are selected by the user. A user may have multiple Apps active simultaneously.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApp := awscdk.Aws_sagemaker.NewCfnApp(this, jsii.String("MyCfnApp"), &cfnAppProps{
//   	appName: jsii.String("appName"),
//   	appType: jsii.String("appType"),
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	resourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the app.
	AppName() *string
	SetAppName(val *string)
	// The type of app.
	//
	// *Allowed Values* : `JupyterServer | KernelGateway | RSessionGateway | RStudioServerPro | TensorBoard | Canvas`.
	AppType() *string
	SetAppType(val *string)
	// The Amazon Resource Name (ARN) of the app, such as `arn:aws:sagemaker:us-west-2:account-id:app/my-app-name` .
	AttrAppArn() *string
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
	// The domain ID.
	DomainId() *string
	SetDomainId(val *string)
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
	// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	ResourceSpec() interface{}
	SetResourceSpec(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user profile name.
	UserProfileName() *string
	SetUserProfileName(val *string)
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

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AppType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) ResourceSpec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UserProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userProfileName",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::App`.
func NewCfnApp(scope awscdk.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::App`.
func NewCfnApp_Override(c CfnApp, scope awscdk.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp) SetAppName(val *string) {
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetAppType(val *string) {
	_jsii_.Set(
		j,
		"appType",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetDomainId(val *string) {
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetResourceSpec(val interface{}) {
	_jsii_.Set(
		j,
		"resourceSpec",
		val,
	)
}

func (j *jsiiProxy_CfnApp) SetUserProfileName(val *string) {
	_jsii_.Set(
		j,
		"userProfileName",
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
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnApp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApp) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApp) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecProperty := &resourceSpecProperty{
//   	instanceType: jsii.String("instanceType"),
//   	sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   	sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   }
//
type CfnApp_ResourceSpecProperty struct {
	// The instance type that the image version runs on.
	//
	// > JupyterServer Apps only support the `system` value. KernelGateway Apps do not support the `system` value, but support all other values for available instance types.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the SageMaker image that the image version belongs to.
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}

// A CloudFormation `AWS::SageMaker::AppImageConfig`.
//
// Creates a configuration for running a SageMaker image as a KernelGateway app. The configuration specifies the Amazon Elastic File System (EFS) storage volume on the image, and a list of the kernels in the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppImageConfig := awscdk.Aws_sagemaker.NewCfnAppImageConfig(this, jsii.String("MyCfnAppImageConfig"), &cfnAppImageConfigProps{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//
//   	// the properties below are optional
//   	kernelGatewayImageConfig: &kernelGatewayImageConfigProperty{
//   		kernelSpecs: []interface{}{
//   			&kernelSpecProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				displayName: jsii.String("displayName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		fileSystemConfig: &fileSystemConfigProperty{
//   			defaultGid: jsii.Number(123),
//   			defaultUid: jsii.Number(123),
//   			mountPath: jsii.String("mountPath"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAppImageConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the AppImageConfig.
	//
	// Must be unique to your account.
	AppImageConfigName() *string
	SetAppImageConfigName(val *string)
	// The Amazon Resource Name (ARN) of the AppImageConfig, such as `arn:aws:sagemaker:us-west-2:account-id:app-image-config/my-app-image-config-name` .
	AttrAppImageConfigArn() *string
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
	// The configuration for the file system and kernels in the SageMaker image.
	KernelGatewayImageConfig() interface{}
	SetKernelGatewayImageConfig(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnAppImageConfig
type jsiiProxy_CfnAppImageConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAppImageConfig) AppImageConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appImageConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) AttrAppImageConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppImageConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) KernelGatewayImageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kernelGatewayImageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::AppImageConfig`.
func NewCfnAppImageConfig(scope awscdk.Construct, id *string, props *CfnAppImageConfigProps) CfnAppImageConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnAppImageConfig{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::AppImageConfig`.
func NewCfnAppImageConfig_Override(c CfnAppImageConfig, scope awscdk.Construct, id *string, props *CfnAppImageConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAppImageConfig) SetAppImageConfigName(val *string) {
	_jsii_.Set(
		j,
		"appImageConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnAppImageConfig) SetKernelGatewayImageConfig(val interface{}) {
	_jsii_.Set(
		j,
		"kernelGatewayImageConfig",
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
func CfnAppImageConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAppImageConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAppImageConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppImageConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnAppImageConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAppImageConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAppImageConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAppImageConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppImageConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &fileSystemConfigProperty{
//   	defaultGid: jsii.Number(123),
//   	defaultUid: jsii.Number(123),
//   	mountPath: jsii.String("mountPath"),
//   }
//
type CfnAppImageConfig_FileSystemConfigProperty struct {
	// The default POSIX group ID (GID).
	//
	// If not specified, defaults to `100` .
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// The default POSIX user ID (UID).
	//
	// If not specified, defaults to `1000` .
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// The path within the image to mount the user's EFS home directory.
	//
	// The directory should be empty. If not specified, defaults to * /home/sagemaker-user* .
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayImageConfigProperty := &kernelGatewayImageConfigProperty{
//   	kernelSpecs: []interface{}{
//   		&kernelSpecProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			displayName: jsii.String("displayName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	fileSystemConfig: &fileSystemConfigProperty{
//   		defaultGid: jsii.Number(123),
//   		defaultUid: jsii.Number(123),
//   		mountPath: jsii.String("mountPath"),
//   	},
//   }
//
type CfnAppImageConfig_KernelGatewayImageConfigProperty struct {
	// The specification of the Jupyter kernels in the image.
	KernelSpecs interface{} `field:"required" json:"kernelSpecs" yaml:"kernelSpecs"`
	// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
	FileSystemConfig interface{} `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

// The specification of a Jupyter kernel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelSpecProperty := &kernelSpecProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   }
//
type CfnAppImageConfig_KernelSpecProperty struct {
	// The name of the Jupyter kernel in the image.
	//
	// This value is case sensitive.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The display name of the kernel.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

// Properties for defining a `CfnAppImageConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppImageConfigProps := &cfnAppImageConfigProps{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//
//   	// the properties below are optional
//   	kernelGatewayImageConfig: &kernelGatewayImageConfigProperty{
//   		kernelSpecs: []interface{}{
//   			&kernelSpecProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				displayName: jsii.String("displayName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		fileSystemConfig: &fileSystemConfigProperty{
//   			defaultGid: jsii.Number(123),
//   			defaultUid: jsii.Number(123),
//   			mountPath: jsii.String("mountPath"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppImageConfigProps struct {
	// The name of the AppImageConfig.
	//
	// Must be unique to your account.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The configuration for the file system and kernels in the SageMaker image.
	KernelGatewayImageConfig interface{} `field:"optional" json:"kernelGatewayImageConfig" yaml:"kernelGatewayImageConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &cfnAppProps{
//   	appName: jsii.String("appName"),
//   	appType: jsii.String("appType"),
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	resourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppProps struct {
	// The name of the app.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The type of app.
	//
	// *Allowed Values* : `JupyterServer | KernelGateway | RSessionGateway | RStudioServerPro | TensorBoard | Canvas`.
	AppType *string `field:"required" json:"appType" yaml:"appType"`
	// The domain ID.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	ResourceSpec interface{} `field:"optional" json:"resourceSpec" yaml:"resourceSpec"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::CodeRepository`.
//
// Creates a Git repository as a resource in your SageMaker account. You can associate the repository with notebook instances so that you can use Git source control for the notebooks you create. The Git repository is a resource in your SageMaker account, so it can be associated with more than one notebook instance, and it persists independently from the lifecycle of any notebook instances it is associated with.
//
// The repository can be hosted either in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeRepository := awscdk.Aws_sagemaker.NewCfnCodeRepository(this, jsii.String("MyCfnCodeRepository"), &cfnCodeRepositoryProps{
//   	gitConfig: &gitConfigProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//
//   		// the properties below are optional
//   		branch: jsii.String("branch"),
//   		secretArn: jsii.String("secretArn"),
//   	},
//
//   	// the properties below are optional
//   	codeRepositoryName: jsii.String("codeRepositoryName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCodeRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the code repository, such as `myCodeRepo` .
	AttrCodeRepositoryName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The name of the Git repository.
	CodeRepositoryName() *string
	SetCodeRepositoryName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Configuration details for the Git repository, including the URL where it is located and the ARN of the AWS Secrets Manager secret that contains the credentials used to access the repository.
	GitConfig() interface{}
	SetGitConfig(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// List of tags for Code Repository.
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

// The jsii proxy struct for CfnCodeRepository
type jsiiProxy_CfnCodeRepository struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCodeRepository) AttrCodeRepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCodeRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) CodeRepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) GitConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeRepository) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::CodeRepository`.
func NewCfnCodeRepository(scope awscdk.Construct, id *string, props *CfnCodeRepositoryProps) CfnCodeRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeRepository{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::CodeRepository`.
func NewCfnCodeRepository_Override(c CfnCodeRepository, scope awscdk.Construct, id *string, props *CfnCodeRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCodeRepository) SetCodeRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"codeRepositoryName",
		val,
	)
}

func (j *jsiiProxy_CfnCodeRepository) SetGitConfig(val interface{}) {
	_jsii_.Set(
		j,
		"gitConfig",
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
func CfnCodeRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCodeRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCodeRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCodeRepository_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnCodeRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCodeRepository) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCodeRepository) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCodeRepository) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCodeRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCodeRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCodeRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCodeRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCodeRepository) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCodeRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCodeRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCodeRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCodeRepository) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCodeRepository) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCodeRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeRepository) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies configuration details for a Git repository in your AWS account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitConfigProperty := &gitConfigProperty{
//   	repositoryUrl: jsii.String("repositoryUrl"),
//
//   	// the properties below are optional
//   	branch: jsii.String("branch"),
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnCodeRepository_GitConfigProperty struct {
	// The URL where the Git repository is located.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The default branch for the Git repository.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository.
	//
	// The secret must have a staging label of `AWSCURRENT` and must be in the following format:
	//
	// `{"username": *UserName* , "password": *Password* }`.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

// Properties for defining a `CfnCodeRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeRepositoryProps := &cfnCodeRepositoryProps{
//   	gitConfig: &gitConfigProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//
//   		// the properties below are optional
//   		branch: jsii.String("branch"),
//   		secretArn: jsii.String("secretArn"),
//   	},
//
//   	// the properties below are optional
//   	codeRepositoryName: jsii.String("codeRepositoryName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCodeRepositoryProps struct {
	// Configuration details for the Git repository, including the URL where it is located and the ARN of the AWS Secrets Manager secret that contains the credentials used to access the repository.
	GitConfig interface{} `field:"required" json:"gitConfig" yaml:"gitConfig"`
	// The name of the Git repository.
	CodeRepositoryName *string `field:"optional" json:"codeRepositoryName" yaml:"codeRepositoryName"`
	// List of tags for Code Repository.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::DataQualityJobDefinition`.
//
// Creates a definition for a job that monitors data quality and drift. For information about model monitor, see [Amazon SageMaker Model Monitor](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataQualityJobDefinition := awscdk.Aws_sagemaker.NewCfnDataQualityJobDefinition(this, jsii.String("MyCfnDataQualityJobDefinition"), &cfnDataQualityJobDefinitionProps{
//   	dataQualityAppSpecification: &dataQualityAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	dataQualityJobInput: &dataQualityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	dataQualityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	dataQualityBaselineConfig: &dataQualityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   		statisticsResource: &statisticsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDataQualityJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the job definition was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the job definition.
	AttrJobDefinitionArn() *string
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
	// Specifies the container that runs the monitoring job.
	DataQualityAppSpecification() interface{}
	SetDataQualityAppSpecification(val interface{})
	// Configures the constraints and baselines for the monitoring job.
	DataQualityBaselineConfig() interface{}
	SetDataQualityBaselineConfig(val interface{})
	// A list of inputs for the monitoring job.
	//
	// Currently endpoints are supported as monitoring inputs.
	DataQualityJobInput() interface{}
	SetDataQualityJobInput(val interface{})
	// The output configuration for monitoring jobs.
	DataQualityJobOutputConfig() interface{}
	SetDataQualityJobOutputConfig(val interface{})
	// The name for the monitoring job definition.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	// Identifies the resources to deploy for a monitoring job.
	JobResources() interface{}
	SetJobResources(val interface{})
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
	// Specifies networking configuration for the monitoring job.
	NetworkConfig() interface{}
	SetNetworkConfig(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition() interface{}
	SetStoppingCondition(val interface{})
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnDataQualityJobDefinition
type jsiiProxy_CfnDataQualityJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) AttrJobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrJobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) DataQualityAppSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) DataQualityBaselineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) DataQualityJobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) DataQualityJobOutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataQualityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) JobResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) NetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) StoppingCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::DataQualityJobDefinition`.
func NewCfnDataQualityJobDefinition(scope awscdk.Construct, id *string, props *CfnDataQualityJobDefinitionProps) CfnDataQualityJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnDataQualityJobDefinition{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::DataQualityJobDefinition`.
func NewCfnDataQualityJobDefinition_Override(c CfnDataQualityJobDefinition, scope awscdk.Construct, id *string, props *CfnDataQualityJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetDataQualityAppSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"dataQualityAppSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetDataQualityBaselineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dataQualityBaselineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetDataQualityJobInput(val interface{}) {
	_jsii_.Set(
		j,
		"dataQualityJobInput",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetDataQualityJobOutputConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dataQualityJobOutputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetJobResources(val interface{}) {
	_jsii_.Set(
		j,
		"jobResources",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetNetworkConfig(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataQualityJobDefinition) SetStoppingCondition(val interface{}) {
	_jsii_.Set(
		j,
		"stoppingCondition",
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
func CfnDataQualityJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataQualityJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataQualityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataQualityJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnDataQualityJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataQualityJobDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration for the cluster of resources used to run the processing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnDataQualityJobDefinition_ClusterConfigProperty struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// `CfnDataQualityJobDefinition.ClusterConfigProperty.InstanceType`.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

// The constraints resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnDataQualityJobDefinition_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Information about the container that a data quality monitoring job runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityAppSpecificationProperty := &dataQualityAppSpecificationProperty{
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	containerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	containerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
type CfnDataQualityJobDefinition_DataQualityAppSpecificationProperty struct {
	// The container image that the data quality monitoring job runs.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The arguments to send to the container that the monitoring job runs.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// The entrypoint for a container used to run a monitoring job.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Sets the environment variables in the container that the monitoring job runs.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

// Configuration for monitoring constraints and monitoring statistics.
//
// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityBaselineConfigProperty := &dataQualityBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   	statisticsResource: &statisticsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_DataQualityBaselineConfigProperty struct {
	// The name of the job that performs baselining for the data quality monitoring job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// Configuration for monitoring constraints and monitoring statistics.
	//
	// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

// The input for the data quality monitoring job.
//
// Currently endpoints are supported for input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityJobInputProperty := &dataQualityJobInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// `CfnDataQualityJobDefinition.DataQualityJobInputProperty.EndpointInput`.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
}

// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnDataQualityJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDataQualityJobDefinition_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &monitoringResourcesProperty{
//   	clusterConfig: &clusterConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		volumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigProperty := &networkConfigProperty{
//   	enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnDataQualityJobDefinition_NetworkConfigProperty struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose `True` to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// The Amazon S3 storage location where the results of a monitoring job are saved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnDataQualityJobDefinition_S3OutputProperty struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

// The statistics resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticsResourceProperty := &statisticsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnDataQualityJobDefinition_StatisticsResourceProperty struct {
	// The Amazon S3 URI for the statistics resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Specifies a limit to how long a model training job or model compilation job can run.
//
// It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.
//
// To stop a training job, SageMaker sends the algorithm the `SIGTERM` signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.
//
// The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with `CreateModel` .
//
// > The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stoppingConditionProperty := &stoppingConditionProperty{
//   	maxRuntimeInSeconds: jsii.Number(123),
//   }
//
type CfnDataQualityJobDefinition_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Properties for defining a `CfnDataQualityJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataQualityJobDefinitionProps := &cfnDataQualityJobDefinitionProps{
//   	dataQualityAppSpecification: &dataQualityAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	dataQualityJobInput: &dataQualityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	dataQualityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	dataQualityBaselineConfig: &dataQualityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   		statisticsResource: &statisticsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataQualityJobDefinitionProps struct {
	// Specifies the container that runs the monitoring job.
	DataQualityAppSpecification interface{} `field:"required" json:"dataQualityAppSpecification" yaml:"dataQualityAppSpecification"`
	// A list of inputs for the monitoring job.
	//
	// Currently endpoints are supported as monitoring inputs.
	DataQualityJobInput interface{} `field:"required" json:"dataQualityJobInput" yaml:"dataQualityJobInput"`
	// The output configuration for monitoring jobs.
	DataQualityJobOutputConfig interface{} `field:"required" json:"dataQualityJobOutputConfig" yaml:"dataQualityJobOutputConfig"`
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configures the constraints and baselines for the monitoring job.
	DataQualityBaselineConfig interface{} `field:"optional" json:"dataQualityBaselineConfig" yaml:"dataQualityBaselineConfig"`
	// The name for the monitoring job definition.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Specifies networking configuration for the monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::Device`.
//
// The `AWS::SageMaker::Device` resource is an Amazon SageMaker resource type that allows you to register your Devices against an existing SageMaker Edge Manager DeviceFleet. Each device must be listed individually in the CFN specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDevice := awscdk.Aws_sagemaker.NewCfnDevice(this, jsii.String("MyCfnDevice"), &cfnDeviceProps{
//   	deviceFleetName: jsii.String("deviceFleetName"),
//
//   	// the properties below are optional
//   	device: &deviceProperty{
//   		deviceName: jsii.String("deviceName"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		iotThingName: jsii.String("iotThingName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDevice interface {
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
	// Edge device you want to create.
	Device() interface{}
	SetDevice(val interface{})
	// The name of the fleet the device belongs to.
	DeviceFleetName() *string
	SetDeviceFleetName(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs that contain metadata to help you categorize and organize your devices.
	//
	// Each tag consists of a key and a value, both of which you define.
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

// The jsii proxy struct for CfnDevice
type jsiiProxy_CfnDevice struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDevice) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) Device() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"device",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) DeviceFleetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceFleetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevice) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Device`.
func NewCfnDevice(scope awscdk.Construct, id *string, props *CfnDeviceProps) CfnDevice {
	_init_.Initialize()

	j := jsiiProxy_CfnDevice{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDevice",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Device`.
func NewCfnDevice_Override(c CfnDevice, scope awscdk.Construct, id *string, props *CfnDeviceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDevice",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDevice) SetDevice(val interface{}) {
	_jsii_.Set(
		j,
		"device",
		val,
	)
}

func (j *jsiiProxy_CfnDevice) SetDeviceFleetName(val *string) {
	_jsii_.Set(
		j,
		"deviceFleetName",
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
func CfnDevice_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDevice",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDevice_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDevice",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDevice_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDevice",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDevice_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnDevice",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDevice) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDevice) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDevice) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDevice) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDevice) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDevice) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDevice) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDevice) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDevice) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDevice) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDevice) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDevice) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDevice) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDevice) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevice) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information of a particular device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	deviceName: jsii.String("deviceName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	iotThingName: jsii.String("iotThingName"),
//   }
//
type CfnDevice_DeviceProperty struct {
	// The name of the device.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Description of the device.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Internet of Things (IoT) object name.
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

// A CloudFormation `AWS::SageMaker::DeviceFleet`.
//
// The `AWS::SageMaker::DeviceFleet` resource is an Amazon SageMaker resource type that allows you to create a DeviceFleet that manages your SageMaker Edge Manager Devices. You must register your devices against the `DeviceFleet` separately.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceFleet := awscdk.Aws_sagemaker.NewCfnDeviceFleet(this, jsii.String("MyCfnDeviceFleet"), &cfnDeviceFleetProps{
//   	deviceFleetName: jsii.String("deviceFleetName"),
//   	outputConfig: &edgeOutputConfigProperty{
//   		s3OutputLocation: jsii.String("s3OutputLocation"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDeviceFleet interface {
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
	// A description of the fleet.
	Description() *string
	SetDescription(val *string)
	// Name of the device fleet.
	DeviceFleetName() *string
	SetDeviceFleetName(val *string)
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
	// The output configuration for storing sample data collected by the fleet.
	OutputConfig() interface{}
	SetOutputConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs that contain metadata to help you categorize and organize your device fleets.
	//
	// Each tag consists of a key and a value, both of which you define.
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

// The jsii proxy struct for CfnDeviceFleet
type jsiiProxy_CfnDeviceFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeviceFleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) DeviceFleetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceFleetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) OutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::DeviceFleet`.
func NewCfnDeviceFleet(scope awscdk.Construct, id *string, props *CfnDeviceFleetProps) CfnDeviceFleet {
	_init_.Initialize()

	j := jsiiProxy_CfnDeviceFleet{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::DeviceFleet`.
func NewCfnDeviceFleet_Override(c CfnDeviceFleet, scope awscdk.Construct, id *string, props *CfnDeviceFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeviceFleet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDeviceFleet) SetDeviceFleetName(val *string) {
	_jsii_.Set(
		j,
		"deviceFleetName",
		val,
	)
}

func (j *jsiiProxy_CfnDeviceFleet) SetOutputConfig(val interface{}) {
	_jsii_.Set(
		j,
		"outputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDeviceFleet) SetRoleArn(val *string) {
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
func CfnDeviceFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeviceFleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeviceFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeviceFleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnDeviceFleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeviceFleet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeviceFleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDeviceFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeviceFleet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The output configuration for storing sample data collected by the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   edgeOutputConfigProperty := &edgeOutputConfigProperty{
//   	s3OutputLocation: jsii.String("s3OutputLocation"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDeviceFleet_EdgeOutputConfigProperty struct {
	// The Amazon Simple Storage (S3) bucket URI.
	S3OutputLocation *string `field:"required" json:"s3OutputLocation" yaml:"s3OutputLocation"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job.
	//
	// If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Properties for defining a `CfnDeviceFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceFleetProps := &cfnDeviceFleetProps{
//   	deviceFleetName: jsii.String("deviceFleetName"),
//   	outputConfig: &edgeOutputConfigProperty{
//   		s3OutputLocation: jsii.String("s3OutputLocation"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeviceFleetProps struct {
	// Name of the device fleet.
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
	// The output configuration for storing sample data collected by the fleet.
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A description of the fleet.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs that contain metadata to help you categorize and organize your device fleets.
	//
	// Each tag consists of a key and a value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &cfnDeviceProps{
//   	deviceFleetName: jsii.String("deviceFleetName"),
//
//   	// the properties below are optional
//   	device: &deviceProperty{
//   		deviceName: jsii.String("deviceName"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		iotThingName: jsii.String("iotThingName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeviceProps struct {
	// The name of the fleet the device belongs to.
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
	// Edge device you want to create.
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// An array of key-value pairs that contain metadata to help you categorize and organize your devices.
	//
	// Each tag consists of a key and a value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::Domain`.
//
// Creates a `Domain` used by Amazon SageMaker Studio. A domain consists of an associated Amazon Elastic File System (EFS) volume, a list of authorized users, and a variety of security, application, policy, and Amazon Virtual Private Cloud (VPC) configurations. An AWS account is limited to one domain per region. Users within a domain can share notebook files and other artifacts with each other.
//
// *EFS storage*
//
// When a domain is created, an EFS volume is created for use by all of the users within the domain. Each user receives a private home directory within the EFS volume for notebooks, Git repositories, and data files.
//
// SageMaker uses the AWS Key Management Service ( AWS KMS) to encrypt the EFS volume attached to the domain with an AWS managed key by default. For more control, you can specify a customer managed key. For more information, see [Protect Data at Rest Using Encryption](https://docs.aws.amazon.com/sagemaker/latest/dg/encryption-at-rest.html) .
//
// *VPC configuration*
//
// All SageMaker Studio traffic between the domain and the EFS volume is through the specified VPC and subnets. For other Studio traffic, you can specify the `AppNetworkAccessType` parameter. `AppNetworkAccessType` corresponds to the network access type that you choose when you onboard to Studio. The following options are available:
//
// - `PublicInternetOnly` - Non-EFS traffic goes through a VPC managed by Amazon SageMaker, which allows internet access. This is the default value.
// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets. Internet access is disabled by default. To allow internet access, you must specify a NAT gateway.
//
// When internet access is disabled, you won't be able to run a Studio notebook or to train or host models unless your VPC has an interface endpoint to the SageMaker API and runtime or a NAT gateway and your security groups allow outbound connections.
//
// > NFS traffic over TCP on port 2049 needs to be allowed in both inbound and outbound rules in order to launch a SageMaker Studio app successfully.
//
// For more information, see [Connect SageMaker Studio Notebooks to Resources in a VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-notebooks-and-internet-access.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomain := awscdk.Aws_sagemaker.NewCfnDomain(this, jsii.String("MyCfnDomain"), &cfnDomainProps{
//   	authMode: jsii.String("authMode"),
//   	defaultUserSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   	domainName: jsii.String("domainName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	appNetworkAccessType: jsii.String("appNetworkAccessType"),
//   	appSecurityGroupManagement: jsii.String("appSecurityGroupManagement"),
//   	domainSettings: &domainSettingsProperty{
//   		rStudioServerProDomainSettings: &rStudioServerProDomainSettingsProperty{
//   			domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   			// the properties below are optional
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   			rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDomain interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly` .
	//
	// - `PublicInternetOnly` - Non-EFS traffic is through a VPC managed by Amazon SageMaker , which allows direct internet access
	// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets
	//
	// *Valid Values* : `PublicInternetOnly | VpcOnly`.
	AppNetworkAccessType() *string
	SetAppNetworkAccessType(val *string)
	// The entity that creates and manages the required security groups for inter-app communication in `VpcOnly` mode.
	//
	// Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and `DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided.
	AppSecurityGroupManagement() *string
	SetAppSecurityGroupManagement(val *string)
	// The Amazon Resource Name (ARN) of the Domain, such as `arn:aws:sagemaker:us-west-2:account-id:domain/my-domain-name` .
	AttrDomainArn() *string
	// The Domain ID.
	AttrDomainId() *string
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	AttrHomeEfsFileSystemId() *string
	// The ID of the security group that authorizes traffic between the `RSessionGateway` apps and the `RStudioServerPro` app.
	AttrSecurityGroupIdForDomainBoundary() *string
	// The AWS SSO managed application instance ID.
	AttrSingleSignOnManagedApplicationInstanceId() *string
	// The URL for the Domain.
	AttrUrl() *string
	// The mode of authentication that members use to access the Domain.
	//
	// *Valid Values* : `SSO | IAM`.
	AuthMode() *string
	SetAuthMode(val *string)
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
	// The default user settings.
	DefaultUserSettings() interface{}
	SetDefaultUserSettings(val interface{})
	// The domain name.
	DomainName() *string
	SetDomainName(val *string)
	// A collection of settings that apply to the `SageMaker Domain` .
	//
	// These settings are specified through the `CreateDomain` API call.
	DomainSettings() interface{}
	SetDomainSettings(val interface{})
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default.
	//
	// For more control, specify a customer managed CMK.
	//
	// *Length Constraints* : Maximum length of 2048.
	//
	// *Pattern* : `.*`
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The VPC subnets that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Array members* : Minimum number of 1 item. Maximum number of 16 items.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// Tags to associated with the Domain.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Tags that you specify for the Domain are also added to all apps that are launched in the Domain.
	//
	// *Array members* : Minimum number of 0 items. Maximum number of 50 items.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	VpcId() *string
	SetVpcId(val *string)
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

// The jsii proxy struct for CfnDomain
type jsiiProxy_CfnDomain struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomain) AppNetworkAccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNetworkAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AppSecurityGroupManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSecurityGroupManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrHomeEfsFileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHomeEfsFileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrSecurityGroupIdForDomainBoundary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSecurityGroupIdForDomainBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrSingleSignOnManagedApplicationInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSingleSignOnManagedApplicationInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) AuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DefaultUserSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultUserSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) DomainSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomain) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Domain`.
func NewCfnDomain(scope awscdk.Construct, id *string, props *CfnDomainProps) CfnDomain {
	_init_.Initialize()

	j := jsiiProxy_CfnDomain{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Domain`.
func NewCfnDomain_Override(c CfnDomain, scope awscdk.Construct, id *string, props *CfnDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnDomain",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomain) SetAppNetworkAccessType(val *string) {
	_jsii_.Set(
		j,
		"appNetworkAccessType",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAppSecurityGroupManagement(val *string) {
	_jsii_.Set(
		j,
		"appSecurityGroupManagement",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetAuthMode(val *string) {
	_jsii_.Set(
		j,
		"authMode",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDefaultUserSettings(val interface{}) {
	_jsii_.Set(
		j,
		"defaultUserSettings",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetDomainSettings(val interface{}) {
	_jsii_.Set(
		j,
		"domainSettings",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnDomain) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
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
func CfnDomain_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDomain",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDomain_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDomain",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomain_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnDomain",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomain) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomain) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomain) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomain) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomain) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomain) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDomain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDomain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomain) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDomain) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomain) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A custom SageMaker image.
//
// For more information, see [Bring your own SageMaker image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customImageProperty := &customImageProperty{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//   	imageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	imageVersionNumber: jsii.Number(123),
//   }
//
type CfnDomain_CustomImageProperty struct {
	// The name of the AppImageConfig.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage.
	//
	// Must be unique to your account.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

// A collection of settings that apply to the `SageMaker Domain` .
//
// These settings are specified through the `CreateDomain` API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSettingsProperty := &domainSettingsProperty{
//   	rStudioServerProDomainSettings: &rStudioServerProDomainSettingsProperty{
//   		domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   		// the properties below are optional
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   		rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type CfnDomain_DomainSettingsProperty struct {
	// A collection of settings that configure the `RStudioServerPro` Domain-level app.
	RStudioServerProDomainSettings interface{} `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the `Domain` uses for communication between Domain-level apps and user apps.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

// The JupyterServer app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jupyterServerAppSettingsProperty := &jupyterServerAppSettingsProperty{
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnDomain_JupyterServerAppSettingsProperty struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
	//
	// If you use the `LifecycleConfigArns` parameter, then this parameter is also required.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

// The KernelGateway app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayAppSettingsProperty := &kernelGatewayAppSettingsProperty{
//   	customImages: []interface{}{
//   		&customImageProperty{
//   			appImageConfigName: jsii.String("appImageConfigName"),
//   			imageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			imageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnDomain_KernelGatewayAppSettingsProperty struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app.
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.
	//
	// > The Amazon SageMaker Studio UI does not use the default instance type value set here. The default instance type set here is used when Apps are created using the AWS Command Line Interface or AWS CloudFormation and the instance type parameter value is not passed.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

// A collection of settings that configure user interaction with the `RStudioServerPro` app.
//
// `RStudioServerProAppSettings` cannot be updated. The `RStudioServerPro` app must be deleted and a new one created to make any changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProAppSettingsProperty := &rStudioServerProAppSettingsProperty{
//   	accessStatus: jsii.String("accessStatus"),
//   	userGroup: jsii.String("userGroup"),
//   }
//
type CfnDomain_RStudioServerProAppSettingsProperty struct {
	// Indicates whether the current user has access to the `RStudioServerPro` app.
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// The level of permissions that the user has within the `RStudioServerPro` app.
	//
	// This value defaults to `User`. The `Admin` value allows the user access to the RStudio Administrative Dashboard.
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}

// A collection of settings that configure the `RStudioServerPro` Domain-level app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProDomainSettingsProperty := &rStudioServerProDomainSettingsProperty{
//   	domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   	// the properties below are optional
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   	rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   }
//
type CfnDomain_RStudioServerProDomainSettingsProperty struct {
	// The ARN of the execution role for the `RStudioServerPro` Domain-level app.
	DomainExecutionRoleArn *string `field:"required" json:"domainExecutionRoleArn" yaml:"domainExecutionRoleArn"`
	// A collection that defines the default `InstanceType` , `SageMakerImageArn` , and `SageMakerImageVersionArn` for the Domain.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A URL pointing to an RStudio Connect server.
	RStudioConnectUrl *string `field:"optional" json:"rStudioConnectUrl" yaml:"rStudioConnectUrl"`
	// A URL pointing to an RStudio Package Manager server.
	RStudioPackageManagerUrl *string `field:"optional" json:"rStudioPackageManagerUrl" yaml:"rStudioPackageManagerUrl"`
}

// Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecProperty := &resourceSpecProperty{
//   	instanceType: jsii.String("instanceType"),
//   	sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   	sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   }
//
type CfnDomain_ResourceSpecProperty struct {
	// The instance type that the image version runs on.
	//
	// > JupyterServer Apps only support the `system` value. KernelGateway Apps do not support the `system` value, but support all other values for available instance types.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the SageMaker image that the image version belongs to.
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}

// Specifies options when sharing an Amazon SageMaker Studio notebook.
//
// These settings are specified as part of `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called, and as part of `UserSettings` when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharingSettingsProperty := &sharingSettingsProperty{
//   	notebookOutputOption: jsii.String("notebookOutputOption"),
//   	s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   	s3OutputPath: jsii.String("s3OutputPath"),
//   }
//
type CfnDomain_SharingSettingsProperty struct {
	// Whether to include the notebook cell output when sharing the notebook.
	//
	// The default is `Disabled` .
	NotebookOutputOption *string `field:"optional" json:"notebookOutputOption" yaml:"notebookOutputOption"`
	// When `NotebookOutputOption` is `Allowed` , the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
	// When `NotebookOutputOption` is `Allowed` , the Amazon S3 bucket used to store the shared notebook snapshots.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

// A collection of settings that apply to users of Amazon SageMaker Studio.
//
// These settings are specified when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called, and as `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called.
//
// `SecurityGroups` is aggregated when specified in both calls. For all other settings in `UserSettings` , the values specified in `CreateUserProfile` take precedence over those specified in `CreateDomain` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingsProperty := &userSettingsProperty{
//   	executionRole: jsii.String("executionRole"),
//   	jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   		customImages: []interface{}{
//   			&customImageProperty{
//   				appImageConfigName: jsii.String("appImageConfigName"),
//   				imageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				imageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   		accessStatus: jsii.String("accessStatus"),
//   		userGroup: jsii.String("userGroup"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	sharingSettings: &sharingSettingsProperty{
//   		notebookOutputOption: jsii.String("notebookOutputOption"),
//   		s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   		s3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   }
//
type CfnDomain_UserSettingsProperty struct {
	// The execution role for the user.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The Jupyter server's app settings.
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A collection of settings that configure user interaction with the `RStudioServerPro` app.
	RStudioServerProAppSettings interface{} `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	//
	// Optional when the `CreateDomain.AppNetworkAccessType` parameter is set to `PublicInternetOnly` .
	//
	// Required when the `CreateDomain.AppNetworkAccessType` parameter is set to `VpcOnly` .
	//
	// Amazon SageMaker adds a security group to allow NFS traffic from SageMaker Studio. Therefore, the number of security groups that you can specify is one less than the maximum number shown.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies options for sharing SageMaker Studio notebooks.
	SharingSettings interface{} `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
}

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &cfnDomainProps{
//   	authMode: jsii.String("authMode"),
//   	defaultUserSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   	domainName: jsii.String("domainName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	appNetworkAccessType: jsii.String("appNetworkAccessType"),
//   	appSecurityGroupManagement: jsii.String("appSecurityGroupManagement"),
//   	domainSettings: &domainSettingsProperty{
//   		rStudioServerProDomainSettings: &rStudioServerProDomainSettingsProperty{
//   			domainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//
//   			// the properties below are optional
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			rStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   			rStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainProps struct {
	// The mode of authentication that members use to access the Domain.
	//
	// *Valid Values* : `SSO | IAM`.
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The default user settings.
	DefaultUserSettings interface{} `field:"required" json:"defaultUserSettings" yaml:"defaultUserSettings"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The VPC subnets that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Array members* : Minimum number of 1 item. Maximum number of 16 items.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly` .
	//
	// - `PublicInternetOnly` - Non-EFS traffic is through a VPC managed by Amazon SageMaker , which allows direct internet access
	// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets
	//
	// *Valid Values* : `PublicInternetOnly | VpcOnly`.
	AppNetworkAccessType *string `field:"optional" json:"appNetworkAccessType" yaml:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in `VpcOnly` mode.
	//
	// Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and `DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided.
	AppSecurityGroupManagement *string `field:"optional" json:"appSecurityGroupManagement" yaml:"appSecurityGroupManagement"`
	// A collection of settings that apply to the `SageMaker Domain` .
	//
	// These settings are specified through the `CreateDomain` API call.
	DomainSettings interface{} `field:"optional" json:"domainSettings" yaml:"domainSettings"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default.
	//
	// For more control, specify a customer managed CMK.
	//
	// *Length Constraints* : Maximum length of 2048.
	//
	// *Pattern* : `.*`
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Tags to associated with the Domain.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Tags that you specify for the Domain are also added to all apps that are launched in the Domain.
	//
	// *Array members* : Minimum number of 0 items. Maximum number of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::Endpoint`.
//
// Use the `AWS::SageMaker::Endpoint` resource to create an endpoint using the specified configuration in the request. Amazon SageMaker uses the endpoint to provision resources and deploy models. You create the endpoint configuration with the [AWS::SageMaker::EndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html) resource. For more information, see [Deploy a Model on Amazon SageMaker Hosting Services](https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works-hosting.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpoint := awscdk.Aws_sagemaker.NewCfnEndpoint(this, jsii.String("MyCfnEndpoint"), &cfnEndpointProps{
//   	endpointConfigName: jsii.String("endpointConfigName"),
//
//   	// the properties below are optional
//   	deploymentConfig: &deploymentConfigProperty{
//   		blueGreenUpdatePolicy: &blueGreenUpdatePolicyProperty{
//   			trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				canarySize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				linearStepSize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				waitIntervalInSeconds: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			terminationWaitInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		autoRollbackConfiguration: &autoRollbackConfigProperty{
//   			alarms: []interface{}{
//   				&alarmProperty{
//   					alarmName: jsii.String("alarmName"),
//   				},
//   			},
//   		},
//   	},
//   	endpointName: jsii.String("endpointName"),
//   	excludeRetainedVariantProperties: []interface{}{
//   		&variantPropertyProperty{
//   			variantPropertyType: jsii.String("variantPropertyType"),
//   		},
//   	},
//   	retainAllVariantProperties: jsii.Boolean(false),
//   	retainDeploymentConfig: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the endpoint, such as `MyEndpoint` .
	AttrEndpointName() *string
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
	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
	DeploymentConfig() interface{}
	SetDeploymentConfig(val interface{})
	// The name of the [AWS::SageMaker::EndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html) resource that specifies the configuration for the endpoint. For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) .
	EndpointConfigName() *string
	SetEndpointConfigName(val *string)
	// The name of the endpoint.The name must be unique within an AWS Region in your AWS account. The name is case-insensitive in `CreateEndpoint` , but the case is preserved and must be matched in [](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html) .
	EndpointName() *string
	SetEndpointName(val *string)
	// When you are updating endpoint resources with [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) whose value is set to `true` , `ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html) to override with the values provided by `EndpointConfig` . If you don't specify a value for `ExcludeAllVariantProperties` , no variant properties are overridden. Don't use this property when creating new endpoint resources or when `RetainAllVariantProperties` is set to `false` .
	ExcludeRetainedVariantProperties() interface{}
	SetExcludeRetainedVariantProperties(val interface{})
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
	// When updating endpoint resources, enables or disables the retention of variant properties, such as the instance count or the variant weight.
	//
	// To retain the variant properties of an endpoint when updating it, set `RetainAllVariantProperties` to `true` . To use the variant properties specified in a new `EndpointConfig` call when updating an endpoint, set `RetainAllVariantProperties` to `false` . Use this property only when updating endpoint resources, not when creating new endpoint resources.
	RetainAllVariantProperties() interface{}
	SetRetainAllVariantProperties(val interface{})
	// Specifies whether to reuse the last deployment configuration.
	//
	// The default value is false (the configuration is not reused).
	RetainDeploymentConfig() interface{}
	SetRetainDeploymentConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
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

// The jsii proxy struct for CfnEndpoint
type jsiiProxy_CfnEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpoint) AttrEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DeploymentConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ExcludeRetainedVariantProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeRetainedVariantProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RetainAllVariantProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainAllVariantProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RetainDeploymentConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainDeploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Endpoint`.
func NewCfnEndpoint(scope awscdk.Construct, id *string, props *CfnEndpointProps) CfnEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpoint{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Endpoint`.
func NewCfnEndpoint_Override(c CfnEndpoint, scope awscdk.Construct, id *string, props *CfnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetDeploymentConfig(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetEndpointConfigName(val *string) {
	_jsii_.Set(
		j,
		"endpointConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetEndpointName(val *string) {
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetExcludeRetainedVariantProperties(val interface{}) {
	_jsii_.Set(
		j,
		"excludeRetainedVariantProperties",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetRetainAllVariantProperties(val interface{}) {
	_jsii_.Set(
		j,
		"retainAllVariantProperties",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint) SetRetainDeploymentConfig(val interface{}) {
	_jsii_.Set(
		j,
		"retainDeploymentConfig",
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
func CfnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An Amazon CloudWatch alarm configured to monitor metrics on an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &alarmProperty{
//   	alarmName: jsii.String("alarmName"),
//   }
//
type CfnEndpoint_AlarmProperty struct {
	// The name of a CloudWatch alarm in your account.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
}

// Automatic rollback configuration for handling endpoint deployment failures and recovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRollbackConfigProperty := &autoRollbackConfigProperty{
//   	alarms: []interface{}{
//   		&alarmProperty{
//   			alarmName: jsii.String("alarmName"),
//   		},
//   	},
//   }
//
type CfnEndpoint_AutoRollbackConfigProperty struct {
	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint.
	//
	// If any alarms are tripped during a deployment, SageMaker rolls back the deployment.
	Alarms interface{} `field:"required" json:"alarms" yaml:"alarms"`
}

// Update policy for a blue/green deployment.
//
// If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueGreenUpdatePolicyProperty := &blueGreenUpdatePolicyProperty{
//   	trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		canarySize: &capacitySizeProperty{
//   			type: jsii.String("type"),
//   			value: jsii.Number(123),
//   		},
//   		linearStepSize: &capacitySizeProperty{
//   			type: jsii.String("type"),
//   			value: jsii.Number(123),
//   		},
//   		waitIntervalInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   	terminationWaitInSeconds: jsii.Number(123),
//   }
//
type CfnEndpoint_BlueGreenUpdatePolicyProperty struct {
	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment.
	TrafficRoutingConfiguration interface{} `field:"required" json:"trafficRoutingConfiguration" yaml:"trafficRoutingConfiguration"`
	// Maximum execution timeout for the deployment.
	//
	// Note that the timeout value should be larger than the total waiting time specified in `TerminationWaitInSeconds` and `WaitIntervalInSeconds` .
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet.
	//
	// Default is 0.
	TerminationWaitInSeconds *float64 `field:"optional" json:"terminationWaitInSeconds" yaml:"terminationWaitInSeconds"`
}

// Specifies the endpoint capacity to activate for production.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacitySizeProperty := &capacitySizeProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnEndpoint_CapacitySizeProperty struct {
	// Specifies the endpoint capacity type.
	//
	// - `INSTANCE_COUNT` : The endpoint activates based on the number of instances.
	// - `CAPACITY_PERCENT` : The endpoint activates based on the specified percentage of capacity.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigProperty := &deploymentConfigProperty{
//   	blueGreenUpdatePolicy: &blueGreenUpdatePolicyProperty{
//   		trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			canarySize: &capacitySizeProperty{
//   				type: jsii.String("type"),
//   				value: jsii.Number(123),
//   			},
//   			linearStepSize: &capacitySizeProperty{
//   				type: jsii.String("type"),
//   				value: jsii.Number(123),
//   			},
//   			waitIntervalInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   		terminationWaitInSeconds: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	autoRollbackConfiguration: &autoRollbackConfigProperty{
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				alarmName: jsii.String("alarmName"),
//   			},
//   		},
//   	},
//   }
//
type CfnEndpoint_DeploymentConfigProperty struct {
	// Update policy for a blue/green deployment.
	//
	// If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default.
	BlueGreenUpdatePolicy interface{} `field:"required" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// Automatic rollback configuration for handling endpoint deployment failures and recovery.
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
}

// Defines the traffic routing strategy during an endpoint deployment to shift traffic from the old fleet to the new fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRoutingConfigProperty := &trafficRoutingConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	canarySize: &capacitySizeProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   	linearStepSize: &capacitySizeProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   	waitIntervalInSeconds: jsii.Number(123),
//   }
//
type CfnEndpoint_TrafficRoutingConfigProperty struct {
	// Traffic routing strategy type.
	//
	// - `ALL_AT_ONCE` : Endpoint traffic shifts to the new fleet in a single step.
	// - `CANARY` : Endpoint traffic shifts to the new fleet in two steps. The first step is the canary, which is a small portion of the traffic. The second step is the remainder of the traffic.
	// - `LINEAR` : Endpoint traffic shifts to the new fleet in n steps of a configurable size.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Batch size for the first step to turn on traffic on the new endpoint fleet.
	//
	// `Value` must be less than or equal to 50% of the variant's total instance count.
	CanarySize interface{} `field:"optional" json:"canarySize" yaml:"canarySize"`
	// Batch size for each step to turn on traffic on the new endpoint fleet.
	//
	// `Value` must be 10-50% of the variant's total instance count.
	LinearStepSize interface{} `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
	// The waiting time (in seconds) between incremental steps to turn on traffic on the new endpoint fleet.
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

// Specifies a production variant property type for an Endpoint.
//
// If you are updating an Endpoint with the [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) option set to `true` , the `VarientProperty` objects listed in [ExcludeRetainedVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-ExcludeRetainedVariantProperties) override the existing variant properties of the Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variantPropertyProperty := &variantPropertyProperty{
//   	variantPropertyType: jsii.String("variantPropertyType"),
//   }
//
type CfnEndpoint_VariantPropertyProperty struct {
	// The type of variant property. The supported values are:.
	//
	// - `DesiredInstanceCount` : Overrides the existing variant instance counts using the [InitialInstanceCount](https://docs.aws.amazon.com/sagemaker/latest/dg/API_ProductionVariant.html#SageMaker-Type-ProductionVariant-InitialInstanceCount) values in the [ProductionVariants](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html#SageMaker-CreateEndpointConfig-request-ProductionVariants) .
	// - `DesiredWeight` : Overrides the existing variant weights using the [InitialVariantWeight](https://docs.aws.amazon.com/sagemaker/latest/dg/API_ProductionVariant.html#SageMaker-Type-ProductionVariant-InitialVariantWeight) values in the [ProductionVariants](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html#SageMaker-CreateEndpointConfig-request-ProductionVariants) .
	// - `DataCaptureConfig` : (Not currently supported.)
	VariantPropertyType *string `field:"optional" json:"variantPropertyType" yaml:"variantPropertyType"`
}

// A CloudFormation `AWS::SageMaker::EndpointConfig`.
//
// The `AWS::SageMaker::EndpointConfig` resource creates a configuration for an Amazon SageMaker endpoint. For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) in the *SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointConfig := awscdk.Aws_sagemaker.NewCfnEndpointConfig(this, jsii.String("MyCfnEndpointConfig"), &cfnEndpointConfigProps{
//   	productionVariants: []interface{}{
//   		&productionVariantProperty{
//   			initialVariantWeight: jsii.Number(123),
//   			modelName: jsii.String("modelName"),
//   			variantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			acceleratorType: jsii.String("acceleratorType"),
//   			initialInstanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			serverlessConfig: &serverlessConfigProperty{
//   				maxConcurrency: jsii.Number(123),
//   				memorySizeInMb: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	asyncInferenceConfig: &asyncInferenceConfigProperty{
//   		outputConfig: &asyncInferenceOutputConfigProperty{
//   			s3OutputPath: jsii.String("s3OutputPath"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			notificationConfig: &asyncInferenceNotificationConfigProperty{
//   				errorTopic: jsii.String("errorTopic"),
//   				successTopic: jsii.String("successTopic"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		clientConfig: &asyncInferenceClientConfigProperty{
//   			maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   	},
//   	dataCaptureConfig: &dataCaptureConfigProperty{
//   		captureOptions: []interface{}{
//   			&captureOptionProperty{
//   				captureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		destinationS3Uri: jsii.String("destinationS3Uri"),
//   		initialSamplingPercentage: jsii.Number(123),
//
//   		// the properties below are optional
//   		captureContentTypeHeader: &captureContentTypeHeaderProperty{
//   			csvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			jsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		enableCapture: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	endpointConfigName: jsii.String("endpointConfigName"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnEndpointConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig() interface{}
	SetAsyncInferenceConfig(val interface{})
	// The name of the endpoint configuration, such as `MyEndpointConfiguration` .
	AttrEndpointConfigName() *string
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
	// Specifies how to capture endpoint data for model monitor.
	//
	// The data capture configuration applies to all production variants hosted at the endpoint.
	DataCaptureConfig() interface{}
	SetDataCaptureConfig(val interface{})
	// The name of the endpoint configuration.
	EndpointConfigName() *string
	SetEndpointConfigName(val *string)
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Alias name: `alias/ExampleAlias`
	// - Alias name ARN: `arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias`
	//
	// The KMS key policy must grant permission to the IAM role that you specify in your `CreateEndpoint` , `UpdateEndpoint` requests. For more information, refer to the AWS Key Management Service section [Using Key Policies in AWS KMS](https://docs.aws.amazon.com//kms/latest/developerguide/key-policies.html)
	//
	// > Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a `KmsKeyId` when using an instance type with local storage. If any of the models that you specify in the `ProductionVariants` parameter use nitro-based instances with local storage, do not specify a value for the `KmsKeyId` parameter. If you specify a value for `KmsKeyId` when using any nitro-based instances with local storage, the call to `CreateEndpointConfig` fails.
	// >
	// > For a list of instance types that support local instance storage, see [Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes) .
	// >
	// > For more information about local instance storage encryption, see [SSD Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html) .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// A list of `ProductionVariant` objects, one for each model that you want to host at this endpoint.
	ProductionVariants() interface{}
	SetProductionVariants(val interface{})
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
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
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

// The jsii proxy struct for CfnEndpointConfig
type jsiiProxy_CfnEndpointConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpointConfig) AsyncInferenceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asyncInferenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) AttrEndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) DataCaptureConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCaptureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) EndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) ProductionVariants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::EndpointConfig`.
func NewCfnEndpointConfig(scope awscdk.Construct, id *string, props *CfnEndpointConfigProps) CfnEndpointConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointConfig{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::EndpointConfig`.
func NewCfnEndpointConfig_Override(c CfnEndpointConfig, scope awscdk.Construct, id *string, props *CfnEndpointConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointConfig) SetAsyncInferenceConfig(val interface{}) {
	_jsii_.Set(
		j,
		"asyncInferenceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig) SetDataCaptureConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dataCaptureConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig) SetEndpointConfigName(val *string) {
	_jsii_.Set(
		j,
		"endpointConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointConfig) SetProductionVariants(val interface{}) {
	_jsii_.Set(
		j,
		"productionVariants",
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
func CfnEndpointConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEndpointConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEndpointConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnEndpointConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpointConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpointConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEndpointConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceClientConfigProperty := &asyncInferenceClientConfigProperty{
//   	maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   }
//
type CfnEndpointConfig_AsyncInferenceClientConfigProperty struct {
	// The maximum number of concurrent requests sent by the SageMaker client to the model container.
	//
	// If no value is provided, SageMaker will choose an optimal value for you.
	MaxConcurrentInvocationsPerInstance *float64 `field:"optional" json:"maxConcurrentInvocationsPerInstance" yaml:"maxConcurrentInvocationsPerInstance"`
}

// Specifies configuration for how an endpoint performs asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceConfigProperty := &asyncInferenceConfigProperty{
//   	outputConfig: &asyncInferenceOutputConfigProperty{
//   		s3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		notificationConfig: &asyncInferenceNotificationConfigProperty{
//   			errorTopic: jsii.String("errorTopic"),
//   			successTopic: jsii.String("successTopic"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	clientConfig: &asyncInferenceClientConfigProperty{
//   		maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   	},
//   }
//
type CfnEndpointConfig_AsyncInferenceConfigProperty struct {
	// Specifies the configuration for asynchronous inference invocation outputs.
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
	ClientConfig interface{} `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

// Specifies the configuration for notifications of inference results for asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceNotificationConfigProperty := &asyncInferenceNotificationConfigProperty{
//   	errorTopic: jsii.String("errorTopic"),
//   	successTopic: jsii.String("successTopic"),
//   }
//
type CfnEndpointConfig_AsyncInferenceNotificationConfigProperty struct {
	// Amazon SNS topic to post a notification to when an inference fails.
	//
	// If no topic is provided, no notification is sent on failure.
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// Amazon SNS topic to post a notification to when an inference completes successfully.
	//
	// If no topic is provided, no notification is sent on success.
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

// Specifies the configuration for asynchronous inference invocation outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceOutputConfigProperty := &asyncInferenceOutputConfigProperty{
//   	s3OutputPath: jsii.String("s3OutputPath"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	notificationConfig: &asyncInferenceNotificationConfigProperty{
//   		errorTopic: jsii.String("errorTopic"),
//   		successTopic: jsii.String("successTopic"),
//   	},
//   }
//
type CfnEndpointConfig_AsyncInferenceOutputConfigProperty struct {
	// The Amazon S3 location to upload inference responses to.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the configuration for notifications of inference results for asynchronous inference.
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
}

// Specifies the JSON and CSV content types of the data that the endpoint captures.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureContentTypeHeaderProperty := &captureContentTypeHeaderProperty{
//   	csvContentTypes: []*string{
//   		jsii.String("csvContentTypes"),
//   	},
//   	jsonContentTypes: []*string{
//   		jsii.String("jsonContentTypes"),
//   	},
//   }
//
type CfnEndpointConfig_CaptureContentTypeHeaderProperty struct {
	// A list of the CSV content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// A list of the JSON content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

// Specifies whether the endpoint captures input data or output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureOptionProperty := &captureOptionProperty{
//   	captureMode: jsii.String("captureMode"),
//   }
//
type CfnEndpointConfig_CaptureOptionProperty struct {
	// Specifies whether the endpoint captures input data or output data.
	CaptureMode *string `field:"required" json:"captureMode" yaml:"captureMode"`
}

// Specifies the configuration of your endpoint for model monitor data capture.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCaptureConfigProperty := &dataCaptureConfigProperty{
//   	captureOptions: []interface{}{
//   		&captureOptionProperty{
//   			captureMode: jsii.String("captureMode"),
//   		},
//   	},
//   	destinationS3Uri: jsii.String("destinationS3Uri"),
//   	initialSamplingPercentage: jsii.Number(123),
//
//   	// the properties below are optional
//   	captureContentTypeHeader: &captureContentTypeHeaderProperty{
//   		csvContentTypes: []*string{
//   			jsii.String("csvContentTypes"),
//   		},
//   		jsonContentTypes: []*string{
//   			jsii.String("jsonContentTypes"),
//   		},
//   	},
//   	enableCapture: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnEndpointConfig_DataCaptureConfigProperty struct {
	// Specifies whether the endpoint captures input data to your model, output data from your model, or both.
	CaptureOptions interface{} `field:"required" json:"captureOptions" yaml:"captureOptions"`
	// The S3 bucket where model monitor stores captured data.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// The percentage of data to capture.
	InitialSamplingPercentage *float64 `field:"required" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// A list of the JSON and CSV content type that the endpoint captures.
	CaptureContentTypeHeader interface{} `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Set to `True` to enable data capture.
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the captured data at rest using Amazon S3 server-side encryption.
	//
	// The KmsKeyId can be any of the following formats: Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab Alias name: alias/ExampleAlias Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account. For more information, see KMS-Managed Encryption Keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the Amazon Simple Storage Service Developer Guide. The KMS key policy must grant permission to the IAM role that you specify in your CreateModel (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html) request. For more information, see Using Key Policies in AWS KMS (http://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the AWS Key Management Service Developer Guide.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Specifies a model that you want to host and the resources to deploy for hosting it.
//
// If you are deploying multiple models, tell Amazon SageMaker how to distribute traffic among the models by specifying the `InitialVariantWeight` objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   productionVariantProperty := &productionVariantProperty{
//   	initialVariantWeight: jsii.Number(123),
//   	modelName: jsii.String("modelName"),
//   	variantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	acceleratorType: jsii.String("acceleratorType"),
//   	initialInstanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	serverlessConfig: &serverlessConfigProperty{
//   		maxConcurrency: jsii.Number(123),
//   		memorySizeInMb: jsii.Number(123),
//   	},
//   }
//
type CfnEndpointConfig_ProductionVariantProperty struct {
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the `VariantWeight` to the sum of all `VariantWeight` values across all ProductionVariants. If unspecified, it defaults to 1.0.
	InitialVariantWeight *float64 `field:"required" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the production variant.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	//
	// EI instances provide on-demand GPU computing for inference. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) . For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// Number of instances to launch initially.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// The ML compute instance type.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The serverless configuration for an endpoint.
	//
	// Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	ServerlessConfig interface{} `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
}

// Specifies the serverless configuration for an endpoint variant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessConfigProperty := &serverlessConfigProperty{
//   	maxConcurrency: jsii.Number(123),
//   	memorySizeInMb: jsii.Number(123),
//   }
//
type CfnEndpointConfig_ServerlessConfigProperty struct {
	// The maximum number of concurrent invocations your serverless endpoint can process.
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The memory size of your serverless endpoint.
	//
	// Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	MemorySizeInMb *float64 `field:"required" json:"memorySizeInMb" yaml:"memorySizeInMb"`
}

// Properties for defining a `CfnEndpointConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointConfigProps := &cfnEndpointConfigProps{
//   	productionVariants: []interface{}{
//   		&productionVariantProperty{
//   			initialVariantWeight: jsii.Number(123),
//   			modelName: jsii.String("modelName"),
//   			variantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			acceleratorType: jsii.String("acceleratorType"),
//   			initialInstanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			serverlessConfig: &serverlessConfigProperty{
//   				maxConcurrency: jsii.Number(123),
//   				memorySizeInMb: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	asyncInferenceConfig: &asyncInferenceConfigProperty{
//   		outputConfig: &asyncInferenceOutputConfigProperty{
//   			s3OutputPath: jsii.String("s3OutputPath"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			notificationConfig: &asyncInferenceNotificationConfigProperty{
//   				errorTopic: jsii.String("errorTopic"),
//   				successTopic: jsii.String("successTopic"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		clientConfig: &asyncInferenceClientConfigProperty{
//   			maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   	},
//   	dataCaptureConfig: &dataCaptureConfigProperty{
//   		captureOptions: []interface{}{
//   			&captureOptionProperty{
//   				captureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		destinationS3Uri: jsii.String("destinationS3Uri"),
//   		initialSamplingPercentage: jsii.Number(123),
//
//   		// the properties below are optional
//   		captureContentTypeHeader: &captureContentTypeHeaderProperty{
//   			csvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			jsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		enableCapture: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	endpointConfigName: jsii.String("endpointConfigName"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEndpointConfigProps struct {
	// A list of `ProductionVariant` objects, one for each model that you want to host at this endpoint.
	ProductionVariants interface{} `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig interface{} `field:"optional" json:"asyncInferenceConfig" yaml:"asyncInferenceConfig"`
	// Specifies how to capture endpoint data for model monitor.
	//
	// The data capture configuration applies to all production variants hosted at the endpoint.
	DataCaptureConfig interface{} `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// The name of the endpoint configuration.
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Alias name: `alias/ExampleAlias`
	// - Alias name ARN: `arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias`
	//
	// The KMS key policy must grant permission to the IAM role that you specify in your `CreateEndpoint` , `UpdateEndpoint` requests. For more information, refer to the AWS Key Management Service section [Using Key Policies in AWS KMS](https://docs.aws.amazon.com//kms/latest/developerguide/key-policies.html)
	//
	// > Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a `KmsKeyId` when using an instance type with local storage. If any of the models that you specify in the `ProductionVariants` parameter use nitro-based instances with local storage, do not specify a value for the `KmsKeyId` parameter. If you specify a value for `KmsKeyId` when using any nitro-based instances with local storage, the call to `CreateEndpointConfig` fails.
	// >
	// > For a list of instance types that support local instance storage, see [Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes) .
	// >
	// > For more information about local instance storage encryption, see [SSD Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html) .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &cfnEndpointProps{
//   	endpointConfigName: jsii.String("endpointConfigName"),
//
//   	// the properties below are optional
//   	deploymentConfig: &deploymentConfigProperty{
//   		blueGreenUpdatePolicy: &blueGreenUpdatePolicyProperty{
//   			trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				canarySize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				linearStepSize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				waitIntervalInSeconds: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			terminationWaitInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		autoRollbackConfiguration: &autoRollbackConfigProperty{
//   			alarms: []interface{}{
//   				&alarmProperty{
//   					alarmName: jsii.String("alarmName"),
//   				},
//   			},
//   		},
//   	},
//   	endpointName: jsii.String("endpointName"),
//   	excludeRetainedVariantProperties: []interface{}{
//   		&variantPropertyProperty{
//   			variantPropertyType: jsii.String("variantPropertyType"),
//   		},
//   	},
//   	retainAllVariantProperties: jsii.Boolean(false),
//   	retainDeploymentConfig: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEndpointProps struct {
	// The name of the [AWS::SageMaker::EndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html) resource that specifies the configuration for the endpoint. For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) .
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
	DeploymentConfig interface{} `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The name of the endpoint.The name must be unique within an AWS Region in your AWS account. The name is case-insensitive in `CreateEndpoint` , but the case is preserved and must be matched in [](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html) .
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// When you are updating endpoint resources with [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) whose value is set to `true` , `ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html) to override with the values provided by `EndpointConfig` . If you don't specify a value for `ExcludeAllVariantProperties` , no variant properties are overridden. Don't use this property when creating new endpoint resources or when `RetainAllVariantProperties` is set to `false` .
	ExcludeRetainedVariantProperties interface{} `field:"optional" json:"excludeRetainedVariantProperties" yaml:"excludeRetainedVariantProperties"`
	// When updating endpoint resources, enables or disables the retention of variant properties, such as the instance count or the variant weight.
	//
	// To retain the variant properties of an endpoint when updating it, set `RetainAllVariantProperties` to `true` . To use the variant properties specified in a new `EndpointConfig` call when updating an endpoint, set `RetainAllVariantProperties` to `false` . Use this property only when updating endpoint resources, not when creating new endpoint resources.
	RetainAllVariantProperties interface{} `field:"optional" json:"retainAllVariantProperties" yaml:"retainAllVariantProperties"`
	// Specifies whether to reuse the last deployment configuration.
	//
	// The default value is false (the configuration is not reused).
	RetainDeploymentConfig interface{} `field:"optional" json:"retainDeploymentConfig" yaml:"retainDeploymentConfig"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::FeatureGroup`.
//
// Create a new `FeatureGroup` . A `FeatureGroup` is a group of `Features` defined in the `FeatureStore` to describe a `Record` .
//
// The `FeatureGroup` defines the schema and features contained in the FeatureGroup. A `FeatureGroup` definition is composed of a list of `Features` , a `RecordIdentifierFeatureName` , an `EventTimeFeatureName` and configurations for its `OnlineStore` and `OfflineStore` . Check [AWS service quotas](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) to see the `FeatureGroup` s quota for your AWS account.
//
// > You must include at least one of `OnlineStoreConfig` and `OfflineStoreConfig` to create a `FeatureGroup` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var offlineStoreConfig interface{}
//   var onlineStoreConfig interface{}
//
//   cfnFeatureGroup := awscdk.Aws_sagemaker.NewCfnFeatureGroup(this, jsii.String("MyCfnFeatureGroup"), &cfnFeatureGroupProps{
//   	eventTimeFeatureName: jsii.String("eventTimeFeatureName"),
//   	featureDefinitions: []interface{}{
//   		&featureDefinitionProperty{
//   			featureName: jsii.String("featureName"),
//   			featureType: jsii.String("featureType"),
//   		},
//   	},
//   	featureGroupName: jsii.String("featureGroupName"),
//   	recordIdentifierFeatureName: jsii.String("recordIdentifierFeatureName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	offlineStoreConfig: offlineStoreConfig,
//   	onlineStoreConfig: onlineStoreConfig,
//   	roleArn: jsii.String("roleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFeatureGroup interface {
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
	// A free form description of a `FeatureGroup` .
	Description() *string
	SetDescription(val *string)
	// The name of the feature that stores the `EventTime` of a Record in a `FeatureGroup` .
	//
	// A `EventTime` is point in time when a new event occurs that corresponds to the creation or update of a `Record` in `FeatureGroup` . All `Records` in the `FeatureGroup` must have a corresponding `EventTime` .
	EventTimeFeatureName() *string
	SetEventTimeFeatureName(val *string)
	// A list of `Feature` s. Each `Feature` must include a `FeatureName` and a `FeatureType` .
	//
	// Valid `FeatureType` s are `Integral` , `Fractional` and `String` .
	//
	// `FeatureName` s cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	//
	// You can create up to 2,500 `FeatureDefinition` s per `FeatureGroup` .
	FeatureDefinitions() interface{}
	SetFeatureDefinitions(val interface{})
	// The name of the `FeatureGroup` .
	FeatureGroupName() *string
	SetFeatureGroupName(val *string)
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
	// The configuration of an `OfflineStore` .
	OfflineStoreConfig() interface{}
	SetOfflineStoreConfig(val interface{})
	// The configuration of an `OnlineStore` .
	OnlineStoreConfig() interface{}
	SetOnlineStoreConfig(val interface{})
	// The name of the `Feature` whose value uniquely identifies a `Record` defined in the `FeatureGroup` `FeatureDefinitions` .
	RecordIdentifierFeatureName() *string
	SetRecordIdentifierFeatureName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM execution role used to create the feature group.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Tags used to define a `FeatureGroup` .
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

// The jsii proxy struct for CfnFeatureGroup
type jsiiProxy_CfnFeatureGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFeatureGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) EventTimeFeatureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTimeFeatureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) FeatureDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) FeatureGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) OfflineStoreConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"offlineStoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) OnlineStoreConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlineStoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) RecordIdentifierFeatureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordIdentifierFeatureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeatureGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::FeatureGroup`.
func NewCfnFeatureGroup(scope awscdk.Construct, id *string, props *CfnFeatureGroupProps) CfnFeatureGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnFeatureGroup{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::FeatureGroup`.
func NewCfnFeatureGroup_Override(c CfnFeatureGroup, scope awscdk.Construct, id *string, props *CfnFeatureGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetEventTimeFeatureName(val *string) {
	_jsii_.Set(
		j,
		"eventTimeFeatureName",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetFeatureDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"featureDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetFeatureGroupName(val *string) {
	_jsii_.Set(
		j,
		"featureGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetOfflineStoreConfig(val interface{}) {
	_jsii_.Set(
		j,
		"offlineStoreConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetOnlineStoreConfig(val interface{}) {
	_jsii_.Set(
		j,
		"onlineStoreConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetRecordIdentifierFeatureName(val *string) {
	_jsii_.Set(
		j,
		"recordIdentifierFeatureName",
		val,
	)
}

func (j *jsiiProxy_CfnFeatureGroup) SetRoleArn(val *string) {
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
func CfnFeatureGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFeatureGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFeatureGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeatureGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnFeatureGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFeatureGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFeatureGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFeatureGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeatureGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A list of features.
//
// You must include `FeatureName` and `FeatureType` . Valid feature `FeatureType` s are `Integral` , `Fractional` and `String` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureDefinitionProperty := &featureDefinitionProperty{
//   	featureName: jsii.String("featureName"),
//   	featureType: jsii.String("featureType"),
//   }
//
type CfnFeatureGroup_FeatureDefinitionProperty struct {
	// The name of a feature.
	//
	// The type must be a string. `FeatureName` cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// The value type of a feature.
	//
	// Valid values are Integral, Fractional, or String.
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
}

// Properties for defining a `CfnFeatureGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var offlineStoreConfig interface{}
//   var onlineStoreConfig interface{}
//
//   cfnFeatureGroupProps := &cfnFeatureGroupProps{
//   	eventTimeFeatureName: jsii.String("eventTimeFeatureName"),
//   	featureDefinitions: []interface{}{
//   		&featureDefinitionProperty{
//   			featureName: jsii.String("featureName"),
//   			featureType: jsii.String("featureType"),
//   		},
//   	},
//   	featureGroupName: jsii.String("featureGroupName"),
//   	recordIdentifierFeatureName: jsii.String("recordIdentifierFeatureName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	offlineStoreConfig: offlineStoreConfig,
//   	onlineStoreConfig: onlineStoreConfig,
//   	roleArn: jsii.String("roleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFeatureGroupProps struct {
	// The name of the feature that stores the `EventTime` of a Record in a `FeatureGroup` .
	//
	// A `EventTime` is point in time when a new event occurs that corresponds to the creation or update of a `Record` in `FeatureGroup` . All `Records` in the `FeatureGroup` must have a corresponding `EventTime` .
	EventTimeFeatureName *string `field:"required" json:"eventTimeFeatureName" yaml:"eventTimeFeatureName"`
	// A list of `Feature` s. Each `Feature` must include a `FeatureName` and a `FeatureType` .
	//
	// Valid `FeatureType` s are `Integral` , `Fractional` and `String` .
	//
	// `FeatureName` s cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	//
	// You can create up to 2,500 `FeatureDefinition` s per `FeatureGroup` .
	FeatureDefinitions interface{} `field:"required" json:"featureDefinitions" yaml:"featureDefinitions"`
	// The name of the `FeatureGroup` .
	FeatureGroupName *string `field:"required" json:"featureGroupName" yaml:"featureGroupName"`
	// The name of the `Feature` whose value uniquely identifies a `Record` defined in the `FeatureGroup` `FeatureDefinitions` .
	RecordIdentifierFeatureName *string `field:"required" json:"recordIdentifierFeatureName" yaml:"recordIdentifierFeatureName"`
	// A free form description of a `FeatureGroup` .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration of an `OfflineStore` .
	OfflineStoreConfig interface{} `field:"optional" json:"offlineStoreConfig" yaml:"offlineStoreConfig"`
	// The configuration of an `OnlineStore` .
	OnlineStoreConfig interface{} `field:"optional" json:"onlineStoreConfig" yaml:"onlineStoreConfig"`
	// The Amazon Resource Name (ARN) of the IAM execution role used to create the feature group.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags used to define a `FeatureGroup` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::Image`.
//
// Creates a custom SageMaker image. A SageMaker image is a set of image versions. Each image version represents a container image stored in Amazon Elastic Container Registry (ECR). For more information, see [Bring your own SageMaker image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImage := awscdk.Aws_sagemaker.NewCfnImage(this, jsii.String("MyCfnImage"), &cfnImageProps{
//   	imageName: jsii.String("imageName"),
//   	imageRoleArn: jsii.String("imageRoleArn"),
//
//   	// the properties below are optional
//   	imageDescription: jsii.String("imageDescription"),
//   	imageDisplayName: jsii.String("imageDisplayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnImage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the image.
	//
	// *Type* : String
	//
	// *Length Constraints* : Maximum length of 256.
	//
	// *Pattern* : `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$`
	AttrImageArn() *string
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
	// The description of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 512.
	//
	// *Pattern* : `.*`
	ImageDescription() *string
	SetImageDescription(val *string)
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName() *string
	SetImageDisplayName(val *string)
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName() *string
	SetImageName(val *string)
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn() *string
	SetImageRoleArn(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs to apply to this resource.
	//
	// *Array Members* : Minimum number of 0 items. Maximum number of 50 items.
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

// The jsii proxy struct for CfnImage
type jsiiProxy_CfnImage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnImage) AttrImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Image`.
func NewCfnImage(scope awscdk.Construct, id *string, props *CfnImageProps) CfnImage {
	_init_.Initialize()

	j := jsiiProxy_CfnImage{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnImage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Image`.
func NewCfnImage_Override(c CfnImage, scope awscdk.Construct, id *string, props *CfnImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnImage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImage) SetImageDescription(val *string) {
	_jsii_.Set(
		j,
		"imageDescription",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetImageDisplayName(val *string) {
	_jsii_.Set(
		j,
		"imageDisplayName",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetImageRoleArn(val *string) {
	_jsii_.Set(
		j,
		"imageRoleArn",
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
func CfnImage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnImage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImage",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnImage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImage) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImage) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImage) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImage) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImage) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImage) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImage) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnImage) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnImage) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImage) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnImage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &cfnImageProps{
//   	imageName: jsii.String("imageName"),
//   	imageRoleArn: jsii.String("imageRoleArn"),
//
//   	// the properties below are optional
//   	imageDescription: jsii.String("imageDescription"),
//   	imageDisplayName: jsii.String("imageDisplayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnImageProps struct {
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn *string `field:"required" json:"imageRoleArn" yaml:"imageRoleArn"`
	// The description of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 512.
	//
	// *Pattern* : `.*`
	ImageDescription *string `field:"optional" json:"imageDescription" yaml:"imageDescription"`
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName *string `field:"optional" json:"imageDisplayName" yaml:"imageDisplayName"`
	// A list of key-value pairs to apply to this resource.
	//
	// *Array Members* : Minimum number of 0 items. Maximum number of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::ImageVersion`.
//
// Creates a version of the SageMaker image specified by `ImageName` . The version represents the Amazon Container Registry (ECR) container image specified by `BaseImage` .
//
// > You can use the `DependsOn` attribute to specify that the creation of a specific resource follows another. You can use it for the following use cases. For more information, see [`DependsOn` attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
// >
// > 1. `DependsOn` can be used to establish a parent/child relationship between `ImageVersion` and `Image` where the `ImageVersion` `DependsOn` the `Image` .
// >
// > 2. `DependsOn` can be used to establish order among `ImageVersion` s within the same `Image` namespace. For example, if ImageVersionB `DependsOn` ImageVersionA and both share the same parent `Image` , then ImageVersionA is version N and ImageVersionB is N+1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageVersion := awscdk.Aws_sagemaker.NewCfnImageVersion(this, jsii.String("MyCfnImageVersion"), &cfnImageVersionProps{
//   	baseImage: jsii.String("baseImage"),
//   	imageName: jsii.String("imageName"),
//   })
//
type CfnImageVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The URI of the container image version referenced by ImageVersion.
	AttrContainerImage() *string
	// The Amazon Resource Name (ARN) of the parent Image.
	AttrImageArn() *string
	// The Amazon Resource Name (ARN) of the image version.
	//
	// *Type* : String
	//
	// *Length Constraints* : Maximum length of 256.
	//
	// *Pattern* : `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])* /[0-9]+$`
	AttrImageVersionArn() *string
	// The version of the image.
	AttrVersion() *float64
	// The container image that the SageMaker image version is based on.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 255.
	//
	// *Pattern* : `.*`
	BaseImage() *string
	SetBaseImage(val *string)
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
	// The name of the parent image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName() *string
	SetImageName(val *string)
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

// The jsii proxy struct for CfnImageVersion
type jsiiProxy_CfnImageVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnImageVersion) AttrContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrContainerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) BaseImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ImageVersion`.
func NewCfnImageVersion(scope awscdk.Construct, id *string, props *CfnImageVersionProps) CfnImageVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnImageVersion{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnImageVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ImageVersion`.
func NewCfnImageVersion_Override(c CfnImageVersion, scope awscdk.Construct, id *string, props *CfnImageVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnImageVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImageVersion) SetBaseImage(val *string) {
	_jsii_.Set(
		j,
		"baseImage",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
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
func CfnImageVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImageVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnImageVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImageVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnImageVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnImageVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImageVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnImageVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImageVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImageVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImageVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnImageVersion) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnImageVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImageVersion) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnImageVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnImageVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnImageVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageVersionProps := &cfnImageVersionProps{
//   	baseImage: jsii.String("baseImage"),
//   	imageName: jsii.String("imageName"),
//   }
//
type CfnImageVersionProps struct {
	// The container image that the SageMaker image version is based on.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 255.
	//
	// *Pattern* : `.*`
	BaseImage *string `field:"required" json:"baseImage" yaml:"baseImage"`
	// The name of the parent image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
}

// A CloudFormation `AWS::SageMaker::Model`.
//
// The `AWS::SageMaker::Model` resource to create a model to host at an Amazon SageMaker endpoint. For more information, see [Deploying a Model on Amazon SageMaker Hosting Services](https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works-hosting.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var environment interface{}
//
//   cfnModel := awscdk.Aws_sagemaker.NewCfnModel(this, jsii.String("MyCfnModel"), &cfnModelProps{
//   	executionRoleArn: jsii.String("executionRoleArn"),
//
//   	// the properties below are optional
//   	containers: []interface{}{
//   		&containerDefinitionProperty{
//   			containerHostname: jsii.String("containerHostname"),
//   			environment: environment,
//   			image: jsii.String("image"),
//   			imageConfig: &imageConfigProperty{
//   				repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   				// the properties below are optional
//   				repositoryAuthConfig: &repositoryAuthConfigProperty{
//   					repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   				},
//   			},
//   			inferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   			mode: jsii.String("mode"),
//   			modelDataUrl: jsii.String("modelDataUrl"),
//   			modelPackageName: jsii.String("modelPackageName"),
//   			multiModelConfig: &multiModelConfigProperty{
//   				modelCacheSetting: jsii.String("modelCacheSetting"),
//   			},
//   		},
//   	},
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	inferenceExecutionConfig: &inferenceExecutionConfigProperty{
//   		mode: jsii.String("mode"),
//   	},
//   	modelName: jsii.String("modelName"),
//   	primaryContainer: &containerDefinitionProperty{
//   		containerHostname: jsii.String("containerHostname"),
//   		environment: environment,
//   		image: jsii.String("image"),
//   		imageConfig: &imageConfigProperty{
//   			repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   			// the properties below are optional
//   			repositoryAuthConfig: &repositoryAuthConfigProperty{
//   				repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   			},
//   		},
//   		inferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   		mode: jsii.String("mode"),
//   		modelDataUrl: jsii.String("modelDataUrl"),
//   		modelPackageName: jsii.String("modelPackageName"),
//   		multiModelConfig: &multiModelConfigProperty{
//   			modelCacheSetting: jsii.String("modelCacheSetting"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   })
//
type CfnModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the model, such as `MyModel` .
	AttrModelName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Specifies the containers in the inference pipeline.
	Containers() interface{}
	SetContainers(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Isolates the model container.
	//
	// No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation() interface{}
	SetEnableNetworkIsolation(val interface{})
	// The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs.
	//
	// Deploying on ML compute instances is part of model hosting. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig() interface{}
	SetInferenceExecutionConfig(val interface{})
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
	// The name of the new model.
	ModelName() *string
	SetModelName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.
	PrimaryContainer() interface{}
	SetPrimaryContainer(val interface{})
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
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_VpcConfig.html) object that specifies the VPC that you want your model to connect to. Control access to and from your model container by configuring the VPC. `VpcConfig` is used in hosting services and in batch transform. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html) .
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnModel
type jsiiProxy_CfnModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModel) AttrModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Containers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) InferenceExecutionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceExecutionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) PrimaryContainer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Model`.
func NewCfnModel(scope awscdk.Construct, id *string, props *CfnModelProps) CfnModel {
	_init_.Initialize()

	j := jsiiProxy_CfnModel{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Model`.
func NewCfnModel_Override(c CfnModel, scope awscdk.Construct, id *string, props *CfnModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModel) SetContainers(val interface{}) {
	_jsii_.Set(
		j,
		"containers",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetEnableNetworkIsolation(val interface{}) {
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetInferenceExecutionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"inferenceExecutionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetModelName(val *string) {
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetPrimaryContainer(val interface{}) {
	_jsii_.Set(
		j,
		"primaryContainer",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetVpcConfig(val interface{}) {
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
// Experimental.
func CfnModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModel) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModel) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the container, as part of model definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var environment interface{}
//
//   containerDefinitionProperty := &containerDefinitionProperty{
//   	containerHostname: jsii.String("containerHostname"),
//   	environment: environment,
//   	image: jsii.String("image"),
//   	imageConfig: &imageConfigProperty{
//   		repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   		// the properties below are optional
//   		repositoryAuthConfig: &repositoryAuthConfigProperty{
//   			repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   		},
//   	},
//   	inferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   	mode: jsii.String("mode"),
//   	modelDataUrl: jsii.String("modelDataUrl"),
//   	modelPackageName: jsii.String("modelPackageName"),
//   	multiModelConfig: &multiModelConfigProperty{
//   		modelCacheSetting: jsii.String("modelCacheSetting"),
//   	},
//   }
//
type CfnModel_ContainerDefinitionProperty struct {
	// This parameter is ignored for models that contain only a `PrimaryContainer` .
	//
	// When a `ContainerDefinition` is part of an inference pipeline, the value of the parameter uniquely identifies the container for the purposes of logging and metrics. For information, see [Use Logs and Metrics to Monitor an Inference Pipeline](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipeline-logs-metrics.html) . If you don't specify a value for this parameter for a `ContainerDefinition` that is part of an inference pipeline, a unique name is automatically assigned based on the position of the `ContainerDefinition` in the pipeline. If you specify a value for the `ContainerHostName` for any `ContainerDefinition` that is part of an inference pipeline, you must specify a value for the `ContainerHostName` parameter of every `ContainerDefinition` in that pipeline.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// The environment variables to set in the Docker container.
	//
	// Each key and value in the `Environment` string to string map can have length of up to 1024. We support up to 16 entries in the map.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The path where inference code is stored.
	//
	// This can be either in Amazon EC2 Container Registry or in a Docker registry that is accessible from the same VPC that you configure for your endpoint. If you are using your own custom algorithm instead of an algorithm provided by SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both `registry/repository[:tag]` and `registry/repository[@digest]` image path formats. For more information, see [Using Your Own Algorithms with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html)
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).
	//
	// For information about storing containers in a private Docker registry, see [Use a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html)
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// `CfnModel.ContainerDefinitionProperty.InferenceSpecificationName`.
	InferenceSpecificationName *string `field:"optional" json:"inferenceSpecificationName" yaml:"inferenceSpecificationName"`
	// Whether the container hosts a single model or multiple models.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix). The S3 path is required for SageMaker built-in algorithms, but not if you use your own algorithms. For more information on built-in algorithms, see [Common Parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html) .
	//
	// > The model artifacts must be in an S3 bucket that is in the same region as the model or endpoint you are creating.
	//
	// If you provide a value for this parameter, SageMaker uses AWS Security Token Service to download model artifacts from the S3 path you provide. AWS STS is activated in your IAM user account by default. If you previously deactivated AWS STS for a region, you need to reactivate AWS STS for that region. For more information, see [Activating and Deactivating AWS STS in an AWS Region](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html) in the *AWS Identity and Access Management User Guide* .
	//
	// > If you use a built-in algorithm to create a model, SageMaker requires that you provide a S3 path to the model artifacts in `ModelDataUrl` .
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// The name or Amazon Resource Name (ARN) of the model package to use to create the model.
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Specifies additional configuration for multi-model endpoints.
	MultiModelConfig interface{} `field:"optional" json:"multiModelConfig" yaml:"multiModelConfig"`
}

// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigProperty := &imageConfigProperty{
//   	repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   	// the properties below are optional
//   	repositoryAuthConfig: &repositoryAuthConfigProperty{
//   		repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   	},
//   }
//
type CfnModel_ImageConfigProperty struct {
	// Set this to one of the following values:.
	//
	// - `Platform` - The model image is hosted in Amazon ECR.
	// - `Vpc` - The model image is hosted in a private Docker registry in your VPC.
	RepositoryAccessMode *string `field:"required" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// (Optional) Specifies an authentication configuration for the private docker registry where your model image is hosted.
	//
	// Specify a value for this property only if you specified `Vpc` as the value for the `RepositoryAccessMode` field, and the private Docker registry where the model image is hosted requires authentication.
	RepositoryAuthConfig interface{} `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}

// Specifies details about how containers in a multi-container endpoint are run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceExecutionConfigProperty := &inferenceExecutionConfigProperty{
//   	mode: jsii.String("mode"),
//   }
//
type CfnModel_InferenceExecutionConfigProperty struct {
	// How containers in a multi-container are run. The following values are valid.
	//
	// - `Serial` - Containers run as a serial pipeline.
	// - `Direct` - Only the individual container that you specify is run.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

// Specifies additional configuration for hosting multi-model endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiModelConfigProperty := &multiModelConfigProperty{
//   	modelCacheSetting: jsii.String("modelCacheSetting"),
//   }
//
type CfnModel_MultiModelConfigProperty struct {
	// Whether to cache models for a multi-model endpoint.
	//
	// By default, multi-model endpoints cache models so that a model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint might perform better if you disable model caching. To disable model caching, set the value of this parameter to Disabled.
	ModelCacheSetting *string `field:"optional" json:"modelCacheSetting" yaml:"modelCacheSetting"`
}

// Specifies an authentication configuration for the private docker registry where your model image is hosted.
//
// Specify a value for this property only if you specified `Vpc` as the value for the `RepositoryAccessMode` field of the `ImageConfig` object that you passed to a call to `CreateModel` and the private Docker registry where the model image is hosted requires authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryAuthConfigProperty := &repositoryAuthConfigProperty{
//   	repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   }
//
type CfnModel_RepositoryAuthConfigProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted.
	//
	// For information about how to create an AWS Lambda function, see [Create a Lambda function with the console](https://docs.aws.amazon.com/lambda/latest/dg/getting-started-create-function.html) in the *AWS Lambda Developer Guide* .
	RepositoryCredentialsProviderArn *string `field:"required" json:"repositoryCredentialsProviderArn" yaml:"repositoryCredentialsProviderArn"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnModel_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// A CloudFormation `AWS::SageMaker::ModelBiasJobDefinition`.
//
// Creates the definition for a model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelBiasJobDefinition := awscdk.Aws_sagemaker.NewCfnModelBiasJobDefinition(this, jsii.String("MyCfnModelBiasJobDefinition"), &cfnModelBiasJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelBiasAppSpecification: &modelBiasAppSpecificationProperty{
//   		configUri: jsii.String("configUri"),
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	modelBiasJobInput: &modelBiasJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	modelBiasJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelBiasBaselineConfig: &modelBiasBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnModelBiasJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the job definition was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the job definition.
	AttrJobDefinitionArn() *string
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
	// The name of the bias job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	// Identifies the resources to deploy for a monitoring job.
	JobResources() interface{}
	SetJobResources(val interface{})
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
	// Configures the model bias job to run a specified Docker container image.
	ModelBiasAppSpecification() interface{}
	SetModelBiasAppSpecification(val interface{})
	// The baseline configuration for a model bias job.
	ModelBiasBaselineConfig() interface{}
	SetModelBiasBaselineConfig(val interface{})
	// Inputs for the model bias job.
	ModelBiasJobInput() interface{}
	SetModelBiasJobInput(val interface{})
	// The output configuration for monitoring jobs.
	ModelBiasJobOutputConfig() interface{}
	SetModelBiasJobOutputConfig(val interface{})
	// Networking options for a model bias job.
	NetworkConfig() interface{}
	SetNetworkConfig(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition() interface{}
	SetStoppingCondition(val interface{})
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnModelBiasJobDefinition
type jsiiProxy_CfnModelBiasJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) AttrJobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrJobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) JobResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) ModelBiasAppSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelBiasAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) ModelBiasBaselineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelBiasBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) ModelBiasJobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelBiasJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) ModelBiasJobOutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelBiasJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) NetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) StoppingCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ModelBiasJobDefinition`.
func NewCfnModelBiasJobDefinition(scope awscdk.Construct, id *string, props *CfnModelBiasJobDefinitionProps) CfnModelBiasJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnModelBiasJobDefinition{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelBiasJobDefinition`.
func NewCfnModelBiasJobDefinition_Override(c CfnModelBiasJobDefinition, scope awscdk.Construct, id *string, props *CfnModelBiasJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetJobResources(val interface{}) {
	_jsii_.Set(
		j,
		"jobResources",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetModelBiasAppSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"modelBiasAppSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetModelBiasBaselineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelBiasBaselineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetModelBiasJobInput(val interface{}) {
	_jsii_.Set(
		j,
		"modelBiasJobInput",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetModelBiasJobOutputConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelBiasJobOutputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetNetworkConfig(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnModelBiasJobDefinition) SetStoppingCondition(val interface{}) {
	_jsii_.Set(
		j,
		"stoppingCondition",
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
func CfnModelBiasJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModelBiasJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModelBiasJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelBiasJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnModelBiasJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelBiasJobDefinition) ValidateProperties(_properties interface{}) {
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
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnModelBiasJobDefinition_ClusterConfigProperty struct {
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.InstanceCount`.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.InstanceType`.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.VolumeSizeInGB`.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// `CfnModelBiasJobDefinition.ClusterConfigProperty.VolumeKmsKeyId`.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

// The constraints resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelBiasJobDefinition_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	endTimeOffset: jsii.String("endTimeOffset"),
//   	featuresAttribute: jsii.String("featuresAttribute"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	probabilityThresholdAttribute: jsii.Number(123),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   	startTimeOffset: jsii.String("startTimeOffset"),
//   }
//
type CfnModelBiasJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// If specified, monitoring jobs substract this time from the end time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The attributes of the input data that are the input features.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// The attribute of the input data that represents the ground truth label.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// The threshold for the class probability to be evaluated as a positive result.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// If specified, monitoring jobs substract this time from the start time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

// Docker container image configuration object for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasAppSpecificationProperty := &modelBiasAppSpecificationProperty{
//   	configUri: jsii.String("configUri"),
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasAppSpecificationProperty struct {
	// JSON formatted S3 file that defines bias parameters.
	//
	// For more information on this JSON configuration file, see [Configure bias parameters](https://docs.aws.amazon.com/sagemaker/latest/json-bias-parameter-config.html) .
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the model bias job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
}

// The configuration for a baseline model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasBaselineConfigProperty := &modelBiasBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasBaselineConfigProperty struct {
	// The name of the baseline model bias job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}

// Inputs for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasJobInputProperty := &modelBiasJobInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		endTimeOffset: jsii.String("endTimeOffset"),
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		probabilityThresholdAttribute: jsii.Number(123),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   		startTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasJobInputProperty struct {
	// Input object for the endpoint.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
	// Location of ground truth labels to use in model bias job.
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
}

// The ground truth labels for the dataset used for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringGroundTruthS3InputProperty := &monitoringGroundTruthS3InputProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelBiasJobDefinition_MonitoringGroundTruthS3InputProperty struct {
	// The address of the Amazon S3 location of the ground truth labels.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnModelBiasJobDefinition_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &monitoringResourcesProperty{
//   	clusterConfig: &clusterConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		volumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigProperty := &networkConfigProperty{
//   	enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnModelBiasJobDefinition_NetworkConfigProperty struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose `True` to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// The Amazon S3 storage location where the results of a monitoring job are saved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnModelBiasJobDefinition_S3OutputProperty struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// `LocalPath` is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

// Specifies a limit to how long a model training job or model compilation job can run.
//
// It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.
//
// To stop a training job, SageMaker sends the algorithm the `SIGTERM` signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.
//
// The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with `CreateModel` .
//
// > The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stoppingConditionProperty := &stoppingConditionProperty{
//   	maxRuntimeInSeconds: jsii.Number(123),
//   }
//
type CfnModelBiasJobDefinition_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Properties for defining a `CfnModelBiasJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelBiasJobDefinitionProps := &cfnModelBiasJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelBiasAppSpecification: &modelBiasAppSpecificationProperty{
//   		configUri: jsii.String("configUri"),
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	modelBiasJobInput: &modelBiasJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	modelBiasJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelBiasBaselineConfig: &modelBiasBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelBiasJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Configures the model bias job to run a specified Docker container image.
	ModelBiasAppSpecification interface{} `field:"required" json:"modelBiasAppSpecification" yaml:"modelBiasAppSpecification"`
	// Inputs for the model bias job.
	ModelBiasJobInput interface{} `field:"required" json:"modelBiasJobInput" yaml:"modelBiasJobInput"`
	// The output configuration for monitoring jobs.
	ModelBiasJobOutputConfig interface{} `field:"required" json:"modelBiasJobOutputConfig" yaml:"modelBiasJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the bias job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The baseline configuration for a model bias job.
	ModelBiasBaselineConfig interface{} `field:"optional" json:"modelBiasBaselineConfig" yaml:"modelBiasBaselineConfig"`
	// Networking options for a model bias job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::ModelExplainabilityJobDefinition`.
//
// Creates the definition for a model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelExplainabilityJobDefinition := awscdk.Aws_sagemaker.NewCfnModelExplainabilityJobDefinition(this, jsii.String("MyCfnModelExplainabilityJobDefinition"), &cfnModelExplainabilityJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelExplainabilityAppSpecification: &modelExplainabilityAppSpecificationProperty{
//   		configUri: jsii.String("configUri"),
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	modelExplainabilityJobInput: &modelExplainabilityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	modelExplainabilityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelExplainabilityBaselineConfig: &modelExplainabilityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnModelExplainabilityJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the job definition was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the job definition.
	AttrJobDefinitionArn() *string
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
	// The name of the model explainability job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	// Identifies the resources to deploy for a monitoring job.
	JobResources() interface{}
	SetJobResources(val interface{})
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
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification() interface{}
	SetModelExplainabilityAppSpecification(val interface{})
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig() interface{}
	SetModelExplainabilityBaselineConfig(val interface{})
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput() interface{}
	SetModelExplainabilityJobInput(val interface{})
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig() interface{}
	SetModelExplainabilityJobOutputConfig(val interface{})
	// Networking options for a model explainability job.
	NetworkConfig() interface{}
	SetNetworkConfig(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::SageMaker::ModelExplainabilityJobDefinition.StoppingCondition`.
	StoppingCondition() interface{}
	SetStoppingCondition(val interface{})
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnModelExplainabilityJobDefinition
type jsiiProxy_CfnModelExplainabilityJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) AttrJobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrJobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) JobResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) ModelExplainabilityAppSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) ModelExplainabilityBaselineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) ModelExplainabilityJobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) ModelExplainabilityJobOutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelExplainabilityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) NetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) StoppingCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ModelExplainabilityJobDefinition`.
func NewCfnModelExplainabilityJobDefinition(scope awscdk.Construct, id *string, props *CfnModelExplainabilityJobDefinitionProps) CfnModelExplainabilityJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnModelExplainabilityJobDefinition{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelExplainabilityJobDefinition`.
func NewCfnModelExplainabilityJobDefinition_Override(c CfnModelExplainabilityJobDefinition, scope awscdk.Construct, id *string, props *CfnModelExplainabilityJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetJobResources(val interface{}) {
	_jsii_.Set(
		j,
		"jobResources",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetModelExplainabilityAppSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"modelExplainabilityAppSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetModelExplainabilityBaselineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelExplainabilityBaselineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetModelExplainabilityJobInput(val interface{}) {
	_jsii_.Set(
		j,
		"modelExplainabilityJobInput",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetModelExplainabilityJobOutputConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelExplainabilityJobOutputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetNetworkConfig(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnModelExplainabilityJobDefinition) SetStoppingCondition(val interface{}) {
	_jsii_.Set(
		j,
		"stoppingCondition",
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
func CfnModelExplainabilityJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModelExplainabilityJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModelExplainabilityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelExplainabilityJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnModelExplainabilityJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelExplainabilityJobDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration for the cluster resources used to run the processing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnModelExplainabilityJobDefinition_ClusterConfigProperty struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

// The constraints resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelExplainabilityJobDefinition_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	featuresAttribute: jsii.String("featuresAttribute"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnModelExplainabilityJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// The attributes of the input data that are the input features.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// The attribute of the input data that represents the ground truth label.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

// Docker container image configuration object for the model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelExplainabilityAppSpecificationProperty := &modelExplainabilityAppSpecificationProperty{
//   	configUri: jsii.String("configUri"),
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityAppSpecificationProperty struct {
	// JSON formatted S3 file that defines explainability parameters.
	//
	// For more information on this JSON configuration file, see [Configure model explainability parameters](https://docs.aws.amazon.com/sagemaker/latest/json-model-explainability-parameter-config.html) .
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the model explainability job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
}

// The configuration for a baseline model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelExplainabilityBaselineConfigProperty := &modelExplainabilityBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityBaselineConfigProperty struct {
	// The name of the baseline model explainability job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a model explainability job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}

// Inputs for the model explainability job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelExplainabilityJobInputProperty := &modelExplainabilityJobInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		featuresAttribute: jsii.String("featuresAttribute"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_ModelExplainabilityJobInputProperty struct {
	// `CfnModelExplainabilityJobDefinition.ModelExplainabilityJobInputProperty.EndpointInput`.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
}

// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnModelExplainabilityJobDefinition_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &monitoringResourcesProperty{
//   	clusterConfig: &clusterConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		volumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigProperty := &networkConfigProperty{
//   	enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_NetworkConfigProperty struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose `True` to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// `CfnModelExplainabilityJobDefinition.NetworkConfigProperty.VpcConfig`.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// The Amazon S3 storage location where the results of a monitoring job are saved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnModelExplainabilityJobDefinition_S3OutputProperty struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

// Specifies a limit to how long a model training job or model compilation job can run.
//
// It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.
//
// To stop a training job, SageMaker sends the algorithm the `SIGTERM` signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.
//
// The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with `CreateModel` .
//
// > The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stoppingConditionProperty := &stoppingConditionProperty{
//   	maxRuntimeInSeconds: jsii.Number(123),
//   }
//
type CfnModelExplainabilityJobDefinition_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnModelExplainabilityJobDefinition_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Properties for defining a `CfnModelExplainabilityJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelExplainabilityJobDefinitionProps := &cfnModelExplainabilityJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelExplainabilityAppSpecification: &modelExplainabilityAppSpecificationProperty{
//   		configUri: jsii.String("configUri"),
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	modelExplainabilityJobInput: &modelExplainabilityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	modelExplainabilityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelExplainabilityBaselineConfig: &modelExplainabilityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelExplainabilityJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification interface{} `field:"required" json:"modelExplainabilityAppSpecification" yaml:"modelExplainabilityAppSpecification"`
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput interface{} `field:"required" json:"modelExplainabilityJobInput" yaml:"modelExplainabilityJobInput"`
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig interface{} `field:"required" json:"modelExplainabilityJobOutputConfig" yaml:"modelExplainabilityJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the model explainability job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig interface{} `field:"optional" json:"modelExplainabilityBaselineConfig" yaml:"modelExplainabilityBaselineConfig"`
	// Networking options for a model explainability job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// `AWS::SageMaker::ModelExplainabilityJobDefinition.StoppingCondition`.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::ModelPackageGroup`.
//
// A group of versioned models in the model registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelPackageGroupPolicy interface{}
//
//   cfnModelPackageGroup := awscdk.Aws_sagemaker.NewCfnModelPackageGroup(this, jsii.String("MyCfnModelPackageGroup"), &cfnModelPackageGroupProps{
//   	modelPackageGroupName: jsii.String("modelPackageGroupName"),
//
//   	// the properties below are optional
//   	modelPackageGroupDescription: jsii.String("modelPackageGroupDescription"),
//   	modelPackageGroupPolicy: modelPackageGroupPolicy,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnModelPackageGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the model group was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the model group.
	AttrModelPackageGroupArn() *string
	// The status of the model group.
	AttrModelPackageGroupStatus() *string
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
	// The description for the model group.
	ModelPackageGroupDescription() *string
	SetModelPackageGroupDescription(val *string)
	// The name of the model group.
	ModelPackageGroupName() *string
	SetModelPackageGroupName(val *string)
	// A resouce policy to control access to a model group.
	//
	// For information about resoure policies, see [Identity-based policies and resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html) in the *AWS Identity and Access Management User Guide.* .
	ModelPackageGroupPolicy() interface{}
	SetModelPackageGroupPolicy(val interface{})
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
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnModelPackageGroup
type jsiiProxy_CfnModelPackageGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModelPackageGroup) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) AttrModelPackageGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelPackageGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) AttrModelPackageGroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelPackageGroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) ModelPackageGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) ModelPackageGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) ModelPackageGroupPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageGroupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackageGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ModelPackageGroup`.
func NewCfnModelPackageGroup(scope awscdk.Construct, id *string, props *CfnModelPackageGroupProps) CfnModelPackageGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnModelPackageGroup{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelPackageGroup`.
func NewCfnModelPackageGroup_Override(c CfnModelPackageGroup, scope awscdk.Construct, id *string, props *CfnModelPackageGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelPackageGroup) SetModelPackageGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"modelPackageGroupDescription",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackageGroup) SetModelPackageGroupName(val *string) {
	_jsii_.Set(
		j,
		"modelPackageGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackageGroup) SetModelPackageGroupPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"modelPackageGroupPolicy",
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
func CfnModelPackageGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModelPackageGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModelPackageGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelPackageGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnModelPackageGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelPackageGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackageGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnModelPackageGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelPackageGroupPolicy interface{}
//
//   cfnModelPackageGroupProps := &cfnModelPackageGroupProps{
//   	modelPackageGroupName: jsii.String("modelPackageGroupName"),
//
//   	// the properties below are optional
//   	modelPackageGroupDescription: jsii.String("modelPackageGroupDescription"),
//   	modelPackageGroupPolicy: modelPackageGroupPolicy,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelPackageGroupProps struct {
	// The name of the model group.
	ModelPackageGroupName *string `field:"required" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The description for the model group.
	ModelPackageGroupDescription *string `field:"optional" json:"modelPackageGroupDescription" yaml:"modelPackageGroupDescription"`
	// A resouce policy to control access to a model group.
	//
	// For information about resoure policies, see [Identity-based policies and resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html) in the *AWS Identity and Access Management User Guide.* .
	ModelPackageGroupPolicy interface{} `field:"optional" json:"modelPackageGroupPolicy" yaml:"modelPackageGroupPolicy"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var environment interface{}
//
//   cfnModelProps := &cfnModelProps{
//   	executionRoleArn: jsii.String("executionRoleArn"),
//
//   	// the properties below are optional
//   	containers: []interface{}{
//   		&containerDefinitionProperty{
//   			containerHostname: jsii.String("containerHostname"),
//   			environment: environment,
//   			image: jsii.String("image"),
//   			imageConfig: &imageConfigProperty{
//   				repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   				// the properties below are optional
//   				repositoryAuthConfig: &repositoryAuthConfigProperty{
//   					repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   				},
//   			},
//   			inferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   			mode: jsii.String("mode"),
//   			modelDataUrl: jsii.String("modelDataUrl"),
//   			modelPackageName: jsii.String("modelPackageName"),
//   			multiModelConfig: &multiModelConfigProperty{
//   				modelCacheSetting: jsii.String("modelCacheSetting"),
//   			},
//   		},
//   	},
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	inferenceExecutionConfig: &inferenceExecutionConfigProperty{
//   		mode: jsii.String("mode"),
//   	},
//   	modelName: jsii.String("modelName"),
//   	primaryContainer: &containerDefinitionProperty{
//   		containerHostname: jsii.String("containerHostname"),
//   		environment: environment,
//   		image: jsii.String("image"),
//   		imageConfig: &imageConfigProperty{
//   			repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   			// the properties below are optional
//   			repositoryAuthConfig: &repositoryAuthConfigProperty{
//   				repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   			},
//   		},
//   		inferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   		mode: jsii.String("mode"),
//   		modelDataUrl: jsii.String("modelDataUrl"),
//   		modelPackageName: jsii.String("modelPackageName"),
//   		multiModelConfig: &multiModelConfigProperty{
//   			modelCacheSetting: jsii.String("modelCacheSetting"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnModelProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs.
	//
	// Deploying on ML compute instances is part of model hosting. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies the containers in the inference pipeline.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Isolates the model container.
	//
	// No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig interface{} `field:"optional" json:"inferenceExecutionConfig" yaml:"inferenceExecutionConfig"`
	// The name of the new model.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.
	PrimaryContainer interface{} `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_VpcConfig.html) object that specifies the VPC that you want your model to connect to. Control access to and from your model container by configuring the VPC. `VpcConfig` is used in hosting services and in batch transform. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html) .
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// A CloudFormation `AWS::SageMaker::ModelQualityJobDefinition`.
//
// Creates a definition for a job that monitors model quality and drift. For information about model monitor, see [Amazon SageMaker Model Monitor](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelQualityJobDefinition := awscdk.Aws_sagemaker.NewCfnModelQualityJobDefinition(this, jsii.String("MyCfnModelQualityJobDefinition"), &cfnModelQualityJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelQualityAppSpecification: &modelQualityAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//   		problemType: jsii.String("problemType"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	modelQualityJobInput: &modelQualityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	modelQualityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelQualityBaselineConfig: &modelQualityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnModelQualityJobDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the job definition was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the job definition.
	AttrJobDefinitionArn() *string
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
	// The name of the monitoring job definition.
	JobDefinitionName() *string
	SetJobDefinitionName(val *string)
	// Identifies the resources to deploy for a monitoring job.
	JobResources() interface{}
	SetJobResources(val interface{})
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
	// Container image configuration object for the monitoring job.
	ModelQualityAppSpecification() interface{}
	SetModelQualityAppSpecification(val interface{})
	// Specifies the constraints and baselines for the monitoring job.
	ModelQualityBaselineConfig() interface{}
	SetModelQualityBaselineConfig(val interface{})
	// A list of the inputs that are monitored.
	//
	// Currently endpoints are supported.
	ModelQualityJobInput() interface{}
	SetModelQualityJobInput(val interface{})
	// The output configuration for monitoring jobs.
	ModelQualityJobOutputConfig() interface{}
	SetModelQualityJobOutputConfig(val interface{})
	// Specifies the network configuration for the monitoring job.
	NetworkConfig() interface{}
	SetNetworkConfig(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition() interface{}
	SetStoppingCondition(val interface{})
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnModelQualityJobDefinition
type jsiiProxy_CfnModelQualityJobDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) AttrJobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrJobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) JobResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) ModelQualityAppSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) ModelQualityBaselineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityBaselineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) ModelQualityJobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) ModelQualityJobOutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityJobOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) NetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) StoppingCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::ModelQualityJobDefinition`.
func NewCfnModelQualityJobDefinition(scope awscdk.Construct, id *string, props *CfnModelQualityJobDefinitionProps) CfnModelQualityJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_CfnModelQualityJobDefinition{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::ModelQualityJobDefinition`.
func NewCfnModelQualityJobDefinition_Override(c CfnModelQualityJobDefinition, scope awscdk.Construct, id *string, props *CfnModelQualityJobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetJobDefinitionName(val *string) {
	_jsii_.Set(
		j,
		"jobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetJobResources(val interface{}) {
	_jsii_.Set(
		j,
		"jobResources",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetModelQualityAppSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"modelQualityAppSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetModelQualityBaselineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelQualityBaselineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetModelQualityJobInput(val interface{}) {
	_jsii_.Set(
		j,
		"modelQualityJobInput",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetModelQualityJobOutputConfig(val interface{}) {
	_jsii_.Set(
		j,
		"modelQualityJobOutputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetNetworkConfig(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnModelQualityJobDefinition) SetStoppingCondition(val interface{}) {
	_jsii_.Set(
		j,
		"stoppingCondition",
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
func CfnModelQualityJobDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModelQualityJobDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnModelQualityJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelQualityJobDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnModelQualityJobDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The configuration for the cluster of resources used to run the processing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnModelQualityJobDefinition_ClusterConfigProperty struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

// The constraints resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelQualityJobDefinition_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	endTimeOffset: jsii.String("endTimeOffset"),
//   	inferenceAttribute: jsii.String("inferenceAttribute"),
//   	probabilityAttribute: jsii.String("probabilityAttribute"),
//   	probabilityThresholdAttribute: jsii.Number(123),
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   	startTimeOffset: jsii.String("startTimeOffset"),
//   }
//
type CfnModelQualityJobDefinition_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// If specified, monitoring jobs substract this time from the end time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// The attribute of the input data that represents the ground truth label.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// In a classification problem, the attribute that represents the class probability.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// The threshold for the class probability to be evaluated as a positive result.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// If specified, monitoring jobs substract this time from the start time.
	//
	// For information about using offsets for scheduling monitoring jobs, see [Schedule Model Quality Monitoring Jobs](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-model-quality-schedule.html) .
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

// Container image configuration object for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityAppSpecificationProperty := &modelQualityAppSpecificationProperty{
//   	imageUri: jsii.String("imageUri"),
//   	problemType: jsii.String("problemType"),
//
//   	// the properties below are optional
//   	containerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	containerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
type CfnModelQualityJobDefinition_ModelQualityAppSpecificationProperty struct {
	// The address of the container image that the monitoring job runs.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The machine learning problem type of the model that the monitoring job monitors.
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// An array of arguments for the container used to run the monitoring job.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container that the monitoring job runs.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Sets the environment variables in the container that the monitoring job runs.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

// Configuration for monitoring constraints and monitoring statistics.
//
// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityBaselineConfigProperty := &modelQualityBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_ModelQualityBaselineConfigProperty struct {
	// The name of the job that performs baselining for the monitoring job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}

// The input for the model quality monitoring job.
//
// Currently endponts are supported for input for model quality monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityJobInputProperty := &modelQualityJobInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		endTimeOffset: jsii.String("endTimeOffset"),
//   		inferenceAttribute: jsii.String("inferenceAttribute"),
//   		probabilityAttribute: jsii.String("probabilityAttribute"),
//   		probabilityThresholdAttribute: jsii.Number(123),
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   		startTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_ModelQualityJobInputProperty struct {
	// Input object for the endpoint.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
	// The ground truth label provided for the model.
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
}

// The ground truth labels for the dataset used for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringGroundTruthS3InputProperty := &monitoringGroundTruthS3InputProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelQualityJobDefinition_MonitoringGroundTruthS3InputProperty struct {
	// The address of the Amazon S3 location of the ground truth labels.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnModelQualityJobDefinition_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &monitoringResourcesProperty{
//   	clusterConfig: &clusterConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		volumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigProperty := &networkConfigProperty{
//   	enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnModelQualityJobDefinition_NetworkConfigProperty struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose `True` to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// The Amazon S3 storage location where the results of a monitoring job are saved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnModelQualityJobDefinition_S3OutputProperty struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

// Specifies a limit to how long a model training job or model compilation job can run.
//
// It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.
//
// To stop a training job, SageMaker sends the algorithm the `SIGTERM` signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.
//
// The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with `CreateModel` .
//
// > The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stoppingConditionProperty := &stoppingConditionProperty{
//   	maxRuntimeInSeconds: jsii.Number(123),
//   }
//
type CfnModelQualityJobDefinition_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Properties for defining a `CfnModelQualityJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelQualityJobDefinitionProps := &cfnModelQualityJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelQualityAppSpecification: &modelQualityAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//   		problemType: jsii.String("problemType"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	modelQualityJobInput: &modelQualityJobInputProperty{
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	modelQualityJobOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelQualityBaselineConfig: &modelQualityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelQualityJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Container image configuration object for the monitoring job.
	ModelQualityAppSpecification interface{} `field:"required" json:"modelQualityAppSpecification" yaml:"modelQualityAppSpecification"`
	// A list of the inputs that are monitored.
	//
	// Currently endpoints are supported.
	ModelQualityJobInput interface{} `field:"required" json:"modelQualityJobInput" yaml:"modelQualityJobInput"`
	// The output configuration for monitoring jobs.
	ModelQualityJobOutputConfig interface{} `field:"required" json:"modelQualityJobOutputConfig" yaml:"modelQualityJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the monitoring job definition.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Specifies the constraints and baselines for the monitoring job.
	ModelQualityBaselineConfig interface{} `field:"optional" json:"modelQualityBaselineConfig" yaml:"modelQualityBaselineConfig"`
	// Specifies the network configuration for the monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::MonitoringSchedule`.
//
// The `AWS::SageMaker::MonitoringSchedule` resource is an Amazon SageMaker resource type that regularly starts SageMaker processing Jobs to monitor the data captured for a SageMaker endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringSchedule := awscdk.Aws_sagemaker.NewCfnMonitoringSchedule(this, jsii.String("MyCfnMonitoringSchedule"), &cfnMonitoringScheduleProps{
//   	monitoringScheduleConfig: &monitoringScheduleConfigProperty{
//   		monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   			monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   				imageUri: jsii.String("imageUri"),
//
//   				// the properties below are optional
//   				containerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				containerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
//   				postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   				recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   			},
//   			monitoringInputs: []interface{}{
//   				&monitoringInputProperty{
//   					endpointInput: &endpointInputProperty{
//   						endpointName: jsii.String("endpointName"),
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			monitoringOutputConfig: &monitoringOutputConfigProperty{
//   				monitoringOutputs: []interface{}{
//   					&monitoringOutputProperty{
//   						s3Output: &s3OutputProperty{
//   							localPath: jsii.String("localPath"),
//   							s3Uri: jsii.String("s3Uri"),
//
//   							// the properties below are optional
//   							s3UploadMode: jsii.String("s3UploadMode"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   			monitoringResources: &monitoringResourcesProperty{
//   				clusterConfig: &clusterConfigProperty{
//   					instanceCount: jsii.Number(123),
//   					instanceType: jsii.String("instanceType"),
//   					volumeSizeInGb: jsii.Number(123),
//
//   					// the properties below are optional
//   					volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			baselineConfig: &baselineConfigProperty{
//   				constraintsResource: &constraintsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   				statisticsResource: &statisticsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			networkConfig: &networkConfigProperty{
//   				enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   				enableNetworkIsolation: jsii.Boolean(false),
//   				vpcConfig: &vpcConfigProperty{
//   					securityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			stoppingCondition: &stoppingConditionProperty{
//   				maxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		monitoringType: jsii.String("monitoringType"),
//   		scheduleConfig: &scheduleConfigProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//   		},
//   	},
//   	monitoringScheduleName: jsii.String("monitoringScheduleName"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	failureReason: jsii.String("failureReason"),
//   	lastMonitoringExecutionSummary: &monitoringExecutionSummaryProperty{
//   		creationTime: jsii.String("creationTime"),
//   		lastModifiedTime: jsii.String("lastModifiedTime"),
//   		monitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		monitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		scheduledTime: jsii.String("scheduledTime"),
//
//   		// the properties below are optional
//   		endpointName: jsii.String("endpointName"),
//   		failureReason: jsii.String("failureReason"),
//   		processingJobArn: jsii.String("processingJobArn"),
//   	},
//   	monitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnMonitoringSchedule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time when the monitoring schedule was created.
	AttrCreationTime() *string
	// The last time that the monitoring schedule was modified.
	AttrLastModifiedTime() *string
	// The Amazon Resource Name (ARN) of the monitoring schedule.
	AttrMonitoringScheduleArn() *string
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
	// The name of the endpoint using the monitoring schedule.
	EndpointName() *string
	SetEndpointName(val *string)
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason() *string
	SetFailureReason(val *string)
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary() interface{}
	SetLastMonitoringExecutionSummary(val interface{})
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
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	MonitoringScheduleConfig() interface{}
	SetMonitoringScheduleConfig(val interface{})
	// The name of the monitoring schedule.
	MonitoringScheduleName() *string
	SetMonitoringScheduleName(val *string)
	// The status of the monitoring schedule.
	MonitoringScheduleStatus() *string
	SetMonitoringScheduleStatus(val *string)
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
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnMonitoringSchedule
type jsiiProxy_CfnMonitoringSchedule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) AttrMonitoringScheduleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMonitoringScheduleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) LastMonitoringExecutionSummary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastMonitoringExecutionSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringScheduleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) MonitoringScheduleStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedule(scope awscdk.Construct, id *string, props *CfnMonitoringScheduleProps) CfnMonitoringSchedule {
	_init_.Initialize()

	j := jsiiProxy_CfnMonitoringSchedule{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedule_Override(c CfnMonitoringSchedule, scope awscdk.Construct, id *string, props *CfnMonitoringScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetEndpointName(val *string) {
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetFailureReason(val *string) {
	_jsii_.Set(
		j,
		"failureReason",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetLastMonitoringExecutionSummary(val interface{}) {
	_jsii_.Set(
		j,
		"lastMonitoringExecutionSummary",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetMonitoringScheduleConfig(val interface{}) {
	_jsii_.Set(
		j,
		"monitoringScheduleConfig",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetMonitoringScheduleName(val *string) {
	_jsii_.Set(
		j,
		"monitoringScheduleName",
		val,
	)
}

func (j *jsiiProxy_CfnMonitoringSchedule) SetMonitoringScheduleStatus(val *string) {
	_jsii_.Set(
		j,
		"monitoringScheduleStatus",
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
func CfnMonitoringSchedule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMonitoringSchedule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMonitoringSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMonitoringSchedule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnMonitoringSchedule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnMonitoringSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineConfigProperty := &baselineConfigProperty{
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   	statisticsResource: &statisticsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnMonitoringSchedule_BaselineConfigProperty struct {
	// The Amazon S3 URI for the constraints resource.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

// Configuration for the cluster used to run model monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	volumeSizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   }
//
type CfnMonitoringSchedule_ClusterConfigProperty struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

// The Amazon S3 URI for the constraints resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnMonitoringSchedule_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnMonitoringSchedule_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

// Container image configuration object for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringAppSpecificationProperty := &monitoringAppSpecificationProperty{
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	containerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	containerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
type CfnMonitoringSchedule_MonitoringAppSpecificationProperty struct {
	// The container image to be run by the monitoring job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// An array of arguments for the container used to run the monitoring job.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container used to run the monitoring job.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

// Summary of information about the last monitoring job to run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringExecutionSummaryProperty := &monitoringExecutionSummaryProperty{
//   	creationTime: jsii.String("creationTime"),
//   	lastModifiedTime: jsii.String("lastModifiedTime"),
//   	monitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   	monitoringScheduleName: jsii.String("monitoringScheduleName"),
//   	scheduledTime: jsii.String("scheduledTime"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	failureReason: jsii.String("failureReason"),
//   	processingJobArn: jsii.String("processingJobArn"),
//   }
//
type CfnMonitoringSchedule_MonitoringExecutionSummaryProperty struct {
	// The time at which the monitoring job was created.
	CreationTime *string `field:"required" json:"creationTime" yaml:"creationTime"`
	// A timestamp that indicates the last time the monitoring job was modified.
	LastModifiedTime *string `field:"required" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// The status of the monitoring job.
	MonitoringExecutionStatus *string `field:"required" json:"monitoringExecutionStatus" yaml:"monitoringExecutionStatus"`
	// The name of the monitoring schedule.
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The time the monitoring job was scheduled.
	ScheduledTime *string `field:"required" json:"scheduledTime" yaml:"scheduledTime"`
	// The name of the endpoint used to run the monitoring job.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// The Amazon Resource Name (ARN) of the monitoring job.
	ProcessingJobArn *string `field:"optional" json:"processingJobArn" yaml:"processingJobArn"`
}

// The inputs for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringInputProperty := &monitoringInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringInputProperty struct {
	// The endpoint for a monitoring job.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
}

// Defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringJobDefinitionProperty := &monitoringJobDefinitionProperty{
//   	monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	monitoringInputs: []interface{}{
//   		&monitoringInputProperty{
//   			endpointInput: &endpointInputProperty{
//   				endpointName: jsii.String("endpointName"),
//   				localPath: jsii.String("localPath"),
//
//   				// the properties below are optional
//   				s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				s3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	monitoringOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	monitoringResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	baselineConfig: &baselineConfigProperty{
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   		statisticsResource: &statisticsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// Configures the monitoring job to run a specified Docker container image.
	MonitoringAppSpecification interface{} `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// The array of inputs for the monitoring job.
	//
	// Currently we support monitoring an Amazon SageMaker Endpoint.
	MonitoringInputs interface{} `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// The array of outputs from the monitoring job to be uploaded to Amazon Simple Storage Service (Amazon S3).
	MonitoringOutputConfig interface{} `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a monitoring job.
	//
	// In distributed processing, you specify more than one instance.
	MonitoringResources interface{} `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	BaselineConfig interface{} `field:"optional" json:"baselineConfig" yaml:"baselineConfig"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Specifies networking options for an monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

// The output configuration for monitoring jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputConfigProperty := &monitoringOutputConfigProperty{
//   	monitoringOutputs: []interface{}{
//   		&monitoringOutputProperty{
//   			s3Output: &s3OutputProperty{
//   				localPath: jsii.String("localPath"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				s3UploadMode: jsii.String("s3UploadMode"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnMonitoringSchedule_MonitoringOutputConfigProperty struct {
	// Monitoring outputs for monitoring jobs.
	//
	// This is where the output of the periodic monitoring jobs is uploaded.
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

// Identifies the resources to deploy for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringResourcesProperty := &monitoringResourcesProperty{
//   	clusterConfig: &clusterConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		volumeSizeInGb: jsii.Number(123),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringResourcesProperty struct {
	// The configuration for the cluster resources used to run the processing job.
	ClusterConfig interface{} `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

// Configures the monitoring schedule and defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringScheduleConfigProperty := &monitoringScheduleConfigProperty{
//   	monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   		monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   			imageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
//   			containerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			containerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   			recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   		},
//   		monitoringInputs: []interface{}{
//   			&monitoringInputProperty{
//   				endpointInput: &endpointInputProperty{
//   					endpointName: jsii.String("endpointName"),
//   					localPath: jsii.String("localPath"),
//
//   					// the properties below are optional
//   					s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   					s3InputMode: jsii.String("s3InputMode"),
//   				},
//   			},
//   		},
//   		monitoringOutputConfig: &monitoringOutputConfigProperty{
//   			monitoringOutputs: []interface{}{
//   				&monitoringOutputProperty{
//   					s3Output: &s3OutputProperty{
//   						localPath: jsii.String("localPath"),
//   						s3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						s3UploadMode: jsii.String("s3UploadMode"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   		monitoringResources: &monitoringResourcesProperty{
//   			clusterConfig: &clusterConfigProperty{
//   				instanceCount: jsii.Number(123),
//   				instanceType: jsii.String("instanceType"),
//   				volumeSizeInGb: jsii.Number(123),
//
//   				// the properties below are optional
//   				volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		baselineConfig: &baselineConfigProperty{
//   			constraintsResource: &constraintsResourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//   			},
//   			statisticsResource: &statisticsResourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		networkConfig: &networkConfigProperty{
//   			enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   			enableNetworkIsolation: jsii.Boolean(false),
//   			vpcConfig: &vpcConfigProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		stoppingCondition: &stoppingConditionProperty{
//   			maxRuntimeInSeconds: jsii.Number(123),
//   		},
//   	},
//   	monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   	monitoringType: jsii.String("monitoringType"),
//   	scheduleConfig: &scheduleConfigProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringScheduleConfigProperty struct {
	// Defines the monitoring job.
	MonitoringJobDefinition interface{} `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// The name of the monitoring job definition to schedule.
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// The type of the monitoring job definition to schedule.
	MonitoringType *string `field:"optional" json:"monitoringType" yaml:"monitoringType"`
	// Configures the monitoring schedule.
	ScheduleConfig interface{} `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigProperty := &networkConfigProperty{
//   	enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   	enableNetworkIsolation: jsii.Boolean(false),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnMonitoringSchedule_NetworkConfigProperty struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose `True` to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

// Information about where and how you want to store the results of a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnMonitoringSchedule_S3OutputProperty struct {
	// The local path to the S3 storage location where SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the S3 storage location where SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

// Configuration details about the monitoring schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigProperty := &scheduleConfigProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnMonitoringSchedule_ScheduleConfigProperty struct {
	// A cron expression that describes details about the monitoring schedule.
	//
	// Currently the only supported cron expressions are:
	//
	// - If you want to set the job to start every hour, please use the following:
	//
	// `Hourly: cron(0 * ? * * *)`
	// - If you want to start the job daily:
	//
	// `cron(0 [00-23] ? * * *)`
	//
	// For example, the following are valid cron expressions:
	//
	// - Daily at noon UTC: `cron(0 12 ? * * *)`
	// - Daily at midnight UTC: `cron(0 0 ? * * *)`
	//
	// To support running every 6, 12 hours, the following are also supported:
	//
	// `cron(0 [00-23]/[01-24] ? * * *)`
	//
	// For example, the following are valid cron expressions:
	//
	// - Every 12 hours, starting at 5pm UTC: `cron(0 17/12 ? * * *)`
	// - Every two hours starting at midnight: `cron(0 0/2 ? * * *)`
	//
	// > - Even though the cron expression is set to start at 5PM UTC, note that there could be a delay of 0-20 minutes from the actual requested time to run the execution.
	// > - We recommend that if you would like a daily schedule, you do not provide this parameter. Amazon SageMaker will pick a time for running every day.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

// The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticsResourceProperty := &statisticsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnMonitoringSchedule_StatisticsResourceProperty struct {
	// The S3 URI for the statistics resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

// Specifies a limit to how long a model training job or model compilation job can run.
//
// It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.
//
// To stop a training job, SageMaker sends the algorithm the `SIGTERM` signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.
//
// The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with `CreateModel` .
//
// > The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stoppingConditionProperty := &stoppingConditionProperty{
//   	maxRuntimeInSeconds: jsii.Number(123),
//   }
//
type CfnMonitoringSchedule_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

// Specifies a VPC that your training jobs and hosted models have access to.
//
// Control access to and from your training and model containers by configuring the VPC. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Training Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnMonitoringSchedule_VpcConfigProperty struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the `Subnets` field.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	//
	// For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html) .
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Properties for defining a `CfnMonitoringSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringScheduleProps := &cfnMonitoringScheduleProps{
//   	monitoringScheduleConfig: &monitoringScheduleConfigProperty{
//   		monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   			monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   				imageUri: jsii.String("imageUri"),
//
//   				// the properties below are optional
//   				containerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				containerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
//   				postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   				recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   			},
//   			monitoringInputs: []interface{}{
//   				&monitoringInputProperty{
//   					endpointInput: &endpointInputProperty{
//   						endpointName: jsii.String("endpointName"),
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			monitoringOutputConfig: &monitoringOutputConfigProperty{
//   				monitoringOutputs: []interface{}{
//   					&monitoringOutputProperty{
//   						s3Output: &s3OutputProperty{
//   							localPath: jsii.String("localPath"),
//   							s3Uri: jsii.String("s3Uri"),
//
//   							// the properties below are optional
//   							s3UploadMode: jsii.String("s3UploadMode"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   			monitoringResources: &monitoringResourcesProperty{
//   				clusterConfig: &clusterConfigProperty{
//   					instanceCount: jsii.Number(123),
//   					instanceType: jsii.String("instanceType"),
//   					volumeSizeInGb: jsii.Number(123),
//
//   					// the properties below are optional
//   					volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			baselineConfig: &baselineConfigProperty{
//   				constraintsResource: &constraintsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   				statisticsResource: &statisticsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			networkConfig: &networkConfigProperty{
//   				enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   				enableNetworkIsolation: jsii.Boolean(false),
//   				vpcConfig: &vpcConfigProperty{
//   					securityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			stoppingCondition: &stoppingConditionProperty{
//   				maxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		monitoringType: jsii.String("monitoringType"),
//   		scheduleConfig: &scheduleConfigProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//   		},
//   	},
//   	monitoringScheduleName: jsii.String("monitoringScheduleName"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	failureReason: jsii.String("failureReason"),
//   	lastMonitoringExecutionSummary: &monitoringExecutionSummaryProperty{
//   		creationTime: jsii.String("creationTime"),
//   		lastModifiedTime: jsii.String("lastModifiedTime"),
//   		monitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		monitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		scheduledTime: jsii.String("scheduledTime"),
//
//   		// the properties below are optional
//   		endpointName: jsii.String("endpointName"),
//   		failureReason: jsii.String("failureReason"),
//   		processingJobArn: jsii.String("processingJobArn"),
//   	},
//   	monitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMonitoringScheduleProps struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	MonitoringScheduleConfig interface{} `field:"required" json:"monitoringScheduleConfig" yaml:"monitoringScheduleConfig"`
	// The name of the monitoring schedule.
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The name of the endpoint using the monitoring schedule.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary interface{} `field:"optional" json:"lastMonitoringExecutionSummary" yaml:"lastMonitoringExecutionSummary"`
	// The status of the monitoring schedule.
	MonitoringScheduleStatus *string `field:"optional" json:"monitoringScheduleStatus" yaml:"monitoringScheduleStatus"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::NotebookInstance`.
//
// The `AWS::SageMaker::NotebookInstance` resource creates an Amazon SageMaker notebook instance. A notebook instance is a machine learning (ML) compute instance running on a Jupyter notebook. For more information, see [Use Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstance := awscdk.Aws_sagemaker.NewCfnNotebookInstance(this, jsii.String("MyCfnNotebookInstance"), &cfnNotebookInstanceProps{
//   	instanceType: jsii.String("instanceType"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	acceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	additionalCodeRepositories: []*string{
//   		jsii.String("additionalCodeRepositories"),
//   	},
//   	defaultCodeRepository: jsii.String("defaultCodeRepository"),
//   	directInternetAccess: jsii.String("directInternetAccess"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	lifecycleConfigName: jsii.String("lifecycleConfigName"),
//   	notebookInstanceName: jsii.String("notebookInstanceName"),
//   	platformIdentifier: jsii.String("platformIdentifier"),
//   	rootAccess: jsii.String("rootAccess"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeSizeInGb: jsii.Number(123),
//   })
//
type CfnNotebookInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance.
	//
	// Currently, only one instance type can be associated with a notebook instance. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	//
	// *Valid Values:* `ml.eia1.medium | ml.eia1.large | ml.eia1.xlarge | ml.eia2.medium | ml.eia2.large | ml.eia2.xlarge` .
	AcceleratorTypes() *[]*string
	SetAcceleratorTypes(val *[]*string)
	// An array of up to three Git repositories associated with the notebook instance.
	//
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	AdditionalCodeRepositories() *[]*string
	SetAdditionalCodeRepositories(val *[]*string)
	// The name of the notebook instance, such as `MyNotebookInstance` .
	AttrNotebookInstanceName() *string
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
	// The Git repository associated with the notebook instance as its default code repository.
	//
	// This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	DefaultCodeRepository() *string
	SetDefaultCodeRepository(val *string)
	// Sets whether SageMaker provides internet access to the notebook instance.
	//
	// If you set this to `Disabled` this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker training and endpoint services unless you configure a NAT Gateway in your VPC.
	//
	// For more information, see [Notebook Instances Are Internet-Enabled by Default](https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access) . You can set the value of this parameter to `Disabled` only if you set a value for the `SubnetId` parameter.
	DirectInternetAccess() *string
	SetDirectInternetAccess(val *string)
	// The type of ML compute instance to launch for the notebook instance.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	InstanceType() *string
	SetInstanceType(val *string)
	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to your notebook instance.
	//
	// The KMS key you provide must be enabled. For information, see [Enabling and Disabling Keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html) in the *AWS Key Management Service Developer Guide* .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The name of a lifecycle configuration to associate with the notebook instance.
	//
	// For information about lifecycle configurations, see [Customize a Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the *Amazon SageMaker Developer Guide* .
	LifecycleConfigName() *string
	SetLifecycleConfigName(val *string)
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
	// The name of the new notebook instance.
	NotebookInstanceName() *string
	SetNotebookInstanceName(val *string)
	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier() *string
	SetPlatformIdentifier(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// When you send any requests to AWS resources from the notebook instance, SageMaker assumes this role to perform tasks on your behalf.
	//
	// You must grant this role necessary permissions so SageMaker can perform these tasks. The policy must allow the SageMaker service principal (sagemaker.amazonaws.com) permissions to assume this role. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	RoleArn() *string
	SetRoleArn(val *string)
	// Whether root access is enabled or disabled for users of the notebook instance. The default value is `Enabled` .
	//
	// > Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users.
	RootAccess() *string
	SetRootAccess(val *string)
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// The security groups must be for the same VPC as specified in the subnet.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.
	SubnetId() *string
	SetSubnetId(val *string)
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	//
	// You can add tags later by using the `CreateTags` API.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	//
	// The default value is 5 GB.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	VolumeSizeInGb() *float64
	SetVolumeSizeInGb(val *float64)
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

// The jsii proxy struct for CfnNotebookInstance
type jsiiProxy_CfnNotebookInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnNotebookInstance) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) AdditionalCodeRepositories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalCodeRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) AttrNotebookInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNotebookInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) DefaultCodeRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCodeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) DirectInternetAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) LifecycleConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) NotebookInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) PlatformIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) RootAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstance) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::NotebookInstance`.
func NewCfnNotebookInstance(scope awscdk.Construct, id *string, props *CfnNotebookInstanceProps) CfnNotebookInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnNotebookInstance{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::NotebookInstance`.
func NewCfnNotebookInstance_Override(c CfnNotebookInstance, scope awscdk.Construct, id *string, props *CfnNotebookInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetAcceleratorTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"acceleratorTypes",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetAdditionalCodeRepositories(val *[]*string) {
	_jsii_.Set(
		j,
		"additionalCodeRepositories",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetDefaultCodeRepository(val *string) {
	_jsii_.Set(
		j,
		"defaultCodeRepository",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetDirectInternetAccess(val *string) {
	_jsii_.Set(
		j,
		"directInternetAccess",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetLifecycleConfigName(val *string) {
	_jsii_.Set(
		j,
		"lifecycleConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetNotebookInstanceName(val *string) {
	_jsii_.Set(
		j,
		"notebookInstanceName",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetPlatformIdentifier(val *string) {
	_jsii_.Set(
		j,
		"platformIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetRootAccess(val *string) {
	_jsii_.Set(
		j,
		"rootAccess",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstance) SetVolumeSizeInGb(val *float64) {
	_jsii_.Set(
		j,
		"volumeSizeInGb",
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
func CfnNotebookInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnNotebookInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnNotebookInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotebookInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnNotebookInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNotebookInstance) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNotebookInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnNotebookInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::SageMaker::NotebookInstanceLifecycleConfig`.
//
// The `AWS::SageMaker::NotebookInstanceLifecycleConfig` resource creates shell scripts that run when you create and/or start a notebook instance. For information about notebook instance lifecycle configurations, see [Customize a Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceLifecycleConfig := awscdk.Aws_sagemaker.NewCfnNotebookInstanceLifecycleConfig(this, jsii.String("MyCfnNotebookInstanceLifecycleConfig"), &cfnNotebookInstanceLifecycleConfigProps{
//   	notebookInstanceLifecycleConfigName: jsii.String("notebookInstanceLifecycleConfigName"),
//   	onCreate: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   	onStart: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   })
//
type CfnNotebookInstanceLifecycleConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the lifecycle configuration, such as `MyLifecycleConfig` .
	AttrNotebookInstanceLifecycleConfigName() *string
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
	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName() *string
	SetNotebookInstanceLifecycleConfigName(val *string)
	// A shell script that runs only once, when you create a notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnCreate() interface{}
	SetOnCreate(val interface{})
	// A shell script that runs every time you start a notebook instance, including when you create the notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnStart() interface{}
	SetOnStart(val interface{})
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

// The jsii proxy struct for CfnNotebookInstanceLifecycleConfig
type jsiiProxy_CfnNotebookInstanceLifecycleConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AttrNotebookInstanceLifecycleConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNotebookInstanceLifecycleConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) NotebookInstanceLifecycleConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookInstanceLifecycleConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OnStart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::NotebookInstanceLifecycleConfig`.
func NewCfnNotebookInstanceLifecycleConfig(scope awscdk.Construct, id *string, props *CfnNotebookInstanceLifecycleConfigProps) CfnNotebookInstanceLifecycleConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnNotebookInstanceLifecycleConfig{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::NotebookInstanceLifecycleConfig`.
func NewCfnNotebookInstanceLifecycleConfig_Override(c CfnNotebookInstanceLifecycleConfig, scope awscdk.Construct, id *string, props *CfnNotebookInstanceLifecycleConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) SetNotebookInstanceLifecycleConfigName(val *string) {
	_jsii_.Set(
		j,
		"notebookInstanceLifecycleConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) SetOnCreate(val interface{}) {
	_jsii_.Set(
		j,
		"onCreate",
		val,
	)
}

func (j *jsiiProxy_CfnNotebookInstanceLifecycleConfig) SetOnStart(val interface{}) {
	_jsii_.Set(
		j,
		"onStart",
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
func CfnNotebookInstanceLifecycleConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnNotebookInstanceLifecycleConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnNotebookInstanceLifecycleConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotebookInstanceLifecycleConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnNotebookInstanceLifecycleConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotebookInstanceLifecycleConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the notebook instance lifecycle configuration script.
//
// Each lifecycle configuration script has a limit of 16384 characters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notebookInstanceLifecycleHookProperty := &notebookInstanceLifecycleHookProperty{
//   	content: jsii.String("content"),
//   }
//
type CfnNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHookProperty struct {
	// A base64-encoded string that contains a shell script for a notebook instance lifecycle configuration.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

// Properties for defining a `CfnNotebookInstanceLifecycleConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceLifecycleConfigProps := &cfnNotebookInstanceLifecycleConfigProps{
//   	notebookInstanceLifecycleConfigName: jsii.String("notebookInstanceLifecycleConfigName"),
//   	onCreate: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   	onStart: []interface{}{
//   		&notebookInstanceLifecycleHookProperty{
//   			content: jsii.String("content"),
//   		},
//   	},
//   }
//
type CfnNotebookInstanceLifecycleConfigProps struct {
	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName *string `field:"optional" json:"notebookInstanceLifecycleConfigName" yaml:"notebookInstanceLifecycleConfigName"`
	// A shell script that runs only once, when you create a notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnCreate interface{} `field:"optional" json:"onCreate" yaml:"onCreate"`
	// A shell script that runs every time you start a notebook instance, including when you create the notebook instance.
	//
	// The shell script must be a base64-encoded string.
	OnStart interface{} `field:"optional" json:"onStart" yaml:"onStart"`
}

// Properties for defining a `CfnNotebookInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotebookInstanceProps := &cfnNotebookInstanceProps{
//   	instanceType: jsii.String("instanceType"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	acceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	additionalCodeRepositories: []*string{
//   		jsii.String("additionalCodeRepositories"),
//   	},
//   	defaultCodeRepository: jsii.String("defaultCodeRepository"),
//   	directInternetAccess: jsii.String("directInternetAccess"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	lifecycleConfigName: jsii.String("lifecycleConfigName"),
//   	notebookInstanceName: jsii.String("notebookInstanceName"),
//   	platformIdentifier: jsii.String("platformIdentifier"),
//   	rootAccess: jsii.String("rootAccess"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeSizeInGb: jsii.Number(123),
//   }
//
type CfnNotebookInstanceProps struct {
	// The type of ML compute instance to launch for the notebook instance.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// When you send any requests to AWS resources from the notebook instance, SageMaker assumes this role to perform tasks on your behalf.
	//
	// You must grant this role necessary permissions so SageMaker can perform these tasks. The policy must allow the SageMaker service principal (sagemaker.amazonaws.com) permissions to assume this role. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance.
	//
	// Currently, only one instance type can be associated with a notebook instance. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	//
	// *Valid Values:* `ml.eia1.medium | ml.eia1.large | ml.eia1.xlarge | ml.eia2.medium | ml.eia2.large | ml.eia2.xlarge` .
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// An array of up to three Git repositories associated with the notebook instance.
	//
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	AdditionalCodeRepositories *[]*string `field:"optional" json:"additionalCodeRepositories" yaml:"additionalCodeRepositories"`
	// The Git repository associated with the notebook instance as its default code repository.
	//
	// This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository. For more information, see [Associating Git Repositories with SageMaker Notebook Instances](https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html) .
	DefaultCodeRepository *string `field:"optional" json:"defaultCodeRepository" yaml:"defaultCodeRepository"`
	// Sets whether SageMaker provides internet access to the notebook instance.
	//
	// If you set this to `Disabled` this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker training and endpoint services unless you configure a NAT Gateway in your VPC.
	//
	// For more information, see [Notebook Instances Are Internet-Enabled by Default](https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access) . You can set the value of this parameter to `Disabled` only if you set a value for the `SubnetId` parameter.
	DirectInternetAccess *string `field:"optional" json:"directInternetAccess" yaml:"directInternetAccess"`
	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to your notebook instance.
	//
	// The KMS key you provide must be enabled. For information, see [Enabling and Disabling Keys](https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html) in the *AWS Key Management Service Developer Guide* .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	//
	// For information about lifecycle configurations, see [Customize a Notebook Instance](https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html) in the *Amazon SageMaker Developer Guide* .
	LifecycleConfigName *string `field:"optional" json:"lifecycleConfigName" yaml:"lifecycleConfigName"`
	// The name of the new notebook instance.
	NotebookInstanceName *string `field:"optional" json:"notebookInstanceName" yaml:"notebookInstanceName"`
	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier *string `field:"optional" json:"platformIdentifier" yaml:"platformIdentifier"`
	// Whether root access is enabled or disabled for users of the notebook instance. The default value is `Enabled` .
	//
	// > Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users.
	RootAccess *string `field:"optional" json:"rootAccess" yaml:"rootAccess"`
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// The security groups must be for the same VPC as specified in the subnet.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	//
	// You can add tags later by using the `CreateTags` API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	//
	// The default value is 5 GB.
	//
	// > Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

// A CloudFormation `AWS::SageMaker::Pipeline`.
//
// The `AWS::SageMaker::Pipeline` resource creates shell scripts that run when you create and/or start a SageMaker Pipeline. For information about SageMaker Pipelines, see [SageMaker Pipelines](https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parallelismConfiguration interface{}
//   var pipelineDefinition interface{}
//
//   cfnPipeline := awscdk.Aws_sagemaker.NewCfnPipeline(this, jsii.String("MyCfnPipeline"), &cfnPipelineProps{
//   	pipelineDefinition: pipelineDefinition,
//   	pipelineName: jsii.String("pipelineName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	parallelismConfiguration: parallelismConfiguration,
//   	pipelineDescription: jsii.String("pipelineDescription"),
//   	pipelineDisplayName: jsii.String("pipelineDisplayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPipeline interface {
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
	// `AWS::SageMaker::Pipeline.ParallelismConfiguration`.
	ParallelismConfiguration() interface{}
	SetParallelismConfiguration(val interface{})
	// The definition of the pipeline.
	//
	// This can be either a JSON string or an Amazon S3 location.
	PipelineDefinition() interface{}
	SetPipelineDefinition(val interface{})
	// The description of the pipeline.
	PipelineDescription() *string
	SetPipelineDescription(val *string)
	// The display name of the pipeline.
	PipelineDisplayName() *string
	SetPipelineDisplayName(val *string)
	// The name of the pipeline.
	PipelineName() *string
	SetPipelineName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role used to execute the pipeline.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags of the pipeline.
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

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPipeline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) ParallelismConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parallelismConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Pipeline`.
func NewCfnPipeline(scope awscdk.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope awscdk.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipeline) SetParallelismConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"parallelismConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineDescription(val *string) {
	_jsii_.Set(
		j,
		"pipelineDescription",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineDisplayName(val *string) {
	_jsii_.Set(
		j,
		"pipelineDisplayName",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineName(val *string) {
	_jsii_.Set(
		j,
		"pipelineName",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetRoleArn(val *string) {
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
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnPipeline",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipeline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipeline) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPipeline) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipeline) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPipeline) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPipeline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPipeline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parallelismConfiguration interface{}
//   var pipelineDefinition interface{}
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	pipelineDefinition: pipelineDefinition,
//   	pipelineName: jsii.String("pipelineName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	parallelismConfiguration: parallelismConfiguration,
//   	pipelineDescription: jsii.String("pipelineDescription"),
//   	pipelineDisplayName: jsii.String("pipelineDisplayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The definition of the pipeline.
	//
	// This can be either a JSON string or an Amazon S3 location.
	PipelineDefinition interface{} `field:"required" json:"pipelineDefinition" yaml:"pipelineDefinition"`
	// The name of the pipeline.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// The Amazon Resource Name (ARN) of the IAM role used to execute the pipeline.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `AWS::SageMaker::Pipeline.ParallelismConfiguration`.
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
	// The description of the pipeline.
	PipelineDescription *string `field:"optional" json:"pipelineDescription" yaml:"pipelineDescription"`
	// The display name of the pipeline.
	PipelineDisplayName *string `field:"optional" json:"pipelineDisplayName" yaml:"pipelineDisplayName"`
	// The tags of the pipeline.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::Project`.
//
// Creates a machine learning (ML) project that can contain one or more templates that set up an ML pipeline from training to deploying an approved model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceCatalogProvisioningDetails interface{}
//
//   cfnProject := awscdk.Aws_sagemaker.NewCfnProject(this, jsii.String("MyCfnProject"), &cfnProjectProps{
//   	projectName: jsii.String("projectName"),
//   	serviceCatalogProvisioningDetails: serviceCatalogProvisioningDetails,
//
//   	// the properties below are optional
//   	projectDescription: jsii.String("projectDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The time that the project was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the project.
	AttrProjectArn() *string
	// The ID of the project.
	//
	// This ID is prepended to all entities associated with this project.
	AttrProjectId() *string
	// The status of the project.
	AttrProjectStatus() *string
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
	// The description of the project.
	ProjectDescription() *string
	SetProjectDescription(val *string)
	// The name of the project.
	ProjectName() *string
	SetProjectName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The product ID and provisioning artifact ID to provision a service catalog.
	//
	// For information, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
	ServiceCatalogProvisioningDetails() interface{}
	SetServiceCatalogProvisioningDetails(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrProjectStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ProjectDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ServiceCatalogProvisioningDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceCatalogProvisioningDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Project`.
func NewCfnProject(scope awscdk.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Project`.
func NewCfnProject_Override(c CfnProject, scope awscdk.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetProjectDescription(val *string) {
	_jsii_.Set(
		j,
		"projectDescription",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetServiceCatalogProvisioningDetails(val interface{}) {
	_jsii_.Set(
		j,
		"serviceCatalogProvisioningDetails",
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
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnProject",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProject) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnProject) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnProject) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceCatalogProvisioningDetails interface{}
//
//   cfnProjectProps := &cfnProjectProps{
//   	projectName: jsii.String("projectName"),
//   	serviceCatalogProvisioningDetails: serviceCatalogProvisioningDetails,
//
//   	// the properties below are optional
//   	projectDescription: jsii.String("projectDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The name of the project.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The product ID and provisioning artifact ID to provision a service catalog.
	//
	// For information, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
	ServiceCatalogProvisioningDetails interface{} `field:"required" json:"serviceCatalogProvisioningDetails" yaml:"serviceCatalogProvisioningDetails"`
	// The description of the project.
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SageMaker::UserProfile`.
//
// Creates a user profile. A user profile represents a single user within a domain, and is the main way to reference a "person" for the purposes of sharing, reporting, and other user-oriented features. This entity is created when a user onboards to Amazon SageMaker Studio. If an administrator invites a person by email or imports them from SSO, a user profile is automatically created. A user profile is the primary holder of settings for an individual user and has a reference to the user's private Amazon Elastic File System (EFS) home directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProfile := awscdk.Aws_sagemaker.NewCfnUserProfile(this, jsii.String("MyCfnUserProfile"), &cfnUserProfileProps{
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	singleSignOnUserIdentifier: jsii.String("singleSignOnUserIdentifier"),
//   	singleSignOnUserValue: jsii.String("singleSignOnUserValue"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   })
//
type CfnUserProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the user profile, such as `arn:aws:sagemaker:us-west-2:account-id:user-profile/my-user-profile` .
	AttrUserProfileArn() *string
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
	// The domain ID.
	DomainId() *string
	SetDomainId(val *string)
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
	// A specifier for the type of value specified in SingleSignOnUserValue.
	//
	// Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserIdentifier() *string
	SetSingleSignOnUserIdentifier(val *string)
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	//
	// If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserValue() *string
	SetSingleSignOnUserValue(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// Tags that you specify for the User Profile are also added to all apps that the User Profile launches.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The user profile name.
	UserProfileName() *string
	SetUserProfileName(val *string)
	// A collection of settings that apply to users of Amazon SageMaker Studio.
	UserSettings() interface{}
	SetUserSettings(val interface{})
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

// The jsii proxy struct for CfnUserProfile
type jsiiProxy_CfnUserProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserProfile) AttrUserProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) SingleSignOnUserIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleSignOnUserIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) SingleSignOnUserValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleSignOnUserValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) UserProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfile) UserSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSettings",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::UserProfile`.
func NewCfnUserProfile(scope awscdk.Construct, id *string, props *CfnUserProfileProps) CfnUserProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnUserProfile{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnUserProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::UserProfile`.
func NewCfnUserProfile_Override(c CfnUserProfile, scope awscdk.Construct, id *string, props *CfnUserProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnUserProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetDomainId(val *string) {
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetSingleSignOnUserIdentifier(val *string) {
	_jsii_.Set(
		j,
		"singleSignOnUserIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetSingleSignOnUserValue(val *string) {
	_jsii_.Set(
		j,
		"singleSignOnUserValue",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetUserProfileName(val *string) {
	_jsii_.Set(
		j,
		"userProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnUserProfile) SetUserSettings(val interface{}) {
	_jsii_.Set(
		j,
		"userSettings",
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
func CfnUserProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnUserProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnUserProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnUserProfile",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnUserProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnUserProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnUserProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserProfile) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserProfile) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserProfile) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnUserProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnUserProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A custom SageMaker image.
//
// For more information, see [Bring your own SageMaker image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customImageProperty := &customImageProperty{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//   	imageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	imageVersionNumber: jsii.Number(123),
//   }
//
type CfnUserProfile_CustomImageProperty struct {
	// The name of the AppImageConfig.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage.
	//
	// Must be unique to your account.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

// The JupyterServer app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jupyterServerAppSettingsProperty := &jupyterServerAppSettingsProperty{
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnUserProfile_JupyterServerAppSettingsProperty struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
	//
	// If you use the `LifecycleConfigArns` parameter, then this parameter is also required.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

// The KernelGateway app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayAppSettingsProperty := &kernelGatewayAppSettingsProperty{
//   	customImages: []interface{}{
//   		&customImageProperty{
//   			appImageConfigName: jsii.String("appImageConfigName"),
//   			imageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			imageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnUserProfile_KernelGatewayAppSettingsProperty struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app.
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.
	//
	// > The Amazon SageMaker Studio UI does not use the default instance type value set here. The default instance type set here is used when Apps are created using the AWS Command Line Interface or AWS CloudFormation and the instance type parameter value is not passed.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

// A collection of settings that configure user interaction with the `RStudioServerPro` app.
//
// `RStudioServerProAppSettings` cannot be updated. The `RStudioServerPro` app must be deleted and a new one created to make any changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProAppSettingsProperty := &rStudioServerProAppSettingsProperty{
//   	accessStatus: jsii.String("accessStatus"),
//   	userGroup: jsii.String("userGroup"),
//   }
//
type CfnUserProfile_RStudioServerProAppSettingsProperty struct {
	// Indicates whether the current user has access to the `RStudioServerPro` app.
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// The level of permissions that the user has within the `RStudioServerPro` app.
	//
	// This value defaults to `User`. The `Admin` value allows the user access to the RStudio Administrative Dashboard.
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}

// Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecProperty := &resourceSpecProperty{
//   	instanceType: jsii.String("instanceType"),
//   	sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   	sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   }
//
type CfnUserProfile_ResourceSpecProperty struct {
	// The instance type that the image version runs on.
	//
	// > JupyterServer Apps only support the `system` value. KernelGateway Apps do not support the `system` value, but support all other values for available instance types.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the SageMaker image that the image version belongs to.
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}

// Specifies options when sharing an Amazon SageMaker Studio notebook.
//
// These settings are specified as part of `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called, and as part of `UserSettings` when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharingSettingsProperty := &sharingSettingsProperty{
//   	notebookOutputOption: jsii.String("notebookOutputOption"),
//   	s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   	s3OutputPath: jsii.String("s3OutputPath"),
//   }
//
type CfnUserProfile_SharingSettingsProperty struct {
	// Whether to include the notebook cell output when sharing the notebook.
	//
	// The default is `Disabled` .
	NotebookOutputOption *string `field:"optional" json:"notebookOutputOption" yaml:"notebookOutputOption"`
	// When `NotebookOutputOption` is `Allowed` , the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
	// When `NotebookOutputOption` is `Allowed` , the Amazon S3 bucket used to store the shared notebook snapshots.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

// A collection of settings that apply to users of Amazon SageMaker Studio.
//
// These settings are specified when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called, and as `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called.
//
// `SecurityGroups` is aggregated when specified in both calls. For all other settings in `UserSettings` , the values specified in `CreateUserProfile` take precedence over those specified in `CreateDomain` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingsProperty := &userSettingsProperty{
//   	executionRole: jsii.String("executionRole"),
//   	jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   		customImages: []interface{}{
//   			&customImageProperty{
//   				appImageConfigName: jsii.String("appImageConfigName"),
//   				imageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				imageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   		accessStatus: jsii.String("accessStatus"),
//   		userGroup: jsii.String("userGroup"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	sharingSettings: &sharingSettingsProperty{
//   		notebookOutputOption: jsii.String("notebookOutputOption"),
//   		s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   		s3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   }
//
type CfnUserProfile_UserSettingsProperty struct {
	// The execution role for the user.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The Jupyter server's app settings.
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A collection of settings that configure user interaction with the `RStudioServerPro` app.
	RStudioServerProAppSettings interface{} `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	//
	// Optional when the `CreateDomain.AppNetworkAccessType` parameter is set to `PublicInternetOnly` .
	//
	// Required when the `CreateDomain.AppNetworkAccessType` parameter is set to `VpcOnly` .
	//
	// Amazon SageMaker adds a security group to allow NFS traffic from SageMaker Studio. Therefore, the number of security groups that you can specify is one less than the maximum number shown.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies options for sharing SageMaker Studio notebooks.
	SharingSettings interface{} `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
}

// Properties for defining a `CfnUserProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProfileProps := &cfnUserProfileProps{
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	singleSignOnUserIdentifier: jsii.String("singleSignOnUserIdentifier"),
//   	singleSignOnUserValue: jsii.String("singleSignOnUserValue"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   }
//
type CfnUserProfileProps struct {
	// The domain ID.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// A specifier for the type of value specified in SingleSignOnUserValue.
	//
	// Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserIdentifier *string `field:"optional" json:"singleSignOnUserIdentifier" yaml:"singleSignOnUserIdentifier"`
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	//
	// If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserValue *string `field:"optional" json:"singleSignOnUserValue" yaml:"singleSignOnUserValue"`
	// An array of key-value pairs to apply to this resource.
	//
	// Tags that you specify for the User Profile are also added to all apps that the User Profile launches.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A collection of settings that apply to users of Amazon SageMaker Studio.
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

// A CloudFormation `AWS::SageMaker::Workteam`.
//
// Creates a new work team for labeling your data. A work team is defined by one or more Amazon Cognito user pools. You must first create the user pools before you can create a work team.
//
// You cannot create more than 25 work teams in an account and region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkteam := awscdk.Aws_sagemaker.NewCfnWorkteam(this, jsii.String("MyCfnWorkteam"), &cfnWorkteamProps{
//   	description: jsii.String("description"),
//   	memberDefinitions: []interface{}{
//   		&memberDefinitionProperty{
//   			cognitoMemberDefinition: &cognitoMemberDefinitionProperty{
//   				cognitoClientId: jsii.String("cognitoClientId"),
//   				cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   				cognitoUserPool: jsii.String("cognitoUserPool"),
//   			},
//   		},
//   	},
//   	notificationConfiguration: &notificationConfigurationProperty{
//   		notificationTopicArn: jsii.String("notificationTopicArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workteamName: jsii.String("workteamName"),
//   })
//
type CfnWorkteam interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the work team.
	AttrWorkteamName() *string
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
	// A description of the work team.
	Description() *string
	SetDescription(val *string)
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
	// A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.
	//
	// Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `CognitoMemberDefinition` . For workforces created using your own OIDC identity provider (IdP) use `OidcMemberDefinition` .
	MemberDefinitions() interface{}
	SetMemberDefinitions(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Configures SNS notifications of available or expiring work items for work teams.
	NotificationConfiguration() interface{}
	SetNotificationConfiguration(val interface{})
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
	// An array of key-value pairs.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The name of the work team.
	WorkteamName() *string
	SetWorkteamName(val *string)
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

// The jsii proxy struct for CfnWorkteam
type jsiiProxy_CfnWorkteam struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkteam) AttrWorkteamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkteamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) MemberDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) NotificationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkteam) WorkteamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamName",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::Workteam`.
func NewCfnWorkteam(scope awscdk.Construct, id *string, props *CfnWorkteamProps) CfnWorkteam {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkteam{}

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnWorkteam",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::Workteam`.
func NewCfnWorkteam_Override(c CfnWorkteam, scope awscdk.Construct, id *string, props *CfnWorkteamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sagemaker.CfnWorkteam",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkteam) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWorkteam) SetMemberDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"memberDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnWorkteam) SetNotificationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnWorkteam) SetWorkteamName(val *string) {
	_jsii_.Set(
		j,
		"workteamName",
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
func CfnWorkteam_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnWorkteam",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWorkteam_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnWorkteam",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWorkteam_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sagemaker.CfnWorkteam",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkteam_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_sagemaker.CfnWorkteam",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkteam) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkteam) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkteam) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkteam) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkteam) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkteam) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkteam) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkteam) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkteam) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkteam) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkteam) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkteam) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkteam) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkteam) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkteam) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Identifies a Amazon Cognito user group.
//
// A user group can be used in on or more work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoMemberDefinitionProperty := &cognitoMemberDefinitionProperty{
//   	cognitoClientId: jsii.String("cognitoClientId"),
//   	cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   	cognitoUserPool: jsii.String("cognitoUserPool"),
//   }
//
type CfnWorkteam_CognitoMemberDefinitionProperty struct {
	// An identifier for an application client.
	//
	// You must create the app client ID using Amazon Cognito.
	CognitoClientId *string `field:"required" json:"cognitoClientId" yaml:"cognitoClientId"`
	// An identifier for a user group.
	CognitoUserGroup *string `field:"required" json:"cognitoUserGroup" yaml:"cognitoUserGroup"`
	// An identifier for a user pool.
	//
	// The user pool must be in the same region as the service that you are calling.
	CognitoUserPool *string `field:"required" json:"cognitoUserPool" yaml:"cognitoUserPool"`
}

// Defines an Amazon Cognito or your own OIDC IdP user group that is part of a work team.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberDefinitionProperty := &memberDefinitionProperty{
//   	cognitoMemberDefinition: &cognitoMemberDefinitionProperty{
//   		cognitoClientId: jsii.String("cognitoClientId"),
//   		cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   		cognitoUserPool: jsii.String("cognitoUserPool"),
//   	},
//   }
//
type CfnWorkteam_MemberDefinitionProperty struct {
	// The Amazon Cognito user group that is part of the work team.
	CognitoMemberDefinition interface{} `field:"required" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
}

// Configures Amazon SNS notifications of available or expiring work items for work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	notificationTopicArn: jsii.String("notificationTopicArn"),
//   }
//
type CfnWorkteam_NotificationConfigurationProperty struct {
	// The ARN for the Amazon SNS topic to which notifications should be published.
	NotificationTopicArn *string `field:"required" json:"notificationTopicArn" yaml:"notificationTopicArn"`
}

// Properties for defining a `CfnWorkteam`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkteamProps := &cfnWorkteamProps{
//   	description: jsii.String("description"),
//   	memberDefinitions: []interface{}{
//   		&memberDefinitionProperty{
//   			cognitoMemberDefinition: &cognitoMemberDefinitionProperty{
//   				cognitoClientId: jsii.String("cognitoClientId"),
//   				cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   				cognitoUserPool: jsii.String("cognitoUserPool"),
//   			},
//   		},
//   	},
//   	notificationConfiguration: &notificationConfigurationProperty{
//   		notificationTopicArn: jsii.String("notificationTopicArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workteamName: jsii.String("workteamName"),
//   }
//
type CfnWorkteamProps struct {
	// A description of the work team.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.
	//
	// Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `CognitoMemberDefinition` . For workforces created using your own OIDC identity provider (IdP) use `OidcMemberDefinition` .
	MemberDefinitions interface{} `field:"optional" json:"memberDefinitions" yaml:"memberDefinitions"`
	// Configures SNS notifications of available or expiring work items for work teams.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// An array of key-value pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the work team.
	WorkteamName *string `field:"optional" json:"workteamName" yaml:"workteamName"`
}

