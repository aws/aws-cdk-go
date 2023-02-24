package awsmediapackage

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaPackage::PackagingConfiguration`.
//
// Creates a packaging configuration in a packaging group.
//
// The packaging configuration represents a single delivery point for an asset. It determines the format and setting for the egressing content. Specify only one package format per configuration, such as `HlsPackage` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingConfiguration := awscdk.Aws_mediapackage.NewCfnPackagingConfiguration(this, jsii.String("MyCfnPackagingConfiguration"), &CfnPackagingConfigurationProps{
//   	Id: jsii.String("id"),
//   	PackagingGroupId: jsii.String("packagingGroupId"),
//
//   	// the properties below are optional
//   	CmafPackage: &CmafPackageProperty{
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				AdMarkers: jsii.String("adMarkers"),
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &CmafEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	DashPackage: &DashPackageProperty{
//   		DashManifests: []interface{}{
//   			&DashManifestProperty{
//   				ManifestLayout: jsii.String("manifestLayout"),
//   				ManifestName: jsii.String("manifestName"),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				Profile: jsii.String("profile"),
//   				ScteMarkersSource: jsii.String("scteMarkersSource"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &DashEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		PeriodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	HlsPackage: &HlsPackageProperty{
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				AdMarkers: jsii.String("adMarkers"),
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &HlsEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   		},
//   		IncludeDvbSubtitles: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		UseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	MssPackage: &MssPackageProperty{
//   		MssManifests: []interface{}{
//   			&MssManifestProperty{
//   				ManifestName: jsii.String("manifestName"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &MssEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPackagingConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the packaging configuration.
	//
	// You can get this from the response to any request to the packaging configuration.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Parameters for CMAF packaging.
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Parameters for DASH-ISO packaging.
	DashPackage() interface{}
	SetDashPackage(val interface{})
	// Parameters for Apple HLS packaging.
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	// Unique identifier that you assign to the packaging configuration.
	Id() *string
	SetId(val *string)
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
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage() interface{}
	SetMssPackage(val interface{})
	// The tree node.
	Node() constructs.Node
	// The ID of the packaging group associated with this packaging configuration.
	PackagingGroupId() *string
	SetPackagingGroupId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the packaging configuration.
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnPackagingConfiguration
type jsiiProxy_CfnPackagingConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPackagingConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) PackagingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration(scope constructs.Construct, id *string, props *CfnPackagingConfigurationProps) CfnPackagingConfiguration {
	_init_.Initialize()

	if err := validateNewCfnPackagingConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPackagingConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration_Override(c CfnPackagingConfiguration, scope constructs.Construct, id *string, props *CfnPackagingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetCmafPackage(val interface{}) {
	if err := j.validateSetCmafPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetDashPackage(val interface{}) {
	if err := j.validateSetDashPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetHlsPackage(val interface{}) {
	if err := j.validateSetHlsPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetMssPackage(val interface{}) {
	if err := j.validateSetMssPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration)SetPackagingGroupId(val *string) {
	if err := j.validateSetPackagingGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packagingGroupId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPackagingConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPackagingConfiguration_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPackagingConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnPackagingConfiguration_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
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
func CfnPackagingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPackagingConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

