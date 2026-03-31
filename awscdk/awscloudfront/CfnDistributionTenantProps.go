package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDistributionTenant`.
//
// Example:
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
//   	TargetOriginId: myBucket.bucketArn,
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
//   			Id: myBucket.bucketArn,
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
//   // Create a connection group so we have access to the RoutingEndpoint associated with the tenant we are about to create
//   connectionGroup := cloudfront.NewCfnConnectionGroup(this, jsii.String("self-hosted-connection-group"), &CfnConnectionGroupProps{
//   	Enabled: jsii.Boolean(true),
//   	Ipv6Enabled: jsii.Boolean(true),
//   	Name: jsii.String("self-hosted-connection-group"),
//   })
//
//   // Export the RoutingEndpoint, skip this step if you'd prefer to fetch it from the CloudFront console or via Cloudfront.ListConnectionGroups API
//   // Export the RoutingEndpoint, skip this step if you'd prefer to fetch it from the CloudFront console or via Cloudfront.ListConnectionGroups API
//   awscdk.NewCfnOutput(this, jsii.String("RoutingEndpoint"), &CfnOutputProps{
//   	Value: connectionGroup.attrRoutingEndpoint,
//   	Description: jsii.String("CloudFront Routing Endpoint to be added to my hosted zone CNAME records"),
//   })
//
//   // Create a distribution tenant with a self-hosted domain.
//   selfHostedTenant := cloudfront.NewCfnDistributionTenant(this, jsii.String("self-hosted-tenant"), &CfnDistributionTenantProps{
//   	DistributionId: myMultiTenantDistribution.DistributionId,
//   	ConnectionGroupId: connectionGroup.attrId,
//   	Name: jsii.String("self-hosted-tenant"),
//   	Domains: []*string{
//   		jsii.String("self-hosted-tenant.my.domain.com"),
//   	},
//   	Enabled: jsii.Boolean(true),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		PrimaryDomainName: jsii.String("self-hosted-tenant.my.domain.com"),
//   		ValidationTokenHost: jsii.String("self-hosted"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html
//
type CfnDistributionTenantProps struct {
	// The ID of the multi-tenant distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-distributionid
	//
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// The domains associated with the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-domains
	//
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// The name of the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the connection group for the distribution tenant.
	//
	// If you don't specify a connection group, CloudFront uses the default connection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-connectiongroupid
	//
	ConnectionGroupId *string `field:"optional" json:"connectionGroupId" yaml:"connectionGroupId"`
	// Customizations for the distribution tenant.
	//
	// For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and AWS WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-customizations
	//
	Customizations interface{} `field:"optional" json:"customizations" yaml:"customizations"`
	// Indicates whether the distribution tenant is in an enabled state.
	//
	// If disabled, the distribution tenant won't serve traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that represents the request for the Amazon CloudFront managed ACM certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-managedcertificaterequest
	//
	ManagedCertificateRequest interface{} `field:"optional" json:"managedCertificateRequest" yaml:"managedCertificateRequest"`
	// A list of parameter values to add to the resource.
	//
	// A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

