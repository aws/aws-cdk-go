# Amazon Managed Streaming for Apache Kafka Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[Amazon MSK](https://aws.amazon.com/msk/) is a fully managed service that makes it easy for you to build and run applications that use Apache Kafka to process streaming data.

The following example creates an MSK Cluster.

```go
var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
})
```

## Allowing Connections

To control who can access the Cluster, use the `.connections` attribute. For a list of ports used by MSK, refer to the [MSK documentation](https://docs.aws.amazon.com/msk/latest/developerguide/client-access.html#port-info).

```go
var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
})

cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(2181)))
cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(9094)))
```

## Cluster Endpoints

You can use the following attributes to get a list of the Kafka broker or ZooKeeper node endpoints

```go
var cluster Cluster

awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokers"), &CfnOutputProps{
	Value: cluster.bootstrapBrokers,
})
awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokersTls"), &CfnOutputProps{
	Value: cluster.bootstrapBrokersTls,
})
awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokersSaslScram"), &CfnOutputProps{
	Value: cluster.bootstrapBrokersSaslScram,
})
awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokerStringSaslIam"), &CfnOutputProps{
	Value: cluster.bootstrapBrokersSaslIam,
})
awscdk.NewCfnOutput(this, jsii.String("ZookeeperConnection"), &CfnOutputProps{
	Value: cluster.zookeeperConnectionString,
})
awscdk.NewCfnOutput(this, jsii.String("ZookeeperConnectionTls"), &CfnOutputProps{
	Value: cluster.zookeeperConnectionStringTls,
})
```

## Importing an existing Cluster

To import an existing MSK cluster into your CDK app use the `.fromClusterArn()` method.

```go
cluster := msk.Cluster_FromClusterArn(this, jsii.String("Cluster"), jsii.String("arn:aws:kafka:us-west-2:1234567890:cluster/a-cluster/11111111-1111-1111-1111-111111111111-1"))
```

## Client Authentication

[MSK supports](https://docs.aws.amazon.com/msk/latest/developerguide/kafka_apis_iam.html) the following authentication mechanisms.

### TLS

To enable client authentication with TLS set the `certificateAuthorityArns` property to reference your ACM Private CA. [More info on Private CAs.](https://docs.aws.amazon.com/msk/latest/developerguide/msk-authentication.html)

```go
import acmpca "github.com/aws/aws-cdk-go/awscdk"

var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Tls(&TlsAuthProps{
		CertificateAuthorities: []ICertificateAuthority{
			acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
		},
	}),
})
```

### SASL/SCRAM

Enable client authentication with [SASL/SCRAM](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html):

```go
var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
		Scram: jsii.Boolean(true),
	}),
})
```

### IAM

Enable client authentication with [IAM](https://docs.aws.amazon.com/msk/latest/developerguide/iam-access-control.html):

```go
var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
		Iam: jsii.Boolean(true),
	}),
})
```

### SASL/IAM + TLS

Enable client authentication with [IAM](https://docs.aws.amazon.com/msk/latest/developerguide/iam-access-control.html)
as well as enable client authentication with TLS by setting the `certificateAuthorityArns` property to reference your ACM Private CA. [More info on Private CAs.](https://docs.aws.amazon.com/msk/latest/developerguide/msk-authentication.html)

```go
import acmpca "github.com/aws/aws-cdk-go/awscdk"

var vpc Vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_SaslTls(&SaslTlsAuthProps{
		Iam: jsii.Boolean(true),
		CertificateAuthorities: []ICertificateAuthority{
			acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
		},
	}),
})
```

## Logging

You can deliver Apache Kafka broker logs to one or more of the following destination types:
Amazon CloudWatch Logs, Amazon S3, Amazon Data Firehose.

To configure logs to be sent to an S3 bucket, provide a bucket in the `logging` config.

```go
var vpc Vpc
var bucket IBucket

cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	Logging: &BrokerLogging{
		S3: &S3LoggingConfiguration{
			Bucket: *Bucket,
		},
	},
})
```

When the S3 destination is configured, AWS will automatically create an S3 bucket policy
that allows the service to write logs to the bucket. This makes it impossible to later update
that bucket policy. To have CDK create the bucket policy so that future updates can be made,
the `@aws-cdk/aws-s3:createDefaultLoggingPolicy` [feature flag](https://docs.aws.amazon.com/cdk/v2/guide/featureflags.html) can be used. This can be set
in the `cdk.json` file.

```json
{
  "context": {
    "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true
  }
}
```

## Storage Mode

You can configure an MSK cluster storage mode using the `storageMode` property.

Tiered storage is a low-cost storage tier for Amazon MSK that scales to virtually unlimited storage,
making it cost-effective to build streaming data applications.

> Visit [Tiered storage](https://docs.aws.amazon.com/msk/latest/developerguide/msk-tiered-storage.html)
> to see the list of compatible Kafka versions and for more details.

```go
var vpc Vpc
var bucket IBucket


cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V4_1_X_KRAFT(),
	Vpc: Vpc,
	StorageMode: msk.StorageMode_TIERED,
})
```

## MSK Serverless

You can also use MSK Serverless by using `ServerlessCluster` class.

MSK Serverless is a cluster type for Amazon MSK that makes it possible for you to run Apache Kafka without having to manage and scale cluster capacity.

MSK Serverless requires IAM access control for all clusters.

For more infomation, see [Use MSK Serverless clusters](https://docs.aws.amazon.com/msk/latest/developerguide/serverless-getting-started.html).

```go
var vpc Vpc


serverlessCluster := msk.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &ServerlessClusterProps{
	ClusterName: jsii.String("MyServerlessCluster"),
	VpcConfigs: []VpcConfig{
		&VpcConfig{
			Vpc: *Vpc,
		},
	},
})
```
