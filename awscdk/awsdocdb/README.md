# Amazon DocumentDB Construct Library

## Starting a Clustered Database

To set up a clustered DocumentDB database, define a `DatabaseCluster`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

```go
var vpc vpc

cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	masterUser: &login{
		username: jsii.String("myuser"),
		 // NOTE: 'admin' is reserved by DocumentDB
		excludeCharacters: jsii.String("\"@/:"),
		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
		secretName: jsii.String("/myapp/mydocdb/masteruser"),
	},
	instanceType: ec2.instanceType.of(ec2.instanceClass_MEMORY5, ec2.instanceSize_LARGE),
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	},
	vpc: vpc,
})
```

By default, the master password will be generated and stored in AWS Secrets Manager with auto-generated description.

Your cluster will be empty by default.

## Connecting

To control who can access the cluster, use the `.connections` attribute. DocumentDB databases have a default port, so
you don't need to specify the port:

```go
var cluster databaseCluster

cluster.connections.allowDefaultPortFromAnyIpv4(jsii.String("Open to the world"))
```

The endpoints to access your database cluster will be available as the `.clusterEndpoint` and `.clusterReadEndpoint`
attributes:

```go
var cluster databaseCluster

writeAddress := cluster.clusterEndpoint.socketAddress
```

If you have existing security groups you would like to add to the cluster, use the `addSecurityGroups` method. Security
groups added in this way will not be managed by the `Connections` object of the cluster.

```go
var vpc vpc
var cluster databaseCluster


securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &securityGroupProps{
	vpc: vpc,
})
cluster.addSecurityGroups(securityGroup)
```

## Deletion protection

Deletion protection can be enabled on an Amazon DocumentDB cluster to prevent accidental deletion of the cluster:

```go
var vpc vpc

cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	masterUser: &login{
		username: jsii.String("myuser"),
	},
	instanceType: ec2.instanceType.of(ec2.instanceClass_MEMORY5, ec2.instanceSize_LARGE),
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	},
	vpc: vpc,
	deletionProtection: jsii.Boolean(true),
})
```

## Rotating credentials

When the master password is generated and stored in AWS Secrets Manager, it can be rotated automatically:

```go
var cluster databaseCluster

cluster.addRotationSingleUser()
```

```go
cluster := docdb.NewDatabaseCluster(stack, jsii.String("Database"), &databaseClusterProps{
	masterUser: &login{
		username: jsii.String("docdb"),
	},
	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
	vpc: vpc,
	removalPolicy: cdk.removalPolicy_DESTROY,
})

cluster.addRotationSingleUser()
```

The multi user rotation scheme is also available:

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"

var myImportedSecret secret
var cluster databaseCluster


cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
	secret: myImportedSecret,
})
```

It's also possible to create user credentials together with the cluster and add rotation:

```go
var cluster databaseCluster

myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
	username: jsii.String("myuser"),
	masterSecret: cluster.secret,
})
myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret

cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
	 // Add rotation using the multi user scheme
	secret: myUserSecretAttached,
})
```

**Note**: This user must be created manually in the database using the master credentials.
The rotation will start as soon as this user exists.

See also [@aws-cdk/aws-secretsmanager](https://github.com/aws/aws-cdk/blob/main/packages/%40aws-cdk/aws-secretsmanager/README.md) for credentials rotation of existing clusters.

## Audit and profiler Logs

Sending audit or profiler needs to be configured in two places:

1. Check / create the needed options in your ParameterGroup for [audit](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing) and
   [profiler](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling) logs.
2. Enable the corresponding option(s) when creating the `DatabaseCluster`:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import logs "github.com/aws/aws-cdk-go/awscdk"

var myLogsPublishingRole role
var vpc vpc


cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	masterUser: &login{
		username: jsii.String("myuser"),
	},
	instanceType: ec2.instanceType.of(ec2.instanceClass_MEMORY5, ec2.instanceSize_LARGE),
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	},
	vpc: vpc,
	exportProfilerLogsToCloudWatch: jsii.Boolean(true),
	 // Enable sending profiler logs
	exportAuditLogsToCloudWatch: jsii.Boolean(true),
	 // Enable sending audit logs
	cloudWatchLogsRetention: logs.retentionDays_THREE_MONTHS,
	 // Optional - default is to never expire logs
	cloudWatchLogsRetentionRole: myLogsPublishingRole,
})
```
