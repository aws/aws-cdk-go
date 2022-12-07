package awslightsail

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
//   	bundleId: jsii.String("bundleId"),
//   	defaultCacheBehavior: &cacheBehaviorProperty{
//   		behavior: jsii.String("behavior"),
//   	},
//   	distributionName: jsii.String("distributionName"),
//   	origin: &inputOriginProperty{
//   		name: jsii.String("name"),
//   		protocolPolicy: jsii.String("protocolPolicy"),
//   		regionName: jsii.String("regionName"),
//   	},
//
//   	// the properties below are optional
//   	cacheBehaviors: []interface{}{
//   		&cacheBehaviorPerPathProperty{
//   			behavior: jsii.String("behavior"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	cacheBehaviorSettings: &cacheSettingsProperty{
//   		allowedHttpMethods: jsii.String("allowedHttpMethods"),
//   		cachedHttpMethods: jsii.String("cachedHttpMethods"),
//   		defaultTtl: jsii.Number(123),
//   		forwardedCookies: &cookieObjectProperty{
//   			cookiesAllowList: []*string{
//   				jsii.String("cookiesAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedHeaders: &headerObjectProperty{
//   			headersAllowList: []*string{
//   				jsii.String("headersAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedQueryStrings: &queryStringObjectProperty{
//   			option: jsii.Boolean(false),
//   			queryStringsAllowList: []*string{
//   				jsii.String("queryStringsAllowList"),
//   			},
//   		},
//   		maximumTtl: jsii.Number(123),
//   		minimumTtl: jsii.Number(123),
//   	},
//   	certificateName: jsii.String("certificateName"),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	isEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDistributionProps struct {
	// The ID of the bundle applied to the distribution.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the default cache behavior of the distribution.
	DefaultCacheBehavior interface{} `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name of the distribution.
	DistributionName *string `field:"required" json:"distributionName" yaml:"distributionName"`
	// An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.
	//
	// The distribution pulls, caches, and serves content from the origin.
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// An array of objects that describe the per-path cache behavior of the distribution.
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// An object that describes the cache behavior settings of the distribution.
	CacheBehaviorSettings interface{} `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The name of the SSL/TLS certificate attached to the distribution.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The IP address type of the distribution.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for IPv4 and IPv6.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A Boolean value indicating whether the distribution is enabled.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

