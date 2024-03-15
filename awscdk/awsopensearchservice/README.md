# Amazon OpenSearch Service Construct Library

See [Migrating to OpenSearch](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-elasticsearch-readme.html#migrating-to-opensearch) for migration instructions from `aws-cdk-lib/aws-elasticsearch` to this module, `aws-cdk-lib/aws-opensearchservice`.

## Quick start

Create a development cluster by simply specifying the version:

```go
devDomain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
})
```

To perform version upgrades without replacing the entire domain, specify the `enableVersionUpgrade` property.

```go
devDomain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	EnableVersionUpgrade: jsii.Boolean(true),
})
```

Create a cluster with GP3 volumes:

```go
gp3Domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_2_5(),
	Ebs: &EbsOptions{
		VolumeSize: jsii.Number(30),
		VolumeType: ec2.EbsDeviceVolumeType_GP3,
		Throughput: jsii.Number(125),
		Iops: jsii.Number(3000),
	},
})
```

Create a production grade cluster by also specifying things like capacity and az distribution

```go
prodDomain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	Capacity: &CapacityConfig{
		MasterNodes: jsii.Number(5),
		DataNodes: jsii.Number(20),
	},
	Ebs: &EbsOptions{
		VolumeSize: jsii.Number(20),
	},
	ZoneAwareness: &ZoneAwarenessConfig{
		AvailabilityZoneCount: jsii.Number(3),
	},
	Logging: &LoggingOptions{
		SlowSearchLogEnabled: jsii.Boolean(true),
		AppLogEnabled: jsii.Boolean(true),
		SlowIndexLogEnabled: jsii.Boolean(true),
	},
})
```

This creates an Amazon OpenSearch Service cluster and automatically sets up log groups for
logging the domain logs and slow search logs.

## A note about SLR

Some cluster configurations (e.g VPC access) require the existence of the [`AWSServiceRoleForAmazonElasticsearchService`](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/slr.html) Service-Linked Role.

When performing such operations via the AWS Console, this SLR is created automatically when needed. However, this is not the behavior when using CloudFormation. If an SLR is needed, but doesn't exist, you will encounter a failure message similar to:

```console
Before you can proceed, you must enable a service-linked role to give Amazon OpenSearch Service...
```

To resolve this, you need to [create](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#create-service-linked-role) the SLR. We recommend using the AWS CLI:

```console
aws iam create-service-linked-role --aws-service-name es.amazonaws.com
```

You can also create it using the CDK, **but note that only the first application deploying this will succeed**:

```go
slr := iam.NewCfnServiceLinkedRole(this, jsii.String("Service Linked Role"), &CfnServiceLinkedRoleProps{
	AwsServiceName: jsii.String("es.amazonaws.com"),
})
```

## Importing existing domains

### Using a known domain endpoint

To import an existing domain into your CDK application, use the `Domain.fromDomainEndpoint` factory method.
This method accepts a domain endpoint of an already existing domain:

```go
domainEndpoint := "https://my-domain-jcjotrt6f7otem4sqcwbch3c4u.us-east-1.es.amazonaws.com"
domain := awscdk.Domain_FromDomainEndpoint(this, jsii.String("ImportedDomain"), domainEndpoint)
```

### Using the output of another CloudFormation stack

To import an existing domain with the help of an exported value from another CloudFormation stack,
use the `Domain.fromDomainAttributes` factory method. This will accept tokens.

```go
domainArn := awscdk.Fn_ImportValue(jsii.String("another-cf-stack-export-domain-arn"))
domainEndpoint := awscdk.Fn_ImportValue(jsii.String("another-cf-stack-export-domain-endpoint"))
domain := awscdk.Domain_FromDomainAttributes(this, jsii.String("ImportedDomain"), &DomainAttributes{
	DomainArn: jsii.String(DomainArn),
	DomainEndpoint: jsii.String(DomainEndpoint),
})
```

## Permissions

### IAM

Helper methods also exist for managing access to the domain.

```go
var fn function
var domain domain


// Grant write access to the app-search index
domain.grantIndexWrite(jsii.String("app-search"), fn)

// Grant read access to the 'app-search/_search' path
domain.grantPathRead(jsii.String("app-search/_search"), fn)
```

## Encryption

The domain can also be created with encryption enabled:

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	Ebs: &EbsOptions{
		VolumeSize: jsii.Number(100),
		VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD,
	},
	NodeToNodeEncryption: jsii.Boolean(true),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
})
```

This sets up the domain with node to node encryption and encryption at
rest. You can also choose to supply your own KMS key to use for encryption at
rest.

## VPC Support

Domains can be placed inside a VPC, providing a secure communication between Amazon OpenSearch Service and other services within the VPC without the need for an internet gateway, NAT device, or VPN connection.

> Visit [VPC Support for Amazon OpenSearch Service Domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) for more details.

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"))
domainProps := &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	Vpc: Vpc,
	// must be enabled since our VPC contains multiple private subnets.
	ZoneAwareness: &ZoneAwarenessConfig{
		Enabled: jsii.Boolean(true),
	},
	Capacity: &CapacityConfig{
		// must be an even number since the default az count is 2.
		DataNodes: jsii.Number(2),
	},
}
awscdk.NewDomain(this, jsii.String("Domain"), domainProps)
```

In addition, you can use the `vpcSubnets` property to control which specific subnets will be used, and the `securityGroups` property to control
which security groups will be attached to the domain. By default, CDK will select all *private* subnets in the VPC, and create one dedicated security group.

## Metrics

Helper methods exist to access common domain metrics for example:

```go
var domain domain

freeStorageSpace := domain.metricFreeStorageSpace()
masterSysMemoryUtilization := domain.metric(jsii.String("MasterSysMemoryUtilization"))
```

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Fine grained access control

The domain can also be created with a master user configured. The password can
be supplied or dynamically created if not supplied.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	EnforceHttps: jsii.Boolean(true),
	NodeToNodeEncryption: jsii.Boolean(true),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	FineGrainedAccessControl: &AdvancedSecurityOptions{
		MasterUserName: jsii.String("master-user"),
	},
})

masterUserPassword := domain.MasterUserPassword
```

## SAML authentication

You can enable SAML authentication to use your existing identity provider
to offer single sign-on (SSO) for dashboards on Amazon OpenSearch Service domains
running OpenSearch or Elasticsearch 6.7 or later.
To use SAML authentication, fine-grained access control must be enabled.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	EnforceHttps: jsii.Boolean(true),
	NodeToNodeEncryption: jsii.Boolean(true),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	FineGrainedAccessControl: &AdvancedSecurityOptions{
		MasterUserName: jsii.String("master-user"),
		SamlAuthenticationEnabled: jsii.Boolean(true),
		SamlAuthenticationOptions: &SAMLOptionsProperty{
			IdpEntityId: jsii.String("entity-id"),
			IdpMetadataContent: jsii.String("metadata-content-with-quotes-escaped"),
		},
	},
})
```

## Using unsigned basic auth

For convenience, the domain can be configured to allow unsigned HTTP requests
that use basic auth. Unless the domain is configured to be part of a VPC this
means anyone can access the domain using the configured master username and
password.

To enable unsigned basic auth access the domain is configured with an access
policy that allows anonymous requests, HTTPS required, node to node encryption,
encryption at rest and fine grained access control.

If the above settings are not set they will be configured as part of enabling
unsigned basic auth. If they are set with conflicting values, an error will be
thrown.

If no master user is configured a default master user is created with the
username `admin`.

If no password is configured a default master user password is created and
stored in the AWS Secrets Manager as secret. The secret has the prefix
`<domain id>MasterUser`.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	UseUnsignedBasicAuth: jsii.Boolean(true),
})

masterUserPassword := domain.MasterUserPassword
```

## Custom access policies

If the domain requires custom access control it can be configured either as a
constructor property, or later by means of a helper method.

For simple permissions the `accessPolicies` constructor may be sufficient:

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	AccessPolicies: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("es:*ESHttpPost"),
				jsii.String("es:ESHttpPut*"),
			},
			Effect: iam.Effect_ALLOW,
			Principals: []iPrincipal{
				iam.NewAccountPrincipal(jsii.String("123456789012")),
			},
			Resources: []*string{
				jsii.String("*"),
			},
		}),
	},
})
```

For more complex use-cases, for example, to set the domain up to receive data from a
[cross-account Kinesis Firehose](https://aws.amazon.com/premiumsupport/knowledge-center/kinesis-firehose-cross-account-streaming/) the `addAccessPolicies` helper method
allows for policies that include the explicit domain ARN.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
})
domain.AddAccessPolicies(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("es:ESHttpPost"),
		jsii.String("es:ESHttpPut"),
	},
	Effect: iam.Effect_ALLOW,
	Principals: []iPrincipal{
		iam.NewAccountPrincipal(jsii.String("123456789012")),
	},
	Resources: []*string{
		domain.DomainArn,
		fmt.Sprintf("%v/*", domain.*DomainArn),
	},
}),
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("es:ESHttpGet"),
	},
	Effect: iam.Effect_ALLOW,
	Principals: []*iPrincipal{
		iam.NewAccountPrincipal(jsii.String("123456789012")),
	},
	Resources: []*string{
		fmt.Sprintf("%v/_all/_settings", domain.*DomainArn),
		fmt.Sprintf("%v/_cluster/stats", domain.*DomainArn),
		fmt.Sprintf("%v/index-name*/_mapping/type-name", domain.*DomainArn),
		fmt.Sprintf("%v/roletest*/_mapping/roletest", domain.*DomainArn),
		fmt.Sprintf("%v/_nodes", domain.*DomainArn),
		fmt.Sprintf("%v/_nodes/stats", domain.*DomainArn),
		fmt.Sprintf("%v/_nodes/*/stats", domain.*DomainArn),
		fmt.Sprintf("%v/_stats", domain.*DomainArn),
		fmt.Sprintf("%v/index-name*/_stats", domain.*DomainArn),
		fmt.Sprintf("%v/roletest*/_stat", domain.*DomainArn),
	},
}))
```

## Audit logs

Audit logs can be enabled for a domain, but only when fine grained access control is enabled.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	EnforceHttps: jsii.Boolean(true),
	NodeToNodeEncryption: jsii.Boolean(true),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	FineGrainedAccessControl: &AdvancedSecurityOptions{
		MasterUserName: jsii.String("master-user"),
	},
	Logging: &LoggingOptions{
		AuditLogEnabled: jsii.Boolean(true),
		SlowSearchLogEnabled: jsii.Boolean(true),
		AppLogEnabled: jsii.Boolean(true),
		SlowIndexLogEnabled: jsii.Boolean(true),
	},
})
```

## Suppress creating CloudWatch Logs resource policy

When logging is enabled for the domain, the CloudWatch Logs resource policy is created by default.
This resource policy is necessary for logging, but since only a maximum of 10 resource policies can be created per region,
the maximum number of resource policies may be a problem when enabling logging for several domains.
By setting the `suppressLogsResourcePolicy` option to true, you can suppress the creation of a CloudWatch Logs resource policy.

If you set the `suppressLogsResourcePolicy` option to true, you must create a resource policy before deployment.
Also, to avoid reaching this limit, consider reusing a broader policy that includes multiple log groups.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	EnforceHttps: jsii.Boolean(true),
	NodeToNodeEncryption: jsii.Boolean(true),
	EncryptionAtRest: &EncryptionAtRestOptions{
		Enabled: jsii.Boolean(true),
	},
	FineGrainedAccessControl: &AdvancedSecurityOptions{
		MasterUserName: jsii.String("master-user"),
	},
	Logging: &LoggingOptions{
		AuditLogEnabled: jsii.Boolean(true),
		SlowSearchLogEnabled: jsii.Boolean(true),
		AppLogEnabled: jsii.Boolean(true),
		SlowIndexLogEnabled: jsii.Boolean(true),
	},
	SuppressLogsResourcePolicy: jsii.Boolean(true),
})
```

> Visit [Monitoring OpenSearch logs with Amazon CloudWatch Logs](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createdomain-configure-slow-logs.html) for more details.

## UltraWarm

UltraWarm nodes can be enabled to provide a cost-effective way to store large amounts of read-only data.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	Capacity: &CapacityConfig{
		MasterNodes: jsii.Number(2),
		WarmNodes: jsii.Number(2),
		WarmInstanceType: jsii.String("ultrawarm1.medium.search"),
	},
})
```

## Cold storage

Cold storage can be enabled on the domain. You must enable UltraWarm storage to enable cold storage.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	Capacity: &CapacityConfig{
		MasterNodes: jsii.Number(2),
		WarmNodes: jsii.Number(2),
		WarmInstanceType: jsii.String("ultrawarm1.medium.search"),
	},
	ColdStorageEnabled: jsii.Boolean(true),
})
```

## Custom endpoint

Custom endpoints can be configured to reach the domain under a custom domain name.

```go
awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	CustomEndpoint: &CustomEndpointOptions{
		DomainName: jsii.String("search.example.com"),
	},
})
```

It is also possible to specify a custom certificate instead of the auto-generated one.

Additionally, an automatic CNAME-Record is created if a hosted zone is provided for the custom endpoint

## Advanced options

[Advanced options](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) can used to configure additional options.

```go
awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	AdvancedOptions: map[string]*string{
		"rest.action.multi.allow_explicit_index": jsii.String("false"),
		"indices.fielddata.cache.size": jsii.String("25"),
		"indices.query.bool.max_clause_count": jsii.String("2048"),
	},
})
```

## Amazon Cognito authentication for OpenSearch Dashboards

The domain can be configured to use Amazon Cognito authentication for OpenSearch Dashboards.

> Visit [Configuring Amazon Cognito authentication for OpenSearch Dashboards](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html) for more details.

```go
var cognitoConfigurationRole role


domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
	CognitoDashboardsAuth: &CognitoOptions{
		Role: cognitoConfigurationRole,
		IdentityPoolId: jsii.String("example-identity-pool-id"),
		UserPoolId: jsii.String("example-user-pool-id"),
	},
})
```

## Enable support for Multi-AZ with Standby deployment

The domain can be configured to use [multi-AZ with standby](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html#managedomains-za-standby).

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
	Ebs: &EbsOptions{
		VolumeSize: jsii.Number(10),
		VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
	},
	ZoneAwareness: &ZoneAwarenessConfig{
		Enabled: jsii.Boolean(true),
		AvailabilityZoneCount: jsii.Number(3),
	},
	Capacity: &CapacityConfig{
		MultiAzWithStandbyEnabled: jsii.Boolean(true),
		MasterNodes: jsii.Number(3),
		DataNodes: jsii.Number(3),
	},
})
```

## Define off-peak windows

The domain can be configured to use a daily 10-hour window considered as off-peak hours.

Off-peak windows were introduced on February 16, 2023.
All domains created before this date have the off-peak window disabled by default.
You must manually enable and configure the off-peak window for these domains.
All domains created after this date will have the off-peak window enabled by default.
You can't disable the off-peak window for a domain after it's enabled.

> Visit [Defining off-peak windows for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html) for more details.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
	OffPeakWindowEnabled: jsii.Boolean(true),
	 // can be omitted if offPeakWindowStart is set
	OffPeakWindowStart: &WindowStartTime{
		Hours: jsii.Number(20),
		Minutes: jsii.Number(0),
	},
})
```

## Configuring service software updates

The domain can be configured to use service software updates.

> Visit [Service software updates in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html) for more details.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
	EnableAutoSoftwareUpdate: jsii.Boolean(true),
})
```

## IP address type

You can specify either dual stack or IPv4 as your IP address type.

```go
domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
	IpAddressType: awscdk.IpAddressType_DUAL_STACK,
})
```
