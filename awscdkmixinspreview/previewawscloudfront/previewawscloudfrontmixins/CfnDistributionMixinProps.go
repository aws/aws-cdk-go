package previewawscloudfrontmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDistributionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionMixinProps := &CfnDistributionMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html
//
type CfnDistributionMixinProps struct {
	// The distribution's configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-distributionconfig
	//
	DistributionConfig interface{} `field:"optional" json:"distributionConfig" yaml:"distributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

