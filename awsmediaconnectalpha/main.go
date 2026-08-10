// The CDK construct library for AWS Elemental MediaConnect
package awsmediaconnectalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.AddFlowOutputOptions",
		reflect.TypeOf((*AddFlowOutputOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.AudioStreamOrderOptions",
		reflect.TypeOf((*AudioStreamOrderOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_AudioStreamOrderOptions{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Bridge",
		reflect.TypeOf((*Bridge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOutput", GoMethod: "AddOutput"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeName", GoGetter: "BridgeName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeRef", GoGetter: "BridgeRef"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeState", GoGetter: "BridgeState"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeType", GoGetter: "BridgeType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isFailoverEnabled", GoGetter: "IsFailoverEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailoverSwitches", GoMethod: "MetricFailoverSwitches"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceBitrate", GoMethod: "MetricSourceBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourcePacketLossPercent", GoMethod: "MetricSourcePacketLossPercent"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Bridge{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBridge)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeAttributes",
		reflect.TypeOf((*BridgeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeConfiguration",
		reflect.TypeOf((*BridgeConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BridgeConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeFailoverConfig",
		reflect.TypeOf((*BridgeFailoverConfig)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BridgeFailoverConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeFailoverOptions",
		reflect.TypeOf((*BridgeFailoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeFlowInput",
		reflect.TypeOf((*BridgeFlowInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeFlowSource",
		reflect.TypeOf((*BridgeFlowSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeNetworkInput",
		reflect.TypeOf((*BridgeNetworkInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeNetworkOutput",
		reflect.TypeOf((*BridgeNetworkOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeNetworkOutputProps",
		reflect.TypeOf((*BridgeNetworkOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeNetworkSource",
		reflect.TypeOf((*BridgeNetworkSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeOutput",
		reflect.TypeOf((*BridgeOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeOutputName", GoGetter: "BridgeOutputName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeOutputRef", GoGetter: "BridgeOutputRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_BridgeOutput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBridgeOutput)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeOutputConfiguration",
		reflect.TypeOf((*BridgeOutputConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BridgeOutputConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeOutputProps",
		reflect.TypeOf((*BridgeOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProps",
		reflect.TypeOf((*BridgeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProtocol",
		reflect.TypeOf((*BridgeProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_BridgeProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSource",
		reflect.TypeOf((*BridgeSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeSourceName", GoGetter: "BridgeSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeSourceRef", GoGetter: "BridgeSourceRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_BridgeSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBridgeSource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceConfiguration",
		reflect.TypeOf((*BridgeSourceConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BridgeSourceConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceProps",
		reflect.TypeOf((*BridgeSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeType",
		reflect.TypeOf((*BridgeType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_BridgeType{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Colorimetry",
		reflect.TypeOf((*Colorimetry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Colorimetry{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.EgressBridgeConfiguration",
		reflect.TypeOf((*EgressBridgeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		reflect.TypeOf((*Encoding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Encoding{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.EncodingConfig",
		reflect.TypeOf((*EncodingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.EncodingProfile",
		reflect.TypeOf((*EncodingProfile)(nil)).Elem(),
		map[string]interface{}{
			"CONTRIBUTION_H264_DEFAULT": EncodingProfile_CONTRIBUTION_H264_DEFAULT,
			"DISTRIBUTION_H264_DEFAULT": EncodingProfile_DISTRIBUTION_H264_DEFAULT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.EncryptionAlgorithm",
		reflect.TypeOf((*EncryptionAlgorithm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_EncryptionAlgorithm{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.EntitlementSource",
		reflect.TypeOf((*EntitlementSource)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.EntitlementStatus",
		reflect.TypeOf((*EntitlementStatus)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": EntitlementStatus_ENABLED,
			"DISABLED": EntitlementStatus_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverConfig",
		reflect.TypeOf((*FailoverConfig)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FailoverConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverConfigurationProps",
		reflect.TypeOf((*FailoverConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverFailoverOptions",
		reflect.TypeOf((*FailoverFailoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverMode",
		reflect.TypeOf((*FailoverMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FailoverMode{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Flow",
		reflect.TypeOf((*Flow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOutput", GoMethod: "AddOutput"},
			_jsii_.MemberMethod{JsiiMethod: "addVpcInterface", GoMethod: "AddVpcInterface"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "egressIp", GoGetter: "EgressIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowArn", GoGetter: "FlowArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowAvailabilityZone", GoGetter: "FlowAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "flowRef", GoGetter: "FlowRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "isFailoverEnabled", GoGetter: "IsFailoverEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceBitrate", GoMethod: "MetricSourceBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceConnected", GoMethod: "MetricSourceConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceDisconnections", GoMethod: "MetricSourceDisconnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceDroppedPackets", GoMethod: "MetricSourceDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceJitter", GoMethod: "MetricSourceJitter"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceNotRecoveredPackets", GoMethod: "MetricSourceNotRecoveredPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourcePacketLossPercent", GoMethod: "MetricSourcePacketLossPercent"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceRoundTripTime", GoMethod: "MetricSourceRoundTripTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceSelected", GoMethod: "MetricSourceSelected"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceTotalPackets", GoMethod: "MetricSourceTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestIp", GoGetter: "SourceIngestIp"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestPort", GoGetter: "SourceIngestPort"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestUrl", GoGetter: "SourceIngestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Flow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowAttributes",
		reflect.TypeOf((*FlowAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FlowEntitlement",
		reflect.TypeOf((*FlowEntitlement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "entitlementArn", GoGetter: "EntitlementArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowEntitlementRef", GoGetter: "FlowEntitlementRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FlowEntitlement{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlowEntitlement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowEntitlementProps",
		reflect.TypeOf((*FlowEntitlementProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FlowGrants",
		reflect.TypeOf((*FlowGrants)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "actions", GoMethod: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "start", GoMethod: "Start"},
			_jsii_.MemberMethod{JsiiMethod: "stop", GoMethod: "Stop"},
		},
		func() interface{} {
			return &jsiiProxy_FlowGrants{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FlowOutput",
		reflect.TypeOf((*FlowOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowOutputArn", GoGetter: "FlowOutputArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowOutputRef", GoGetter: "FlowOutputRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FlowOutput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlowOutput)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowOutputProps",
		reflect.TypeOf((*FlowOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowProps",
		reflect.TypeOf((*FlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSize",
		reflect.TypeOf((*FlowSize)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FlowSize{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSource",
		reflect.TypeOf((*FlowSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceArn", GoGetter: "FlowSourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceName", GoGetter: "FlowSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceRef", GoGetter: "FlowSourceRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ingestIp", GoGetter: "IngestIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestPort", GoGetter: "SourceIngestPort"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_FlowSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlowSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSourceAttributes",
		reflect.TypeOf((*FlowSourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSourceProps",
		reflect.TypeOf((*FlowSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.FmtpVideo",
		reflect.TypeOf((*FmtpVideo)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.ForwardErrorCorrection",
		reflect.TypeOf((*ForwardErrorCorrection)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": ForwardErrorCorrection_ENABLED,
			"DISABLED": ForwardErrorCorrection_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		reflect.TypeOf((*Framerate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_Framerate{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Gateway",
		reflect.TypeOf((*Gateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addNetwork", GoMethod: "AddNetwork"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayArn", GoGetter: "GatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayState", GoGetter: "GatewayState"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBridgeDroppedPackets", GoMethod: "MetricEgressBridgeDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBridgeTotalPackets", GoMethod: "MetricEgressBridgeTotalPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBridgeDroppedPackets", GoMethod: "MetricIngressBridgeDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBridgeTotalPackets", GoMethod: "MetricIngressBridgeTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Gateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGateway)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.GatewayBridgeSource",
		reflect.TypeOf((*GatewayBridgeSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.GatewayNetwork",
		reflect.TypeOf((*GatewayNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_GatewayNetwork{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.GatewayNetworkDefineProps",
		reflect.TypeOf((*GatewayNetworkDefineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.GatewayProps",
		reflect.TypeOf((*GatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IBridge",
		reflect.TypeOf((*IBridge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOutput", GoMethod: "AddOutput"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeName", GoGetter: "BridgeName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeRef", GoGetter: "BridgeRef"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeState", GoGetter: "BridgeState"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeType", GoGetter: "BridgeType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "isFailoverEnabled", GoGetter: "IsFailoverEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailoverSwitches", GoMethod: "MetricFailoverSwitches"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceBitrate", GoMethod: "MetricSourceBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourcePacketLossPercent", GoMethod: "MetricSourcePacketLossPercent"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBridge{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIBridgeRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IBridgeOutput",
		reflect.TypeOf((*IBridgeOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeOutputName", GoGetter: "BridgeOutputName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeOutputRef", GoGetter: "BridgeOutputRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBridgeOutput{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIBridgeOutputRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IBridgeSource",
		reflect.TypeOf((*IBridgeSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeArn", GoGetter: "BridgeArn"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeSourceName", GoGetter: "BridgeSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "bridgeSourceRef", GoGetter: "BridgeSourceRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IBridgeSource{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIBridgeSourceRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IFlow",
		reflect.TypeOf((*IFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOutput", GoMethod: "AddOutput"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "egressIp", GoGetter: "EgressIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowArn", GoGetter: "FlowArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowAvailabilityZone", GoGetter: "FlowAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "flowRef", GoGetter: "FlowRef"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "isFailoverEnabled", GoGetter: "IsFailoverEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceBitrate", GoMethod: "MetricSourceBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceConnected", GoMethod: "MetricSourceConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceDisconnections", GoMethod: "MetricSourceDisconnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceDroppedPackets", GoMethod: "MetricSourceDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceJitter", GoMethod: "MetricSourceJitter"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceNotRecoveredPackets", GoMethod: "MetricSourceNotRecoveredPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourcePacketLossPercent", GoMethod: "MetricSourcePacketLossPercent"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceRoundTripTime", GoMethod: "MetricSourceRoundTripTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceSelected", GoMethod: "MetricSourceSelected"},
			_jsii_.MemberMethod{JsiiMethod: "metricSourceTotalPackets", GoMethod: "MetricSourceTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestIp", GoGetter: "SourceIngestIp"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestPort", GoGetter: "SourceIngestPort"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestUrl", GoGetter: "SourceIngestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFlow{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIFlowRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IFlowEntitlement",
		reflect.TypeOf((*IFlowEntitlement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "entitlementArn", GoGetter: "EntitlementArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowEntitlementRef", GoGetter: "FlowEntitlementRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFlowEntitlement{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIFlowEntitlementRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IFlowOutput",
		reflect.TypeOf((*IFlowOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowOutputArn", GoGetter: "FlowOutputArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowOutputRef", GoGetter: "FlowOutputRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFlowOutput{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIFlowOutputRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IFlowSource",
		reflect.TypeOf((*IFlowSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceArn", GoGetter: "FlowSourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceName", GoGetter: "FlowSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "flowSourceRef", GoGetter: "FlowSourceRef"},
			_jsii_.MemberProperty{JsiiProperty: "ingestIp", GoGetter: "IngestIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIngestPort", GoGetter: "SourceIngestPort"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IFlowSource{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIFlowSourceRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IGateway",
		reflect.TypeOf((*IGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayArn", GoGetter: "GatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayRef", GoGetter: "GatewayRef"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayState", GoGetter: "GatewayState"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBridgeDroppedPackets", GoMethod: "MetricEgressBridgeDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBridgeTotalPackets", GoMethod: "MetricEgressBridgeTotalPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBridgeDroppedPackets", GoMethod: "MetricIngressBridgeDroppedPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBridgeTotalPackets", GoMethod: "MetricIngressBridgeTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGateway{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIGatewayRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IRouterInput",
		reflect.TypeOf((*IRouterInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBitrate", GoMethod: "MetricBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnected", GoMethod: "MetricConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricContinuityCounterErrors", GoMethod: "MetricContinuityCounterErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailoverSwitches", GoMethod: "MetricFailoverSwitches"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricNotRecoveredPackets", GoMethod: "MetricNotRecoveredPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalPackets", GoMethod: "MetricTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputArn", GoGetter: "RouterInputArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputId", GoGetter: "RouterInputId"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputRef", GoGetter: "RouterInputRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "transitEncryptionSecret", GoGetter: "TransitEncryptionSecret"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRouterInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIRouterInputRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IRouterNetworkInterface",
		reflect.TypeOf((*IRouterNetworkInterface)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceArn", GoGetter: "RouterNetworkInterfaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceId", GoGetter: "RouterNetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceRef", GoGetter: "RouterNetworkInterfaceRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRouterNetworkInterface{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIRouterNetworkInterfaceRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediaconnect-alpha.IRouterOutput",
		reflect.TypeOf((*IRouterOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricArqRequests", GoMethod: "MetricArqRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricBitrate", GoMethod: "MetricBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnected", GoMethod: "MetricConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalPackets", GoMethod: "MetricTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputArn", GoGetter: "RouterOutputArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputId", GoGetter: "RouterOutputId"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputName", GoGetter: "RouterOutputName"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputRef", GoGetter: "RouterOutputRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRouterOutput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediaconnectIRouterOutputRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.IngressBridgeConfiguration",
		reflect.TypeOf((*IngressBridgeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.KeyType",
		reflect.TypeOf((*KeyType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_KeyType{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MaintenanceConfiguration",
		reflect.TypeOf((*MaintenanceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.MaintenanceDay",
		reflect.TypeOf((*MaintenanceDay)(nil)).Elem(),
		map[string]interface{}{
			"MONDAY": MaintenanceDay_MONDAY,
			"TUESDAY": MaintenanceDay_TUESDAY,
			"WEDNESDAY": MaintenanceDay_WEDNESDAY,
			"THURSDAY": MaintenanceDay_THURSDAY,
			"FRIDAY": MaintenanceDay_FRIDAY,
			"SATURDAY": MaintenanceDay_SATURDAY,
			"SUNDAY": MaintenanceDay_SUNDAY,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MaintenanceWindow",
		reflect.TypeOf((*MaintenanceWindow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaConnectFlowConfigurationProps",
		reflect.TypeOf((*MediaConnectFlowConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaConnectFlowConfigurationWithoutConnectionProps",
		reflect.TypeOf((*MediaConnectFlowConfigurationWithoutConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaConnectFlowConnectionProps",
		reflect.TypeOf((*MediaConnectFlowConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaConnectFlowNoConnectionProps",
		reflect.TypeOf((*MediaConnectFlowNoConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaLiveChannelConfigurationProps",
		reflect.TypeOf((*MediaLiveChannelConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaLiveChannelConfigurationWithoutConnectionProps",
		reflect.TypeOf((*MediaLiveChannelConfigurationWithoutConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaLiveInputConnectionProps",
		reflect.TypeOf((*MediaLiveInputConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaLiveNoInputConnectionProps",
		reflect.TypeOf((*MediaLiveNoInputConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.MediaLivePipeline",
		reflect.TypeOf((*MediaLivePipeline)(nil)).Elem(),
		map[string]interface{}{
			"PIPELINE_0": MediaLivePipeline_PIPELINE_0,
			"PIPELINE_1": MediaLivePipeline_PIPELINE_1,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStream",
		reflect.TypeOf((*MediaStream)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MediaStream{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamAncillaryData",
		reflect.TypeOf((*MediaStreamAncillaryData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamAudio",
		reflect.TypeOf((*MediaStreamAudio)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamBase",
		reflect.TypeOf((*MediaStreamBase)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamSourceConfigurationCdi",
		reflect.TypeOf((*MediaStreamSourceConfigurationCdi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamSourceConfigurationJpegXs",
		reflect.TypeOf((*MediaStreamSourceConfigurationJpegXs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStreamVideo",
		reflect.TypeOf((*MediaStreamVideo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		reflect.TypeOf((*MediaVideoFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MediaVideoFormat{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MergeConfigurationProps",
		reflect.TypeOf((*MergeConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MergeFailoverOptions",
		reflect.TypeOf((*MergeFailoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.MonitoringMetric",
		reflect.TypeOf((*MonitoringMetric)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.NdiConfig",
		reflect.TypeOf((*NdiConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.NdiDiscoveryServerConfig",
		reflect.TypeOf((*NdiDiscoveryServerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.NdiOutputConfig",
		reflect.TypeOf((*NdiOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkConfiguration",
		reflect.TypeOf((*NetworkConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowlistCidr", GoGetter: "AllowlistCidr"},
			_jsii_.MemberProperty{JsiiProperty: "vpcInterfaceName", GoGetter: "VpcInterfaceName"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkInterface",
		reflect.TypeOf((*NetworkInterface)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkInterface{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.OutputConfiguration",
		reflect.TypeOf((*OutputConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OutputConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.PixelAspectRatio",
		reflect.TypeOf((*PixelAspectRatio)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_PixelAspectRatio{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.PrimarySource",
		reflect.TypeOf((*PrimarySource)(nil)).Elem(),
		map[string]interface{}{
			"FIRST_SOURCE": PrimarySource_FIRST_SOURCE,
			"SECOND_SOURCE": PrimarySource_SECOND_SOURCE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.PublicNetworkConfigurationProps",
		reflect.TypeOf((*PublicNetworkConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RistOutputConfig",
		reflect.TypeOf((*RistOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RistOutputProtocolProps",
		reflect.TypeOf((*RistOutputProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RistProtocolProps",
		reflect.TypeOf((*RistProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInput",
		reflect.TypeOf((*RouterInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBitrate", GoMethod: "MetricBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnected", GoMethod: "MetricConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricContinuityCounterErrors", GoMethod: "MetricContinuityCounterErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailoverSwitches", GoMethod: "MetricFailoverSwitches"},
			_jsii_.MemberMethod{JsiiMethod: "metricLatency", GoMethod: "MetricLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricNotRecoveredPackets", GoMethod: "MetricNotRecoveredPackets"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalPackets", GoMethod: "MetricTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputArn", GoGetter: "RouterInputArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputId", GoGetter: "RouterInputId"},
			_jsii_.MemberProperty{JsiiProperty: "routerInputRef", GoGetter: "RouterInputRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitEncryptionSecret", GoGetter: "TransitEncryptionSecret"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RouterInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouterInput)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputAttributes",
		reflect.TypeOf((*RouterInputAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		reflect.TypeOf((*RouterInputConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RouterInputConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputEndpoint",
		reflect.TypeOf((*RouterInputEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputGrants",
		reflect.TypeOf((*RouterInputGrants)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "actions", GoMethod: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "restart", GoMethod: "Restart"},
			_jsii_.MemberMethod{JsiiMethod: "start", GoMethod: "Start"},
			_jsii_.MemberMethod{JsiiMethod: "stop", GoMethod: "Stop"},
		},
		func() interface{} {
			return &jsiiProxy_RouterInputGrants{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProps",
		reflect.TypeOf((*RouterInputProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocol",
		reflect.TypeOf((*RouterInputProtocol)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RouterInputProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocolOptions",
		reflect.TypeOf((*RouterInputProtocolOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RouterInputProtocolOptions{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputTier",
		reflect.TypeOf((*RouterInputTier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RouterInputTier{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkConfiguration",
		reflect.TypeOf((*RouterNetworkConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RouterNetworkConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkInterface",
		reflect.TypeOf((*RouterNetworkInterface)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceArn", GoGetter: "RouterNetworkInterfaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceId", GoGetter: "RouterNetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "routerNetworkInterfaceRef", GoGetter: "RouterNetworkInterfaceRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RouterNetworkInterface{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouterNetworkInterface)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkInterfaceAttributes",
		reflect.TypeOf((*RouterNetworkInterfaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkInterfaceProps",
		reflect.TypeOf((*RouterNetworkInterfaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutput",
		reflect.TypeOf((*RouterOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricArqRequests", GoMethod: "MetricArqRequests"},
			_jsii_.MemberMethod{JsiiMethod: "metricBitrate", GoMethod: "MetricBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnected", GoMethod: "MetricConnected"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalPackets", GoMethod: "MetricTotalPackets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputArn", GoGetter: "RouterOutputArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputId", GoGetter: "RouterOutputId"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputName", GoGetter: "RouterOutputName"},
			_jsii_.MemberProperty{JsiiProperty: "routerOutputRef", GoGetter: "RouterOutputRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RouterOutput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouterOutput)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputAttributes",
		reflect.TypeOf((*RouterOutputAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		reflect.TypeOf((*RouterOutputConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RouterOutputConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProps",
		reflect.TypeOf((*RouterOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocol",
		reflect.TypeOf((*RouterOutputProtocol)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RouterOutputProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocolOptions",
		reflect.TypeOf((*RouterOutputProtocolOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RouterOutputProtocolOptions{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputTier",
		reflect.TypeOf((*RouterOutputTier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RouterOutputTier{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterSource",
		reflect.TypeOf((*RouterSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterSrtEncryption",
		reflect.TypeOf((*RouterSrtEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RouterTransitConfig",
		reflect.TypeOf((*RouterTransitConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.RoutingScope",
		reflect.TypeOf((*RoutingScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_RoutingScope{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RtpFecOutputConfig",
		reflect.TypeOf((*RtpFecOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RtpOutputConfig",
		reflect.TypeOf((*RtpOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RtpOutputProtocolProps",
		reflect.TypeOf((*RtpOutputProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.RtpProtocolProps",
		reflect.TypeOf((*RtpProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.ScanMode",
		reflect.TypeOf((*ScanMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ScanMode{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceBase",
		reflect.TypeOf((*SourceBase)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceCdi",
		reflect.TypeOf((*SourceCdi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.SourceConfiguration",
		reflect.TypeOf((*SourceConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "flowSourceName", GoGetter: "FlowSourceName"},
		},
		func() interface{} {
			return &jsiiProxy_SourceConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceJpegXs",
		reflect.TypeOf((*SourceJpegXs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceMonitoringConfig",
		reflect.TypeOf((*SourceMonitoringConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceNdi",
		reflect.TypeOf((*SourceNdi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.SourcePriorityConfig",
		reflect.TypeOf((*SourcePriorityConfig)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SourcePriorityConfig{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		reflect.TypeOf((*SourceProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_SourceProtocol{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceRist",
		reflect.TypeOf((*SourceRist)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceRtp",
		reflect.TypeOf((*SourceRtp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceSrtCaller",
		reflect.TypeOf((*SourceSrtCaller)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceSrtListener",
		reflect.TypeOf((*SourceSrtListener)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SourceZixiPush",
		reflect.TypeOf((*SourceZixiPush)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtCallerOutputConfig",
		reflect.TypeOf((*SrtCallerOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtCallerOutputProtocolProps",
		reflect.TypeOf((*SrtCallerOutputProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtCallerProtocolProps",
		reflect.TypeOf((*SrtCallerProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtListenerOutputConfig",
		reflect.TypeOf((*SrtListenerOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtListenerOutputProtocolProps",
		reflect.TypeOf((*SrtListenerOutputProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtListenerProtocolProps",
		reflect.TypeOf((*SrtListenerProtocolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.SrtPasswordEncryption",
		reflect.TypeOf((*SrtPasswordEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.StandardConfigurationProps",
		reflect.TypeOf((*StandardConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.StandardOutputConfigurationProps",
		reflect.TypeOf((*StandardOutputConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediaconnect-alpha.State",
		reflect.TypeOf((*State)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": State_ENABLED,
			"DISABLED": State_DISABLED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.StaticKeyEncryption",
		reflect.TypeOf((*StaticKeyEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		reflect.TypeOf((*Tcs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Tcs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.TransitEncryption",
		reflect.TypeOf((*TransitEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.VideoRange",
		reflect.TypeOf((*VideoRange)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_VideoRange{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterface",
		reflect.TypeOf((*VpcInterface)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VpcInterface{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterfaceConfig",
		reflect.TypeOf((*VpcInterfaceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterfaceDefineProps",
		reflect.TypeOf((*VpcInterfaceDefineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterfaceFromNetworkInterfacesProps",
		reflect.TypeOf((*VpcInterfaceFromNetworkInterfacesProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.VpcNetworkConfigurationProps",
		reflect.TypeOf((*VpcNetworkConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.ZixiPullOutputConfig",
		reflect.TypeOf((*ZixiPullOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediaconnect-alpha.ZixiPushOutputConfig",
		reflect.TypeOf((*ZixiPushOutputConfig)(nil)).Elem(),
	)
}
