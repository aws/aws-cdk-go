package mixinsawslightsail

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
//   	BundleId: jsii.String("bundleId"),
//   	CacheBehaviors: []interface{}{
//   		&CacheBehaviorPerPathProperty{
//   			Behavior: jsii.String("behavior"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	CacheBehaviorSettings: &CacheSettingsProperty{
//   		AllowedHttpMethods: jsii.String("allowedHttpMethods"),
//   		CachedHttpMethods: jsii.String("cachedHttpMethods"),
//   		DefaultTtl: jsii.Number(123),
//   		ForwardedCookies: &CookieObjectProperty{
//   			CookiesAllowList: []*string{
//   				jsii.String("cookiesAllowList"),
//   			},
//   			Option: jsii.String("option"),
//   		},
//   		ForwardedHeaders: &HeaderObjectProperty{
//   			HeadersAllowList: []*string{
//   				jsii.String("headersAllowList"),
//   			},
//   			Option: jsii.String("option"),
//   		},
//   		ForwardedQueryStrings: &QueryStringObjectProperty{
//   			Option: jsii.Boolean(false),
//   			QueryStringsAllowList: []*string{
//   				jsii.String("queryStringsAllowList"),
//   			},
//   		},
//   		MaximumTtl: jsii.Number(123),
//   		MinimumTtl: jsii.Number(123),
//   	},
//   	CertificateName: jsii.String("certificateName"),
//   	DefaultCacheBehavior: &CacheBehaviorProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   	DistributionName: jsii.String("distributionName"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	IsEnabled: jsii.Boolean(false),
//   	Origin: &InputOriginProperty{
//   		Name: jsii.String("name"),
//   		ProtocolPolicy: jsii.String("protocolPolicy"),
//   		RegionName: jsii.String("regionName"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html
//
type CfnDistributionMixinProps struct {
	// The ID of the bundle applied to the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-bundleid
	//
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// An array of objects that describe the per-path cache behavior of the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-cachebehaviors
	//
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// An object that describes the cache behavior settings of the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-cachebehaviorsettings
	//
	CacheBehaviorSettings interface{} `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The name of the SSL/TLS certificate attached to the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-certificatename
	//
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// An object that describes the default cache behavior of the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-defaultcachebehavior
	//
	DefaultCacheBehavior interface{} `field:"optional" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name of the distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-distributionname
	//
	DistributionName *string `field:"optional" json:"distributionName" yaml:"distributionName"`
	// The IP address type of the distribution.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for IPv4 and IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A Boolean value indicating whether the distribution is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.
	//
	// The distribution pulls, caches, and serves content from the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-origin
	//
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

