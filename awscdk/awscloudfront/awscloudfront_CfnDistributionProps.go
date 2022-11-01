package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDistribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistributionProps := &cfnDistributionProps{
//   	distributionConfig: &distributionConfigProperty{
//   		defaultCacheBehavior: &defaultCacheBehaviorProperty{
//   			targetOriginId: jsii.String("targetOriginId"),
//   			viewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   			// the properties below are optional
//   			allowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			cachedMethods: []*string{
//   				jsii.String("cachedMethods"),
//   			},
//   			cachePolicyId: jsii.String("cachePolicyId"),
//   			compress: jsii.Boolean(false),
//   			defaultTtl: jsii.Number(123),
//   			fieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   			forwardedValues: &forwardedValuesProperty{
//   				queryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				cookies: &cookiesProperty{
//   					forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					whitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				queryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			functionAssociations: []interface{}{
//   				&functionAssociationProperty{
//   					eventType: jsii.String("eventType"),
//   					functionArn: jsii.String("functionArn"),
//   				},
//   			},
//   			lambdaFunctionAssociations: []interface{}{
//   				&lambdaFunctionAssociationProperty{
//   					eventType: jsii.String("eventType"),
//   					includeBody: jsii.Boolean(false),
//   					lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   				},
//   			},
//   			maxTtl: jsii.Number(123),
//   			minTtl: jsii.Number(123),
//   			originRequestPolicyId: jsii.String("originRequestPolicyId"),
//   			realtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   			responseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   			smoothStreaming: jsii.Boolean(false),
//   			trustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			trustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		cacheBehaviors: []interface{}{
//   			&cacheBehaviorProperty{
//   				pathPattern: jsii.String("pathPattern"),
//   				targetOriginId: jsii.String("targetOriginId"),
//   				viewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   				// the properties below are optional
//   				allowedMethods: []*string{
//   					jsii.String("allowedMethods"),
//   				},
//   				cachedMethods: []*string{
//   					jsii.String("cachedMethods"),
//   				},
//   				cachePolicyId: jsii.String("cachePolicyId"),
//   				compress: jsii.Boolean(false),
//   				defaultTtl: jsii.Number(123),
//   				fieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   				forwardedValues: &forwardedValuesProperty{
//   					queryString: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					cookies: &cookiesProperty{
//   						forward: jsii.String("forward"),
//
//   						// the properties below are optional
//   						whitelistedNames: []*string{
//   							jsii.String("whitelistedNames"),
//   						},
//   					},
//   					headers: []*string{
//   						jsii.String("headers"),
//   					},
//   					queryStringCacheKeys: []*string{
//   						jsii.String("queryStringCacheKeys"),
//   					},
//   				},
//   				functionAssociations: []interface{}{
//   					&functionAssociationProperty{
//   						eventType: jsii.String("eventType"),
//   						functionArn: jsii.String("functionArn"),
//   					},
//   				},
//   				lambdaFunctionAssociations: []interface{}{
//   					&lambdaFunctionAssociationProperty{
//   						eventType: jsii.String("eventType"),
//   						includeBody: jsii.Boolean(false),
//   						lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   					},
//   				},
//   				maxTtl: jsii.Number(123),
//   				minTtl: jsii.Number(123),
//   				originRequestPolicyId: jsii.String("originRequestPolicyId"),
//   				realtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   				responseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   				smoothStreaming: jsii.Boolean(false),
//   				trustedKeyGroups: []*string{
//   					jsii.String("trustedKeyGroups"),
//   				},
//   				trustedSigners: []*string{
//   					jsii.String("trustedSigners"),
//   				},
//   			},
//   		},
//   		cnamEs: []*string{
//   			jsii.String("cnamEs"),
//   		},
//   		comment: jsii.String("comment"),
//   		customErrorResponses: []interface{}{
//   			&customErrorResponseProperty{
//   				errorCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				errorCachingMinTtl: jsii.Number(123),
//   				responseCode: jsii.Number(123),
//   				responsePagePath: jsii.String("responsePagePath"),
//   			},
//   		},
//   		customOrigin: &legacyCustomOriginProperty{
//   			dnsName: jsii.String("dnsName"),
//   			originProtocolPolicy: jsii.String("originProtocolPolicy"),
//   			originSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//
//   			// the properties below are optional
//   			httpPort: jsii.Number(123),
//   			httpsPort: jsii.Number(123),
//   		},
//   		defaultRootObject: jsii.String("defaultRootObject"),
//   		httpVersion: jsii.String("httpVersion"),
//   		ipv6Enabled: jsii.Boolean(false),
//   		logging: &loggingProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			includeCookies: jsii.Boolean(false),
//   			prefix: jsii.String("prefix"),
//   		},
//   		originGroups: &originGroupsProperty{
//   			quantity: jsii.Number(123),
//
//   			// the properties below are optional
//   			items: []interface{}{
//   				&originGroupProperty{
//   					failoverCriteria: &originGroupFailoverCriteriaProperty{
//   						statusCodes: &statusCodesProperty{
//   							items: []interface{}{
//   								jsii.Number(123),
//   							},
//   							quantity: jsii.Number(123),
//   						},
//   					},
//   					id: jsii.String("id"),
//   					members: &originGroupMembersProperty{
//   						items: []interface{}{
//   							&originGroupMemberProperty{
//   								originId: jsii.String("originId"),
//   							},
//   						},
//   						quantity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		origins: []interface{}{
//   			&originProperty{
//   				domainName: jsii.String("domainName"),
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				connectionAttempts: jsii.Number(123),
//   				connectionTimeout: jsii.Number(123),
//   				customOriginConfig: &customOriginConfigProperty{
//   					originProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   					// the properties below are optional
//   					httpPort: jsii.Number(123),
//   					httpsPort: jsii.Number(123),
//   					originKeepaliveTimeout: jsii.Number(123),
//   					originReadTimeout: jsii.Number(123),
//   					originSslProtocols: []*string{
//   						jsii.String("originSslProtocols"),
//   					},
//   				},
//   				originAccessControlId: jsii.String("originAccessControlId"),
//   				originCustomHeaders: []interface{}{
//   					&originCustomHeaderProperty{
//   						headerName: jsii.String("headerName"),
//   						headerValue: jsii.String("headerValue"),
//   					},
//   				},
//   				originPath: jsii.String("originPath"),
//   				originShield: &originShieldProperty{
//   					enabled: jsii.Boolean(false),
//   					originShieldRegion: jsii.String("originShieldRegion"),
//   				},
//   				s3OriginConfig: &s3OriginConfigProperty{
//   					originAccessIdentity: jsii.String("originAccessIdentity"),
//   				},
//   			},
//   		},
//   		priceClass: jsii.String("priceClass"),
//   		restrictions: &restrictionsProperty{
//   			geoRestriction: &geoRestrictionProperty{
//   				restrictionType: jsii.String("restrictionType"),
//
//   				// the properties below are optional
//   				locations: []*string{
//   					jsii.String("locations"),
//   				},
//   			},
//   		},
//   		s3Origin: &legacyS3OriginProperty{
//   			dnsName: jsii.String("dnsName"),
//
//   			// the properties below are optional
//   			originAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		viewerCertificate: &viewerCertificateProperty{
//   			acmCertificateArn: jsii.String("acmCertificateArn"),
//   			cloudFrontDefaultCertificate: jsii.Boolean(false),
//   			iamCertificateId: jsii.String("iamCertificateId"),
//   			minimumProtocolVersion: jsii.String("minimumProtocolVersion"),
//   			sslSupportMethod: jsii.String("sslSupportMethod"),
//   		},
//   		webAclId: jsii.String("webAclId"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDistributionProps struct {
	// The current configuration information for the distribution.
	//
	// Send a `GET` request to the `/ *CloudFront API version* /distribution ID/config` resource.
	DistributionConfig interface{} `field:"required" json:"distributionConfig" yaml:"distributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

