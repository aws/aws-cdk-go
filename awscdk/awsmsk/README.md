# Amazon Managed Streaming for Apache Kafka Construct Library

[Amazon MSK](https://aws.amazon.com/msk/) is a fully managed service that makes it easy for you to build and run applications that use Apache Kafka to process streaming data.

The following example creates an MSK Cluster.

```go
var vpc vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V2_8_1(),
	Vpc: Vpc,
})
```

## Allowing Connections

To control who can access the Cluster, use the `.connections` attribute. For a list of ports used by MSK, refer to the [MSK documentation](https://docs.aws.amazon.com/msk/latest/developerguide/client-access.html#port-info).

```go
var vpc vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V2_8_1(),
	Vpc: Vpc,
})

cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(2181)))
cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(9094)))
```

## Cluster Endpoints

You can use the following attributes to get a list of the Kafka broker or ZooKeeper node endpoints

```go
var cluster cluster

awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokers"), &CfnOutputProps{
	Value: cluster.bootstrapBrokers,
})
awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokersTls"), &CfnOutputProps{
	Value: cluster.bootstrapBrokersTls,
})
awscdk.NewCfnOutput(this, jsii.String("BootstrapBrokersSaslScram"), &CfnOutputProps{
	Value: cluster.bootstrapBrokersSaslScram,
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

> Only one authentication method can be enabled.

### TLS

To enable client authentication with TLS set the `certificateAuthorityArns` property to reference your ACM Private CA. [More info on Private CAs.](https://docs.aws.amazon.com/msk/latest/developerguide/msk-authentication.html)

```go
import acmpca "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V2_8_1(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Tls(&TlsAuthProps{
		CertificateAuthorities: []iCertificateAuthority{
			acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
		},
	}),
})
```

### SASL/SCRAM

Enable client authentication with [SASL/SCRAM](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html):

```go
var vpc vpc

cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V2_8_1(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
		Scram: jsii.Boolean(true),
	}),
})
```

### SASL/IAM

Enable client authentication with [IAM](https://docs.aws.amazon.com/msk/latest/developerguide/iam-access-control.html):

```go
var vpc vpc

cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
	ClusterName: jsii.String("myCluster"),
	KafkaVersion: msk.KafkaVersion_V2_8_1(),
	Vpc: Vpc,
	EncryptionInTransit: &EncryptionInTransitConfig{
		ClientBroker: msk.ClientBrokerEncryption_TLS,
	},
	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
		Iam: jsii.Boolean(true),
	}),
})
```
