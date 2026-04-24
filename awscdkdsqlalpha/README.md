# Amazon Aurora DSQL Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

Amazon Aurora DSQL is a serverless, distributed relational database service optimized for transactional workloads. Aurora DSQL offers virtually unlimited scale and doesn't require you to manage infrastructure. The active-active highly available architecture provides 99.99% single-Region and 99.999% multi-Region availability.

The `@aws-cdk/aws-dsql-alpha` package contains primitives for setting up Aurora DSQL clusters.

```go
import dsql "github.com/aws/aws-cdk-go/awscdkdsqlalpha"
```

## Creating an Aurora DSQL Cluster

To create an Aurora DSQL cluster, define a `Cluster`:

```go
cluster := dsql.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("my-dsql-cluster"),
	DeletionProtection: jsii.Boolean(true),
})
```

## Granting Connect Access

You can grant IAM principals the ability to connect to the cluster:

```go
role := iam.NewRole(this, jsii.String("DBRole"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(this.Account),
})
cluster := dsql.NewCluster(this, jsii.String("Cluster"))

// Use one of the following statements to grant the role the necessary permissions
cluster.grantConnect(role) // Grant the role dsql:DbConnect
cluster.grantConnectAdmin(role) // Grant the role dsql:DbConnectAdmin
cluster.grant(role, jsii.String("dsql:DbConnect"))
```
