package awsmediapackagev2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the configuration parameters for a MediaPackage V2 origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpoint := awscdk.Aws_mediapackagev2.NewCfnOriginEndpoint(this, jsii.String("MyCfnOriginEndpoint"), &CfnOriginEndpointProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	ContainerType: jsii.String("containerType"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//
//   	// the properties below are optional
//   	DashManifests: []interface{}{
//   		&DashManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			BaseUrls: []interface{}{
//   				&DashBaseUrlProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					DvbPriority: jsii.Number(123),
//   					DvbWeight: jsii.Number(123),
//   					ServiceLocation: jsii.String("serviceLocation"),
//   				},
//   			},
//   			Compactness: jsii.String("compactness"),
//   			DrmSignaling: jsii.String("drmSignaling"),
//   			DvbSettings: &DashDvbSettingsProperty{
//   				ErrorMetrics: []interface{}{
//   					&DashDvbMetricsReportingProperty{
//   						ReportingUrl: jsii.String("reportingUrl"),
//
//   						// the properties below are optional
//   						Probability: jsii.Number(123),
//   					},
//   				},
//   				FontDownload: &DashDvbFontDownloadProperty{
//   					FontFamily: jsii.String("fontFamily"),
//   					MimeType: jsii.String("mimeType"),
//   					Url: jsii.String("url"),
//   				},
//   			},
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			MinBufferTimeSeconds: jsii.Number(123),
//   			MinUpdatePeriodSeconds: jsii.Number(123),
//   			PeriodTriggers: []*string{
//   				jsii.String("periodTriggers"),
//   			},
//   			Profiles: []*string{
//   				jsii.String("profiles"),
//   			},
//   			ProgramInformation: &DashProgramInformationProperty{
//   				Copyright: jsii.String("copyright"),
//   				LanguageCode: jsii.String("languageCode"),
//   				MoreInformationUrl: jsii.String("moreInformationUrl"),
//   				Source: jsii.String("source"),
//   				Title: jsii.String("title"),
//   			},
//   			ScteDash: &ScteDashProperty{
//   				AdMarkerDash: jsii.String("adMarkerDash"),
//   			},
//   			SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   			SubtitleConfiguration: &DashSubtitleConfigurationProperty{
//   				TtmlConfiguration: &DashTtmlConfigurationProperty{
//   					TtmlProfile: jsii.String("ttmlProfile"),
//   				},
//   			},
//   			SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			UtcTiming: &DashUtcTimingProperty{
//   				TimingMode: jsii.String("timingMode"),
//   				TimingSource: jsii.String("timingSource"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ForceEndpointErrorConfiguration: &ForceEndpointErrorConfigurationProperty{
//   		EndpointErrorConditions: []*string{
//   			jsii.String("endpointErrorConditions"),
//   		},
//   	},
//   	HlsManifests: []interface{}{
//   		&HlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			StartTag: &StartTagProperty{
//   				TimeOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Precise: jsii.Boolean(false),
//   			},
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	LowLatencyHlsManifests: []interface{}{
//   		&LowLatencyHlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			StartTag: &StartTagProperty{
//   				TimeOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Precise: jsii.Boolean(false),
//   			},
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	MssManifests: []interface{}{
//   		&MssManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestLayout: jsii.String("manifestLayout"),
//   			ManifestWindowSeconds: jsii.Number(123),
//   		},
//   	},
//   	Segment: &SegmentProperty{
//   		Encryption: &EncryptionProperty{
//   			EncryptionMethod: &EncryptionMethodProperty{
//   				CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   				IsmEncryptionMethod: jsii.String("ismEncryptionMethod"),
//   				TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   			},
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				DrmSystems: []*string{
//   					jsii.String("drmSystems"),
//   				},
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				Url: jsii.String("url"),
//   			},
//
//   			// the properties below are optional
//   			CmafExcludeSegmentDrmMetadata: jsii.Boolean(false),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		IncludeIframeOnlyStreams: jsii.Boolean(false),
//   		Scte: &ScteProperty{
//   			ScteFilter: []*string{
//   				jsii.String("scteFilter"),
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		TsIncludeDvbSubtitles: jsii.Boolean(false),
//   		TsUseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	StartoverWindowSeconds: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html
//
type CfnOriginEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsmediapackagev2.IOriginEndpointRef
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the origin endpoint.
	AttrArn() *string
	// The timestamp of the creation of the origin endpoint.
	AttrCreatedAt() *string
	// The egress domain URL for stream delivery from MediaPackage.
	AttrDashManifestUrls() *[]*string
	// The egress domain URL for stream delivery from MediaPackage.
	AttrHlsManifestUrls() *[]*string
	// The egress domain URL for stream delivery from MediaPackage.
	AttrLowLatencyHlsManifestUrls() *[]*string
	// The timestamp of the modification of the origin endpoint.
	AttrModifiedAt() *string
	AttrMssManifestUrls() *[]*string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name of the channel group associated with the origin endpoint configuration.
	ChannelGroupName() *string
	SetChannelGroupName(val *string)
	// The channel name associated with the origin endpoint.
	ChannelName() *string
	SetChannelName(val *string)
	// The container type associated with the origin endpoint configuration.
	ContainerType() *string
	SetContainerType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A DASH manifest configuration.
	DashManifests() interface{}
	SetDashManifests(val interface{})
	// The description associated with the origin endpoint.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// The failover settings for the endpoint.
	ForceEndpointErrorConfiguration() interface{}
	SetForceEndpointErrorConfiguration(val interface{})
	// The HLS manifests associated with the origin endpoint configuration.
	HlsManifests() interface{}
	SetHlsManifests(val interface{})
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
	// The low-latency HLS (LL-HLS) manifests associated with the origin endpoint.
	LowLatencyHlsManifests() interface{}
	SetLowLatencyHlsManifests(val interface{})
	// A list of Microsoft Smooth Streaming (MSS) manifest configurations associated with the origin endpoint.
	MssManifests() interface{}
	SetMssManifests(val interface{})
	// The tree node.
	Node() constructs.Node
	// The name of the origin endpoint associated with the origin endpoint configuration.
	OriginEndpointName() *string
	SetOriginEndpointName(val *string)
	// A reference to a OriginEndpoint resource.
	OriginEndpointRef() *interfacesawsmediapackagev2.OriginEndpointReference
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The segment associated with the origin endpoint.
	Segment() interface{}
	SetSegment(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The size of the window (in seconds) to specify a window of the live stream that's available for on-demand viewing.
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	// The tags associated with the origin endpoint.
	Tags() *[]*awscdk.CfnTag
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

// The jsii proxy struct for CfnOriginEndpoint
type jsiiProxy_CfnOriginEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsmediapackagev2IOriginEndpointRef
	internal.Type__awscdkITaggableV2
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

func (j *jsiiProxy_CfnOriginEndpoint) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrDashManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDashManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrHlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrHlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrLowLatencyHlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrLowLatencyHlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrMssManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrMssManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
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

func (j *jsiiProxy_CfnOriginEndpoint) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
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

func (j *jsiiProxy_CfnOriginEndpoint) DashManifests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashManifests",
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

func (j *jsiiProxy_CfnOriginEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ForceEndpointErrorConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEndpointErrorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) HlsManifests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsManifests",
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

func (j *jsiiProxy_CfnOriginEndpoint) LowLatencyHlsManifests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowLatencyHlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) MssManifests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) OriginEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) OriginEndpointRef() *interfacesawsmediapackagev2.OriginEndpointReference {
	var returns *interfacesawsmediapackagev2.OriginEndpointReference
	_jsii_.Get(
		j,
		"originEndpointRef",
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

func (j *jsiiProxy_CfnOriginEndpoint) Segment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segment",
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

func (j *jsiiProxy_CfnOriginEndpoint) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
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

func (j *jsiiProxy_CfnOriginEndpoint) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackageV2::OriginEndpoint`.
func NewCfnOriginEndpoint(scope constructs.Construct, id *string, props *CfnOriginEndpointProps) CfnOriginEndpoint {
	_init_.Initialize()

	if err := validateNewCfnOriginEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackageV2::OriginEndpoint`.
func NewCfnOriginEndpoint_Override(c CfnOriginEndpoint, scope constructs.Construct, id *string, props *CfnOriginEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetChannelGroupName(val *string) {
	if err := j.validateSetChannelGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetContainerType(val *string) {
	if err := j.validateSetContainerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerType",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetDashManifests(val interface{}) {
	if err := j.validateSetDashManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashManifests",
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

func (j *jsiiProxy_CfnOriginEndpoint)SetForceEndpointErrorConfiguration(val interface{}) {
	if err := j.validateSetForceEndpointErrorConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceEndpointErrorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetHlsManifests(val interface{}) {
	if err := j.validateSetHlsManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsManifests",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetLowLatencyHlsManifests(val interface{}) {
	if err := j.validateSetLowLatencyHlsManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lowLatencyHlsManifests",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetMssManifests(val interface{}) {
	if err := j.validateSetMssManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mssManifests",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetOriginEndpointName(val *string) {
	if err := j.validateSetOriginEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originEndpointName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint)SetSegment(val interface{}) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
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

func (j *jsiiProxy_CfnOriginEndpoint)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOriginEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnOriginEndpoint_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
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
func CfnOriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
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
		"aws-cdk-lib.aws_mediapackagev2.CfnOriginEndpoint",
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

func (c *jsiiProxy_CfnOriginEndpoint) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnOriginEndpoint) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnOriginEndpoint) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnOriginEndpoint) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnOriginEndpoint) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

