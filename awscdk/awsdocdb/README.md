# Amazon DocumentDB Construct Library

## Starting a Clustered Database

To set up a clustered DocumentDB database, define a `DatabaseCluster`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

```go
var vpc Vpc

cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
		 // NOTE: 'admin' is reserved by DocumentDB
		ExcludeCharacters: jsii.String("\"@/:"),
		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
		SecretName: jsii.String("/myapp/mydocdb/masteruser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	CopyTagsToSnapshot: jsii.Boolean(true),
})
```

By default, the master password will be generated and stored in AWS Secrets Manager with auto-generated description.

Your cluster will be empty by default.

## Serverless Clusters

DocumentDB supports serverless clusters that automatically scale capacity based on your application's needs.
To create a serverless cluster, specify the `serverlessV2ScalingConfiguration` instead of `instanceType`:

```go
var vpc Vpc

cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	Vpc: Vpc,
	ServerlessV2ScalingConfiguration: &ServerlessV2ScalingConfiguration{
		MinCapacity: jsii.Number(0.5),
		MaxCapacity: jsii.Number(2),
	},
	EngineVersion: jsii.String("5.0.0"),
})
```

**Note**: DocumentDB serverless requires engine version 5.0.0 or higher and is not compatible with all features. See the [AWS documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/docdb-serverless-limitations.html) for limitations.

## Connecting

To control who can access the cluster, use the `.connections` attribute. DocumentDB databases have a default port, so
you don't need to specify the port:

```go
var cluster DatabaseCluster

cluster.Connections.AllowDefaultPortFromAnyIpv4(jsii.String("Open to the world"))
```

The endpoints to access your database cluster will be available as the `.clusterEndpoint` and `.clusterReadEndpoint`
attributes:

```go
var cluster DatabaseCluster

writeAddress := cluster.ClusterEndpoint.SocketAddress
```

If you have existing security groups you would like to add to the cluster, use the `addSecurityGroups` method. Security
groups added in this way will not be managed by the `Connections` object of the cluster.

```go
var vpc Vpc
var cluster DatabaseCluster


securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
	Vpc: Vpc,
})
cluster.AddSecurityGroups(securityGroup)
```

## Deletion protection

Deletion protection can be enabled on an Amazon DocumentDB cluster to prevent accidental deletion of the cluster:

```go
var vpc Vpc

cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	DeletionProtection: jsii.Boolean(true),
})
```

## Rotating credentials

When the master password is generated and stored in AWS Secrets Manager, it can be rotated automatically:

```go
var cluster DatabaseCluster

cluster.AddRotationSingleUser()
```

```go
cluster := docdb.NewDatabaseCluster(stack, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("docdb"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R5, ec2.InstanceSize_LARGE),
	Vpc: Vpc,
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
})

cluster.AddRotationSingleUser()
```

The multi user rotation scheme is also available:

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"

var myImportedSecret Secret
var cluster DatabaseCluster


cluster.AddRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
	Secret: myImportedSecret,
})
```

It's also possible to create user credentials together with the cluster and add rotation:

```go
var cluster DatabaseCluster

myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &DatabaseSecretProps{
	Username: jsii.String("myuser"),
	MasterSecret: cluster.Secret,
})
myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret

cluster.AddRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
	 // Add rotation using the multi user scheme
	Secret: myUserSecretAttached,
})
```

**Note**: This user must be created manually in the database using the master credentials.
The rotation will start as soon as this user exists.

See also [aws-cdk-lib/aws-secretsmanager](https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/aws-secretsmanager/README.md) for credentials rotation of existing clusters.

## Audit and profiler Logs

Sending audit or profiler needs to be configured in two places:

1. Check / create the needed options in your ParameterGroup for [audit](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing) and
   [profiler](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling) logs.
2. Enable the corresponding option(s) when creating the `DatabaseCluster`:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import logs "github.com/aws/aws-cdk-go/awscdk"

var myLogsPublishingRole Role
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	ExportProfilerLogsToCloudWatch: jsii.Boolean(true),
	 // Enable sending profiler logs
	ExportAuditLogsToCloudWatch: jsii.Boolean(true),
	 // Enable sending audit logs
	CloudWatchLogsRetention: logs.RetentionDays_THREE_MONTHS,
	 // Optional - default is to never expire logs
	CloudWatchLogsRetentionRole: myLogsPublishingRole,
})
```

## Enable Performance Insights

By enabling this feature it will be cascaded and enabled in all instances inside the cluster:

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	EnablePerformanceInsights: jsii.Boolean(true),
})
```

## Removal Policy

This resource supports the snapshot removal policy.
To specify it use the `removalPolicy` property:

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	RemovalPolicy: awscdk.RemovalPolicy_SNAPSHOT,
})
```

**Note**: A `RemovalPolicy.DESTROY` removal policy will be applied to the
cluster's instances and security group by default as they don't support the snapshot
removal policy.

> Visit [DeletionPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) for more details.

To specify a custom removal policy for the cluster's instances, use the
`instanceRemovalPolicy` property:

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	RemovalPolicy: awscdk.RemovalPolicy_SNAPSHOT,
	InstanceRemovalPolicy: awscdk.RemovalPolicy_RETAIN,
})
```

To specify a custom removal policy for the cluster's security group, use the
`securityGroupRemovalPolicy` property:

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	RemovalPolicy: awscdk.RemovalPolicy_SNAPSHOT,
	SecurityGroupRemovalPolicy: awscdk.RemovalPolicy_RETAIN,
})
```

## CA certificate

Use the `caCertificate` property to specify the [CA certificate](https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html) to use for all instances inside the cluster:

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	CaCertificate: docdb.CaCertificate_RDS_CA_RSA4096_G1(),
})
```

## Storage Type

You can specify [storage type](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-storage-configs.html) for the cluster.

```go
var vpc Vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	MasterUser: &Login{
		Username: jsii.String("myuser"),
	},
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
	Vpc: Vpc,
	StorageType: docdb.StorageType_IOPT1,
})
```

**Note**: `StorageType.IOPT1` is supported starting with engine version 5.0.0.

**Note**: For serverless clusters, storage type is managed automatically and cannot be specified.
