# Amazon Neptune Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

Amazon Neptune is a fast, reliable, fully managed graph database service that makes it easy to build and run applications that work with highly connected datasets. The core of Neptune is a purpose-built, high-performance graph database engine. This engine is optimized for storing billions of relationships and querying the graph with milliseconds latency. Neptune supports the popular graph query languages Apache TinkerPop Gremlin and W3C’s SPARQL, enabling you to build queries that efficiently navigate highly connected datasets.

The `@aws-cdk/aws-neptune` package contains primitives for setting up Neptune database clusters and instances.

```go
import neptune "github.com/aws/aws-cdk-go/awscdkneptunealpha"
```

## Starting a Neptune Database

To set up a Neptune database, define a `DatabaseCluster`. You must always launch a database in a VPC.

```go
cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	vpc: vpc,
	instanceType: neptune.instanceType_R5_LARGE(),
})
```

By default only writer instance is provisioned with this construct.

## Connecting

To control who can access the cluster, use the `.connections` attribute. Neptune databases have a default port, so
you don't need to specify the port:

```go
cluster.connections.allowDefaultPortFromAnyIpv4(jsii.String("Open to the world"))
```

The endpoints to access your database cluster will be available as the `.clusterEndpoint` and `.clusterReadEndpoint`
attributes:

```go
writeAddress := cluster.clusterEndpoint.socketAddress
```

## IAM Authentication

You can also authenticate to a database cluster using AWS Identity and Access Management (IAM) database authentication;
See [https://docs.aws.amazon.com/neptune/latest/userguide/iam-auth.html](https://docs.aws.amazon.com/neptune/latest/userguide/iam-auth.html) for more information and a list of supported
versions and limitations.

The following example shows enabling IAM authentication for a database cluster and granting connection access to an IAM role.

```go
cluster := neptune.NewDatabaseCluster(this, jsii.String("Cluster"), &databaseClusterProps{
	vpc: vpc,
	instanceType: neptune.instanceType_R5_LARGE(),
	iamAuthentication: jsii.Boolean(true),
})
role := iam.NewRole(this, jsii.String("DBRole"), &roleProps{
	assumedBy: iam.NewAccountPrincipal(this.account),
})
// Use one of the following statements to grant the role the necessary permissions
cluster.grantConnect(role) // Grant the role neptune-db:* access to the DB
cluster.grant(role, jsii.String("neptune-db:ReadDataViaQuery"), jsii.String("neptune-db:WriteDataViaQuery"))
```

## Customizing parameters

Neptune allows configuring database behavior by supplying custom parameter groups.  For more details, refer to the
following link: [https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html](https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html)

```go
clusterParams := neptune.NewClusterParameterGroup(this, jsii.String("ClusterParams"), &clusterParameterGroupProps{
	description: jsii.String("Cluster parameter group"),
	parameters: map[string]*string{
		"neptune_enable_audit_log": jsii.String("1"),
	},
})

dbParams := neptune.NewParameterGroup(this, jsii.String("DbParams"), &parameterGroupProps{
	description: jsii.String("Db parameter group"),
	parameters: map[string]*string{
		"neptune_query_timeout": jsii.String("120000"),
	},
})

cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	vpc: vpc,
	instanceType: neptune.instanceType_R5_LARGE(),
	clusterParameterGroup: clusterParams,
	parameterGroup: dbParams,
})
```

Note: if you want to use Neptune engine `1.2.0.0` or later, you need to specify the corresponding `engineVersion` prop to `neptune.DatabaseCluster` and `family` prop of `ParameterGroupFamily.NEPTUNE_1_2` to `neptune.ClusterParameterGroup` and `neptune.ParameterGroup`.

## Adding replicas

`DatabaseCluster` allows launching replicas along with the writer instance. This can be specified using the `instanceCount`
attribute.

```go
cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	vpc: vpc,
	instanceType: neptune.instanceType_R5_LARGE(),
	instances: jsii.Number(2),
})
```

Additionally it is also possible to add replicas using `DatabaseInstance` for an existing cluster.

```go
replica1 := neptune.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	cluster: cluster,
	instanceType: neptune.*instanceType_R5_LARGE(),
})
```

## Automatic minor version upgrades

By setting `autoMinorVersionUpgrade` to true, Neptune will automatically update
the engine of the entire cluster to the latest minor version after a stabilization
window of 2 to 3 weeks.

```go
neptune.NewDatabaseCluster(this, jsii.String("Cluster"), &databaseClusterProps{
	vpc: vpc,
	instanceType: neptune.instanceType_R5_LARGE(),
	autoMinorVersionUpgrade: jsii.Boolean(true),
})
```
