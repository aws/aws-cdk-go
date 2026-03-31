package awscloudfront


// The level of permissions granted to the CloudFront Distribution when configuring OAC.
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
//   myMultiTenantDistribution := cloudfront.NewDistribution(this, jsii.String("distribution"), &DistributionProps{
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
//   // Create a distribution tenant using an existing ACM certificate
//   cfnDistributionTenant := cloudfront.NewCfnDistributionTenant(this, jsii.String("distribution-tenant"), &CfnDistributionTenantProps{
//   	DistributionId: myMultiTenantDistribution.DistributionId,
//   	Domains: []*string{
//   		jsii.String("my-tenant.my.domain.com"),
//   	},
//   	Name: jsii.String("my-tenant"),
//   	Enabled: jsii.Boolean(true),
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			Name: jsii.String("tenantName"),
//   			Value: jsii.String("app"),
//   		},
//   	},
//   	Customizations: &CustomizationsProperty{
//   		Certificate: &CertificateProperty{
//   			Arn: jsii.String("REPLACE_WITH_ARN"),
//   		},
//   	},
//   })
//
type AccessLevel string

const (
	// Grants read permissions to CloudFront Distribution.
	AccessLevel_READ AccessLevel = "READ"
	// Grants versioned read permissions to CloudFront Distribution.
	AccessLevel_READ_VERSIONED AccessLevel = "READ_VERSIONED"
	// Grants list permissions to CloudFront Distribution.
	AccessLevel_LIST AccessLevel = "LIST"
	// Grants write permission to CloudFront Distribution.
	AccessLevel_WRITE AccessLevel = "WRITE"
	// Grants delete permission to CloudFront Distribution.
	AccessLevel_DELETE AccessLevel = "DELETE"
)

