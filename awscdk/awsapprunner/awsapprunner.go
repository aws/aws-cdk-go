package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppRunner::ObservabilityConfiguration`.
//
// Specify an AWS App Runner observability configuration by using the `AWS::AppRunner::ObservabilityConfiguration` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::ObservabilityConfiguration` resource is an AWS App Runner resource type that specifies an App Runner observability configuration.
//
// App Runner requires this resource when you specify App Runner services and you want to enable non-default observability features. You can share an observability configuration across multiple services.
//
// Create multiple revisions of a configuration by specifying this resource multiple times using the same `ObservabilityConfigurationName` . App Runner creates multiple resources with incremental `ObservabilityConfigurationRevision` values. When you specify a service and configure an observability configuration resource, the service uses the latest active revision of the observability configuration by default. You can optionally configure the service to use a specific revision.
//
// The observability configuration resource is designed to configure multiple features (currently one feature, tracing). This resource takes optional parameters that describe the configuration of these features (currently one parameter, `TraceConfiguration` ). If you don't specify a feature parameter, App Runner doesn't enable the feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnObservabilityConfiguration := awscdk.Aws_apprunner.NewCfnObservabilityConfiguration(this, jsii.String("MyCfnObservabilityConfiguration"), &cfnObservabilityConfigurationProps{
//   	observabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceConfiguration: &traceConfigurationProperty{
//   		vendor: jsii.String("vendor"),
//   	},
//   })
//
type CfnObservabilityConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// It's set to `true` for the configuration with the highest `Revision` among all configurations that share the same `ObservabilityConfigurationName` .
	//
	// It's set to `false` otherwise.
	AttrLatest() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of this observability configuration.
	AttrObservabilityConfigurationArn() *string
	// The revision of this observability configuration.
	//
	// It's unique among all the active configurations ( `"Status": "ACTIVE"` ) that share the same `ObservabilityConfigurationName` .
	AttrObservabilityConfigurationRevision() *float64
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region , App Runner creates revision number `1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// > The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a revision of it.
	// >
	// > When you want to use your own observability configuration for your App Runner service, *create a configuration with a different name* , and then provide it when you create or update your service.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your observability configuration.
	ObservabilityConfigurationName() *string
	SetObservabilityConfigurationName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	Tags() awscdk.TagManager
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration() interface{}
	SetTraceConfiguration(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnObservabilityConfiguration
type jsiiProxy_CfnObservabilityConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrLatest() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrObservabilityConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrObservabilityConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) AttrObservabilityConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrObservabilityConfigurationRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) ObservabilityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) TraceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfiguration(scope constructs.Construct, id *string, props *CfnObservabilityConfigurationProps) CfnObservabilityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnObservabilityConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfiguration_Override(c CfnObservabilityConfiguration, scope constructs.Construct, id *string, props *CfnObservabilityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnObservabilityConfiguration) SetObservabilityConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"observabilityConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnObservabilityConfiguration) SetTraceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"traceConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnObservabilityConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnObservabilityConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
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
func CfnObservabilityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObservabilityConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apprunner.CfnObservabilityConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnObservabilityConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnObservabilityConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the configuration of the tracing feature within an AWS App Runner observability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   traceConfigurationProperty := &traceConfigurationProperty{
//   	vendor: jsii.String("vendor"),
//   }
//
type CfnObservabilityConfiguration_TraceConfigurationProperty struct {
	// The implementation provider chosen for tracing App Runner services.
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
}

// Properties for defining a `CfnObservabilityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnObservabilityConfigurationProps := &cfnObservabilityConfigurationProps{
//   	observabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceConfiguration: &traceConfigurationProperty{
//   		vendor: jsii.String("vendor"),
//   	},
//   }
//
type CfnObservabilityConfigurationProps struct {
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region , App Runner creates revision number `1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// > The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a revision of it.
	// >
	// > When you want to use your own observability configuration for your App Runner service, *create a configuration with a different name* , and then provide it when you create or update your service.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your observability configuration.
	ObservabilityConfigurationName *string `field:"optional" json:"observabilityConfigurationName" yaml:"observabilityConfigurationName"`
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration interface{} `field:"optional" json:"traceConfiguration" yaml:"traceConfiguration"`
}

// A CloudFormation `AWS::AppRunner::Service`.
//
// Specify an AWS App Runner service by using the `AWS::AppRunner::Service` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::Service` resource is an AWS App Runner resource type that specifies an App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnService := awscdk.Aws_apprunner.NewCfnService(this, jsii.String("MyCfnService"), &cfnServiceProps{
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		authenticationConfiguration: &authenticationConfigurationProperty{
//   			accessRoleArn: jsii.String("accessRoleArn"),
//   			connectionArn: jsii.String("connectionArn"),
//   		},
//   		autoDeploymentsEnabled: jsii.Boolean(false),
//   		codeRepository: &codeRepositoryProperty{
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   			sourceCodeVersion: &sourceCodeVersionProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			codeConfiguration: &codeConfigurationProperty{
//   				configurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				codeConfigurationValues: &codeConfigurationValuesProperty{
//   					runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					buildCommand: jsii.String("buildCommand"),
//   					port: jsii.String("port"),
//   					runtimeEnvironmentVariables: []interface{}{
//   						&keyValuePairProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					startCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		imageRepository: &imageRepositoryProperty{
//   			imageIdentifier: jsii.String("imageIdentifier"),
//   			imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			imageConfiguration: &imageConfigurationProperty{
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	autoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	healthCheckConfiguration: &healthCheckConfigurationProperty{
//   		healthyThreshold: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		path: jsii.String("path"),
//   		protocol: jsii.String("protocol"),
//   		timeout: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		instanceRoleArn: jsii.String("instanceRoleArn"),
//   		memory: jsii.String("memory"),
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		egressConfiguration: &egressConfigurationProperty{
//   			egressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   	},
//   	observabilityConfiguration: &serviceObservabilityConfigurationProperty{
//   		observabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnService interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of this service.
	AttrServiceArn() *string
	// An ID that App Runner generated for this service.
	//
	// It's unique within the AWS Region .
	AttrServiceId() *string
	// A subdomain URL that App Runner generated for this service.
	//
	// You can use this URL to access your service web application.
	AttrServiceUrl() *string
	// The current state of the App Runner service. These particular values mean the following.
	//
	// - `CREATE_FAILED` – The service failed to create. To troubleshoot this failure, read the failure events and logs, change any parameters that need to be fixed, and retry the call to create the service.
	//
	// The failed service isn't usable, and still counts towards your service quota. When you're done analyzing the failure, delete the service.
	// - `DELETE_FAILED` – The service failed to delete and can't be successfully recovered. Retry the service deletion call to ensure that all related resources are removed.
	AttrStatus() *string
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	AutoScalingConfigurationArn() *string
	SetAutoScalingConfigurationArn(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration() interface{}
	SetHealthCheckConfiguration(val interface{})
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration() interface{}
	SetInstanceConfiguration(val interface{})
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
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// The observability configuration of your service.
	ObservabilityConfiguration() interface{}
	SetObservabilityConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your service.
	ServiceName() *string
	SetServiceName(val *string)
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration() interface{}
	SetSourceConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnService
type jsiiProxy_CfnService struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnService) AttrServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) AutoScalingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) EncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) HealthCheckConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) InstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ObservabilityConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"observabilityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) SourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnService) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::Service`.
func NewCfnService(scope constructs.Construct, id *string, props *CfnServiceProps) CfnService {
	_init_.Initialize()

	j := jsiiProxy_CfnService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::Service`.
func NewCfnService_Override(c CfnService, scope constructs.Construct, id *string, props *CfnServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnService) SetAutoScalingConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"autoScalingConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetHealthCheckConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetInstanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"instanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetObservabilityConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"observabilityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_CfnService) SetSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sourceConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnService_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnService_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
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
func CfnService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnService_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apprunner.CfnService",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnService) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnService) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnService) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnService) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnService) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnService) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnService) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes resources needed to authenticate access to some source repositories.
//
// The specific resource depends on the repository provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationProperty := &authenticationConfigurationProperty{
//   	accessRoleArn: jsii.String("accessRoleArn"),
//   	connectionArn: jsii.String("connectionArn"),
//   }
//
type CfnService_AuthenticationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository.
	//
	// It's required for GitHub code repositories.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationProperty := &codeConfigurationProperty{
//   	configurationSource: jsii.String("configurationSource"),
//
//   	// the properties below are optional
//   	codeConfigurationValues: &codeConfigurationValuesProperty{
//   		runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		port: jsii.String("port"),
//   		runtimeEnvironmentVariables: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
type CfnService_CodeConfigurationProperty struct {
	// The source of the App Runner configuration. Values are interpreted as follows:.
	//
	// - `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and ignores `CodeConfigurationValues` .
	// - `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a `apprunner.yaml` file in the source code repository (or ignoring the file if it exists).
	CodeConfigurationValues interface{} `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationValuesProperty := &codeConfigurationValuesProperty{
//   	runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	buildCommand: jsii.String("buildCommand"),
//   	port: jsii.String("port"),
//   	runtimeEnvironmentVariables: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	startCommand: jsii.String("startCommand"),
//   }
//
type CfnService_CodeConfigurationValuesProperty struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents a programming language runtime.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// The command App Runner runs to start your application.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

// Describes a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryProperty := &codeRepositoryProperty{
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersionProperty{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	codeConfiguration: &codeConfigurationProperty{
//   		configurationSource: jsii.String("configurationSource"),
//
//   		// the properties below are optional
//   		codeConfigurationValues: &codeConfigurationValuesProperty{
//   			runtime: jsii.String("runtime"),
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			port: jsii.String("port"),
//   			runtimeEnvironmentVariables: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
type CfnService_CodeRepositoryProperty struct {
	// The location of the repository that contains the source code.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	SourceCodeVersion interface{} `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Configuration for building and running the service from a source code repository.
	//
	// > `CodeConfiguration` is required only for `CreateService` request.
	CodeConfiguration interface{} `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
}

// Describes configuration settings related to outbound network traffic of an AWS App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressConfigurationProperty := &egressConfigurationProperty{
//   	egressType: jsii.String("egressType"),
//
//   	// the properties below are optional
//   	vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   }
//
type CfnService_EgressConfigurationProperty struct {
	// The type of egress configuration.
	//
	// Set to `DEFAULT` for access to resources hosted on public networks.
	//
	// Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn` .
	EgressType *string `field:"required" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service.
	//
	// Only valid when `EgressType = VPC` .
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

// Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	kmsKey: jsii.String("kmsKey"),
//   }
//
type CfnService_EncryptionConfigurationProperty struct {
	// The ARN of the KMS key that's used for encryption.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

// Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigurationProperty := &healthCheckConfigurationProperty{
//   	healthyThreshold: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	path: jsii.String("path"),
//   	protocol: jsii.String("protocol"),
//   	timeout: jsii.Number(123),
//   	unhealthyThreshold: jsii.Number(123),
//   }
//
type CfnService_HealthCheckConfigurationProperty struct {
	// The number of consecutive checks that must succeed before App Runner decides that the service is healthy.
	//
	// Default: `1`.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time interval, in seconds, between health checks.
	//
	// Default: `5`.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The URL that health check requests are sent to.
	//
	// `Path` is only applicable when you set `Protocol` to `HTTP` .
	//
	// Default: `"/"`.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The IP protocol that App Runner uses to perform health checks for your service.
	//
	// If you set `Protocol` to `HTTP` , App Runner sends health check requests to the HTTP path specified by `Path` .
	//
	// Default: `TCP`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The time, in seconds, to wait for a health check response before deciding it failed.
	//
	// Default: `2`.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.
	//
	// Default: `5`.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigurationProperty := &imageConfigurationProperty{
//   	port: jsii.String("port"),
//   	runtimeEnvironmentVariables: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	startCommand: jsii.String("startCommand"),
//   }
//
type CfnService_ImageConfigurationProperty struct {
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

// Describes a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageRepositoryProperty := &imageRepositoryProperty{
//   	imageIdentifier: jsii.String("imageIdentifier"),
//   	imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   	// the properties below are optional
//   	imageConfiguration: &imageConfigurationProperty{
//   		port: jsii.String("port"),
//   		runtimeEnvironmentVariables: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
type CfnService_ImageRepositoryProperty struct {
	// The identifier of an image.
	//
	// For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html) in the *Amazon ECR User Guide* .
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether the repository is private or public.
	ImageRepositoryType *string `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

// Describes the runtime configuration of an AWS App Runner service instance (scaling unit).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConfigurationProperty := &instanceConfigurationProperty{
//   	cpu: jsii.String("cpu"),
//   	instanceRoleArn: jsii.String("instanceRoleArn"),
//   	memory: jsii.String("memory"),
//   }
//
type CfnService_InstanceConfigurationProperty struct {
	// The number of CPU units reserved for each instance of your App Runner service.
	//
	// Default: `1 vCPU`.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The amount of memory, in MB or GB, reserved for each instance of your App Runner service.
	//
	// Default: `2 GB`.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

// Describes a key-value pair, which is a string-to-string mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValuePairProperty := &keyValuePairProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnService_KeyValuePairProperty struct {
	// The key name string to map to a value.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value string to which the key name is mapped.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Describes configuration settings related to network traffic of an AWS App Runner service.
//
// Consists of embedded objects for each configurable network feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	egressConfiguration: &egressConfigurationProperty{
//   		egressType: jsii.String("egressType"),
//
//   		// the properties below are optional
//   		vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	},
//   }
//
type CfnService_NetworkConfigurationProperty struct {
	// Network configuration settings for outbound message traffic.
	EgressConfiguration interface{} `field:"required" json:"egressConfiguration" yaml:"egressConfiguration"`
}

// Describes the observability configuration of an AWS App Runner service.
//
// These are additional observability features, like tracing, that you choose to enable. They're configured in a separate resource that you associate with your service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceObservabilityConfigurationProperty := &serviceObservabilityConfigurationProperty{
//   	observabilityEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   }
//
type CfnService_ServiceObservabilityConfigurationProperty struct {
	// When `true` , an observability configuration resource is associated with the service, and an `ObservabilityConfigurationArn` is specified.
	ObservabilityEnabled interface{} `field:"required" json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// The Amazon Resource Name (ARN) of the observability configuration that is associated with the service.
	//
	// Specified only when `ObservabilityEnabled` is `true` .
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing`
	ObservabilityConfigurationArn *string `field:"optional" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceCodeVersionProperty := &sourceCodeVersionProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnService_SourceCodeVersionProperty struct {
	// The type of version identifier.
	//
	// For a git-based repository, branches represent versions.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A source code version.
	//
	// For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Describes the source deployed to an AWS App Runner service.
//
// It can be a code or an image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &sourceConfigurationProperty{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		accessRoleArn: jsii.String("accessRoleArn"),
//   		connectionArn: jsii.String("connectionArn"),
//   	},
//   	autoDeploymentsEnabled: jsii.Boolean(false),
//   	codeRepository: &codeRepositoryProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersionProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//
//   		// the properties below are optional
//   		codeConfiguration: &codeConfigurationProperty{
//   			configurationSource: jsii.String("configurationSource"),
//
//   			// the properties below are optional
//   			codeConfigurationValues: &codeConfigurationValuesProperty{
//   				runtime: jsii.String("runtime"),
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//   	imageRepository: &imageRepositoryProperty{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfigurationProperty{
//   			port: jsii.String("port"),
//   			runtimeEnvironmentVariables: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
type CfnService_SourceConfigurationProperty struct {
	// Describes the resources that are needed to authenticate access to some source repositories.
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// If `true` , continuous integration from the source repository is enabled for the App Runner service.
	//
	// Each repository change (including any source code commit or new image version) starts a deployment.
	//
	// Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The description of a source code repository.
	//
	// You must provide either this member or `ImageRepository` (but not both).
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// The description of a source image repository.
	//
	// You must provide either this member or `CodeRepository` (but not both).
	ImageRepository interface{} `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		authenticationConfiguration: &authenticationConfigurationProperty{
//   			accessRoleArn: jsii.String("accessRoleArn"),
//   			connectionArn: jsii.String("connectionArn"),
//   		},
//   		autoDeploymentsEnabled: jsii.Boolean(false),
//   		codeRepository: &codeRepositoryProperty{
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   			sourceCodeVersion: &sourceCodeVersionProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			codeConfiguration: &codeConfigurationProperty{
//   				configurationSource: jsii.String("configurationSource"),
//
//   				// the properties below are optional
//   				codeConfigurationValues: &codeConfigurationValuesProperty{
//   					runtime: jsii.String("runtime"),
//
//   					// the properties below are optional
//   					buildCommand: jsii.String("buildCommand"),
//   					port: jsii.String("port"),
//   					runtimeEnvironmentVariables: []interface{}{
//   						&keyValuePairProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					startCommand: jsii.String("startCommand"),
//   				},
//   			},
//   		},
//   		imageRepository: &imageRepositoryProperty{
//   			imageIdentifier: jsii.String("imageIdentifier"),
//   			imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   			// the properties below are optional
//   			imageConfiguration: &imageConfigurationProperty{
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	autoScalingConfigurationArn: jsii.String("autoScalingConfigurationArn"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	healthCheckConfiguration: &healthCheckConfigurationProperty{
//   		healthyThreshold: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		path: jsii.String("path"),
//   		protocol: jsii.String("protocol"),
//   		timeout: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		instanceRoleArn: jsii.String("instanceRoleArn"),
//   		memory: jsii.String("memory"),
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		egressConfiguration: &egressConfigurationProperty{
//   			egressType: jsii.String("egressType"),
//
//   			// the properties below are optional
//   			vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   		},
//   	},
//   	observabilityConfiguration: &serviceObservabilityConfigurationProperty{
//   		observabilityEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		observabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   	},
//   	serviceName: jsii.String("serviceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceProps struct {
	// The source to deploy to the App Runner service.
	//
	// It can be a code or an image repository.
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.
	//
	// If not provided, App Runner associates the latest revision of a default auto scaling configuration.
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`
	AutoScalingConfigurationArn *string `field:"optional" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.
	//
	// By default, App Runner uses an AWS managed key .
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.
	HealthCheckConfiguration interface{} `field:"optional" json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Configuration settings related to network traffic of the web application that the App Runner service runs.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The observability configuration of your service.
	ObservabilityConfiguration interface{} `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// A name for the App Runner service.
	//
	// It must be unique across all the running App Runner services in your AWS account in the AWS Region .
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// An optional list of metadata items that you can associate with the App Runner service resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::AppRunner::VpcConnector`.
//
// Specify an AWS App Runner VPC connector by using the `AWS::AppRunner::VpcConnector` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::VpcConnector` resource is an AWS App Runner resource type that specifies an App Runner VPC connector.
//
// App Runner requires this resource when you want to associate your App Runner service to a custom Amazon Virtual Private Cloud ( Amazon VPC ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcConnector := awscdk.Aws_apprunner.NewCfnVpcConnector(this, jsii.String("MyCfnVpcConnector"), &cfnVpcConnectorProps{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   })
//
type CfnVpcConnector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of this VPC connector.
	AttrVpcConnectorArn() *string
	// The revision of this VPC connector.
	//
	// It's unique among all the active connectors ( `"Status": "ACTIVE"` ) that share the same `Name` .
	//
	// > At this time, App Runner supports only one revision per name.
	AttrVpcConnectorRevision() *float64
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner currently only provides support for IPv4.
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// A name for the VPC connector.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
	VpcConnectorName() *string
	SetVpcConnectorName(val *string)
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

// The jsii proxy struct for CfnVpcConnector
type jsiiProxy_CfnVpcConnector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVpcConnector) AttrVpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) AttrVpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrVpcConnectorRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppRunner::VpcConnector`.
func NewCfnVpcConnector(scope constructs.Construct, id *string, props *CfnVpcConnectorProps) CfnVpcConnector {
	_init_.Initialize()

	j := jsiiProxy_CfnVpcConnector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppRunner::VpcConnector`.
func NewCfnVpcConnector_Override(c CfnVpcConnector, scope constructs.Construct, id *string, props *CfnVpcConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CfnVpcConnector) SetVpcConnectorName(val *string) {
	_jsii_.Set(
		j,
		"vpcConnectorName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVpcConnector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnVpcConnector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
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
func CfnVpcConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVpcConnector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apprunner.CfnVpcConnector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVpcConnector) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVpcConnector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVpcConnector) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVpcConnector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVpcConnector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVpcConnector) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnVpcConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcConnectorProps := &cfnVpcConnectorProps{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   }
//
type CfnVpcConnectorProps struct {
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner currently only provides support for IPv4.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A name for the VPC connector.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
	VpcConnectorName *string `field:"optional" json:"vpcConnectorName" yaml:"vpcConnectorName"`
}

