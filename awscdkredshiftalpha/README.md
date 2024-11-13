# Amazon Redshift Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## Starting a Redshift Cluster Database

To set up a Redshift cluster, define a `Cluster`. It will be launched in a VPC.
You can specify a VPC, otherwise one will be created. The nodes are always launched in private subnets and are encrypted by default.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"))
cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
})
```

By default, the master password will be generated and stored in AWS Secrets Manager.
You can specify characters to not include in generated passwords by setting `excludeCharacters` property.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"))
cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
		ExcludeCharacters: jsii.String("\"@/\\ '`"),
	},
	Vpc: Vpc,
})
```

A default database named `default_db` will be created in the cluster. To change the name of this database set the `defaultDatabaseName` attribute in the constructor properties.

By default, the cluster will not be publicly accessible.
Depending on your use case, you can make the cluster publicly accessible with the `publiclyAccessible` property.

## Adding a logging bucket for database audit logging to S3

Amazon Redshift logs information about connections and user activities in your database. These logs help you to monitor the database for security and troubleshooting purposes, a process called database auditing. To send these logs to an S3 bucket, specify the `loggingProperties` when creating a new cluster.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"))
bucket := s3.Bucket_FromBucketName(this, jsii.String("bucket"), jsii.String("amzn-s3-demo-bucket"))

cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
	LoggingProperties: &LoggingProperties{
		LoggingBucket: bucket,
		LoggingKeyPrefix: jsii.String("prefix"),
	},
})
```

## Availability Zone Relocation

By using [relocation in Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-recovery.html), you allow Amazon Redshift to move a cluster to another Availability Zone (AZ) without any loss of data or changes to your applications.
This feature can be applied to both new and existing clusters.

To enable this feature, set the `availabilityZoneRelocation` property to `true`.

```go
// Example automatically generated from non-compiling source. May contain errors.
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var vpc iVpc


cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
	NodeType: nodeType_RA3_XLPLUS,
	AvailabilityZoneRelocation: jsii.Boolean(true),
})
```

**Note**: The `availabilityZoneRelocation` property is only available for RA3 node types.

## Connecting

To control who can access the cluster, use the `.connections` attribute. Redshift Clusters have
a default port, so you don't need to specify the port:

```go
cluster.Connections.AllowDefaultPortFromAnyIpv4(jsii.String("Open to the world"))
```

The endpoint to access your database cluster will be available as the `.clusterEndpoint` attribute:

```go
cluster.ClusterEndpoint.SocketAddress
```

## Database Resources

This module allows for the creation of non-CloudFormation database resources such as users
and tables. This allows you to manage identities, permissions, and stateful resources
within your Redshift cluster from your CDK application.

Because these resources are not available in CloudFormation, this library leverages
[custom
resources](https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html)
to manage them. In addition to the IAM permissions required to make Redshift service
calls, the execution role for the custom resource handler requires database credentials to
create resources within the cluster.

These database credentials can be supplied explicitly through the `adminUser` properties
of the various database resource constructs. Alternatively, the credentials can be
automatically pulled from the Redshift cluster's default administrator
credentials. However, this option is only available if the password for the credentials
was generated by the CDK application (ie., no value vas provided for [the `masterPassword`
property](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-redshift.Login.html#masterpasswordspan-classapi-icon-api-icon-experimental-titlethis-api-element-is-experimental-it-may-change-without-noticespan)
of
[`Cluster.masterUser`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-redshift.Cluster.html#masteruserspan-classapi-icon-api-icon-experimental-titlethis-api-element-is-experimental-it-may-change-without-noticespan)).

### Creating Users

Create a user within a Redshift cluster database by instantiating a `User` construct. This
will generate a username and password, store the credentials in a [AWS Secrets Manager
`Secret`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-secretsmanager.Secret.html),
and make a query to the Redshift cluster to create a new database user with the
credentials.

```go
awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

By default, the user credentials are encrypted with your AWS account's default Secrets
Manager encryption key. You can specify the encryption key used for this purpose by
supplying a key in the `encryptionKey` property.

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


encryptionKey := kms.NewKey(this, jsii.String("Key"))
awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	EncryptionKey: encryptionKey,
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

By default, a username is automatically generated from the user construct ID and its path
in the construct tree. You can specify a particular username by providing a value for the
`username` property. Usernames must be valid identifiers; see: [Names and
identifiers](https://docs.aws.amazon.com/redshift/latest/dg/r_names.html) in the *Amazon
Redshift Database Developer Guide*.

```go
awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Username: jsii.String("myuser"),
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

The user password is generated by AWS Secrets Manager using the default configuration
found in
[`secretsmanager.SecretStringGenerator`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-secretsmanager.SecretStringGenerator.html),
except with password length `30` and some SQL-incompliant characters excluded. The
plaintext for the password will never be present in the CDK application; instead, a
[CloudFormation Dynamic
Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html)
will be used wherever the password value is required.

You can specify characters to not include in generated passwords by setting `excludeCharacters` property.

```go
awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
	ExcludeCharacters: jsii.String("\"@/\\ '`"),
})
```

### Creating Tables

Create a table within a Redshift cluster database by instantiating a `Table`
construct. This will make a query to the Redshift cluster to create a new database table
with the supplied schema.

```go
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

Tables greater than v2.114.1 can have their table name changed, for versions <= v2.114.1, this would not be possible.
Therefore, changing of table names for <= v2.114.1 have been disabled.

```go
// Example automatically generated from non-compiling source. May contain errors.
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableName: jsii.String("oldTableName"),
	 // This value can be change for versions greater than v2.114.1
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

The table can be configured to have distStyle attribute and a distKey column:

```go
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
			DistKey: jsii.Boolean(true),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
	DistStyle: awscdkredshiftalpha.TableDistStyle_KEY,
})
```

The table can also be configured to have sortStyle attribute and sortKey columns:

```go
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
			SortKey: jsii.Boolean(true),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
			SortKey: jsii.Boolean(true),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
	SortStyle: awscdkredshiftalpha.TableSortStyle_COMPOUND,
})
```

Tables and their respective columns can be configured to contain comments:

```go
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
			Comment: jsii.String("This is a column comment"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
			Comment: jsii.String("This is a another column comment"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
	TableComment: jsii.String("This is a table comment"),
})
```

Table columns can be configured to use a specific compression encoding:

```go
import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"


awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
			Encoding: awscdkredshiftalpha.ColumnEncoding_TEXT32K,
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
			Encoding: awscdkredshiftalpha.ColumnEncoding_DELTA32K,
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

Table columns can also contain an `id` attribute, which can allow table columns to be renamed.

**NOTE** To use the `id` attribute, you must also enable the `@aws-cdk/aws-redshift:columnId` feature flag.

```go
awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Id: jsii.String("col1"),
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Id: jsii.String("col2"),
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
```

Query execution duration is limited to 1 minute by default. You can change this by setting the `timeout` property.

Valid timeout values are between 1 seconds and 15 minutes.

```go
import "github.com/aws/aws-cdk-go/awscdk"


awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Id: jsii.String("col1"),
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Id: jsii.String("col2"),
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
	Timeout: awscdk.Duration_Minutes(jsii.Number(15)),
})
```

### Granting Privileges

You can give a user privileges to perform certain actions on a table by using the
`Table.grant()` method.

```go
user := awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
table := awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})

table.grant(user, awscdkredshiftalpha.TableAction_DROP, awscdkredshiftalpha.TableAction_SELECT)
```

Take care when managing privileges via the CDK, as attempting to manage a user's
privileges on the same table in multiple CDK applications could lead to accidentally
overriding these permissions. Consider the following two CDK applications which both refer
to the same user and table. In application 1, the resources are created and the user is
given `INSERT` permissions on the table:

```go
databaseName := "databaseName"
username := "myuser"
tableName := "mytable"

user := awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Username: username,
	Cluster: cluster,
	DatabaseName: databaseName,
})
table := awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: databaseName,
})
table.grant(user, awscdkredshiftalpha.TableAction_INSERT)
```

In application 2, the resources are imported and the user is given `INSERT` permissions on
the table:

```go
databaseName := "databaseName"
username := "myuser"
tableName := "mytable"

user := awscdkredshiftalpha.User_FromUserAttributes(this, jsii.String("User"), &UserAttributes{
	Username: username,
	Password: awscdk.SecretValue_UnsafePlainText(jsii.String("NOT_FOR_PRODUCTION")),
	Cluster: cluster,
	DatabaseName: databaseName,
})
table := awscdkredshiftalpha.Table_FromTableAttributes(this, jsii.String("Table"), &TableAttributes{
	TableName: tableName,
	TableColumns: []column{
		&column{
			Name: jsii.String("col1"),
			DataType: jsii.String("varchar(4)"),
		},
		&column{
			Name: jsii.String("col2"),
			DataType: jsii.String("float"),
		},
	},
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
table.Grant(user, awscdkredshiftalpha.TableAction_INSERT)
```

Both applications attempt to grant the user the appropriate privilege on the table by
submitting a `GRANT USER` SQL query to the Redshift cluster. Note that the latter of these
two calls will have no effect since the user has already been granted the privilege.

Now, if application 1 were to remove the call to `grant`, a `REVOKE USER` SQL query is
submitted to the Redshift cluster. In general, application 1 does not know that
application 2 has also granted this permission and thus cannot decide not to issue the
revocation. This leads to the undesirable state where application 2 still contains the
call to `grant` but the user does not have the specified permission.

Note that this does not occur when duplicate privileges are granted within the same
application, as such privileges are de-duplicated before any SQL query is submitted.

## Rotating credentials

When the master password is generated and stored in AWS Secrets Manager, it can be rotated automatically:

```go
cluster.AddRotationSingleUser()
```

The multi user rotation scheme is also available:

```go
user := awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
	Cluster: cluster,
	DatabaseName: jsii.String("databaseName"),
})
cluster.AddRotationMultiUser(jsii.String("MultiUserRotation"), &RotationMultiUserOptions{
	Secret: user.Secret,
})
```

## Adding Parameters

You can add a parameter to a parameter group with`ClusterParameterGroup.addParameter()`.

```go
import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"


params := awscdkredshiftalpha.NewClusterParameterGroup(this, jsii.String("Params"), &ClusterParameterGroupProps{
	Description: jsii.String("desc"),
	Parameters: map[string]*string{
		"require_ssl": jsii.String("true"),
	},
})

params.AddParameter(jsii.String("enable_user_activity_logging"), jsii.String("true"))
```

Additionally, you can add a parameter to the cluster's associated parameter group with `Cluster.addToParameterGroup()`. If the cluster does not have an associated parameter group, a new parameter group is created.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
		MasterPassword: cdk.SecretValue_UnsafePlainText(jsii.String("tooshort")),
	},
	Vpc: Vpc,
})

cluster.AddToParameterGroup(jsii.String("enable_user_activity_logging"), jsii.String("true"))
```

## Rebooting for Parameter Updates

In most cases, existing clusters [must be manually rebooted](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html) to apply parameter changes. You can automate parameter related reboots by setting the cluster's `rebootForParameterChanges` property to `true` , or by using `Cluster.enableRebootForParameterChanges()`.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
		MasterPassword: cdk.SecretValue_UnsafePlainText(jsii.String("tooshort")),
	},
	Vpc: Vpc,
})

cluster.AddToParameterGroup(jsii.String("enable_user_activity_logging"), jsii.String("true"))
cluster.EnableRebootForParameterChanges()
```

## Elastic IP

If you configure your cluster to be publicly accessible, you can optionally select an *elastic IP address* to use for the external IP address. An elastic IP address is a static IP address that is associated with your AWS account. You can use an elastic IP address to connect to your cluster from outside the VPC. An elastic IP address gives you the ability to change your underlying configuration without affecting the IP address that clients use to connect to your cluster. This approach can be helpful for situations such as recovery after a failure.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
		MasterPassword: cdk.SecretValue_UnsafePlainText(jsii.String("tooshort")),
	},
	Vpc: Vpc,
	PubliclyAccessible: jsii.Boolean(true),
	ElasticIp: jsii.String("10.123.123.255"),
})
```

If the Cluster is in a VPC and you want to connect to it using the private IP address from within the cluster, it is important to enable *DNS resolution* and *DNS hostnames* in the VPC config. If these parameters would not be set, connections from within the VPC would connect to the elastic IP address and not the private IP address.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

vpc := ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
	EnableDnsSupport: jsii.Boolean(true),
	EnableDnsHostnames: jsii.Boolean(true),
})
```

Note that if there is already an existing, public accessible Cluster, which VPC configuration is changed to use *DNS hostnames* and *DNS resolution*, connections still use the elastic IP address until the cluster is resized.

### Elastic IP vs. Cluster node public IP

The elastic IP address is an external IP address for accessing the cluster outside of a VPC. It's not related to the cluster node public IP addresses and private IP addresses that are accessible via the `clusterEndpoint` property. The public and private cluster node IP addresses appear regardless of whether the cluster is publicly accessible or not. They are used only in certain circumstances to configure ingress rules on the remote host. These circumstances occur when you load data from an Amazon EC2 instance or other remote host using a Secure Shell (SSH) connection.

### Attach Elastic IP after Cluster creation

In some cases, you might want to associate the cluster with an elastic IP address or change an elastic IP address that is associated with the cluster. To attach an elastic IP address after the cluster is created, first update the cluster so that it is not publicly accessible, then make it both publicly accessible and add an Elastic IP address in the same operation.

## Enhanced VPC Routing

When you use Amazon Redshift enhanced VPC routing, Amazon Redshift forces all COPY and UNLOAD traffic between your cluster and your data repositories through your virtual private cloud (VPC) based on the Amazon VPC service. By using enhanced VPC routing, you can use standard VPC features, such as VPC security groups, network access control lists (ACLs), VPC endpoints, VPC endpoint policies, internet gateways, and Domain Name System (DNS) servers, as described in the Amazon VPC User Guide. You use these features to tightly manage the flow of data between your Amazon Redshift cluster and other resources. When you use enhanced VPC routing to route traffic through your VPC, you can also use VPC flow logs to monitor COPY and UNLOAD traffic.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
		MasterPassword: cdk.SecretValue_UnsafePlainText(jsii.String("tooshort")),
	},
	Vpc: Vpc,
	EnhancedVpcRouting: jsii.Boolean(true),
})
```

If enhanced VPC routing is not enabled, Amazon Redshift routes traffic through the internet, including traffic to other services within the AWS network.

## Default IAM role

Some Amazon Redshift features require Amazon Redshift to access other AWS services on your behalf. For your Amazon Redshift clusters to act on your behalf, you supply security credentials to your clusters. The preferred method to supply security credentials is to specify an AWS Identity and Access Management (IAM) role.

When you create an IAM role and set it as the default for the cluster using console, you don't have to provide the IAM role's Amazon Resource Name (ARN) to perform authentication and authorization.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


defaultRole := iam.NewRole(this, jsii.String("DefaultRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("redshift.amazonaws.com")),
})

awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
	Roles: []iRole{
		defaultRole,
	},
	DefaultRole: defaultRole,
})
```

A default role can also be added to a cluster using the `addDefaultIamRole` method.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


defaultRole := iam.NewRole(this, jsii.String("DefaultRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("redshift.amazonaws.com")),
})

redshiftCluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
	Roles: []iRole{
		defaultRole,
	},
})

redshiftCluster.AddDefaultIamRole(defaultRole)
```

## IAM roles

Attaching IAM roles to a Redshift Cluster grants permissions to the Redshift service to perform actions on your behalf.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("redshift.amazonaws.com")),
})
cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
	Roles: []iRole{
		role,
	},
})
```

Additional IAM roles can be attached to a cluster using the `addIamRole` method.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("redshift.amazonaws.com")),
})
cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
	MasterUser: &Login{
		MasterUsername: jsii.String("admin"),
	},
	Vpc: Vpc,
})
cluster.AddIamRole(role)
```

## Multi-AZ

Amazon Redshift supports [multiple Availability Zones (Multi-AZ) deployments]((https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-multi-az.html)) for provisioned RA3 clusters.
By using Multi-AZ deployments, your Amazon Redshift data warehouse can continue operating in failure scenarios when an unexpected event happens in an Availability Zone.

To create a Multi-AZ cluster, set the `multiAz` property to `true` when creating the cluster.

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc ec2.IVpc


redshift.NewCluster(stack, jsii.String("Cluster"), map[string]interface{}{
	"masterUser": map[string]*string{
		"masterUsername": jsii.String("admin"),
	},
	"vpc": vpc,
	 // 3 AZs are required for Multi-AZ
	"nodeType": redshift.NodeType_RA3_XLPLUS,
	 // must be RA3 node type
	"clusterType": redshift.ClusterType_MULTI_NODE,
	 // must be MULTI_NODE
	"numberOfNodes": jsii.Number(2),
	 // must be 2 or more
	"multiAz": jsii.Boolean(true),
})
```

## Resizing

As your data warehousing needs change, it's possible to resize your Redshift cluster. If the cluster was deployed via CDK,
it's important to resize it via CDK so the change is registered in the AWS CloudFormation template.
There are two types of resize operations:

* Elastic resize - Number of nodes and node type can be changed, but not at the same time. Elastic resize is the default behavior,
  as it's a fast operation and typically completes in minutes. Elastic resize is only supported on clusters of the following types:

  * dc1.large (if your cluster is in a VPC)
  * dc1.8xlarge (if your cluster is in a VPC)
  * dc2.large
  * dc2.8xlarge
  * ds2.xlarge
  * ds2.8xlarge
  * ra3.xlplus
  * ra3.4xlarge
  * ra3.16xlarge
* Classic resize - Number of nodes, node type, or both, can be changed. This operation takes longer to complete,
  but is useful when the resize operation doesn't meet the criteria of an elastic resize. If you prefer classic resizing,
  you can set the `classicResizing` flag when creating the cluster.

There are other constraints to be aware of, for example, elastic resizing does not support single-node clusters and there are
limits on the number of nodes you can add to a cluster. See the [AWS Redshift Documentation](https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-operations.html#rs-resize-tutorial) and [AWS API Documentation](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ResizeCluster.html) for more details.
