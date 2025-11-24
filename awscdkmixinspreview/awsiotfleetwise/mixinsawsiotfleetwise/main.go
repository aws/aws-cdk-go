package mixinsawsiotfleetwise

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignMixinProps",
		reflect.TypeOf((*CfnCampaignMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin",
		reflect.TypeOf((*CfnCampaignPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaignPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.CollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.ConditionBasedCollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ConditionBasedCollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.ConditionBasedSignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ConditionBasedSignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.DataDestinationConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.DataPartitionProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.DataPartitionStorageOptionsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.DataPartitionUploadOptionsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DataPartitionUploadOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.MqttTopicConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_MqttTopicConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.S3ConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_S3ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.SignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.SignalFetchInformationProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalFetchInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.SignalInformationProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SignalInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.StorageMaximumSizeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_StorageMaximumSizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.StorageMinimumTimeToLiveProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_StorageMinimumTimeToLiveProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.TimeBasedCollectionSchemeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeBasedCollectionSchemeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.TimeBasedSignalFetchConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeBasedSignalFetchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignPropsMixin.TimestreamConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimestreamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestMixinProps",
		reflect.TypeOf((*CfnDecoderManifestMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDecoderManifestPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.CanInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.CanNetworkInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanNetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.CanSignalDecoderProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanSignalDecoderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.CanSignalProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_CanSignalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.NetworkInterfacesItemsProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_NetworkInterfacesItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.ObdInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.ObdNetworkInterfaceProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdNetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.ObdSignalDecoderProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdSignalDecoderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.ObdSignalProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_ObdSignalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin.SignalDecodersItemsProperty",
		reflect.TypeOf((*CfnDecoderManifestPropsMixin_SignalDecodersItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnModelManifestMixinProps",
		reflect.TypeOf((*CfnModelManifestMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnModelManifestPropsMixin",
		reflect.TypeOf((*CfnModelManifestPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModelManifestPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogMixinProps",
		reflect.TypeOf((*CfnSignalCatalogMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSignalCatalogPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.ActuatorProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_ActuatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.AttributeProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_AttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.BranchProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_BranchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.NodeCountsProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_NodeCountsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.NodeProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_NodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnSignalCatalogPropsMixin.SensorProperty",
		reflect.TypeOf((*CfnSignalCatalogPropsMixin_SensorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnStateTemplateMixinProps",
		reflect.TypeOf((*CfnStateTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnStateTemplatePropsMixin",
		reflect.TypeOf((*CfnStateTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleMixinProps",
		reflect.TypeOf((*CfnVehicleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin",
		reflect.TypeOf((*CfnVehiclePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVehiclePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin.PeriodicStateTemplateUpdateStrategyProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_PeriodicStateTemplateUpdateStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin.StateTemplateAssociationProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_StateTemplateAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin.StateTemplateUpdateStrategyProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_StateTemplateUpdateStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehiclePropsMixin.TimePeriodProperty",
		reflect.TypeOf((*CfnVehiclePropsMixin_TimePeriodProperty)(nil)).Elem(),
	)
}
