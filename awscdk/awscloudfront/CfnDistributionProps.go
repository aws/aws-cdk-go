package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDistribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistributionProps := &CfnDistributionProps{
//   	DistributionConfig: &DistributionConfigProperty{
//   		DefaultCacheBehavior: &DefaultCacheBehaviorProperty{
//   			TargetOriginId: jsii.String("targetOriginId"),
//   			ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   			// the properties below are optional
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
//   				QueryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Cookies: &CookiesProperty{
//   					Forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					WhitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
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
//   			TrustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			TrustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		AnycastIpListId: jsii.String("anycastIpListId"),
//   		CacheBehaviors: []interface{}{
//   			&CacheBehaviorProperty{
//   				PathPattern: jsii.String("pathPattern"),
//   				TargetOriginId: jsii.String("targetOriginId"),
//   				ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   				// the properties below are optional
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
//   					QueryString: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					Cookies: &CookiesProperty{
//   						Forward: jsii.String("forward"),
//
//   						// the properties below are optional
//   						WhitelistedNames: []*string{
//   							jsii.String("whitelistedNames"),
//   						},
//   					},
//   					Headers: []*string{
//   						jsii.String("headers"),
//   					},
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
//   				RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   				ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   				SmoothStreaming: jsii.Boolean(false),
//   				TrustedKeyGroups: []*string{
//   					jsii.String("trustedKeyGroups"),
//   				},
//   				TrustedSigners: []*string{
//   					jsii.String("trustedSigners"),
//   				},
//   			},
//   		},
//   		CnamEs: []*string{
//   			jsii.String("cnamEs"),
//   		},
//   		Comment: jsii.String("comment"),
//   		ConnectionMode: jsii.String("connectionMode"),
//   		ContinuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   		CustomErrorResponses: []interface{}{
//   			&CustomErrorResponseProperty{
//   				ErrorCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				ErrorCachingMinTtl: jsii.Number(123),
//   				ResponseCode: jsii.Number(123),
//   				ResponsePagePath: jsii.String("responsePagePath"),
//   			},
//   		},
//   		CustomOrigin: &LegacyCustomOriginProperty{
//   			DnsName: jsii.String("dnsName"),
//   			OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   			OriginSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//
//   			// the properties below are optional
//   			HttpPort: jsii.Number(123),
//   			HttpsPort: jsii.Number(123),
//   		},
//   		DefaultRootObject: jsii.String("defaultRootObject"),
//   		HttpVersion: jsii.String("httpVersion"),
//   		Ipv6Enabled: jsii.Boolean(false),
//   		Logging: &LoggingProperty{
//   			Bucket: jsii.String("bucket"),
//   			IncludeCookies: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		OriginGroups: &OriginGroupsProperty{
//   			Quantity: jsii.Number(123),
//
//   			// the properties below are optional
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
//
//   					// the properties below are optional
//   					SelectionCriteria: jsii.String("selectionCriteria"),
//   				},
//   			},
//   		},
//   		Origins: []interface{}{
//   			&OriginProperty{
//   				DomainName: jsii.String("domainName"),
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				ConnectionAttempts: jsii.Number(123),
//   				ConnectionTimeout: jsii.Number(123),
//   				CustomOriginConfig: &CustomOriginConfigProperty{
//   					OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   					// the properties below are optional
//   					HttpPort: jsii.Number(123),
//   					HttpsPort: jsii.Number(123),
//   					IpAddressType: jsii.String("ipAddressType"),
//   					OriginKeepaliveTimeout: jsii.Number(123),
//   					OriginReadTimeout: jsii.Number(123),
//   					OriginSslProtocols: []*string{
//   						jsii.String("originSslProtocols"),
//   					},
//   				},
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
//   					VpcOriginId: jsii.String("vpcOriginId"),
//
//   					// the properties below are optional
//   					OriginKeepaliveTimeout: jsii.Number(123),
//   					OriginReadTimeout: jsii.Number(123),
//   				},
//   			},
//   		},
//   		PriceClass: jsii.String("priceClass"),
//   		Restrictions: &RestrictionsProperty{
//   			GeoRestriction: &GeoRestrictionProperty{
//   				RestrictionType: jsii.String("restrictionType"),
//
//   				// the properties below are optional
//   				Locations: []*string{
//   					jsii.String("locations"),
//   				},
//   			},
//   		},
//   		S3Origin: &LegacyS3OriginProperty{
//   			DnsName: jsii.String("dnsName"),
//
//   			// the properties below are optional
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		Staging: jsii.Boolean(false),
//   		TenantConfig: &TenantConfigProperty{
//   			ParameterDefinitions: []interface{}{
//   				&ParameterDefinitionProperty{
//   					Definition: &DefinitionProperty{
//   						StringSchema: &StringSchemaProperty{
//   							Required: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							Comment: jsii.String("comment"),
//   							DefaultValue: jsii.String("defaultValue"),
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
//   		WebAclId: jsii.String("webAclId"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html
//
type CfnDistributionProps struct {
	// The distribution's configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-distributionconfig
	//
	DistributionConfig interface{} `field:"required" json:"distributionConfig" yaml:"distributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

