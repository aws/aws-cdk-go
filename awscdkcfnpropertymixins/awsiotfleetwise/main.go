package awsiotfleetwise

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignMixinProps",
		reflect.TypeOf((*CfnCampaignMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin",
		reflect.TypeOf((*CfnCampaignPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaignPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.CollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.ConditionBasedCollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ConditionBasedCollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.ConditionBasedSignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ConditionBasedSignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.DataDestinationConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.DataPartitionProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.DataPartitionStorageOptionsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.DataPartitionUploadOptionsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionUploadOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.MqttTopicConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_MqttTopicConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.S3ConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_S3ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.SignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.SignalFetchInformationProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalFetchInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.SignalInformationProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.StorageMaximumSizeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_StorageMaximumSizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.StorageMinimumTimeToLiveProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_StorageMinimumTimeToLiveProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.TimeBasedCollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeBasedCollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.TimeBasedSignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeBasedSignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnCampaignPropsMixin.TimestreamConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimestreamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestMixinProps",
		reflect.TypeOf((*CfnDecoderManifestMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDecoderManifestPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.CanInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.CanNetworkInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanNetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.CanSignalDecoderProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanSignalDecoderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.CanSignalProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanSignalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.NetworkInterfacesItemsProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_NetworkInterfacesItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.ObdInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.ObdNetworkInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdNetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.ObdSignalDecoderProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdSignalDecoderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.ObdSignalProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdSignalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnDecoderManifestPropsMixin.SignalDecodersItemsProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_SignalDecodersItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestMixinProps",
		reflect.TypeOf((*CfnModelManifestMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnModelManifestPropsMixin",
		reflect.TypeOf((*CfnModelManifestPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelManifestPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogMixinProps",
		reflect.TypeOf((*CfnSignalCatalogMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSignalCatalogPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.ActuatorProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_ActuatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.AttributeProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_AttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.BranchProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_BranchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.NodeCountsProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_NodeCountsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.NodeProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_NodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnSignalCatalogPropsMixin.SensorProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_SensorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnStateTemplateMixinProps",
		reflect.TypeOf((*CfnStateTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnStateTemplatePropsMixin",
		reflect.TypeOf((*CfnStateTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehicleMixinProps",
		reflect.TypeOf((*CfnVehicleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehiclePropsMixin",
		reflect.TypeOf((*CfnVehiclePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVehiclePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehiclePropsMixin.PeriodicStateTemplateUpdateStrategyProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_PeriodicStateTemplateUpdateStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehiclePropsMixin.StateTemplateAssociationProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_StateTemplateAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehiclePropsMixin.StateTemplateUpdateStrategyProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_StateTemplateUpdateStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotfleetwise.CfnVehiclePropsMixin.TimePeriodProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_TimePeriodProperty)(nil)).Elem(),
	)
}
