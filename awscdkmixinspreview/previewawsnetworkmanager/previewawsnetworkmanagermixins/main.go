package previewawsnetworkmanagermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectAttachmentMixinProps",
		reflect.TypeOf((*CfnConnectAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectAttachmentPropsMixin",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectAttachmentPropsMixin.ConnectAttachmentOptionsProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ConnectAttachmentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectPeerMixinProps",
		reflect.TypeOf((*CfnConnectPeerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectPeerPropsMixin",
		reflect.TypeOf((*CfnConnectPeerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectPeerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectPeerPropsMixin.BgpOptionsProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_BgpOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectPeerPropsMixin.ConnectPeerBgpConfigurationProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_ConnectPeerBgpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnConnectPeerPropsMixin.ConnectPeerConfigurationProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_ConnectPeerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkMixinProps",
		reflect.TypeOf((*CfnCoreNetworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkPropsMixin",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCoreNetworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkPropsMixin.CoreNetworkEdgeProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkEdgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkPropsMixin.CoreNetworkNetworkFunctionGroupProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkNetworkFunctionGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkPropsMixin.CoreNetworkSegmentProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkSegmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCoreNetworkPropsMixin.SegmentsProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_SegmentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationMixinProps",
		reflect.TypeOf((*CfnCustomerGatewayAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationPropsMixin",
		reflect.TypeOf((*CfnCustomerGatewayAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomerGatewayAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDeviceMixinProps",
		reflect.TypeOf((*CfnDeviceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDevicePropsMixin",
		reflect.TypeOf((*CfnDevicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDevicePropsMixin.AWSLocationProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_AWSLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDevicePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentMixinProps",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkMixinProps",
		reflect.TypeOf((*CfnGlobalNetworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkPropsMixin",
		reflect.TypeOf((*CfnGlobalNetworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalNetworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnLinkAssociationMixinProps",
		reflect.TypeOf((*CfnLinkAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnLinkAssociationPropsMixin",
		reflect.TypeOf((*CfnLinkAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnLinkMixinProps",
		reflect.TypeOf((*CfnLinkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnLinkPropsMixin",
		reflect.TypeOf((*CfnLinkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnLinkPropsMixin.BandwidthProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_BandwidthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSiteMixinProps",
		reflect.TypeOf((*CfnSiteMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSitePropsMixin",
		reflect.TypeOf((*CfnSitePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSitePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSitePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnSitePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSiteToSiteVpnAttachmentMixinProps",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSiteToSiteVpnAttachmentPropsMixin",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSiteToSiteVpnAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSiteToSiteVpnAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnSiteToSiteVpnAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayPeeringMixinProps",
		reflect.TypeOf((*CfnTransitGatewayPeeringMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayPeeringPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayPeeringPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayPeeringPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRegistrationMixinProps",
		reflect.TypeOf((*CfnTransitGatewayRegistrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRegistrationPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayRegistrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRegistrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRouteTableAttachmentMixinProps",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRouteTableAttachmentPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRouteTableAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRouteTableAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnTransitGatewayRouteTableAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnVpcAttachmentMixinProps",
		reflect.TypeOf((*CfnVpcAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnVpcAttachmentPropsMixin",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnVpcAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnVpcAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnVpcAttachmentPropsMixin.VpcOptionsProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_VpcOptionsProperty)(nil)).Elem(),
	)
}
