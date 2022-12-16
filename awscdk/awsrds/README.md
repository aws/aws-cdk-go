# Amazon Relational Database Service Construct Library

```go
import rds "github.com/aws/aws-cdk-go/awscdk"
```

## Starting a clustered database

To set up a clustered database (like Aurora), define a `DatabaseCluster`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
		version: rds.auroraMysqlEngineVersion_VER_2_08_1(),
	}),
	credentials: rds.credentials.fromGeneratedSecret(jsii.String("clusteradmin")),
	 // Optional - will default to 'admin' username and generated password
	instanceProps: &instanceProps{
		// optional , defaults to t3.medium
		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_SMALL),
		vpcSubnets: &subnetSelection{
			subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
		},
		vpc: vpc,
	},
})
```

If there isn't a constant for the exact version you want to use,
all of the `Version` classes have a static `of` method that can be used to create an arbitrary version.

```go
customEngineVersion := rds.auroraMysqlEngineVersion.of(jsii.String("5.7.mysql_aurora.2.08.1"))
```

By default, the master password will be generated and stored in AWS Secrets Manager with auto-generated description.

Your cluster will be empty by default. To add a default database upon construction, specify the
`defaultDatabaseName` attribute.

To use dual-stack mode, specify `NetworkType.DUAL` on the `networkType` property:

```go
var vpc vpc
// VPC and subnets must have IPv6 CIDR blocks
cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
		version: rds.auroraMysqlEngineVersion_VER_3_02_1(),
	}),
	instanceProps: &instanceProps{
		vpc: vpc,
		publiclyAccessible: jsii.Boolean(false),
	},
	networkType: rds.networkType_DUAL,
})
```

For more information about dual-stack mode, see [Working with a DB cluster in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html).

Use `DatabaseClusterFromSnapshot` to create a cluster from a snapshot:

```go
var vpc vpc

rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &databaseClusterFromSnapshotProps{
	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
		version: rds.auroraEngineVersion_VER_1_22_2(),
	}),
	instanceProps: &instanceProps{
		vpc: vpc,
	},
	snapshotIdentifier: jsii.String("mySnapshot"),
})
```

### Updating the database instances in a cluster

Database cluster instances may be updated in bulk or on a rolling basis.

An update to all instances in a cluster may cause significant downtime. To reduce the downtime, set the `instanceUpdateBehavior` property in `DatabaseClusterBaseProps` to `InstanceUpdateBehavior.ROLLING`. This adds a dependency between each instance so the update is performed on only one instance at a time.

Use `InstanceUpdateBehavior.BULK` to update all instances at once.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
		version: rds.auroraMysqlEngineVersion_VER_3_01_0(),
	}),
	instances: jsii.Number(2),
	instanceProps: &instanceProps{
		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_SMALL),
		vpc: vpc,
	},
	instanceUpdateBehaviour: rds.instanceUpdateBehaviour_ROLLING,
})
```

## Starting an instance database

To set up an instance database, define a `DatabaseInstance`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	// optional, defaults to m5.large
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_SMALL),
	credentials: rds.credentials.fromGeneratedSecret(jsii.String("syscdk")),
	 // Optional - will default to 'admin' username and generated password
	vpc: vpc,
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
	},
})
```

If there isn't a constant for the exact engine version you want to use,
all of the `Version` classes have a static `of` method that can be used to create an arbitrary version.

```go
customEngineVersion := rds.oracleEngineVersion.of(jsii.String("19.0.0.0.ru-2020-04.rur-2020-04.r1"), jsii.String("19"))
```

By default, the master password will be generated and stored in AWS Secrets Manager.

To use the storage auto scaling option of RDS you can specify the maximum allocated storage.
This is the upper limit to which RDS can automatically scale the storage. More info can be found
[here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
Example for max storage configuration:

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
		version: rds.postgresEngineVersion_VER_12_3(),
	}),
	// optional, defaults to m5.large
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_SMALL),
	vpc: vpc,
	maxAllocatedStorage: jsii.Number(200),
})
```

To use dual-stack mode, specify `NetworkType.DUAL` on the `networkType` property:

```go
var vpc vpc
// VPC and subnets must have IPv6 CIDR blocks
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
		version: rds.postgresEngineVersion_VER_14_4(),
	}),
	vpc: vpc,
	networkType: rds.networkType_DUAL,
	publiclyAccessible: jsii.Boolean(false),
})
```

For more information about dual-stack mode, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html).

Use `DatabaseInstanceFromSnapshot` and `DatabaseInstanceReadReplica` to create an instance from snapshot or
a source database respectively:

```go
var vpc vpc

var sourceInstance databaseInstance

rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &databaseInstanceFromSnapshotProps{
	snapshotIdentifier: jsii.String("my-snapshot"),
	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
		version: rds.postgresEngineVersion_VER_12_3(),
	}),
	// optional, defaults to m5.large
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_LARGE),
	vpc: vpc,
})
rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &databaseInstanceReadReplicaProps{
	sourceDatabaseInstance: sourceInstance,
	instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE2, ec2.*instanceSize_LARGE),
	vpc: vpc,
})
```

Automatic backups of read replica instances are only supported for MySQL and MariaDB. By default,
automatic backups are disabled for read replicas and can only be enabled (using `backupRetention`)
if also enabled on the source instance.

Creating a "production" Oracle database instance with option and parameter groups:

```go
// Set open cursors with parameter group
parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	parameters: map[string]*string{
		"open_cursors": jsii.String("2500"),
	},
})

optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &optionGroupProps{
	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	configurations: []optionConfiguration{
		&optionConfiguration{
			name: jsii.String("LOCATOR"),
		},
		&optionConfiguration{
			name: jsii.String("OEM"),
			port: jsii.Number(1158),
			vpc: vpc,
		},
	},
})

// Allow connections to OEM
optionGroup.optionConnections.oEM.connections.allowDefaultPortFromAnyIpv4()

// Database instance with production values
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	licenseModel: rds.licenseModel_BRING_YOUR_OWN_LICENSE,
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_MEDIUM),
	multiAz: jsii.Boolean(true),
	storageType: rds.storageType_IO1,
	credentials: rds.credentials.fromUsername(jsii.String("syscdk")),
	vpc: vpc,
	databaseName: jsii.String("ORCL"),
	storageEncrypted: jsii.Boolean(true),
	backupRetention: cdk.duration.days(jsii.Number(7)),
	monitoringInterval: cdk.*duration.seconds(jsii.Number(60)),
	enablePerformanceInsights: jsii.Boolean(true),
	cloudwatchLogsExports: []*string{
		jsii.String("trace"),
		jsii.String("audit"),
		jsii.String("alert"),
		jsii.String("listener"),
	},
	cloudwatchLogsRetention: logs.retentionDays_ONE_MONTH,
	autoMinorVersionUpgrade: jsii.Boolean(true),
	 // required to be true if LOCATOR is used in the option group
	optionGroup: optionGroup,
	parameterGroup: parameterGroup,
	removalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Allow connections on default port from any IPV4
instance.connections.allowDefaultPortFromAnyIpv4()

// Rotate the master user password every 30 days
instance.addRotationSingleUser()

// Add alarm for high CPU
// Add alarm for high CPU
cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &alarmProps{
	metric: instance.metricCPUUtilization(),
	threshold: jsii.Number(90),
	evaluationPeriods: jsii.Number(1),
})

// Trigger Lambda function on instance availability events
fn := lambda.NewFunction(this, jsii.String("Function"), &functionProps{
	code: lambda.code.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	handler: jsii.String("index.handler"),
	runtime: lambda.runtime_NODEJS_14_X(),
})

availabilityRule := instance.onEvent(jsii.String("Availability"), &onEventOptions{
	target: targets.NewLambdaFunction(fn),
})
availabilityRule.addEventPattern(&eventPattern{
	detail: map[string]interface{}{
		"EventCategories": []interface{}{
			jsii.String("availability"),
		},
	},
})
```

Add XMLDB and OEM with option group

```go
// Set open cursors with parameter group
parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	parameters: map[string]*string{
		"open_cursors": jsii.String("2500"),
	},
})

optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &optionGroupProps{
	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	configurations: []optionConfiguration{
		&optionConfiguration{
			name: jsii.String("LOCATOR"),
		},
		&optionConfiguration{
			name: jsii.String("OEM"),
			port: jsii.Number(1158),
			vpc: vpc,
		},
	},
})

// Allow connections to OEM
optionGroup.optionConnections.oEM.connections.allowDefaultPortFromAnyIpv4()

// Database instance with production values
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	licenseModel: rds.licenseModel_BRING_YOUR_OWN_LICENSE,
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_MEDIUM),
	multiAz: jsii.Boolean(true),
	storageType: rds.storageType_IO1,
	credentials: rds.credentials.fromUsername(jsii.String("syscdk")),
	vpc: vpc,
	databaseName: jsii.String("ORCL"),
	storageEncrypted: jsii.Boolean(true),
	backupRetention: cdk.duration.days(jsii.Number(7)),
	monitoringInterval: cdk.*duration.seconds(jsii.Number(60)),
	enablePerformanceInsights: jsii.Boolean(true),
	cloudwatchLogsExports: []*string{
		jsii.String("trace"),
		jsii.String("audit"),
		jsii.String("alert"),
		jsii.String("listener"),
	},
	cloudwatchLogsRetention: logs.retentionDays_ONE_MONTH,
	autoMinorVersionUpgrade: jsii.Boolean(true),
	 // required to be true if LOCATOR is used in the option group
	optionGroup: optionGroup,
	parameterGroup: parameterGroup,
	removalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Allow connections on default port from any IPV4
instance.connections.allowDefaultPortFromAnyIpv4()

// Rotate the master user password every 30 days
instance.addRotationSingleUser()

// Add alarm for high CPU
// Add alarm for high CPU
cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &alarmProps{
	metric: instance.metricCPUUtilization(),
	threshold: jsii.Number(90),
	evaluationPeriods: jsii.Number(1),
})

// Trigger Lambda function on instance availability events
fn := lambda.NewFunction(this, jsii.String("Function"), &functionProps{
	code: lambda.code.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	handler: jsii.String("index.handler"),
	runtime: lambda.runtime_NODEJS_14_X(),
})

availabilityRule := instance.onEvent(jsii.String("Availability"), &onEventOptions{
	target: targets.NewLambdaFunction(fn),
})
availabilityRule.addEventPattern(&eventPattern{
	detail: map[string]interface{}{
		"EventCategories": []interface{}{
			jsii.String("availability"),
		},
	},
})
```

Use the `storageType` property to specify the [type of storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)
to use for the instance:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
		version: mysqlEngineVersion_VER_8_0_30,
	}),
	vpc: vpc,
	storageType: rds.storageType_IO1,
	iops: jsii.Number(5000),
})

gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &databaseInstanceProps{
	engine: rds.*databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
		version: *mysqlEngineVersion_VER_8_0_30,
	}),
	vpc: vpc,
	allocatedStorage: jsii.Number(500),
	storageType: rds.*storageType_GP3,
	storageThroughput: jsii.Number(500),
})
```

## Setting Public Accessibility

You can set public accessibility for the database instance or cluster using the `publiclyAccessible` property.
If you specify `true`, it creates an instance with a publicly resolvable DNS name, which resolves to a public IP address.
If you specify `false`, it creates an internal instance with a DNS name that resolves to a private IP address.
The default value depends on `vpcSubnets`.
It will be `true` if `vpcSubnets` is `subnetType: SubnetType.PUBLIC`, `false` otherwise.

```go
var vpc vpc

// Setting public accessibility for DB instance
// Setting public accessibility for DB instance
rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
		version: rds.mysqlEngineVersion_VER_8_0_19(),
	}),
	vpc: vpc,
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
	},
	publiclyAccessible: jsii.Boolean(true),
})

// Setting public accessibility for DB cluster
// Setting public accessibility for DB cluster
rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &databaseClusterProps{
	engine: rds.databaseClusterEngine_AURORA(),
	instanceProps: &instanceProps{
		vpc: vpc,
		vpcSubnets: &subnetSelection{
			subnetType: ec2.*subnetType_PRIVATE_WITH_EGRESS,
		},
		publiclyAccessible: jsii.Boolean(true),
	},
})
```

## Instance events

To define Amazon CloudWatch event rules for database instances, use the `onEvent`
method:

```go
var instance databaseInstance
var fn function

rule := instance.onEvent(jsii.String("InstanceEvent"), &onEventOptions{
	target: targets.NewLambdaFunction(fn),
})
```

## Login credentials

By default, database instances and clusters (with the exception of `DatabaseInstanceFromSnapshot` and `ServerlessClusterFromSnapshot`) will have `admin` user with an auto-generated password.
An alternative username (and password) may be specified for the admin user instead of the default.

The following examples use a `DatabaseInstance`, but the same usage is applicable to `DatabaseCluster`.

```go
var vpc vpc

engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
	version: rds.postgresEngineVersion_VER_12_3(),
})
rds.NewDatabaseInstance(this, jsii.String("InstanceWithUsername"), &databaseInstanceProps{
	engine: engine,
	vpc: vpc,
	credentials: rds.credentials.fromGeneratedSecret(jsii.String("postgres")),
})

rds.NewDatabaseInstance(this, jsii.String("InstanceWithUsernameAndPassword"), &databaseInstanceProps{
	engine: engine,
	vpc: vpc,
	credentials: rds.*credentials.fromPassword(jsii.String("postgres"), awscdk.SecretValue.ssmSecure(jsii.String("/dbPassword"), jsii.String("1"))),
})

mySecret := secretsmanager.secret.fromSecretName(this, jsii.String("DBSecret"), jsii.String("myDBLoginInfo"))
rds.NewDatabaseInstance(this, jsii.String("InstanceWithSecretLogin"), &databaseInstanceProps{
	engine: engine,
	vpc: vpc,
	credentials: rds.*credentials.fromSecret(mySecret),
})
```

Secrets generated by `fromGeneratedSecret()` can be customized:

```go
var vpc vpc

engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
	version: rds.postgresEngineVersion_VER_12_3(),
})
myKey := kms.NewKey(this, jsii.String("MyKey"))

rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &databaseInstanceProps{
	engine: engine,
	vpc: vpc,
	credentials: rds.credentials.fromGeneratedSecret(jsii.String("postgres"), &credentialsBaseOptions{
		secretName: jsii.String("my-cool-name"),
		encryptionKey: myKey,
		excludeCharacters: jsii.String("!&*^#@()"),
		replicaRegions: []replicaRegion{
			&replicaRegion{
				region: jsii.String("eu-west-1"),
			},
			&replicaRegion{
				region: jsii.String("eu-west-2"),
			},
		},
	}),
})
```

### Snapshot credentials

As noted above, Databases created with `DatabaseInstanceFromSnapshot` or `ServerlessClusterFromSnapshot` will not create user and auto-generated password by default because it's not possible to change the master username for a snapshot. Instead, they will use the existing username and password from the snapshot. You can still generate a new password - to generate a secret similarly to the other constructs, pass in credentials with `fromGeneratedSecret()` or `fromGeneratedPassword()`.

```go
var vpc vpc

engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
	version: rds.postgresEngineVersion_VER_12_3(),
})
myKey := kms.NewKey(this, jsii.String("MyKey"))

rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("InstanceFromSnapshotWithCustomizedSecret"), &databaseInstanceFromSnapshotProps{
	engine: engine,
	vpc: vpc,
	snapshotIdentifier: jsii.String("mySnapshot"),
	credentials: rds.snapshotCredentials.fromGeneratedSecret(jsii.String("username"), &snapshotCredentialsFromGeneratedPasswordOptions{
		encryptionKey: myKey,
		excludeCharacters: jsii.String("!&*^#@()"),
		replicaRegions: []replicaRegion{
			&replicaRegion{
				region: jsii.String("eu-west-1"),
			},
			&replicaRegion{
				region: jsii.String("eu-west-2"),
			},
		},
	}),
})
```

## Connecting

To control who can access the cluster or instance, use the `.connections` attribute. RDS databases have
a default port, so you don't need to specify the port:

```go
var cluster databaseCluster

cluster.connections.allowFromAnyIpv4(ec2.port.allTraffic(), jsii.String("Open to the world"))
```

The endpoints to access your database cluster will be available as the `.clusterEndpoint` and `.readerEndpoint`
attributes:

```go
var cluster databaseCluster

writeAddress := cluster.clusterEndpoint.socketAddress
```

For an instance database:

```go
var instance databaseInstance

address := instance.instanceEndpoint.socketAddress
```

## Rotating credentials

When the master password is generated and stored in AWS Secrets Manager, it can be rotated automatically:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"

var instance databaseInstance
var mySecurityGroup securityGroup

instance.addRotationSingleUser(&rotationSingleUserOptions{
	automaticallyAfter: cdk.duration.days(jsii.Number(7)),
	 // defaults to 30 days
	excludeCharacters: jsii.String("!@#$%^&*"),
	 // defaults to the set " %+~`#/// here*()|[]{}:;<>?!'/@\"\\"
	securityGroup: mySecurityGroup,
})
```

```go
cluster := rds.NewDatabaseCluster(stack, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine_AURORA(),
	instanceProps: &instanceProps{
		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_SMALL),
		vpc: vpc,
	},
})

cluster.addRotationSingleUser()

clusterWithCustomRotationOptions := rds.NewDatabaseCluster(stack, jsii.String("CustomRotationOptions"), &databaseClusterProps{
	engine: rds.*databaseClusterEngine_AURORA(),
	instanceProps: &instanceProps{
		instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE3, ec2.*instanceSize_SMALL),
		vpc: vpc,
	},
})
clusterWithCustomRotationOptions.addRotationSingleUser(&rotationSingleUserOptions{
	automaticallyAfter: cdk.duration.days(jsii.Number(7)),
	excludeCharacters: jsii.String("!@#$%^&*"),
	securityGroup: securityGroup,
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
	},
	endpoint: endpoint,
})
```

The multi user rotation scheme is also available:

```go
var instance databaseInstance
var myImportedSecret databaseSecret

instance.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
	secret: myImportedSecret,
})
```

It's also possible to create user credentials together with the instance/cluster and add rotation:

```go
var instance databaseInstance

myUserSecret := rds.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
	username: jsii.String("myuser"),
	secretName: jsii.String("my-user-secret"),
	 // optional, defaults to a CloudFormation-generated name
	masterSecret: instance.secret,
	excludeCharacters: jsii.String("{}[]()'\"/\\"),
})
myUserSecretAttached := myUserSecret.attach(instance) // Adds DB connections information in the secret

instance.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
	 // Add rotation using the multi user scheme
	secret: myUserSecretAttached,
})
```

**Note**: This user must be created manually in the database using the master credentials.
The rotation will start as soon as this user exists.

Access to the Secrets Manager API is required for the secret rotation. This can be achieved either with
internet connectivity (through NAT) or with a VPC interface endpoint. By default, the rotation Lambda function
is deployed in the same subnets as the instance/cluster. If access to the Secrets Manager API is not possible from
those subnets or using the default API endpoint, use the `vpcSubnets` and/or `endpoint` options:

```go
var instance databaseInstance
var myEndpoint interfaceVpcEndpoint


instance.addRotationSingleUser(&rotationSingleUserOptions{
	vpcSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
	},
	 // Place rotation Lambda in private subnets
	endpoint: myEndpoint,
})
```

See also [@aws-cdk/aws-secretsmanager](https://github.com/aws/aws-cdk/blob/main/packages/%40aws-cdk/aws-secretsmanager/README.md) for credentials rotation of existing clusters/instances.

## IAM Authentication

You can also authenticate to a database instance using AWS Identity and Access Management (IAM) database authentication;
See [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) for more information
and a list of supported versions and limitations.

**Note**: `grantConnect()` does not currently work - see [this GitHub issue](https://github.com/aws/aws-cdk/issues/11851).

The following example shows enabling IAM authentication for a database instance and granting connection access to an IAM role.

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
		version: rds.mysqlEngineVersion_VER_8_0_19(),
	}),
	vpc: vpc,
	iamAuthentication: jsii.Boolean(true),
})
role := iam.NewRole(this, jsii.String("DBRole"), &roleProps{
	assumedBy: iam.NewAccountPrincipal(this.account),
})
instance.grantConnect(role)
```

The following example shows granting connection access for RDS Proxy to an IAM role.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine_AURORA(),
	instanceProps: &instanceProps{
		vpc: vpc,
	},
})

proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &databaseProxyProps{
	proxyTarget: rds.proxyTarget.fromCluster(cluster),
	secrets: []iSecret{
		cluster.secret,
	},
	vpc: vpc,
})

role := iam.NewRole(this, jsii.String("DBProxyRole"), &roleProps{
	assumedBy: iam.NewAccountPrincipal(this.account),
})
proxy.grantConnect(role, jsii.String("admin"))
```

**Note**: In addition to the setup above, a database user will need to be created to support IAM auth.
See [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.DBAccounts.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.DBAccounts.html) for setup instructions.

## Kerberos Authentication

You can also authenticate using Kerberos to a database instance using AWS Managed Microsoft AD for authentication;
See [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) for more information
and a list of supported versions and limitations.

The following example shows enabling domain support for a database instance and creating an IAM role to access
Directory Services.

```go
var vpc vpc

role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
	managedPolicies: []iManagedPolicy{
		iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
	},
})
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
		version: rds.mysqlEngineVersion_VER_8_0_19(),
	}),
	vpc: vpc,
	domain: jsii.String("d-????????"),
	 // The ID of the domain for the instance to join.
	domainRole: role,
})
```

**Note**: In addition to the setup above, you need to make sure that the database instance has network connectivity
to the domain controllers. This includes enabling cross-VPC traffic if in a different VPC and setting up the
appropriate security groups/network ACL to allow traffic between the database instance and domain controllers.
Once configured, see [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) for details
on configuring users for each available database engine.

## Metrics

Database instances and clusters both expose metrics (`cloudwatch.Metric`):

```go
// The number of database connections in use (average over 5 minutes)
var instance databaseInstance

// Average CPU utilization over 5 minutes
var cluster databaseCluster

dbConnections := instance.metricDatabaseConnections()
cpuUtilization := cluster.metricCPUUtilization()

// The average amount of time taken per disk I/O operation (average over 1 minute)
readLatency := instance.metric(jsii.String("ReadLatency"), &metricOptions{
	statistic: jsii.String("Average"),
	period: awscdk.Duration.seconds(jsii.Number(60)),
})
```

## Enabling S3 integration

Data in S3 buckets can be imported to and exported from certain database engines using SQL queries. To enable this
functionality, set the `s3ImportBuckets` and `s3ExportBuckets` properties for import and export respectively. When
configured, the CDK automatically creates and configures IAM roles as required.
Additionally, the `s3ImportRole` and `s3ExportRole` properties can be used to set this role directly.

You can read more about loading data to (or from) S3 here:

* Aurora MySQL - [import](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.LoadFromS3.html)
  and [export](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.SaveIntoS3.html).
* Aurora PostgreSQL - [import](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html#USER_PostgreSQL.S3Import)
  and [export](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html).
* Microsoft SQL Server - [import and export](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.html)
* PostgreSQL - [import](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html)
  and [export](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/postgresql-s3-export.html)
* Oracle - [import and export](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html)

The following snippet sets up a database cluster with different S3 buckets where the data is imported and exported -

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

importBucket := s3.NewBucket(this, jsii.String("importbucket"))
exportBucket := s3.NewBucket(this, jsii.String("exportbucket"))
rds.NewDatabaseCluster(this, jsii.String("dbcluster"), &databaseClusterProps{
	engine: rds.databaseClusterEngine_AURORA(),
	instanceProps: &instanceProps{
		vpc: vpc,
	},
	s3ImportBuckets: []iBucket{
		importBucket,
	},
	s3ExportBuckets: []*iBucket{
		exportBucket,
	},
})
```

## Creating a Database Proxy

Amazon RDS Proxy sits between your application and your relational database to efficiently manage
connections to the database and improve scalability of the application. Learn more about at [Amazon RDS Proxy](https://aws.amazon.com/rds/proxy/)

The following code configures an RDS Proxy for a `DatabaseInstance`.

```go
var vpc vpc
var securityGroup securityGroup
var secrets []secret
var dbInstance databaseInstance


proxy := dbInstance.addProxy(jsii.String("proxy"), &databaseProxyOptions{
	borrowTimeout: awscdk.Duration.seconds(jsii.Number(30)),
	maxConnectionsPercent: jsii.Number(50),
	secrets: secrets,
	vpc: vpc,
})
```

## Exporting Logs

You can publish database logs to Amazon CloudWatch Logs. With CloudWatch Logs, you can perform real-time analysis of the log data,
store the data in highly durable storage, and manage the data with the CloudWatch Logs Agent. This is available for both database
instances and clusters; the types of logs available depend on the database type and engine being used.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"
var myLogsPublishingRole role
var vpc vpc


// Exporting logs from a cluster
cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
		version: rds.auroraEngineVersion_VER_1_17_9(),
	}),
	instanceProps: &instanceProps{
		vpc: vpc,
	},
	cloudwatchLogsExports: []*string{
		jsii.String("error"),
		jsii.String("general"),
		jsii.String("slowquery"),
		jsii.String("audit"),
	},
	 // Export all available MySQL-based logs
	cloudwatchLogsRetention: logs.retentionDays_THREE_MONTHS,
	 // Optional - default is to never expire logs
	cloudwatchLogsRetentionRole: myLogsPublishingRole,
})

// Exporting logs from an instance
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
		version: rds.postgresEngineVersion_VER_12_3(),
	}),
	vpc: vpc,
	cloudwatchLogsExports: []*string{
		jsii.String("postgresql"),
	},
})
```

## Option Groups

Some DB engines offer additional features that make it easier to manage data and databases, and to provide additional security for your database.
Amazon RDS uses option groups to enable and configure these features. An option group can specify features, called options,
that are available for a particular Amazon RDS DB instance.

```go
var vpc vpc
var securityGroup securityGroup


rds.NewOptionGroup(this, jsii.String("Options"), &optionGroupProps{
	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
		version: rds.oracleEngineVersion_VER_19(),
	}),
	configurations: []optionConfiguration{
		&optionConfiguration{
			name: jsii.String("OEM"),
			port: jsii.Number(5500),
			vpc: vpc,
			securityGroups: []iSecurityGroup{
				securityGroup,
			},
		},
	},
})
```

## Parameter Groups

Database parameters specify how the database is configured.
For example, database parameters can specify the amount of resources, such as memory, to allocate to a database.
You manage your database configuration by associating your DB instances with parameter groups.
Amazon RDS defines parameter groups with default settings.

You can create your own parameter group for your cluster or instance and associate it with your database:

```go
var vpc vpc


parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
	engine: rds.databaseInstanceEngine.sqlServerEe(&sqlServerEeInstanceEngineProps{
		version: rds.sqlServerEngineVersion_VER_11(),
	}),
	parameters: map[string]*string{
		"locks": jsii.String("100"),
	},
})

rds.NewDatabaseInstance(this, jsii.String("Database"), &databaseInstanceProps{
	engine: rds.*databaseInstanceEngine_SQL_SERVER_EE(),
	vpc: vpc,
	parameterGroup: parameterGroup,
})
```

Another way to specify parameters is to use the inline field `parameters` that creates an RDS parameter group for you.
You can use this if you do not want to reuse the parameter group instance for different instances:

```go
var vpc vpc


rds.NewDatabaseInstance(this, jsii.String("Database"), &databaseInstanceProps{
	engine: rds.databaseInstanceEngine.sqlServerEe(&sqlServerEeInstanceEngineProps{
		version: rds.sqlServerEngineVersion_VER_11(),
	}),
	vpc: vpc,
	parameters: map[string]*string{
		"locks": jsii.String("100"),
	},
})
```

You cannot specify a parameter map and a parameter group at the same time.

## Serverless

[Amazon Aurora Serverless](https://aws.amazon.com/rds/aurora/serverless/) is an on-demand, auto-scaling configuration for Amazon
Aurora. The database will automatically start up, shut down, and scale capacity
up or down based on your application's needs. It enables you to run your database
in the cloud without managing any database instances.

The following example initializes an Aurora Serverless PostgreSql cluster.
Aurora Serverless clusters can specify scaling properties which will be used to
automatically scale the database cluster seamlessly based on the workload.

```go
var vpc vpc


cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &serverlessClusterProps{
	engine: rds.databaseClusterEngine_AURORA_POSTGRESQL(),
	copyTagsToSnapshot: jsii.Boolean(true),
	 // whether to save the cluster tags when creating the snapshot. Default is 'true'
	parameterGroup: rds.parameterGroup.fromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
	vpc: vpc,
	scaling: &serverlessScalingOptions{
		autoPause: awscdk.Duration.minutes(jsii.Number(10)),
		 // default is to pause after 5 minutes of idle time
		minCapacity: rds.auroraCapacityUnit_ACU_8,
		 // default is 2 Aurora capacity units (ACUs)
		maxCapacity: rds.*auroraCapacityUnit_ACU_32,
	},
})
```

Aurora Serverless Clusters do not support the following features:

* Loading data from an Amazon S3 bucket
* Saving data to an Amazon S3 bucket
* Invoking an AWS Lambda function with an Aurora MySQL native function
* Aurora replicas
* Backtracking
* Multi-master clusters
* Database cloning
* IAM database cloning
* IAM database authentication
* Restoring a snapshot from MySQL DB instance
* Performance Insights
* RDS Proxy

Read more about the [limitations of Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)

Learn more about using Amazon Aurora Serverless by reading the [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)

Use `ServerlessClusterFromSnapshot` to create a serverless cluster from a snapshot:

```go
var vpc vpc

rds.NewServerlessClusterFromSnapshot(this, jsii.String("Cluster"), &serverlessClusterFromSnapshotProps{
	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
	vpc: vpc,
	snapshotIdentifier: jsii.String("mySnapshot"),
})
```

### Data API

You can access your Aurora Serverless DB cluster using the built-in Data API. The Data API doesn't require a persistent connection to the DB cluster. Instead, it provides a secure HTTP endpoint and integration with AWS SDKs.

The following example shows granting Data API access to a Lamba function.

```go
var vpc vpc

var code code


cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &serverlessClusterProps{
	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
	vpc: vpc,
	 // this parameter is optional for serverless Clusters
	enableDataApi: jsii.Boolean(true),
})
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: code,
	environment: map[string]*string{
		"CLUSTER_ARN": cluster.clusterArn,
		"SECRET_ARN": cluster.secret.secretArn,
	},
})
cluster.grantDataApiAccess(fn)
```

**Note**: To invoke the Data API, the resource will need to read the secret associated with the cluster.

To learn more about using the Data API, see the [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html).

### Default VPC

The `vpc` parameter is optional.

If not provided, the cluster will be created in the default VPC of the account and region.
As this VPC is not deployed with AWS CDK, you can't configure the `vpcSubnets`, `subnetGroup` or `securityGroups` of the Aurora Serverless Cluster.
If you want to provide one of `vpcSubnets`, `subnetGroup` or `securityGroups` parameter, please provide a `vpc`.
