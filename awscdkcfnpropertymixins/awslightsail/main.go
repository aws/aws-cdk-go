package awslightsail

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnAlarmMixinProps",
		reflect.TypeOf((*CfnAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnAlarmPropsMixin",
		reflect.TypeOf((*CfnAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnBucketMixinProps",
		reflect.TypeOf((*CfnBucketMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnBucketPropsMixin",
		reflect.TypeOf((*CfnBucketPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnBucketPropsMixin.AccessRulesProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AccessRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnCertificateMixinProps",
		reflect.TypeOf((*CfnCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnCertificatePropsMixin",
		reflect.TypeOf((*CfnCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerMixinProps",
		reflect.TypeOf((*CfnContainerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin",
		reflect.TypeOf((*CfnContainerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.ContainerProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_ContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.ContainerServiceDeploymentProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_ContainerServiceDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.EcrImagePullerRoleProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_EcrImagePullerRoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.HealthCheckConfigProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_HealthCheckConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.PortInfoProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PortInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.PrivateRegistryAccessProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PrivateRegistryAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.PublicDomainNameProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PublicDomainNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnContainerPropsMixin.PublicEndpointProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PublicEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabasePropsMixin",
		reflect.TypeOf((*CfnDatabasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabasePropsMixin.RelationalDatabaseParameterProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_RelationalDatabaseParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabaseSnapshotMixinProps",
		reflect.TypeOf((*CfnDatabaseSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabaseSnapshotPropsMixin",
		reflect.TypeOf((*CfnDatabaseSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabaseSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDatabaseSnapshotPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDatabaseSnapshotPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskMixinProps",
		reflect.TypeOf((*CfnDiskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskPropsMixin",
		reflect.TypeOf((*CfnDiskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskPropsMixin.AddOnProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskPropsMixin.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskSnapshotMixinProps",
		reflect.TypeOf((*CfnDiskSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskSnapshotPropsMixin",
		reflect.TypeOf((*CfnDiskSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiskSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDiskSnapshotPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDiskSnapshotPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionMixinProps",
		reflect.TypeOf((*CfnDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin",
		reflect.TypeOf((*CfnDistributionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.CacheBehaviorPerPathProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorPerPathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.CacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.CacheSettingsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.CookieObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CookieObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.HeaderObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_HeaderObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.InputOriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_InputOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDistributionPropsMixin.QueryStringObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_QueryStringObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDomainPropsMixin.DomainEntryProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnDomainPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceMixinProps",
		reflect.TypeOf((*CfnInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin",
		reflect.TypeOf((*CfnInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.AddOnProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.DiskProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_DiskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.HardwareProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_HardwareProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.MonthlyTransferProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_MonthlyTransferProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.NetworkingProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.PortProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_PortProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstancePropsMixin.StateProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_StateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotMixinProps",
		reflect.TypeOf((*CfnInstanceSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin",
		reflect.TypeOf((*CfnInstanceSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnInstanceSnapshotPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnLoadBalancerPropsMixin",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnLoadBalancerTlsCertificateMixinProps",
		reflect.TypeOf((*CfnLoadBalancerTlsCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnLoadBalancerTlsCertificatePropsMixin",
		reflect.TypeOf((*CfnLoadBalancerTlsCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnStaticIpMixinProps",
		reflect.TypeOf((*CfnStaticIpMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnStaticIpPropsMixin",
		reflect.TypeOf((*CfnStaticIpPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStaticIpPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
