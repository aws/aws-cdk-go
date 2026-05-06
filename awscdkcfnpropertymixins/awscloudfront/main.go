package awscloudfront

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListMixinProps",
		reflect.TypeOf((*CfnAnycastIpListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnycastIpListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin.AnycastIpListProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_AnycastIpListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin.IpamCidrConfigProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_IpamCidrConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin.IpamCidrConfigResultProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_IpamCidrConfigResultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin.TagsProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_TagsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyMixinProps",
		reflect.TypeOf((*CfnCachePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin",
		reflect.TypeOf((*CfnCachePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCachePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin.CachePolicyConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_CachePolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin.CookiesConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin.HeadersConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin.ParametersInCacheKeyAndForwardedToOriginProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_ParametersInCacheKeyAndForwardedToOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCachePolicyPropsMixin.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCloudFrontOriginAccessIdentityMixinProps",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCloudFrontOriginAccessIdentityPropsMixin",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnCloudFrontOriginAccessIdentityPropsMixin.CloudFrontOriginAccessIdentityConfigProperty",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityPropsMixin_CloudFrontOriginAccessIdentityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionFunctionMixinProps",
		reflect.TypeOf((*CfnConnectionFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionFunctionPropsMixin",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionFunctionPropsMixin.ConnectionFunctionConfigProperty",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin_ConnectionFunctionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionFunctionPropsMixin.KeyValueStoreAssociationProperty",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin_KeyValueStoreAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionGroupMixinProps",
		reflect.TypeOf((*CfnConnectionGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnConnectionGroupPropsMixin",
		reflect.TypeOf((*CfnConnectionGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyMixinProps",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.ContinuousDeploymentPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_ContinuousDeploymentPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.SessionStickinessConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SessionStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.SingleHeaderConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.SingleHeaderPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleHeaderPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.SingleWeightConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleWeightConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.SingleWeightPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleWeightPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnContinuousDeploymentPolicyPropsMixin.TrafficConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_TrafficConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionMixinProps",
		reflect.TypeOf((*CfnDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.CacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.CacheTagConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheTagConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.ConnectionFunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ConnectionFunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.CookiesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CookiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.CustomErrorResponseProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CustomErrorResponseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.CustomOriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CustomOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.DefaultCacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DefaultCacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.DefinitionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.DistributionConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.ForwardedValuesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ForwardedValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.FunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_FunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.GeoRestrictionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_GeoRestrictionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.GrpcConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_GrpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.LambdaFunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LambdaFunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.LegacyCustomOriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LegacyCustomOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.LegacyS3OriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LegacyS3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginCustomHeaderProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginCustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginGroupFailoverCriteriaProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupFailoverCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginGroupMemberProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupMemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginGroupMembersProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupMembersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginGroupProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginGroupsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginMtlsConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginMtlsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.OriginShieldProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginShieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.ParameterDefinitionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ParameterDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.RestrictionsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_RestrictionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.S3OriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_S3OriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.StatusCodesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_StatusCodesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.StringSchemaProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_StringSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.TenantConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_TenantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.TrustStoreConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_TrustStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.ViewerCertificateProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ViewerCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.ViewerMtlsConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ViewerMtlsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionPropsMixin.VpcOriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_VpcOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantMixinProps",
		reflect.TypeOf((*CfnDistributionTenantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionTenantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.CustomizationsProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_CustomizationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.DomainResultProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_DomainResultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.GeoRestrictionCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_GeoRestrictionCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.ManagedCertificateRequestProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_ManagedCertificateRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.ParameterProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_ParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnDistributionTenantPropsMixin.WebAclCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_WebAclCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFunctionPropsMixin.FunctionConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFunctionPropsMixin.FunctionMetadataProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnFunctionPropsMixin.KeyValueStoreAssociationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KeyValueStoreAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyGroupMixinProps",
		reflect.TypeOf((*CfnKeyGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyGroupPropsMixin",
		reflect.TypeOf((*CfnKeyGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyGroupPropsMixin.KeyGroupConfigProperty",
		reflect.TypeOf((*CfnKeyGroupPropsMixin_KeyGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStoreMixinProps",
		reflect.TypeOf((*CfnKeyValueStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin",
		reflect.TypeOf((*CfnKeyValueStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyValueStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin.ImportSourceProperty",
		reflect.TypeOf((*CfnKeyValueStorePropsMixin_ImportSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnMonitoringSubscriptionMixinProps",
		reflect.TypeOf((*CfnMonitoringSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnMonitoringSubscriptionPropsMixin",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitoringSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnMonitoringSubscriptionPropsMixin.MonitoringSubscriptionProperty",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin_MonitoringSubscriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnMonitoringSubscriptionPropsMixin.RealtimeMetricsSubscriptionConfigProperty",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin_RealtimeMetricsSubscriptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginAccessControlMixinProps",
		reflect.TypeOf((*CfnOriginAccessControlMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginAccessControlPropsMixin",
		reflect.TypeOf((*CfnOriginAccessControlPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginAccessControlPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginAccessControlPropsMixin.OriginAccessControlConfigProperty",
		reflect.TypeOf((*CfnOriginAccessControlPropsMixin_OriginAccessControlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyMixinProps",
		reflect.TypeOf((*CfnOriginRequestPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyPropsMixin",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginRequestPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyPropsMixin.CookiesConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyPropsMixin.HeadersConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyPropsMixin.OriginRequestPolicyConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_OriginRequestPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnOriginRequestPolicyPropsMixin.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnPublicKeyMixinProps",
		reflect.TypeOf((*CfnPublicKeyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnPublicKeyPropsMixin",
		reflect.TypeOf((*CfnPublicKeyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPublicKeyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnPublicKeyPropsMixin.PublicKeyConfigProperty",
		reflect.TypeOf((*CfnPublicKeyPropsMixin_PublicKeyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnRealtimeLogConfigMixinProps",
		reflect.TypeOf((*CfnRealtimeLogConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnRealtimeLogConfigPropsMixin",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRealtimeLogConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnRealtimeLogConfigPropsMixin.EndPointProperty",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin_EndPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnRealtimeLogConfigPropsMixin.KinesisStreamConfigProperty",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin_KinesisStreamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyMixinProps",
		reflect.TypeOf((*CfnResponseHeadersPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResponseHeadersPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowMethodsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowMethodsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowOriginsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowOriginsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.AccessControlExposeHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlExposeHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.ContentSecurityPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ContentSecurityPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.ContentTypeOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ContentTypeOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.CorsConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CorsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.CustomHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.CustomHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CustomHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.FrameOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_FrameOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.ReferrerPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ReferrerPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.RemoveHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_RemoveHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.RemoveHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_RemoveHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.ResponseHeadersPolicyConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ResponseHeadersPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.SecurityHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_SecurityHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.ServerTimingHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ServerTimingHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.StrictTransportSecurityProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_StrictTransportSecurityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnResponseHeadersPolicyPropsMixin.XSSProtectionProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_XSSProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionMixinProps",
		reflect.TypeOf((*CfnStreamingDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionPropsMixin",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamingDistributionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionPropsMixin.S3OriginProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_S3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionPropsMixin.StreamingDistributionConfigProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_StreamingDistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnStreamingDistributionPropsMixin.TrustedSignersProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_TrustedSignersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnTrustStoreMixinProps",
		reflect.TypeOf((*CfnTrustStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnTrustStorePropsMixin",
		reflect.TypeOf((*CfnTrustStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnTrustStorePropsMixin.CaCertificatesBundleS3LocationProperty",
		reflect.TypeOf((*CfnTrustStorePropsMixin_CaCertificatesBundleS3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnTrustStorePropsMixin.CaCertificatesBundleSourceProperty",
		reflect.TypeOf((*CfnTrustStorePropsMixin_CaCertificatesBundleSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnVpcOriginMixinProps",
		reflect.TypeOf((*CfnVpcOriginMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnVpcOriginPropsMixin",
		reflect.TypeOf((*CfnVpcOriginPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcOriginPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnVpcOriginPropsMixin.VpcOriginEndpointConfigProperty",
		reflect.TypeOf((*CfnVpcOriginPropsMixin_VpcOriginEndpointConfigProperty)(nil)).Elem(),
	)
}
