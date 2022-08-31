package awsmediapackage

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::MediaPackage::OriginEndpoint`.
//
// Create an endpoint on an AWS Elemental MediaPackage channel.
//
// An endpoint represents a single delivery point of a channel, and defines content output handling through various components, such as packaging protocols, DRM and encryption integration, and more.
//
// After it's created, an endpoint provides a fixed public URL. This URL remains the same throughout the lifetime of the endpoint, regardless of any failures or upgrades that might occur. Integrate the URL with a downstream CDN (such as Amazon CloudFront) or playback device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpoint := awscdk.Aws_mediapackage.NewCfnOriginEndpoint(this, jsii.String("MyCfnOriginEndpoint"), &cfnOriginEndpointProps{
//   	channelId: jsii.String("channelId"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	cmafPackage: &cmafPackageProperty{
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				adMarkers: jsii.String("adMarkers"),
//   				adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				adTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				playlistType: jsii.String("playlistType"),
//   				playlistWindowSeconds: jsii.Number(123),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentPrefix: jsii.String("segmentPrefix"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	dashPackage: &dashPackageProperty{
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		manifestLayout: jsii.String("manifestLayout"),
//   		manifestWindowSeconds: jsii.Number(123),
//   		minBufferTimeSeconds: jsii.Number(123),
//   		minUpdatePeriodSeconds: jsii.Number(123),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		profile: jsii.String("profile"),
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		suggestedPresentationDelaySeconds: jsii.Number(123),
//   		utcTiming: jsii.String("utcTiming"),
//   		utcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	description: jsii.String("description"),
//   	hlsPackage: &hlsPackageProperty{
//   		adMarkers: jsii.String("adMarkers"),
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   		},
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		playlistType: jsii.String("playlistType"),
//   		playlistWindowSeconds: jsii.Number(123),
//   		programDateTimeIntervalSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	manifestName: jsii.String("manifestName"),
//   	mssPackage: &mssPackageProperty{
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		manifestWindowSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	origination: jsii.String("origination"),
//   	startoverWindowSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeDelaySeconds: jsii.Number(123),
//   	whitelist: []*string{
//   		jsii.String("whitelist"),
//   	},
//   })
//
type CfnOriginEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The endpoint's unique system-generated resource name, based on the AWS record.
	AttrArn() *string
	// URL for the key provider’s key retrieval API endpoint.
	//
	// Must start with https://.
	AttrUrl() *string
	// Parameters for CDN authorization.
	Authorization() interface{}
	SetAuthorization(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The ID of the channel associated with this endpoint.
	ChannelId() *string
	SetChannelId(val *string)
	// Parameters for Common Media Application Format (CMAF) packaging.
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Parameters for DASH packaging.
	DashPackage() interface{}
	SetDashPackage(val interface{})
	// Any descriptive information that you want to add to the endpoint for future identification purposes.
	Description() *string
	SetDescription(val *string)
	// Parameters for Apple HLS packaging.
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
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
	// Experimental.
	LogicalId() *string
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	ManifestName() *string
	SetManifestName(val *string)
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage() interface{}
	SetMssPackage(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Controls video origination from this endpoint.
	//
	// - `ALLOW` - enables this endpoint to serve content to requesting devices.
	// - `DENY` - prevents this endpoint from serving content. Denying origination is helpful for harvesting live-to-VOD assets. For more information about harvesting and origination, see [Live-to-VOD Requirements](https://docs.aws.amazon.com/mediapackage/latest/ug/ltov-reqmts.html) .
	Origination() *string
	SetOrigination(val *string)
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
	// Maximum duration (seconds) of content to retain for startover playback.
	//
	// Omit this attribute or enter `0` to indicate that startover playback is disabled for this endpoint.
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	// The tags to assign to the endpoint.
	Tags() awscdk.TagManager
	// Minimum duration (seconds) of delay to enforce on the playback of live content.
	//
	// Omit this attribute or enter `0` to indicate that there is no time delay in effect for this endpoint.
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The IP addresses that can access this endpoint.
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
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

// The jsii proxy struct for CfnOriginEndpoint
type jsiiProxy_CfnOriginEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Authorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Origination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) StartoverWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint(scope awscdk.Construct, id *string, props *CfnOriginEndpointProps) CfnOriginEndpoint {
	_init_.Initialize()

	if err := validateNewCfnOriginEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginEndpoint{}

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint_Override(c CfnOriginEndpoint, scope awscdk.Construct, id *string, props *CfnOriginEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetAuthorization(val interface{}) {
	if err := j.validateSetAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetCmafPackage(val interface{}) {
	if err := j.validateSetCmafPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetDashPackage(val interface{}) {
	if err := j.validateSetDashPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetHlsPackage(val interface{}) {
	if err := j.validateSetHlsPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetManifestName(val *string) {
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetMssPackage(val interface{}) {
	if err := j.validateSetMssPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetOrigination(val *string) {
	_jsii_.Set(
		j,
		"origination",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetStartoverWindowSeconds(val *float64) {
	_jsii_.Set(
		j,
		"startoverWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetTimeDelaySeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetWhitelist(val *[]*string) {
	_jsii_.Set(
		j,
		"whitelist",
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
func CfnOriginEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnOriginEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnOriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_mediapackage.CfnOriginEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnOriginEndpoint) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnOriginEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

