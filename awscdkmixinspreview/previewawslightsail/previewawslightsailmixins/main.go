package previewawslightsailmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmMixinProps",
		reflect.TypeOf((*CfnAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmPropsMixin",
		reflect.TypeOf((*CfnAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnBucketMixinProps",
		reflect.TypeOf((*CfnBucketMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnBucketPropsMixin",
		reflect.TypeOf((*CfnBucketPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnBucketPropsMixin.AccessRulesProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AccessRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnCertificateMixinProps",
		reflect.TypeOf((*CfnCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnCertificatePropsMixin",
		reflect.TypeOf((*CfnCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerMixinProps",
		reflect.TypeOf((*CfnContainerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin",
		reflect.TypeOf((*CfnContainerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.ContainerProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_ContainerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.ContainerServiceDeploymentProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_ContainerServiceDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.EcrImagePullerRoleProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_EcrImagePullerRoleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.HealthCheckConfigProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_HealthCheckConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.PortInfoProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PortInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.PrivateRegistryAccessProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PrivateRegistryAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.PublicDomainNameProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PublicDomainNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnContainerPropsMixin.PublicEndpointProperty",
		reflect.TypeOf((*CfnContainerPropsMixin_PublicEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabaseMixinProps",
		reflect.TypeOf((*CfnDatabaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin",
		reflect.TypeOf((*CfnDatabasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDatabasePropsMixin.RelationalDatabaseParameterProperty",
		reflect.TypeOf((*CfnDatabasePropsMixin_RelationalDatabaseParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskMixinProps",
		reflect.TypeOf((*CfnDiskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskPropsMixin",
		reflect.TypeOf((*CfnDiskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskPropsMixin.AddOnProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskPropsMixin.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDiskPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskSnapshotMixinProps",
		reflect.TypeOf((*CfnDiskSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskSnapshotPropsMixin",
		reflect.TypeOf((*CfnDiskSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiskSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDiskSnapshotPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDiskSnapshotPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionMixinProps",
		reflect.TypeOf((*CfnDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin",
		reflect.TypeOf((*CfnDistributionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.CacheBehaviorPerPathProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorPerPathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.CacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.CacheSettingsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.CookieObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CookieObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.HeaderObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_HeaderObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.InputOriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_InputOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDistributionPropsMixin.QueryStringObjectProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_QueryStringObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDomainPropsMixin.DomainEntryProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnDomainPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstanceMixinProps",
		reflect.TypeOf((*CfnInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin",
		reflect.TypeOf((*CfnInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.AddOnProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_AddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.AutoSnapshotAddOnProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_AutoSnapshotAddOnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.DiskProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_DiskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.HardwareProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_HardwareProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.LocationProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.MonthlyTransferProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_MonthlyTransferProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.NetworkingProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.PortProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_PortProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstancePropsMixin.StateProperty",
		reflect.TypeOf((*CfnInstancePropsMixin_StateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstanceSnapshotMixinProps",
		reflect.TypeOf((*CfnInstanceSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstanceSnapshotPropsMixin",
		reflect.TypeOf((*CfnInstanceSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnInstanceSnapshotPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnInstanceSnapshotPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerPropsMixin",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificateMixinProps",
		reflect.TypeOf((*CfnLoadBalancerTlsCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnLoadBalancerTlsCertificatePropsMixin",
		reflect.TypeOf((*CfnLoadBalancerTlsCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerTlsCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnStaticIpMixinProps",
		reflect.TypeOf((*CfnStaticIpMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnStaticIpPropsMixin",
		reflect.TypeOf((*CfnStaticIpPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStaticIpPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
