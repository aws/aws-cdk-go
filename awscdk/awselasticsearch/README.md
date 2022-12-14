# Amazon Elasticsearch Service Construct Library

> Instead of this module, we recommend using the [@aws-cdk/aws-opensearchservice](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-opensearchservice-readme.html) module. See [Amazon OpenSearch Service FAQs](https://aws.amazon.com/opensearch-service/faqs/#Name_change) for details. See [Migrating to OpenSearch](#migrating-to-opensearch) for migration instructions.

## Quick start

Create a development cluster by simply specifying the version:

```go
devDomain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
})
```

To perform version upgrades without replacing the entire domain, specify the `enableVersionUpgrade` property.

```go
devDomain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_10(),
	enableVersionUpgrade: jsii.Boolean(true),
})
```

Create a production grade cluster by also specifying things like capacity and az distribution

```go
prodDomain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	capacity: &capacityConfig{
		masterNodes: jsii.Number(5),
		dataNodes: jsii.Number(20),
	},
	ebs: &ebsOptions{
		volumeSize: jsii.Number(20),
	},
	zoneAwareness: &zoneAwarenessConfig{
		availabilityZoneCount: jsii.Number(3),
	},
	logging: &loggingOptions{
		slowSearchLogEnabled: jsii.Boolean(true),
		appLogEnabled: jsii.Boolean(true),
		slowIndexLogEnabled: jsii.Boolean(true),
	},
})
```

This creates an Elasticsearch cluster and automatically sets up log groups for
logging the domain logs and slow search logs.

## A note about SLR

Some cluster configurations (e.g VPC access) require the existence of the [`AWSServiceRoleForAmazonElasticsearchService`](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/slr-es.html) Service-Linked Role.

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
slr := iam.NewCfnServiceLinkedRole(this, jsii.String("ElasticSLR"), &cfnServiceLinkedRoleProps{
	awsServiceName: jsii.String("es.amazonaws.com"),
})
```

## Importing existing domains

To import an existing domain into your CDK application, use the `Domain.fromDomainEndpoint` factory method.
This method accepts a domain endpoint of an already existing domain:

```go
domainEndpoint := "https://my-domain-jcjotrt6f7otem4sqcwbch3c4u.us-east-1.es.amazonaws.com"
domain := es.domain.fromDomainEndpoint(this, jsii.String("ImportedDomain"), domainEndpoint)
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
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_4(),
	ebs: &ebsOptions{
		volumeSize: jsii.Number(100),
		volumeType: ec2.ebsDeviceVolumeType_GENERAL_PURPOSE_SSD,
	},
	nodeToNodeEncryption: jsii.Boolean(true),
	encryptionAtRest: &encryptionAtRestOptions{
		enabled: jsii.Boolean(true),
	},
})
```

This sets up the domain with node to node encryption and encryption at
rest. You can also choose to supply your own KMS key to use for encryption at
rest.

## VPC Support

Elasticsearch domains can be placed inside a VPC, providing a secure communication between Amazon ES and other services within the VPC without the need for an internet gateway, NAT device, or VPN connection.

> Visit [VPC Support for Amazon Elasticsearch Service Domains](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html) for more details.

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"))
domainProps := &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	removalPolicy: awscdk.RemovalPolicy_DESTROY,
	vpc: vpc,
	// must be enabled since our VPC contains multiple private subnets.
	zoneAwareness: &zoneAwarenessConfig{
		enabled: jsii.Boolean(true),
	},
	capacity: &capacityConfig{
		// must be an even number since the default az count is 2.
		dataNodes: jsii.Number(2),
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
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	enforceHttps: jsii.Boolean(true),
	nodeToNodeEncryption: jsii.Boolean(true),
	encryptionAtRest: &encryptionAtRestOptions{
		enabled: jsii.Boolean(true),
	},
	fineGrainedAccessControl: &advancedSecurityOptions{
		masterUserName: jsii.String("master-user"),
	},
})

masterUserPassword := domain.masterUserPassword
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
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	useUnsignedBasicAuth: jsii.Boolean(true),
})

masterUserPassword := domain.masterUserPassword
```

## Custom access policies

If the domain requires custom access control it can be configured either as a
constructor property, or later by means of a helper method.

For simple permissions the `accessPolicies` constructor may be sufficient:

```go
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	accessPolicies: []policyStatement{
		iam.NewPolicyStatement(&policyStatementProps{
			actions: []*string{
				jsii.String("es:*ESHttpPost"),
				jsii.String("es:ESHttpPut*"),
			},
			effect: iam.effect_ALLOW,
			principals: []iPrincipal{
				iam.NewAccountPrincipal(jsii.String("123456789012")),
			},
			resources: []*string{
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
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
})

domain.addAccessPolicies(
iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("es:ESHttpPost"),
		jsii.String("es:ESHttpPut"),
	},
	effect: iam.effect_ALLOW,
	principals: []iPrincipal{
		iam.NewAccountPrincipal(jsii.String("123456789012")),
	},
	resources: []*string{
		domain.domainArn,
		fmt.Sprintf("%v/*", domain.domainArn),
	},
}),
iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("es:ESHttpGet"),
	},
	effect: iam.*effect_ALLOW,
	principals: []*iPrincipal{
		iam.NewAccountPrincipal(jsii.String("123456789012")),
	},
	resources: []*string{
		fmt.Sprintf("%v/_all/_settings", domain.domainArn),
		fmt.Sprintf("%v/_cluster/stats", domain.domainArn),
		fmt.Sprintf("%v/index-name*/_mapping/type-name", domain.domainArn),
		fmt.Sprintf("%v/roletest*/_mapping/roletest", domain.domainArn),
		fmt.Sprintf("%v/_nodes", domain.domainArn),
		fmt.Sprintf("%v/_nodes/stats", domain.domainArn),
		fmt.Sprintf("%v/_nodes/*/stats", domain.domainArn),
		fmt.Sprintf("%v/_stats", domain.domainArn),
		fmt.Sprintf("%v/index-name*/_stats", domain.domainArn),
		fmt.Sprintf("%v/roletest*/_stat", domain.domainArn),
	},
}))
```

## Audit logs

Audit logs can be enabled for a domain, but only when fine grained access control is enabled.

```go
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_1(),
	enforceHttps: jsii.Boolean(true),
	nodeToNodeEncryption: jsii.Boolean(true),
	encryptionAtRest: &encryptionAtRestOptions{
		enabled: jsii.Boolean(true),
	},
	fineGrainedAccessControl: &advancedSecurityOptions{
		masterUserName: jsii.String("master-user"),
	},
	logging: &loggingOptions{
		auditLogEnabled: jsii.Boolean(true),
		slowSearchLogEnabled: jsii.Boolean(true),
		appLogEnabled: jsii.Boolean(true),
		slowIndexLogEnabled: jsii.Boolean(true),
	},
})
```

## UltraWarm

UltraWarm nodes can be enabled to provide a cost-effective way to store large amounts of read-only data.

```go
domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_10(),
	capacity: &capacityConfig{
		masterNodes: jsii.Number(2),
		warmNodes: jsii.Number(2),
		warmInstanceType: jsii.String("ultrawarm1.medium.elasticsearch"),
	},
})
```

## Custom endpoint

Custom endpoints can be configured to reach the ES domain under a custom domain name.

```go
es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_7(),
	customEndpoint: &customEndpointOptions{
		domainName: jsii.String("search.example.com"),
	},
})
```

It is also possible to specify a custom certificate instead of the auto-generated one.

Additionally, an automatic CNAME-Record is created if a hosted zone is provided for the custom endpoint

## Advanced options

[Advanced options](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options) can used to configure additional options.

```go
es.NewDomain(this, jsii.String("Domain"), &domainProps{
	version: es.elasticsearchVersion_V7_7(),
	advancedOptions: map[string]*string{
		"rest.action.multi.allow_explicit_index": jsii.String("false"),
		"indices.fielddata.cache.size": jsii.String("25"),
		"indices.query.bool.max_clause_count": jsii.String("2048"),
	},
})
```

## Migrating to OpenSearch

To migrate from this module (`@aws-cdk/aws-elasticsearch`) to the new `@aws-cdk/aws-opensearchservice` module, you must modify your CDK application to refer to the new module (including some associated changes) and then perform a CloudFormation resource deletion/import.

### Necessary CDK Modifications

Make the following modifications to your CDK application to migrate to the `@aws-cdk/aws-opensearchservice` module.

* Rewrite module imports to use `'@aws-cdk/aws-opensearchservice` to `'@aws-cdk/aws-elasticsearch`.
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
  // Example automatically generated from non-compiling source. May contain errors.
  version := es.elasticsearchVersion_V7_1()
  ```

  ...becomes...

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  version := opensearch.engineVersion_ELASTICSEARCH_7_1()
  ```
* Replace the `cognitoKibanaAuth` property of `DomainProps` with `cognitoDashboardsAuth`.
  For example:

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  es.NewDomain(this, jsii.String("Domain"), &domainProps{
  	cognitoKibanaAuth: &cognitoOptions{
  		identityPoolId: jsii.String("test-identity-pool-id"),
  		userPoolId: jsii.String("test-user-pool-id"),
  		role: role,
  	},
  	version: elasticsearchVersion,
  })
  ```

  ...becomes...

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
  	cognitoDashboardsAuth: &cognitoOptions{
  		identityPoolId: jsii.String("test-identity-pool-id"),
  		userPoolId: jsii.String("test-user-pool-id"),
  		role: role,
  	},
  	version: openSearchVersion,
  })
  ```
* Rewrite instance type suffixes from `.elasticsearch` to `.search`.
  For example:

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  es.NewDomain(this, jsii.String("Domain"), &domainProps{
  	capacity: &capacityConfig{
  		masterNodeInstanceType: jsii.String("r5.large.elasticsearch"),
  	},
  	version: elasticsearchVersion,
  })
  ```

  ...becomes...

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
  	capacity: &capacityConfig{
  		masterNodeInstanceType: jsii.String("r5.large.search"),
  	},
  	version: openSearchVersion,
  })
  ```
* Any `CfnInclude`'d domains will need to be re-written in their original template in
  order to be successfully included as a `opensearch.CfnDomain`

### CloudFormation Migration

Follow these steps to migrate your application without data loss:

* Ensure that the [removal policy](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_core.RemovalPolicy.html) on your domains are set to `RemovalPolicy.RETAIN`. This is the default for the domain construct, so nothing is required unless you have specifically set the removal policy to some other value.
* Remove the domain resource from your CloudFormation stacks by manually modifying the synthesized templates used to create the CloudFormation stacks. This may also involve modifying or deleting dependent resources, such as the custom resources that CDK creates to manage the domain's access policy or any other resource you have connected to the domain. You will need to search for references to each domain's logical ID to determine which other resources refer to it and replace or delete those references. Do not remove resources that are dependencies of the domain or you will have to recreate or import them before importing the domain. After modification, deploy the stacks through the AWS Management Console or using the AWS CLI.
* Migrate your CDK application to use the new `@aws-cdk/aws-opensearchservice` module by applying the necessary modifications listed above. Synthesize your application and obtain the resulting stack templates.
* Copy just the definition of the domain from the "migrated" templates to the corresponding "stripped" templates that you deployed above. [Import](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-existing-stack.html) the orphaned domains into your CloudFormation stacks using these templates.
* Synthesize and deploy your CDK application to reconfigure/recreate the modified dependent resources. The CloudFormation stacks should now contain the same resources as existed prior to migration.
* Proceed with development as normal!
