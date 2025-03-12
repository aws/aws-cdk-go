# Amazon OpenSearch Service Construct Library

> Instead of this module, we recommend using the [aws-cdk-lib/aws-opensearchservice](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_opensearchservice-readme.html) module. See [Amazon OpenSearch Service FAQs](https://aws.amazon.com/opensearch-service/faqs/#Name_change) for details. See [Migrating to OpenSearch](#migrating-to-opensearch) for migration instructions.

## Quick start

Create a development cluster by simply specifying the version:

```go
devDomain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
})
```

To perform version upgrades without replacing the entire domain, specify the `enableVersionUpgrade` property.

```go
devDomain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_10(),
	EnableVersionUpgrade: jsii.Boolean(true),
})
```

Create a production grade cluster by also specifying things like capacity and az distribution

```go
prodDomain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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

This creates an Elasticsearch cluster and automatically sets up log groups for
logging the domain logs and slow search logs.

## A note about SLR

Some cluster configurations (e.g VPC access) require the existence of the [`AWSServiceRoleForAmazonElasticsearchService`](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/slr.html) service-linked role.

When performing such operations via the AWS Console, this SLR is created automatically when needed. However, this is not the behavior when using CloudFormation. If an SLR is needed, but doesn't exist, you will encounter a failure message similar to:

```console
Before you can proceed, you must enable a service-linked role to give Amazon ES...
```

To resolve this, you need to [create](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#create-service-linked-role) the SLR. We recommend using the AWS CLI:

```console
aws iam create-service-linked-role --aws-service-name es.amazonaws.com
```

You can also create it using the CDK, **but note that only the first application deploying this will succeed**:

```go
slr := iam.NewCfnServiceLinkedRole(this, jsii.String("ElasticSLR"), &CfnServiceLinkedRoleProps{
	AwsServiceName: jsii.String("es.amazonaws.com"),
})
```

## Importing existing domains

To import an existing domain into your CDK application, use the `Domain.fromDomainEndpoint` factory method.
This method accepts a domain endpoint of an already existing domain:

```go
domainEndpoint := "https://my-domain-jcjotrt6f7otem4sqcwbch3c4u.us-east-1.es.amazonaws.com"
domain := es.Domain_FromDomainEndpoint(this, jsii.String("ImportedDomain"), domainEndpoint)
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
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_4(),
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

Elasticsearch domains can be placed inside a VPC, providing a secure communication between Amazon ES and other services within the VPC without the need for an internet gateway, NAT device, or VPN connection.

> See [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) for more details.

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"))
domainProps := &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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
es.NewDomain(this, jsii.String("Domain"), domainProps)
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
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
	UseUnsignedBasicAuth: jsii.Boolean(true),
})

masterUserPassword := domain.MasterUserPassword
```

## Custom access policies

If the domain requires custom access control it can be configured either as a
constructor property, or later by means of a helper method.

For simple permissions the `accessPolicies` constructor may be sufficient:

```go
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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

Audit logs can be enabled for a domain, but only when fine-grained access control is enabled.

```go
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_1(),
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

## UltraWarm

UltraWarm nodes can be enabled to provide a cost-effective way to store large amounts of read-only data.

```go
domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_10(),
	Capacity: &CapacityConfig{
		MasterNodes: jsii.Number(2),
		WarmNodes: jsii.Number(2),
		WarmInstanceType: jsii.String("ultrawarm1.medium.elasticsearch"),
	},
})
```

## Custom endpoint

Custom endpoints can be configured to reach the ES domain under a custom domain name.

```go
es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_7(),
	CustomEndpoint: &CustomEndpointOptions{
		DomainName: jsii.String("search.example.com"),
	},
})
```

It is also possible to specify a custom certificate instead of the auto-generated one.

Additionally, an automatic CNAME-Record is created if a hosted zone is provided for the custom endpoint

## Advanced options

[Advanced cluster settings](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) can used to configure additional options.

```go
es.NewDomain(this, jsii.String("Domain"), &DomainProps{
	Version: es.ElasticsearchVersion_V7_7(),
	AdvancedOptions: map[string]*string{
		"rest.action.multi.allow_explicit_index": jsii.String("false"),
		"indices.fielddata.cache.size": jsii.String("25"),
		"indices.query.bool.max_clause_count": jsii.String("2048"),
	},
})
```

## Migrating to OpenSearch

To migrate from this module (`aws-cdk-lib/aws-elasticsearch`) to the new `aws-cdk-lib/aws-opensearchservice` module, you must modify your CDK application to refer to the new module (including some associated changes) and then perform a CloudFormation resource deletion/import.

### Necessary CDK Modifications

Make the following modifications to your CDK application to migrate to the `aws-cdk-lib/aws-opensearchservice` module.

* Rewrite module imports to use `'aws-cdk-lib/aws-opensearchservice` to `'aws-cdk-lib/aws-elasticsearch`.
  For example:

  ```go
  import es "github.com/aws/aws-cdk-go/awscdk"
  import "github.com/aws/aws-cdk-go/awscdk"
  ```

  ...becomes...

  ```go
  import opensearch "github.com/aws/aws-cdk-go/awscdk"
  import "github.com/aws/aws-cdk-go/awscdk"
  ```
* Replace instances of `es.ElasticsearchVersion` with `opensearch.EngineVersion`.
  For example:

  ```go
  version := es.elasticsearchVersion_V7_1()
  ```

  ...becomes...

  ```go
  version := opensearch.engineVersion_ELASTICSEARCH_7_1()
  ```
* Replace the `cognitoKibanaAuth` property of `DomainProps` with `cognitoDashboardsAuth`.
  For example:

  ```go
  es.NewDomain(this, jsii.String("Domain"), &DomainProps{
  	CognitoKibanaAuth: &CognitoOptions{
  		IdentityPoolId: jsii.String("test-identity-pool-id"),
  		UserPoolId: jsii.String("test-user-pool-id"),
  		Role: role,
  	},
  	Version: elasticsearchVersion,
  })
  ```

  ...becomes...

  ```go
  opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
  	CognitoDashboardsAuth: &CognitoOptions{
  		IdentityPoolId: jsii.String("test-identity-pool-id"),
  		UserPoolId: jsii.String("test-user-pool-id"),
  		Role: role,
  	},
  	Version: openSearchVersion,
  })
  ```
* Rewrite instance type suffixes from `.elasticsearch` to `.search`.
  For example:

  ```go
  es.NewDomain(this, jsii.String("Domain"), &DomainProps{
  	Capacity: &CapacityConfig{
  		MasterNodeInstanceType: jsii.String("r5.large.elasticsearch"),
  	},
  	Version: elasticsearchVersion,
  })
  ```

  ...becomes...

  ```go
  opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
  	Capacity: &CapacityConfig{
  		MasterNodeInstanceType: jsii.String("r5.large.search"),
  	},
  	Version: openSearchVersion,
  })
  ```
* Any `CfnInclude`'d domains will need to be re-written in their original template in
  order to be successfully included as a `opensearch.CfnDomain`

### CloudFormation Migration

Follow these steps to migrate your application without data loss:

* Ensure that the [removal policy](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_core.RemovalPolicy.html) on your domains are set to `RemovalPolicy.RETAIN`. This is the default for the domain construct, so nothing is required unless you have specifically set the removal policy to some other value.
* Remove the domain resource from your CloudFormation stacks by manually modifying the synthesized templates used to create the CloudFormation stacks. This may also involve modifying or deleting dependent resources, such as the custom resources that CDK creates to manage the domain's access policy or any other resource you have connected to the domain. You will need to search for references to each domain's logical ID to determine which other resources refer to it and replace or delete those references. Do not remove resources that are dependencies of the domain or you will have to recreate or import them before importing the domain. After modification, deploy the stacks through the AWS Management Console or using the AWS CLI.
* Migrate your CDK application to use the new `aws-cdk-lib/aws-opensearchservice` module by applying the necessary modifications listed above. Synthesize your application and obtain the resulting stack templates.
* Copy just the definition of the domain from the "migrated" templates to the corresponding "stripped" templates that you deployed above. [Import](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-existing-stack.html) the orphaned domains into your CloudFormation stacks using these templates.
* Synthesize and deploy your CDK application to reconfigure/recreate the modified dependent resources. The CloudFormation stacks should now contain the same resources as existed prior to migration.
* Proceed with development as normal!
