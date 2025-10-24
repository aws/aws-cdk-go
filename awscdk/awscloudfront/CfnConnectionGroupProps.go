package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnectionGroup`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // Create the simple Origin
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
//   	OriginAccessLevels: []AccessLevel{
//   		cloudfront.AccessLevel_READ,
//   		cloudfront.AccessLevel_LIST,
//   	},
//   })
//
//   // Create the Distribution construct
//   myMultiTenantDistribution := cloudfront.NewDistribution(this, jsii.String("cf-hosted-distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: s3Origin,
//   	},
//   	DefaultRootObject: jsii.String("index.html"),
//   })
//
//   // Access the underlying L1 CfnDistribution to configure SaaS Manager properties which are not yet available in the L2 Distribution construct
//   cfnDistribution := myMultiTenantDistribution.Node.defaultChild.(CfnDistribution)
//
//   defaultCacheBehavior := &DefaultCacheBehaviorProperty{
//   	TargetOriginId: myBucket.BucketArn,
//   	ViewerProtocolPolicy: jsii.String("allow-all"),
//   	Compress: jsii.Boolean(false),
//   	AllowedMethods: []*string{
//   		jsii.String("GET"),
//   		jsii.String("HEAD"),
//   	},
//   	CachePolicyId: cloudfront.CachePolicy_CACHING_OPTIMIZED().CachePolicyId,
//   }
//   // Create the updated distributionConfig
//   distributionConfig := &DistributionConfigProperty{
//   	DefaultCacheBehavior: defaultCacheBehavior,
//   	Enabled: jsii.Boolean(true),
//   	// the properties below are optional
//   	ConnectionMode: jsii.String("tenant-only"),
//   	Origins: []interface{}{
//   		&OriginProperty{
//   			Id: myBucket.*BucketArn,
//   			DomainName: myBucket.BucketDomainName,
//   			S3OriginConfig: &S3OriginConfigProperty{
//   			},
//   			OriginPath: jsii.String("/{{tenantName}}"),
//   		},
//   	},
//   	TenantConfig: &TenantConfigProperty{
//   		ParameterDefinitions: []interface{}{
//   			&ParameterDefinitionProperty{
//   				Definition: &DefinitionProperty{
//   					StringSchema: &StringSchemaProperty{
//   						Required: jsii.Boolean(false),
//   						// the properties below are optional
//   						Comment: jsii.String("tenantName"),
//   						DefaultValue: jsii.String("root"),
//   					},
//   				},
//   				Name: jsii.String("tenantName"),
//   			},
//   		},
//   	},
//   }
//
//   // Override the distribution configuration to enable multi-tenancy.
//   cfnDistribution.DistributionConfig = distributionConfig
//
//   // Create a connection group and a cname record in an existing hosted zone to validate domain ownership
//   connectionGroup := cloudfront.NewCfnConnectionGroup(this, jsii.String("cf-hosted-connection-group"), &CfnConnectionGroupProps{
//   	Enabled: jsii.Boolean(true),
//   	Ipv6Enabled: jsii.Boolean(true),
//   	Name: jsii.String("my-connection-group"),
//   })
//
//   // Import the existing hosted zone info, replacing with your hostedZoneId and zoneName
//   hostedZoneId := "YOUR_HOSTED_ZONE_ID"
//   zoneName := "my.domain.com"
//   hostedZone := route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("hosted-zone"), &HostedZoneAttributes{
//   	HostedZoneId: jsii.String(HostedZoneId),
//   	ZoneName: jsii.String(ZoneName),
//   })
//
//   record := route53.NewCnameRecord(this, jsii.String("cname-record"), &CnameRecordProps{
//   	DomainName: connectionGroup.AttrRoutingEndpoint,
//   	Zone: hostedZone,
//   	RecordName: jsii.String("cf-hosted-tenant.my.domain.com"),
//   })
//
//   // Create the cloudfront-hosted tenant, passing in the previously created connection group
//   cloudfrontHostedTenant := cloudfront.NewCfnDistributionTenant(this, jsii.String("cf-hosted-tenant"), &CfnDistributionTenantProps{
//   	DistributionId: myMultiTenantDistribution.DistributionId,
//   	Name: jsii.String("cf-hosted-tenant"),
//   	Domains: []*string{
//   		jsii.String("cf-hosted-tenant.my.domain.com"),
//   	},
//   	ConnectionGroupId: connectionGroup.AttrId,
//   	Enabled: jsii.Boolean(true),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		ValidationTokenHost: jsii.String("cloudfront"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html
//
type CfnConnectionGroupProps struct {
	// The name of the connection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Anycast static IP list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-anycastiplistid
	//
	AnycastIpListId *string `field:"optional" json:"anycastIpListId" yaml:"anycastIpListId"`
	// Whether the connection group is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// IPv6 is enabled for the connection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-ipv6enabled
	//
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectiongroup.html#cfn-cloudfront-connectiongroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

