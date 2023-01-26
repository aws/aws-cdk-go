package awsgroundstation

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_groundstation.CfnConfig",
		reflect.TypeOf((*CfnConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configData", GoGetter: "ConfigData"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.AntennaDownlinkConfigProperty",
		reflect.TypeOf((*CfnConfig_AntennaDownlinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.AntennaDownlinkDemodDecodeConfigProperty",
		reflect.TypeOf((*CfnConfig_AntennaDownlinkDemodDecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.AntennaUplinkConfigProperty",
		reflect.TypeOf((*CfnConfig_AntennaUplinkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.ConfigDataProperty",
		reflect.TypeOf((*CfnConfig_ConfigDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.DataflowEndpointConfigProperty",
		reflect.TypeOf((*CfnConfig_DataflowEndpointConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.DecodeConfigProperty",
		reflect.TypeOf((*CfnConfig_DecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.DemodulationConfigProperty",
		reflect.TypeOf((*CfnConfig_DemodulationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.EirpProperty",
		reflect.TypeOf((*CfnConfig_EirpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.FrequencyBandwidthProperty",
		reflect.TypeOf((*CfnConfig_FrequencyBandwidthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.FrequencyProperty",
		reflect.TypeOf((*CfnConfig_FrequencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.S3RecordingConfigProperty",
		reflect.TypeOf((*CfnConfig_S3RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.SpectrumConfigProperty",
		reflect.TypeOf((*CfnConfig_SpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.TrackingConfigProperty",
		reflect.TypeOf((*CfnConfig_TrackingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.UplinkEchoConfigProperty",
		reflect.TypeOf((*CfnConfig_UplinkEchoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfig.UplinkSpectrumConfigProperty",
		reflect.TypeOf((*CfnConfig_UplinkSpectrumConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnConfigProps",
		reflect.TypeOf((*CfnConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup",
		reflect.TypeOf((*CfnDataflowEndpointGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointDetails", GoGetter: "EndpointDetails"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataflowEndpointGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup.DataflowEndpointProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroup_DataflowEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup.EndpointDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroup_EndpointDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup.SecurityDetailsProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroup_SecurityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroup.SocketAddressProperty",
		reflect.TypeOf((*CfnDataflowEndpointGroup_SocketAddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupProps",
		reflect.TypeOf((*CfnDataflowEndpointGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile",
		reflect.TypeOf((*CfnMissionProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrRegion", GoGetter: "AttrRegion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contactPostPassDurationSeconds", GoGetter: "ContactPostPassDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "contactPrePassDurationSeconds", GoGetter: "ContactPrePassDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataflowEdges", GoGetter: "DataflowEdges"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "minimumViableContactDurationSeconds", GoGetter: "MinimumViableContactDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trackingConfigArn", GoGetter: "TrackingConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMissionProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfile.DataflowEdgeProperty",
		reflect.TypeOf((*CfnMissionProfile_DataflowEdgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_groundstation.CfnMissionProfileProps",
		reflect.TypeOf((*CfnMissionProfileProps)(nil)).Elem(),
	)
}
