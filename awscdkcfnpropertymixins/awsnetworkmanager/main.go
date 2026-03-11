package awsnetworkmanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectAttachmentMixinProps",
		reflect.TypeOf((*CfnConnectAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectAttachmentPropsMixin",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectAttachmentPropsMixin.ConnectAttachmentOptionsProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ConnectAttachmentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnConnectAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectPeerMixinProps",
		reflect.TypeOf((*CfnConnectPeerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectPeerPropsMixin",
		reflect.TypeOf((*CfnConnectPeerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectPeerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectPeerPropsMixin.BgpOptionsProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_BgpOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectPeerPropsMixin.ConnectPeerBgpConfigurationProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_ConnectPeerBgpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnConnectPeerPropsMixin.ConnectPeerConfigurationProperty",
		reflect.TypeOf((*CfnConnectPeerPropsMixin_ConnectPeerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkMixinProps",
		reflect.TypeOf((*CfnCoreNetworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPrefixListAssociationMixinProps",
		reflect.TypeOf((*CfnCoreNetworkPrefixListAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPrefixListAssociationPropsMixin",
		reflect.TypeOf((*CfnCoreNetworkPrefixListAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCoreNetworkPrefixListAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPropsMixin",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCoreNetworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPropsMixin.CoreNetworkEdgeProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkEdgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPropsMixin.CoreNetworkNetworkFunctionGroupProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkNetworkFunctionGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPropsMixin.CoreNetworkSegmentProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_CoreNetworkSegmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCoreNetworkPropsMixin.SegmentsProperty",
		reflect.TypeOf((*CfnCoreNetworkPropsMixin_SegmentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCustomerGatewayAssociationMixinProps",
		reflect.TypeOf((*CfnCustomerGatewayAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnCustomerGatewayAssociationPropsMixin",
		reflect.TypeOf((*CfnCustomerGatewayAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomerGatewayAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDeviceMixinProps",
		reflect.TypeOf((*CfnDeviceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin",
		reflect.TypeOf((*CfnDevicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDevicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin.AWSLocationProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_AWSLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDevicePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDevicePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDirectConnectGatewayAttachmentMixinProps",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDirectConnectGatewayAttachmentPropsMixin",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDirectConnectGatewayAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnDirectConnectGatewayAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnDirectConnectGatewayAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnGlobalNetworkMixinProps",
		reflect.TypeOf((*CfnGlobalNetworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnGlobalNetworkPropsMixin",
		reflect.TypeOf((*CfnGlobalNetworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalNetworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnLinkAssociationMixinProps",
		reflect.TypeOf((*CfnLinkAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnLinkAssociationPropsMixin",
		reflect.TypeOf((*CfnLinkAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnLinkMixinProps",
		reflect.TypeOf((*CfnLinkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnLinkPropsMixin",
		reflect.TypeOf((*CfnLinkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLinkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnLinkPropsMixin.BandwidthProperty",
		reflect.TypeOf((*CfnLinkPropsMixin_BandwidthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSiteMixinProps",
		reflect.TypeOf((*CfnSiteMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSitePropsMixin",
		reflect.TypeOf((*CfnSitePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSitePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSitePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnSitePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSiteToSiteVpnAttachmentMixinProps",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSiteToSiteVpnAttachmentPropsMixin",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSiteToSiteVpnAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSiteToSiteVpnAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnSiteToSiteVpnAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnSiteToSiteVpnAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayPeeringMixinProps",
		reflect.TypeOf((*CfnTransitGatewayPeeringMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayPeeringPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayPeeringPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayPeeringPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRegistrationMixinProps",
		reflect.TypeOf((*CfnTransitGatewayRegistrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRegistrationPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayRegistrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRegistrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRouteTableAttachmentMixinProps",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRouteTableAttachmentPropsMixin",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRouteTableAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRouteTableAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnTransitGatewayRouteTableAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnVpcAttachmentMixinProps",
		reflect.TypeOf((*CfnVpcAttachmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnVpcAttachmentPropsMixin",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcAttachmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnVpcAttachmentPropsMixin.ProposedNetworkFunctionGroupChangeProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_ProposedNetworkFunctionGroupChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnVpcAttachmentPropsMixin.ProposedSegmentChangeProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_ProposedSegmentChangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkmanager.CfnVpcAttachmentPropsMixin.VpcOptionsProperty",
		reflect.TypeOf((*CfnVpcAttachmentPropsMixin_VpcOptionsProperty)(nil)).Elem(),
	)
}
