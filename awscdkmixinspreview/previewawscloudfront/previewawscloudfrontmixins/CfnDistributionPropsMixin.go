package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionPropsMixin := awscdkmixinspreview.Mixins.NewCfnDistributionPropsMixin(&CfnDistributionMixinProps{
//   	DistributionConfig: &DistributionConfigProperty{
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		AnycastIpListId: jsii.String("anycastIpListId"),
//   		CacheBehaviors: []interface{}{
//   			&CacheBehaviorProperty{
//   				AllowedMethods: []*string{
//   					jsii.String("allowedMethods"),
//   				},
//   				CachedMethods: []*string{
//   					jsii.String("cachedMethods"),
//   				},
//   				CachePolicyId: jsii.String("cachePolicyId"),
//   				Compress: jsii.Boolean(false),
//   				DefaultTtl: jsii.Number(123),
//   				FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   				ForwardedValues: &ForwardedValuesProperty{
//   					Cookies: &CookiesProperty{
//   						Forward: jsii.String("forward"),
//   						WhitelistedNames: []*string{
//   							jsii.String("whitelistedNames"),
//   						},
//   					},
//   					Headers: []*string{
//   						jsii.String("headers"),
//   					},
//   					QueryString: jsii.Boolean(false),
//   					QueryStringCacheKeys: []*string{
//   						jsii.String("queryStringCacheKeys"),
//   					},
//   				},
//   				FunctionAssociations: []interface{}{
//   					&FunctionAssociationProperty{
//   						EventType: jsii.String("eventType"),
//   						FunctionArn: jsii.String("functionArn"),
//   					},
//   				},
//   				GrpcConfig: &GrpcConfigProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   				LambdaFunctionAssociations: []interface{}{
//   					&LambdaFunctionAssociationProperty{
//   						EventType: jsii.String("eventType"),
//   						IncludeBody: jsii.Boolean(false),
//   						LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   					},
//   				},
//   				MaxTtl: jsii.Number(123),
//   				MinTtl: jsii.Number(123),
//   				OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   				PathPattern: jsii.String("pathPattern"),
//   				RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   				ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   				SmoothStreaming: jsii.Boolean(false),
//   				TargetOriginId: jsii.String("targetOriginId"),
//   				TrustedKeyGroups: []*string{
//   					jsii.String("trustedKeyGroups"),
//   				},
//   				TrustedSigners: []*string{
//   					jsii.String("trustedSigners"),
//   				},
//   				ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//   			},
//   		},
//   		CnamEs: []*string{
//   			jsii.String("cnamEs"),
//   		},
//   		Comment: jsii.String("comment"),
//   		ConnectionFunctionAssociation: &ConnectionFunctionAssociationProperty{
//   			Id: jsii.String("id"),
//   		},
//   		ConnectionMode: jsii.String("connectionMode"),
//   		ContinuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   		CustomErrorResponses: []interface{}{
//   			&CustomErrorResponseProperty{
//   				ErrorCachingMinTtl: jsii.Number(123),
//   				ErrorCode: jsii.Number(123),
//   				ResponseCode: jsii.Number(123),
//   				ResponsePagePath: jsii.String("responsePagePath"),
//   			},
//   		},
//   		CustomOrigin: &LegacyCustomOriginProperty{
//   			DnsName: jsii.String("dnsName"),
//   			HttpPort: jsii.Number(123),
//   			HttpsPort: jsii.Number(123),
//   			OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   			OriginSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//   		},
//   		DefaultCacheBehavior: &DefaultCacheBehaviorProperty{
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			CachedMethods: []*string{
//   				jsii.String("cachedMethods"),
//   			},
//   			CachePolicyId: jsii.String("cachePolicyId"),
//   			Compress: jsii.Boolean(false),
//   			DefaultTtl: jsii.Number(123),
//   			FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   			ForwardedValues: &ForwardedValuesProperty{
//   				Cookies: &CookiesProperty{
//   					Forward: jsii.String("forward"),
//   					WhitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				QueryString: jsii.Boolean(false),
//   				QueryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			FunctionAssociations: []interface{}{
//   				&FunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   			},
//   			GrpcConfig: &GrpcConfigProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			LambdaFunctionAssociations: []interface{}{
//   				&LambdaFunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					IncludeBody: jsii.Boolean(false),
//   					LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   				},
//   			},
//   			MaxTtl: jsii.Number(123),
//   			MinTtl: jsii.Number(123),
//   			OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   			RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   			ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   			SmoothStreaming: jsii.Boolean(false),
//   			TargetOriginId: jsii.String("targetOriginId"),
//   			TrustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			TrustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   			ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//   		},
//   		DefaultRootObject: jsii.String("defaultRootObject"),
//   		Enabled: jsii.Boolean(false),
//   		HttpVersion: jsii.String("httpVersion"),
//   		Ipv6Enabled: jsii.Boolean(false),
//   		Logging: &LoggingProperty{
//   			Bucket: jsii.String("bucket"),
//   			IncludeCookies: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		OriginGroups: &OriginGroupsProperty{
//   			Items: []interface{}{
//   				&OriginGroupProperty{
//   					FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   						StatusCodes: &StatusCodesProperty{
//   							Items: []interface{}{
//   								jsii.Number(123),
//   							},
//   							Quantity: jsii.Number(123),
//   						},
//   					},
//   					Id: jsii.String("id"),
//   					Members: &OriginGroupMembersProperty{
//   						Items: []interface{}{
//   							&OriginGroupMemberProperty{
//   								OriginId: jsii.String("originId"),
//   							},
//   						},
//   						Quantity: jsii.Number(123),
//   					},
//   					SelectionCriteria: jsii.String("selectionCriteria"),
//   				},
//   			},
//   			Quantity: jsii.Number(123),
//   		},
//   		Origins: []interface{}{
//   			&OriginProperty{
//   				ConnectionAttempts: jsii.Number(123),
//   				ConnectionTimeout: jsii.Number(123),
//   				CustomOriginConfig: &CustomOriginConfigProperty{
//   					HttpPort: jsii.Number(123),
//   					HttpsPort: jsii.Number(123),
//   					IpAddressType: jsii.String("ipAddressType"),
//   					OriginKeepaliveTimeout: jsii.Number(123),
//   					OriginMtlsConfig: &OriginMtlsConfigProperty{
//   						ClientCertificateArn: jsii.String("clientCertificateArn"),
//   					},
//   					OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   					OriginReadTimeout: jsii.Number(123),
//   					OriginSslProtocols: []*string{
//   						jsii.String("originSslProtocols"),
//   					},
//   				},
//   				DomainName: jsii.String("domainName"),
//   				Id: jsii.String("id"),
//   				OriginAccessControlId: jsii.String("originAccessControlId"),
//   				OriginCustomHeaders: []interface{}{
//   					&OriginCustomHeaderProperty{
//   						HeaderName: jsii.String("headerName"),
//   						HeaderValue: jsii.String("headerValue"),
//   					},
//   				},
//   				OriginPath: jsii.String("originPath"),
//   				OriginShield: &OriginShieldProperty{
//   					Enabled: jsii.Boolean(false),
//   					OriginShieldRegion: jsii.String("originShieldRegion"),
//   				},
//   				ResponseCompletionTimeout: jsii.Number(123),
//   				S3OriginConfig: &S3OriginConfigProperty{
//   					OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   					OriginReadTimeout: jsii.Number(123),
//   				},
//   				VpcOriginConfig: &VpcOriginConfigProperty{
//   					OriginKeepaliveTimeout: jsii.Number(123),
//   					OriginReadTimeout: jsii.Number(123),
//   					OwnerAccountId: jsii.String("ownerAccountId"),
//   					VpcOriginId: jsii.String("vpcOriginId"),
//   				},
//   			},
//   		},
//   		PriceClass: jsii.String("priceClass"),
//   		Restrictions: &RestrictionsProperty{
//   			GeoRestriction: &GeoRestrictionProperty{
//   				Locations: []*string{
//   					jsii.String("locations"),
//   				},
//   				RestrictionType: jsii.String("restrictionType"),
//   			},
//   		},
//   		S3Origin: &LegacyS3OriginProperty{
//   			DnsName: jsii.String("dnsName"),
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		Staging: jsii.Boolean(false),
//   		TenantConfig: &TenantConfigProperty{
//   			ParameterDefinitions: []interface{}{
//   				&ParameterDefinitionProperty{
//   					Definition: &DefinitionProperty{
//   						StringSchema: &StringSchemaProperty{
//   							Comment: jsii.String("comment"),
//   							DefaultValue: jsii.String("defaultValue"),
//   							Required: jsii.Boolean(false),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		ViewerCertificate: &ViewerCertificateProperty{
//   			AcmCertificateArn: jsii.String("acmCertificateArn"),
//   			CloudFrontDefaultCertificate: jsii.Boolean(false),
//   			IamCertificateId: jsii.String("iamCertificateId"),
//   			MinimumProtocolVersion: jsii.String("minimumProtocolVersion"),
//   			SslSupportMethod: jsii.String("sslSupportMethod"),
//   		},
//   		ViewerMtlsConfig: &ViewerMtlsConfigProperty{
//   			Mode: jsii.String("mode"),
//   			TrustStoreConfig: &TrustStoreConfigProperty{
//   				AdvertiseTrustStoreCaNames: jsii.Boolean(false),
//   				IgnoreCertificateExpiry: jsii.Boolean(false),
//   				TrustStoreId: jsii.String("trustStoreId"),
//   			},
//   		},
//   		WebAclId: jsii.String("webAclId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html
//
type CfnDistributionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDistributionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDistributionPropsMixin
type jsiiProxy_CfnDistributionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDistributionPropsMixin) Props() *CfnDistributionMixinProps {
	var returns *CfnDistributionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::Distribution`.
func NewCfnDistributionPropsMixin(props *CfnDistributionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDistributionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDistributionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDistributionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::Distribution`.
func NewCfnDistributionPropsMixin_Override(c CfnDistributionPropsMixin, props *CfnDistributionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDistributionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDistributionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

