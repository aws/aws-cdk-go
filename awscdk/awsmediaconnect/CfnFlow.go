package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MediaConnect::Flow` resource defines a connection between one or more video sources and one or more outputs.
//
// For each flow, you specify the transport protocol to use, encryption information, and details for any outputs or entitlements that you want. AWS Elemental MediaConnect returns an ingest endpoint where you can send your live video as a single unicast stream. The service replicates and distributes the video to every output that you specify, whether inside or outside the AWS Cloud. You can also set up entitlements on a flow to allow other AWS accounts to access your content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlow := awscdk.Aws_mediaconnect.NewCfnFlow(this, jsii.String("MyCfnFlow"), &CfnFlowProps{
//   	Name: jsii.String("name"),
//   	Source: &SourceProperty{
//   		Decryption: &EncryptionProperty{
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			Algorithm: jsii.String("algorithm"),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			DeviceId: jsii.String("deviceId"),
//   			KeyType: jsii.String("keyType"),
//   			Region: jsii.String("region"),
//   			ResourceId: jsii.String("resourceId"),
//   			SecretArn: jsii.String("secretArn"),
//   			Url: jsii.String("url"),
//   		},
//   		Description: jsii.String("description"),
//   		EntitlementArn: jsii.String("entitlementArn"),
//   		GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   			BridgeArn: jsii.String("bridgeArn"),
//
//   			// the properties below are optional
//   			VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   				VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   			},
//   		},
//   		IngestIp: jsii.String("ingestIp"),
//   		IngestPort: jsii.Number(123),
//   		MaxBitrate: jsii.Number(123),
//   		MaxLatency: jsii.Number(123),
//   		MaxSyncBuffer: jsii.Number(123),
//   		MediaStreamSourceConfigurations: []interface{}{
//   			&MediaStreamSourceConfigurationProperty{
//   				EncodingName: jsii.String("encodingName"),
//   				MediaStreamName: jsii.String("mediaStreamName"),
//
//   				// the properties below are optional
//   				InputConfigurations: []interface{}{
//   					&InputConfigurationProperty{
//   						InputPort: jsii.Number(123),
//   						Interface: &InterfaceProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		MinLatency: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		Protocol: jsii.String("protocol"),
//   		SenderControlPort: jsii.Number(123),
//   		SenderIpAddress: jsii.String("senderIpAddress"),
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceIngestPort: jsii.String("sourceIngestPort"),
//   		SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   		SourceListenerPort: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		WhitelistCidr: jsii.String("whitelistCidr"),
//   	},
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	FlowSize: jsii.String("flowSize"),
//   	Maintenance: &MaintenanceProperty{
//   		MaintenanceDay: jsii.String("maintenanceDay"),
//   		MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   	},
//   	MediaStreams: []interface{}{
//   		&MediaStreamProperty{
//   			MediaStreamId: jsii.Number(123),
//   			MediaStreamName: jsii.String("mediaStreamName"),
//   			MediaStreamType: jsii.String("mediaStreamType"),
//
//   			// the properties below are optional
//   			Attributes: &MediaStreamAttributesProperty{
//   				Fmtp: &FmtpProperty{
//   					ChannelOrder: jsii.String("channelOrder"),
//   					Colorimetry: jsii.String("colorimetry"),
//   					ExactFramerate: jsii.String("exactFramerate"),
//   					Par: jsii.String("par"),
//   					Range: jsii.String("range"),
//   					ScanMode: jsii.String("scanMode"),
//   					Tcs: jsii.String("tcs"),
//   				},
//   				Lang: jsii.String("lang"),
//   			},
//   			ClockRate: jsii.Number(123),
//   			Description: jsii.String("description"),
//   			Fmt: jsii.Number(123),
//   			VideoFormat: jsii.String("videoFormat"),
//   		},
//   	},
//   	NdiConfig: &NdiConfigProperty{
//   		MachineName: jsii.String("machineName"),
//   		NdiDiscoveryServers: []interface{}{
//   			&NdiDiscoveryServerConfigProperty{
//   				DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   				VpcInterfaceAdapter: jsii.String("vpcInterfaceAdapter"),
//
//   				// the properties below are optional
//   				DiscoveryServerPort: jsii.Number(123),
//   			},
//   		},
//   		NdiState: jsii.String("ndiState"),
//   	},
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		RecoveryWindow: jsii.Number(123),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	SourceMonitoringConfig: &SourceMonitoringConfigProperty{
//   		AudioMonitoringSettings: []interface{}{
//   			&AudioMonitoringSettingProperty{
//   				SilentAudio: &SilentAudioProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ContentQualityAnalysisState: jsii.String("contentQualityAnalysisState"),
//   		ThumbnailState: jsii.String("thumbnailState"),
//   		VideoMonitoringSettings: []interface{}{
//   			&VideoMonitoringSettingProperty{
//   				BlackFrames: &BlackFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   				FrozenFrames: &FrozenFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	VpcInterfaces: []interface{}{
//   		&VpcInterfaceProperty{
//   			Name: jsii.String("name"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: jsii.String("networkInterfaceType"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html
//
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The IP address from which video will be sent to output destinations.
	AttrEgressIp() *string
	// The Amazon Resource Name (ARN) of the flow.
	AttrFlowArn() *string
	// The Availability Zone that the flow was created in.
	//
	// These options are limited to the Availability Zones within the current AWS Region.
	AttrFlowAvailabilityZone() *string
	// This read-only value represents the automatically-generated NDI machine name that MediaConnect generated for this flow.
	//
	// These NDI machine names are only generated when you don't specify your own custom name.
	AttrFlowNdiMachineName() *string
	// The IP address that the flow listens on for incoming content.
	AttrSourceIngestIp() *string
	// The ARN of the source.
	AttrSourceSourceArn() *string
	// The port that the flow listens on for incoming content.
	//
	// If the protocol of the source is Zixi, the port must be set to 2088.
	AttrSourceSourceIngestPort() *string
	// The Availability Zone that you want to create the flow in.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Determines the processing capacity and feature set of the flow.
	FlowSize() *string
	SetFlowSize(val *string)
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
	// The maintenance settings you want to use for the flow.
	Maintenance() interface{}
	SetMaintenance(val interface{})
	// The media streams that are associated with the flow.
	MediaStreams() interface{}
	SetMediaStreams(val interface{})
	// The name of the flow.
	Name() *string
	SetName(val *string)
	// Specifies the configuration settings for NDI outputs.
	NdiConfig() interface{}
	SetNdiConfig(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The settings for the source that you want to use for the new flow.
	Source() interface{}
	SetSource(val interface{})
	// The settings for source failover.
	SourceFailoverConfig() interface{}
	SetSourceFailoverConfig(val interface{})
	// The settings for source monitoring.
	SourceMonitoringConfig() interface{}
	SetSourceMonitoringConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// The VPC Interfaces for this flow.
	VpcInterfaces() interface{}
	SetVpcInterfaces(val interface{})
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

// The jsii proxy struct for CfnFlow
type jsiiProxy_CfnFlow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlow) AttrEgressIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEgressIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrFlowAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrFlowNdiMachineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowNdiMachineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceIngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceIngestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AttrSourceSourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) FlowSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Maintenance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) MediaStreams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) NdiConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ndiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceFailoverConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFailoverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceMonitoringConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceMonitoringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) VpcInterfaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInterfaces",
		&returns,
	)
	return returns
}


func NewCfnFlow(scope constructs.Construct, id *string, props *CfnFlowProps) CfnFlow {
	_init_.Initialize()

	if err := validateNewCfnFlowParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlow{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnFlow_Override(c CfnFlow, scope constructs.Construct, id *string, props *CfnFlowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlow)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetFlowSize(val *string) {
	_jsii_.Set(
		j,
		"flowSize",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetMaintenance(val interface{}) {
	if err := j.validateSetMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenance",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetMediaStreams(val interface{}) {
	if err := j.validateSetMediaStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStreams",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetNdiConfig(val interface{}) {
	if err := j.validateSetNdiConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ndiConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetSource(val interface{}) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetSourceFailoverConfig(val interface{}) {
	if err := j.validateSetSourceFailoverConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFailoverConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetSourceMonitoringConfig(val interface{}) {
	if err := j.validateSetSourceMonitoringConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMonitoringConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetVpcInterfaces(val interface{}) {
	if err := j.validateSetVpcInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcInterfaces",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnFlow_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
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
func CfnFlow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconnect.CfnFlow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFlow) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFlow) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlow) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlow) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFlow) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFlow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

