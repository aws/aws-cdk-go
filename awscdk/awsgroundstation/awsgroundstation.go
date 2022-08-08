package awsgroundstation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::GroundStation::Config`.
//
// Creates a `Config` with the specified parameters.
//
// Config objects provide Ground Station with the details necessary in order to schedule and execute satellite contacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfig := awscdk.Aws_groundstation.NewCfnConfig(this, jsii.String("MyCfnConfig"), &cfnConfigProps{
//   	configData: &configDataProperty{
//   		antennaDownlinkConfig: &antennaDownlinkConfigProperty{
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaDownlinkDemodDecodeConfig: &antennaDownlinkDemodDecodeConfigProperty{
//   			decodeConfig: &decodeConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			demodulationConfig: &demodulationConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaUplinkConfig: &antennaUplinkConfigProperty{
//   			spectrumConfig: &uplinkSpectrumConfigProperty{
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   			targetEirp: &eirpProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			transmitDisabled: jsii.Boolean(false),
//   		},
//   		dataflowEndpointConfig: &dataflowEndpointConfigProperty{
//   			dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   			dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   		},
//   		s3RecordingConfig: &s3RecordingConfigProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			prefix: jsii.String("prefix"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		trackingConfig: &trackingConfigProperty{
//   			autotrack: jsii.String("autotrack"),
//   		},
//   		uplinkEchoConfig: &uplinkEchoConfigProperty{
//   			antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the config, such as `arn:aws:groundstation:us-east-2:1234567890:config/tracking/9940bf3b-d2ba-427e-9906-842b5e5d2296` .
	AttrArn() *string
	// The ID of the config, such as `9940bf3b-d2ba-427e-9906-842b5e5d2296` .
	AttrId() *string
	// The type of the config, such as `tracking` .
	AttrType() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Object containing the parameters of a config.
	//
	// Only one subtype may be specified per config. See the subtype definitions for a description of each config subtype.
	ConfigData() interface{}
	SetConfigData(val interface{})
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
	// The name of the config object.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags assigned to a resource.
	Tags() awscdk.TagManager
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
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

// The jsii proxy struct for CfnConfig
type jsiiProxy_CfnConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfig) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) ConfigData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfig) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::GroundStation::Config`.
func NewCfnConfig(scope constructs.Construct, id *string, props *CfnConfigProps) CfnConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GroundStation::Config`.
func NewCfnConfig_Override(c CfnConfig, scope constructs.Construct, id *string, props *CfnConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfig) SetConfigData(val interface{}) {
	_jsii_.Set(
		j,
		"configData",
		val,
	)
}

func (j *jsiiProxy_CfnConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
//
// Use an antenna downlink config in a mission profile to receive the downlink data in raw DigIF format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   antennaDownlinkConfigProperty := &antennaDownlinkConfigProperty{
//   	spectrumConfig: &spectrumConfigProperty{
//   		bandwidth: &frequencyBandwidthProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		centerFrequency: &frequencyProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		polarization: jsii.String("polarization"),
//   	},
//   }
//
type CfnConfig_AntennaDownlinkConfigProperty struct {
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

// Provides information about how AWS Ground Station should configure an antenna for downlink during a contact.
//
// Use an antenna downlink demod decode config in a mission profile to receive the downlink data that has been demodulated and decoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   antennaDownlinkDemodDecodeConfigProperty := &antennaDownlinkDemodDecodeConfigProperty{
//   	decodeConfig: &decodeConfigProperty{
//   		unvalidatedJson: jsii.String("unvalidatedJson"),
//   	},
//   	demodulationConfig: &demodulationConfigProperty{
//   		unvalidatedJson: jsii.String("unvalidatedJson"),
//   	},
//   	spectrumConfig: &spectrumConfigProperty{
//   		bandwidth: &frequencyBandwidthProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		centerFrequency: &frequencyProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		polarization: jsii.String("polarization"),
//   	},
//   }
//
type CfnConfig_AntennaDownlinkDemodDecodeConfigProperty struct {
	// Defines how the RF signal will be decoded.
	DecodeConfig interface{} `field:"optional" json:"decodeConfig" yaml:"decodeConfig"`
	// Defines how the RF signal will be demodulated.
	DemodulationConfig interface{} `field:"optional" json:"demodulationConfig" yaml:"demodulationConfig"`
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

// Provides information about how AWS Ground Station should configure an antenna for uplink during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   antennaUplinkConfigProperty := &antennaUplinkConfigProperty{
//   	spectrumConfig: &uplinkSpectrumConfigProperty{
//   		centerFrequency: &frequencyProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		polarization: jsii.String("polarization"),
//   	},
//   	targetEirp: &eirpProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	transmitDisabled: jsii.Boolean(false),
//   }
//
type CfnConfig_AntennaUplinkConfigProperty struct {
	// Defines the spectrum configuration.
	SpectrumConfig interface{} `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
	// The equivalent isotropically radiated power (EIRP) to use for uplink transmissions.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	TargetEirp interface{} `field:"optional" json:"targetEirp" yaml:"targetEirp"`
	// Whether or not uplink transmit is disabled.
	TransmitDisabled interface{} `field:"optional" json:"transmitDisabled" yaml:"transmitDisabled"`
}

// Config objects provide information to Ground Station about how to configure the antenna and how data flows during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configDataProperty := &configDataProperty{
//   	antennaDownlinkConfig: &antennaDownlinkConfigProperty{
//   		spectrumConfig: &spectrumConfigProperty{
//   			bandwidth: &frequencyBandwidthProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   	},
//   	antennaDownlinkDemodDecodeConfig: &antennaDownlinkDemodDecodeConfigProperty{
//   		decodeConfig: &decodeConfigProperty{
//   			unvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		demodulationConfig: &demodulationConfigProperty{
//   			unvalidatedJson: jsii.String("unvalidatedJson"),
//   		},
//   		spectrumConfig: &spectrumConfigProperty{
//   			bandwidth: &frequencyBandwidthProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   	},
//   	antennaUplinkConfig: &antennaUplinkConfigProperty{
//   		spectrumConfig: &uplinkSpectrumConfigProperty{
//   			centerFrequency: &frequencyProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			polarization: jsii.String("polarization"),
//   		},
//   		targetEirp: &eirpProperty{
//   			units: jsii.String("units"),
//   			value: jsii.Number(123),
//   		},
//   		transmitDisabled: jsii.Boolean(false),
//   	},
//   	dataflowEndpointConfig: &dataflowEndpointConfigProperty{
//   		dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   		dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   	},
//   	s3RecordingConfig: &s3RecordingConfigProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		prefix: jsii.String("prefix"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	trackingConfig: &trackingConfigProperty{
//   		autotrack: jsii.String("autotrack"),
//   	},
//   	uplinkEchoConfig: &uplinkEchoConfigProperty{
//   		antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnConfig_ConfigDataProperty struct {
	// Provides information for an antenna downlink config object.
	//
	// Antenna downlink config objects are used to provide parameters for downlinks where no demodulation or decoding is performed by Ground Station (RF over IP downlinks).
	AntennaDownlinkConfig interface{} `field:"optional" json:"antennaDownlinkConfig" yaml:"antennaDownlinkConfig"`
	// Provides information for a downlink demod decode config object.
	//
	// Downlink demod decode config objects are used to provide parameters for downlinks where the Ground Station service will demodulate and decode the downlinked data.
	AntennaDownlinkDemodDecodeConfig interface{} `field:"optional" json:"antennaDownlinkDemodDecodeConfig" yaml:"antennaDownlinkDemodDecodeConfig"`
	// Provides information for an uplink config object.
	//
	// Uplink config objects are used to provide parameters for uplink contacts.
	AntennaUplinkConfig interface{} `field:"optional" json:"antennaUplinkConfig" yaml:"antennaUplinkConfig"`
	// Provides information for a dataflow endpoint config object.
	//
	// Dataflow endpoint config objects are used to provide parameters about which IP endpoint(s) to use during a contact. Dataflow endpoints are where Ground Station sends data during a downlink contact and where Ground Station receives data to send to the satellite during an uplink contact.
	DataflowEndpointConfig interface{} `field:"optional" json:"dataflowEndpointConfig" yaml:"dataflowEndpointConfig"`
	// Provides information for an S3 recording config object.
	//
	// S3 recording config objects are used to provide parameters for S3 recording during downlink contacts.
	S3RecordingConfig interface{} `field:"optional" json:"s3RecordingConfig" yaml:"s3RecordingConfig"`
	// Provides information for a tracking config object.
	//
	// Tracking config objects are used to provide parameters about how to track the satellite through the sky during a contact.
	TrackingConfig interface{} `field:"optional" json:"trackingConfig" yaml:"trackingConfig"`
	// Provides information for an uplink echo config object.
	//
	// Uplink echo config objects are used to provide parameters for uplink echo during uplink contacts.
	UplinkEchoConfig interface{} `field:"optional" json:"uplinkEchoConfig" yaml:"uplinkEchoConfig"`
}

// Provides information to AWS Ground Station about which IP endpoints to use during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointConfigProperty := &dataflowEndpointConfigProperty{
//   	dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   	dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   }
//
type CfnConfig_DataflowEndpointConfigProperty struct {
	// The name of the dataflow endpoint to use during contacts.
	DataflowEndpointName *string `field:"optional" json:"dataflowEndpointName" yaml:"dataflowEndpointName"`
	// The region of the dataflow endpoint to use during contacts.
	//
	// When omitted, Ground Station will use the region of the contact.
	DataflowEndpointRegion *string `field:"optional" json:"dataflowEndpointRegion" yaml:"dataflowEndpointRegion"`
}

// Defines decoding settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decodeConfigProperty := &decodeConfigProperty{
//   	unvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
type CfnConfig_DecodeConfigProperty struct {
	// The decoding settings are in JSON format and define a set of steps to perform to decode the data.
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

// Defines demodulation settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   demodulationConfigProperty := &demodulationConfigProperty{
//   	unvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
type CfnConfig_DemodulationConfigProperty struct {
	// The demodulation settings are in JSON format and define parameters for demodulation, for example which modulation scheme (e.g. PSK, QPSK, etc.) and matched filter to use.
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

// Defines an equivalent isotropically radiated power (EIRP).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eirpProperty := &eirpProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_EirpProperty struct {
	// The units of the EIRP.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the EIRP.
	//
	// Valid values are between 20.0 to 50.0 dBW.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// Defines a bandwidth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyBandwidthProperty := &frequencyBandwidthProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_FrequencyBandwidthProperty struct {
	// The units of the bandwidth.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the bandwidth. AWS Ground Station currently has the following bandwidth limitations:.
	//
	// - For `AntennaDownlinkDemodDecodeconfig` , valid values are between 125 kHz to 650 MHz.
	// - For `AntennaDownlinkconfig` , valid values are between 10 kHz to 54 MHz.
	// - For `AntennaUplinkConfig` , valid values are between 10 kHz to 54 MHz.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// Defines a frequency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyProperty := &frequencyProperty{
//   	units: jsii.String("units"),
//   	value: jsii.Number(123),
//   }
//
type CfnConfig_FrequencyProperty struct {
	// The units of the frequency.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The value of the frequency.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// Provides information about how AWS Ground Station should save downlink data to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3RecordingConfigProperty := &s3RecordingConfigProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	prefix: jsii.String("prefix"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnConfig_S3RecordingConfigProperty struct {
	// S3 Bucket where the data is written.
	//
	// The name of the S3 Bucket provided must begin with `aws-groundstation` .
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The prefix of the S3 data object.
	//
	// If you choose to use any optional keys for substitution, these values will be replaced with the corresponding information from your contact details. For example, a prefix of `{satellite_id}/{year}/{month}/{day}/` will replaced with `fake_satellite_id/2021/01/10/`
	//
	// *Optional keys for substitution* : `{satellite_id}` | `{config-name}` | `{config-id}` | `{year}` | `{month}` | `{day}`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Defines the ARN of the role assumed for putting archives to S3.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

// Defines a spectrum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spectrumConfigProperty := &spectrumConfigProperty{
//   	bandwidth: &frequencyBandwidthProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	centerFrequency: &frequencyProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	polarization: jsii.String("polarization"),
//   }
//
type CfnConfig_SpectrumConfigProperty struct {
	// The bandwidth of the spectrum. AWS Ground Station currently has the following bandwidth limitations:.
	//
	// - For `AntennaDownlinkDemodDecodeconfig` , valid values are between 125 kHz to 650 MHz.
	// - For `AntennaDownlinkconfig` , valid values are between 10 kHz to 54 MHz.
	// - For `AntennaUplinkConfig` , valid values are between 10 kHz to 54 MHz.
	Bandwidth interface{} `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// The center frequency of the spectrum.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	CenterFrequency interface{} `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// The polarization of the spectrum.
	//
	// Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"` . Capturing both `"RIGHT_HAND"` and `"LEFT_HAND"` polarization requires two separate configs.
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

// Provides information about how AWS Ground Station should track the satellite through the sky during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackingConfigProperty := &trackingConfigProperty{
//   	autotrack: jsii.String("autotrack"),
//   }
//
type CfnConfig_TrackingConfigProperty struct {
	// Specifies whether or not to use autotrack.
	//
	// `REMOVED` specifies that program track should only be used during the contact. `PREFERRED` specifies that autotracking is preferred during the contact but fallback to program track if the signal is lost. `REQUIRED` specifies that autotracking is required during the contact and not to use program track if the signal is lost.
	Autotrack *string `field:"optional" json:"autotrack" yaml:"autotrack"`
}

// Provides information about how AWS Ground Station should echo back uplink transmissions to a dataflow endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkEchoConfigProperty := &uplinkEchoConfigProperty{
//   	antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnConfig_UplinkEchoConfigProperty struct {
	// Defines the ARN of the uplink config to echo back to a dataflow endpoint.
	AntennaUplinkConfigArn *string `field:"optional" json:"antennaUplinkConfigArn" yaml:"antennaUplinkConfigArn"`
	// Whether or not uplink echo is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Defines a uplink spectrum.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkSpectrumConfigProperty := &uplinkSpectrumConfigProperty{
//   	centerFrequency: &frequencyProperty{
//   		units: jsii.String("units"),
//   		value: jsii.Number(123),
//   	},
//   	polarization: jsii.String("polarization"),
//   }
//
type CfnConfig_UplinkSpectrumConfigProperty struct {
	// The center frequency of the spectrum.
	//
	// Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	CenterFrequency interface{} `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// The polarization of the spectrum.
	//
	// Valid values are `"RIGHT_HAND"` and `"LEFT_HAND"` .
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

// Properties for defining a `CfnConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigProps := &cfnConfigProps{
//   	configData: &configDataProperty{
//   		antennaDownlinkConfig: &antennaDownlinkConfigProperty{
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaDownlinkDemodDecodeConfig: &antennaDownlinkDemodDecodeConfigProperty{
//   			decodeConfig: &decodeConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			demodulationConfig: &demodulationConfigProperty{
//   				unvalidatedJson: jsii.String("unvalidatedJson"),
//   			},
//   			spectrumConfig: &spectrumConfigProperty{
//   				bandwidth: &frequencyBandwidthProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   		},
//   		antennaUplinkConfig: &antennaUplinkConfigProperty{
//   			spectrumConfig: &uplinkSpectrumConfigProperty{
//   				centerFrequency: &frequencyProperty{
//   					units: jsii.String("units"),
//   					value: jsii.Number(123),
//   				},
//   				polarization: jsii.String("polarization"),
//   			},
//   			targetEirp: &eirpProperty{
//   				units: jsii.String("units"),
//   				value: jsii.Number(123),
//   			},
//   			transmitDisabled: jsii.Boolean(false),
//   		},
//   		dataflowEndpointConfig: &dataflowEndpointConfigProperty{
//   			dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   			dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   		},
//   		s3RecordingConfig: &s3RecordingConfigProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			prefix: jsii.String("prefix"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		trackingConfig: &trackingConfigProperty{
//   			autotrack: jsii.String("autotrack"),
//   		},
//   		uplinkEchoConfig: &uplinkEchoConfigProperty{
//   			antennaUplinkConfigArn: jsii.String("antennaUplinkConfigArn"),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfigProps struct {
	// Object containing the parameters of a config.
	//
	// Only one subtype may be specified per config. See the subtype definitions for a description of each config subtype.
	ConfigData interface{} `field:"required" json:"configData" yaml:"configData"`
	// The name of the config object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags assigned to a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::GroundStation::DataflowEndpointGroup`.
//
// Creates a Dataflow Endpoint Group request.
//
// Dataflow endpoint groups contain a list of endpoints. When the name of a dataflow endpoint group is specified in a mission profile, the Ground Station service will connect to the endpoints and flow data during a contact.
//
// For more information about dataflow endpoint groups, see [Dataflow Endpoint Groups](https://docs.aws.amazon.com/ground-station/latest/ug/dataflowendpointgroups.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroup := awscdk.Aws_groundstation.NewCfnDataflowEndpointGroup(this, jsii.String("MyCfnDataflowEndpointGroup"), &cfnDataflowEndpointGroupProps{
//   	endpointDetails: []interface{}{
//   		&endpointDetailsProperty{
//   			endpoint: &dataflowEndpointProperty{
//   				address: &socketAddressProperty{
//   					name: jsii.String("name"),
//   					port: jsii.Number(123),
//   				},
//   				mtu: jsii.Number(123),
//   				name: jsii.String("name"),
//   			},
//   			securityDetails: &securityDetailsProperty{
//   				roleArn: jsii.String("roleArn"),
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDataflowEndpointGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the dataflow endpoint group, such as `arn:aws:groundstation:us-east-2:1234567890:dataflow-endpoint-group/9940bf3b-d2ba-427e-9906-842b5e5d2296` .
	AttrArn() *string
	// UUID of a dataflow endpoint group.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// List of Endpoint Details, containing address and port for each endpoint.
	EndpointDetails() interface{}
	SetEndpointDetails(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags assigned to a resource.
	Tags() awscdk.TagManager
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
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

// The jsii proxy struct for CfnDataflowEndpointGroup
type jsiiProxy_CfnDataflowEndpointGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) EndpointDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::GroundStation::DataflowEndpointGroup`.
func NewCfnDataflowEndpointGroup(scope constructs.Construct, id *string, props *CfnDataflowEndpointGroupProps) CfnDataflowEndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDataflowEndpointGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GroundStation::DataflowEndpointGroup`.
func NewCfnDataflowEndpointGroup_Override(c CfnDataflowEndpointGroup, scope constructs.Construct, id *string, props *CfnDataflowEndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataflowEndpointGroup) SetEndpointDetails(val interface{}) {
	_jsii_.Set(
		j,
		"endpointDetails",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataflowEndpointGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataflowEndpointGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnDataflowEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataflowEndpointGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information such as socket address and name that defines an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointProperty := &dataflowEndpointProperty{
//   	address: &socketAddressProperty{
//   		name: jsii.String("name"),
//   		port: jsii.Number(123),
//   	},
//   	mtu: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnDataflowEndpointGroup_DataflowEndpointProperty struct {
	// The address and port of an endpoint.
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Valid values are between 1400 and 1500. A default value of 1500 is used if not set.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The endpoint name.
	//
	// When listing available contacts for a satellite, Ground Station searches for a dataflow endpoint whose name matches the value specified by the dataflow endpoint config of the selected mission profile. If no matching dataflow endpoints are found then Ground Station will not display any available contacts for the satellite.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// The security details and endpoint information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &endpointDetailsProperty{
//   	endpoint: &dataflowEndpointProperty{
//   		address: &socketAddressProperty{
//   			name: jsii.String("name"),
//   			port: jsii.Number(123),
//   		},
//   		mtu: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	securityDetails: &securityDetailsProperty{
//   		roleArn: jsii.String("roleArn"),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataflowEndpointGroup_EndpointDetailsProperty struct {
	// Information about the endpoint such as name and the endpoint address.
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The role ARN, and IDs for security groups and subnets.
	SecurityDetails interface{} `field:"optional" json:"securityDetails" yaml:"securityDetails"`
}

// Information about IAM roles, subnets, and security groups needed for this DataflowEndpointGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityDetailsProperty := &securityDetailsProperty{
//   	roleArn: jsii.String("roleArn"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDataflowEndpointGroup_SecurityDetailsProperty struct {
	// The ARN of a role which Ground Station has permission to assume, such as `arn:aws:iam::1234567890:role/DataDeliveryServiceRole` .
	//
	// Ground Station will assume this role and create an ENI in your VPC on the specified subnet upon creation of a dataflow endpoint group. This ENI is used as the ingress/egress point for data streamed during a satellite contact.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The security group Ids of the security role, such as `sg-1234567890abcdef0` .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet Ids of the security details, such as `subnet-12345678` .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

// The address of the endpoint, such as `192.168.1.1` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   socketAddressProperty := &socketAddressProperty{
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataflowEndpointGroup_SocketAddressProperty struct {
	// The name of the endpoint, such as `Endpoint 1` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port of the endpoint, such as `55888` .
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Properties for defining a `CfnDataflowEndpointGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroupProps := &cfnDataflowEndpointGroupProps{
//   	endpointDetails: []interface{}{
//   		&endpointDetailsProperty{
//   			endpoint: &dataflowEndpointProperty{
//   				address: &socketAddressProperty{
//   					name: jsii.String("name"),
//   					port: jsii.Number(123),
//   				},
//   				mtu: jsii.Number(123),
//   				name: jsii.String("name"),
//   			},
//   			securityDetails: &securityDetailsProperty{
//   				roleArn: jsii.String("roleArn"),
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataflowEndpointGroupProps struct {
	// List of Endpoint Details, containing address and port for each endpoint.
	EndpointDetails interface{} `field:"required" json:"endpointDetails" yaml:"endpointDetails"`
	// Tags assigned to a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::GroundStation::MissionProfile`.
//
// Mission profiles specify parameters and provide references to config objects to define how Ground Station lists and executes contacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMissionProfile := awscdk.Aws_groundstation.NewCfnMissionProfile(this, jsii.String("MyCfnMissionProfile"), &cfnMissionProfileProps{
//   	dataflowEdges: []interface{}{
//   		&dataflowEdgeProperty{
//   			destination: jsii.String("destination"),
//   			source: jsii.String("source"),
//   		},
//   	},
//   	minimumViableContactDurationSeconds: jsii.Number(123),
//   	name: jsii.String("name"),
//   	trackingConfigArn: jsii.String("trackingConfigArn"),
//
//   	// the properties below are optional
//   	contactPostPassDurationSeconds: jsii.Number(123),
//   	contactPrePassDurationSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnMissionProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the mission profile, such as `arn:aws:groundstation:us-east-2:1234567890:mission-profile/9940bf3b-d2ba-427e-9906-842b5e5d2296` .
	AttrArn() *string
	// The ID of the mission profile, such as `9940bf3b-d2ba-427e-9906-842b5e5d2296` .
	AttrId() *string
	// The region of the mission profile.
	AttrRegion() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Amount of time in seconds after a contact ends that you’d like to receive a CloudWatch Event indicating the pass has finished.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPostPassDurationSeconds() *float64
	SetContactPostPassDurationSeconds(val *float64)
	// Amount of time in seconds prior to contact start that you'd like to receive a CloudWatch Event indicating an upcoming pass.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPrePassDurationSeconds() *float64
	SetContactPrePassDurationSeconds(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A list containing lists of config ARNs.
	//
	// Each list of config ARNs is an edge, with a "from" config and a "to" config.
	DataflowEdges() interface{}
	SetDataflowEdges(val interface{})
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
	// Minimum length of a contact in seconds that Ground Station will return when listing contacts.
	//
	// Ground Station will not return contacts shorter than this duration.
	MinimumViableContactDurationSeconds() *float64
	SetMinimumViableContactDurationSeconds(val *float64)
	// The name of the mission profile.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags assigned to the mission profile.
	Tags() awscdk.TagManager
	// The ARN of a tracking config objects that defines how to track the satellite through the sky during a contact.
	TrackingConfigArn() *string
	SetTrackingConfigArn(val *string)
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
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

// The jsii proxy struct for CfnMissionProfile
type jsiiProxy_CfnMissionProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMissionProfile) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) AttrRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) ContactPostPassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPostPassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) ContactPrePassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPrePassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) DataflowEdges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataflowEdges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) MinimumViableContactDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumViableContactDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) TrackingConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfile) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::GroundStation::MissionProfile`.
func NewCfnMissionProfile(scope constructs.Construct, id *string, props *CfnMissionProfileProps) CfnMissionProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnMissionProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GroundStation::MissionProfile`.
func NewCfnMissionProfile_Override(c CfnMissionProfile, scope constructs.Construct, id *string, props *CfnMissionProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetContactPostPassDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"contactPostPassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetContactPrePassDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"contactPrePassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetDataflowEdges(val interface{}) {
	_jsii_.Set(
		j,
		"dataflowEdges",
		val,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetMinimumViableContactDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"minimumViableContactDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMissionProfile) SetTrackingConfigArn(val *string) {
	_jsii_.Set(
		j,
		"trackingConfigArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMissionProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMissionProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnMissionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMissionProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMissionProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMissionProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMissionProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMissionProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMissionProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMissionProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMissionProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMissionProfile) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMissionProfile) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMissionProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMissionProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMissionProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMissionProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMissionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMissionProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A dataflow edge defines from where and to where data will flow during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEdgeProperty := &dataflowEdgeProperty{
//   	destination: jsii.String("destination"),
//   	source: jsii.String("source"),
//   }
//
type CfnMissionProfile_DataflowEdgeProperty struct {
	// The ARN of the destination for this dataflow edge.
	//
	// For example, specify the ARN of a dataflow endpoint config for a downlink edge or an antenna uplink config for an uplink edge.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The ARN of the source for this dataflow edge.
	//
	// For example, specify the ARN of an antenna downlink config for a downlink edge or a dataflow endpoint config for an uplink edge.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

// Properties for defining a `CfnMissionProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMissionProfileProps := &cfnMissionProfileProps{
//   	dataflowEdges: []interface{}{
//   		&dataflowEdgeProperty{
//   			destination: jsii.String("destination"),
//   			source: jsii.String("source"),
//   		},
//   	},
//   	minimumViableContactDurationSeconds: jsii.Number(123),
//   	name: jsii.String("name"),
//   	trackingConfigArn: jsii.String("trackingConfigArn"),
//
//   	// the properties below are optional
//   	contactPostPassDurationSeconds: jsii.Number(123),
//   	contactPrePassDurationSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMissionProfileProps struct {
	// A list containing lists of config ARNs.
	//
	// Each list of config ARNs is an edge, with a "from" config and a "to" config.
	DataflowEdges interface{} `field:"required" json:"dataflowEdges" yaml:"dataflowEdges"`
	// Minimum length of a contact in seconds that Ground Station will return when listing contacts.
	//
	// Ground Station will not return contacts shorter than this duration.
	MinimumViableContactDurationSeconds *float64 `field:"required" json:"minimumViableContactDurationSeconds" yaml:"minimumViableContactDurationSeconds"`
	// The name of the mission profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of a tracking config objects that defines how to track the satellite through the sky during a contact.
	TrackingConfigArn *string `field:"required" json:"trackingConfigArn" yaml:"trackingConfigArn"`
	// Amount of time in seconds after a contact ends that you’d like to receive a CloudWatch Event indicating the pass has finished.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Amount of time in seconds prior to contact start that you'd like to receive a CloudWatch Event indicating an upcoming pass.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// Tags assigned to the mission profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

