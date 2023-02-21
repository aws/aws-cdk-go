package awsmediatailor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		reflect.TypeOf((*CfnPlaybackConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "adDecisionServerUrl", GoGetter: "AdDecisionServerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDashConfigurationManifestEndpointPrefix", GoGetter: "AttrDashConfigurationManifestEndpointPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "attrHlsConfigurationManifestEndpointPrefix", GoGetter: "AttrHlsConfigurationManifestEndpointPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "attrPlaybackConfigurationArn", GoGetter: "AttrPlaybackConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPlaybackEndpointPrefix", GoGetter: "AttrPlaybackEndpointPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "attrSessionInitializationEndpointPrefix", GoGetter: "AttrSessionInitializationEndpointPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "availSuppression", GoGetter: "AvailSuppression"},
			_jsii_.MemberProperty{JsiiProperty: "bumper", GoGetter: "Bumper"},
			_jsii_.MemberProperty{JsiiProperty: "cdnConfiguration", GoGetter: "CdnConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationAliases", GoGetter: "ConfigurationAliases"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dashConfiguration", GoGetter: "DashConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hlsConfiguration", GoGetter: "HlsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "livePreRollConfiguration", GoGetter: "LivePreRollConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "manifestProcessingRules", GoGetter: "ManifestProcessingRules"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "personalizationThresholdSeconds", GoGetter: "PersonalizationThresholdSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "slateAdUrl", GoGetter: "SlateAdUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transcodeProfileName", GoGetter: "TranscodeProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "videoContentSourceUrl", GoGetter: "VideoContentSourceUrl"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.AdMarkerPassthroughProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_AdMarkerPassthroughProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.AvailSuppressionProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_AvailSuppressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.BumperProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_BumperProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.CdnConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_CdnConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.DashConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_DashConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.HlsConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_HlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.LivePreRollConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_LivePreRollConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration.ManifestProcessingRulesProperty",
		reflect.TypeOf((*CfnPlaybackConfiguration_ManifestProcessingRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfigurationProps",
		reflect.TypeOf((*CfnPlaybackConfigurationProps)(nil)).Elem(),
	)
}
