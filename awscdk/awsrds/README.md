# Amazon Relational Database Service Construct Library

```go
import rds "github.com/aws/aws-cdk-go/awscdk"
```

## Starting a clustered database

To set up a clustered database (like Aurora), define a `DatabaseCluster`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

You must specify the instance to use as the writer, along with an optional list
of readers (up to 15).

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
	 // Optional - will default to 'admin' username and generated password
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
		PubliclyAccessible: jsii.Boolean(false),
	}),
	Readers: []iClusterInstance{
		rds.ClusterInstance_*Provisioned(jsii.String("reader1"), &ProvisionedClusterInstanceProps{
			PromotionTier: jsii.Number(1),
		}),
		rds.ClusterInstance_ServerlessV2(jsii.String("reader2")),
	},
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	Vpc: Vpc,
})
```

To adopt Aurora I/O-Optimized, specify `DBClusterStorageType.AURORA_IOPT1` on the `storageType` property.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
		Version: rds.AuroraPostgresEngineVersion_VER_15_2(),
	}),
	Credentials: rds.Credentials_FromUsername(jsii.String("adminuser"), &CredentialsFromUsernameOptions{
		Password: awscdk.SecretValue_UnsafePlainText(jsii.String("7959866cacc02c2d243ecfe177464fe6")),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
		PubliclyAccessible: jsii.Boolean(false),
	}),
	Readers: []iClusterInstance{
		rds.ClusterInstance_*Provisioned(jsii.String("reader")),
	},
	StorageType: rds.DBClusterStorageType_AURORA_IOPT1,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	Vpc: Vpc,
})
```

If there isn't a constant for the exact version you want to use,
all of the `Version` classes have a static `of` method that can be used to create an arbitrary version.

```go
customEngineVersion := rds.AuroraMysqlEngineVersion_Of(jsii.String("5.7.mysql_aurora.2.08.1"))
```

By default, the master password will be generated and stored in AWS Secrets Manager with auto-generated description.

Your cluster will be empty by default. To add a default database upon construction, specify the
`defaultDatabaseName` attribute.

To use dual-stack mode, specify `NetworkType.DUAL` on the `networkType` property:

```go
var vpc vpc
// VPC and subnets must have IPv6 CIDR blocks
cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_02_1(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
		PubliclyAccessible: jsii.Boolean(false),
	}),
	Vpc: Vpc,
	NetworkType: rds.NetworkType_DUAL,
})
```

For more information about dual-stack mode, see [Working with a DB cluster in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html).

Use `DatabaseClusterFromSnapshot` to create a cluster from a snapshot:

```go
var vpc vpc

rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &DatabaseClusterFromSnapshotProps{
	Engine: rds.DatabaseClusterEngine_Aurora(&AuroraClusterEngineProps{
		Version: rds.AuroraEngineVersion_VER_1_22_2(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
	SnapshotIdentifier: jsii.String("mySnapshot"),
})
```

### Updating the database instances in a cluster

Database cluster instances may be updated in bulk or on a rolling basis.

An update to all instances in a cluster may cause significant downtime. To reduce the downtime, set the
`instanceUpdateBehavior` property in `DatabaseClusterBaseProps` to `InstanceUpdateBehavior.ROLLING`.
This adds a dependency between each instance so the update is performed on only one instance at a time.

Use `InstanceUpdateBehavior.BULK` to update all instances at once.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("Instance"), &ProvisionedClusterInstanceProps{
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
	}),
	Readers: []iClusterInstance{
		rds.ClusterInstance_*Provisioned(jsii.String("reader")),
	},
	InstanceUpdateBehaviour: rds.InstanceUpdateBehaviour_ROLLING,
	 // Optional - defaults to rds.InstanceUpdateBehaviour.BULK
	Vpc: Vpc,
})
```

### Serverless V2 instances in a Cluster

It is possible to create an RDS cluster with *both* serverlessV2 and provisioned
instances. For example, this will create a cluster with a provisioned writer and
a serverless v2 reader.

> *Note* Before getting starting with this type of cluster it is
> highly recommended that you read through the [Developer Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html)
> which goes into much more detail on the things you need to take into
> consideration.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Readers: []iClusterInstance{
		rds.ClusterInstance_ServerlessV2(jsii.String("reader")),
	},
	Vpc: Vpc,
})
```

### Monitoring

There are some CloudWatch metrics that are [important for Aurora Serverless
v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.viewing.monitoring).

* `ServerlessDatabaseCapacity`: An instance-level metric that can also be
  evaluated at the cluster level. At the cluster-level it represents the average
  capacity of all the instances in the cluster.
* `ACUUtilization`: Value of the `ServerlessDatabaseCapacity`/ max ACU of the
  cluster.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Readers: []iClusterInstance{
		rds.ClusterInstance_ServerlessV2(jsii.String("reader")),
	},
	Vpc: Vpc,
})

cluster.metricServerlessDatabaseCapacity(&MetricOptions{
	Period: awscdk.Duration_Minutes(jsii.Number(10)),
}).CreateAlarm(this, jsii.String("capacity"), &CreateAlarmOptions{
	Threshold: jsii.Number(1.5),
	EvaluationPeriods: jsii.Number(3),
})
cluster.metricACUUtilization(&MetricOptions{
	Period: awscdk.Duration_*Minutes(jsii.Number(10)),
}).CreateAlarm(this, jsii.String("alarm"), &CreateAlarmOptions{
	EvaluationPeriods: jsii.Number(3),
	Threshold: jsii.Number(90),
})
```

#### Capacity & Scaling

There are some things to take into consideration with Aurora Serverless v2.

To create a cluster that can support serverless v2 instances you configure a
minimum and maximum capacity range on the cluster. This is an example showing
the default values:

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
	ServerlessV2MinCapacity: jsii.Number(0.5),
	ServerlessV2MaxCapacity: jsii.Number(2),
	Vpc: Vpc,
})
```

The capacity is defined as a number of Aurora capacity units (ACUs). You can
specify in half-step increments (40, 40.5, 41, etc). Each serverless instance in
the cluster inherits the capacity that is defined on the cluster. It is not
possible to configure separate capacity at the instance level.

The maximum capacity is mainly used for budget control since it allows you to
set a cap on how high your instance can scale.

The minimum capacity is a little more involved. This controls a couple different
things.

* The scale-up rate is proportional to the current capacity (larger instances
  scale up faster)

  * Adjust the minimum capacity to obtain a suitable scaling rate
* Network throughput is proportional to capacity

> *Info* More complete details can be found [in the docs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2-examples-setting-capacity-range-for-cluster)

Another way that you control the capacity/scaling of your serverless v2 reader
instances is based on the [promotion tier](https://aws.amazon.com/blogs/aws/additional-failover-control-for-amazon-aurora/)
which can be between 0-15. Any serverless v2 instance in the 0-1 tiers will scale alongside the
writer even if the current read load does not require the capacity. This is
because instances in the 0-1 tier are first priority for failover and Aurora
wants to ensure that in the event of a failover the reader that gets promoted is
scaled to handle the write load.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
	Readers: []iClusterInstance{
		rds.ClusterInstance_*ServerlessV2(jsii.String("reader1"), &ServerlessV2ClusterInstanceProps{
			ScaleWithWriter: jsii.Boolean(true),
		}),
		rds.ClusterInstance_*ServerlessV2(jsii.String("reader2")),
	},
	Vpc: Vpc,
})
```

* When the writer scales up, any readers in tier 0-1 will scale up to match
* Scaling for tier 2-15 is independent of what is happening on the writer
* Readers in tier 2-15 scale up based on read load against the individual reader

When configuring your cluster it is important to take this into consideration
and ensure that in the event of a failover there is an instance that is scaled
up to take over.

### Mixing Serverless v2 and Provisioned instances

You are able to create a cluster that has both provisioned and serverless
instances. [This blog post](https://aws.amazon.com/blogs/database/evaluate-amazon-aurora-serverless-v2-for-your-provisioned-aurora-clusters/)
has an excellent guide on choosing between serverless and provisioned instances
based on use case.

There are a couple of high level differences:

* Engine Version (serverless only supports MySQL 8+ & PostgreSQL 13+)
* Memory up to 256GB can be supported by serverless

#### Provisioned writer

With a provisioned writer and serverless v2 readers, some of the serverless
readers will need to be configured to scale with the writer so they can act as
failover targets. You will need to determine the correct capacity based on the
provisioned instance type and it's utilization.

As an example, if the CPU utilization for a db.r6g.4xlarge (128 GB) instance
stays at 10% most times, then the minimum ACUs may be set at 6.5 ACUs
(10% of 128 GB) and maximum may be set at 64 ACUs (64x2GB=128GB). Keep in mind
that the speed at which the serverless instance can scale up is determined by
the minimum capacity so if your cluster has spiky workloads you may need to set
a higher minimum capacity.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R6G, ec2.InstanceSize_XLARGE4),
	}),
	ServerlessV2MinCapacity: jsii.Number(6.5),
	ServerlessV2MaxCapacity: jsii.Number(64),
	Readers: []iClusterInstance{
		rds.ClusterInstance_ServerlessV2(jsii.String("reader1"), &ServerlessV2ClusterInstanceProps{
			ScaleWithWriter: jsii.Boolean(true),
		}),
		rds.ClusterInstance_*ServerlessV2(jsii.String("reader2")),
	},
	Vpc: Vpc,
})
```

In the above example `reader1` will scale with the writer based on the writer's
utilization. So if the writer were to go to `50%` utilization then `reader1`
would scale up to use `32` ACUs. If the read load stayed consistent then
`reader2` may remain at `6.5` since it is not configured to scale with the
writer.

If one of your Aurora Serverless v2 DB instances consistently reaches the
limit of its maximum capacity, Aurora indicates this condition by setting the
DB instance to a status of `incompatible-parameters`. While the DB instance has
the incompatible-parameters status, some operations are blocked. For example,
you can't upgrade the engine version.

#### CA certificate

Use the `caCertificate` property to specify the [CA certificates](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html)
to use for a cluster instances:

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
		CaCertificate: rds.CaCertificate_RDS_CA_RDS2048_G1(),
	}),
	Readers: []iClusterInstance{
		rds.ClusterInstance_ServerlessV2(jsii.String("reader"), &ServerlessV2ClusterInstanceProps{
			CaCertificate: rds.CaCertificate_Of(jsii.String("custom-ca")),
		}),
	},
	Vpc: Vpc,
})
```

### Migrating from instanceProps

Creating instances in a `DatabaseCluster` using `instanceProps` & `instances` is
deprecated. To migrate to the new properties you can provide the
`isFromLegacyInstanceProps` property.

For example, in order to migrate from this deprecated config:

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Instances: jsii.Number(2),
	InstanceProps: &InstanceProps{
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
		VpcSubnets: &SubnetSelection{
			SubnetType: ec2.SubnetType_PUBLIC,
		},
		Vpc: *Vpc,
	},
})
```

You would need to migrate to this. The old method of providing `instanceProps`
and `instances` will create the number of `instances` that you provide. The
first instance will be the writer and the rest will be the readers. It's
important that the `id` that you provide is `Instance{NUMBER}`. The writer
should always be `Instance1` and the readers will increment from there.

Make sure to run a `cdk diff` before deploying to make sure that all changes are
expected. **Always test the migration in a non-production environment first.**

```go
var vpc vpc
instanceProps := map[string]interface{}{
	"instanceType": ec2.InstanceType_of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
	"isFromLegacyInstanceProps": jsii.Boolean(true),
}
cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	Vpc: Vpc,
	Writer: rds.ClusterInstance_Provisioned(jsii.String("Instance1"), &ProvisionedClusterInstanceProps{
		InstanceType: instanceProps.instanceType,
		IsFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
	}),
	Readers: []iClusterInstance{
		rds.ClusterInstance_*Provisioned(jsii.String("Instance2"), &ProvisionedClusterInstanceProps{
			InstanceType: instanceProps.instanceType,
			IsFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
		}),
	},
})
```

## Starting an instance database

To set up an instance database, define a `DatabaseInstance`. You must
always launch a database in a VPC. Use the `vpcSubnets` attribute to control whether
your instances will be launched privately or publicly:

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	// optional, defaults to m5.large
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("syscdk")),
	 // Optional - will default to 'admin' username and generated password
	Vpc: Vpc,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
})
```

If there isn't a constant for the exact engine version you want to use,
all of the `Version` classes have a static `of` method that can be used to create an arbitrary version.

```go
customEngineVersion := rds.OracleEngineVersion_Of(jsii.String("19.0.0.0.ru-2020-04.rur-2020-04.r1"), jsii.String("19"))
```

By default, the master password will be generated and stored in AWS Secrets Manager.

To use the storage auto scaling option of RDS you can specify the maximum allocated storage.
This is the upper limit to which RDS can automatically scale the storage. More info can be found
[here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
Example for max storage configuration:

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
		Version: rds.PostgresEngineVersion_VER_15_2(),
	}),
	// optional, defaults to m5.large
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_SMALL),
	Vpc: Vpc,
	MaxAllocatedStorage: jsii.Number(200),
})
```

To use dual-stack mode, specify `NetworkType.DUAL` on the `networkType` property:

```go
var vpc vpc
// VPC and subnets must have IPv6 CIDR blocks
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
		Version: rds.PostgresEngineVersion_VER_15_2(),
	}),
	Vpc: Vpc,
	NetworkType: rds.NetworkType_DUAL,
	PubliclyAccessible: jsii.Boolean(false),
})
```

For more information about dual-stack mode, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html).

Use `DatabaseInstanceFromSnapshot` and `DatabaseInstanceReadReplica` to create an instance from snapshot or
a source database respectively:

```go
var vpc vpc

var sourceInstance databaseInstance

rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &DatabaseInstanceFromSnapshotProps{
	SnapshotIdentifier: jsii.String("my-snapshot"),
	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
		Version: rds.PostgresEngineVersion_VER_15_2(),
	}),
	// optional, defaults to m5.large
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_LARGE),
	Vpc: Vpc,
})
rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &DatabaseInstanceReadReplicaProps{
	SourceDatabaseInstance: sourceInstance,
	InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_LARGE),
	Vpc: Vpc,
})
```

Automatic backups of read replica instances are only supported for MySQL and MariaDB. By default,
automatic backups are disabled for read replicas and can only be enabled (using `backupRetention`)
if also enabled on the source instance.

Creating a "production" Oracle database instance with option and parameter groups:

```go
// Set open cursors with parameter group
parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	Parameters: map[string]*string{
		"open_cursors": jsii.String("2500"),
	},
})

optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &OptionGroupProps{
	Engine: rds.DatabaseInstanceEngine_*OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	Configurations: []optionConfiguration{
		&optionConfiguration{
			Name: jsii.String("LOCATOR"),
		},
		&optionConfiguration{
			Name: jsii.String("OEM"),
			Port: jsii.Number(1158),
			Vpc: *Vpc,
		},
	},
})

// Allow connections to OEM
optionGroup.OptionConnections.oEM.Connections.AllowDefaultPortFromAnyIpv4()

// Database instance with production values
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_*OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	LicenseModel: rds.LicenseModel_BRING_YOUR_OWN_LICENSE,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_MEDIUM),
	MultiAz: jsii.Boolean(true),
	StorageType: rds.StorageType_IO1,
	Credentials: rds.Credentials_FromUsername(jsii.String("syscdk")),
	Vpc: Vpc,
	DatabaseName: jsii.String("ORCL"),
	StorageEncrypted: jsii.Boolean(true),
	BackupRetention: cdk.Duration_Days(jsii.Number(7)),
	MonitoringInterval: cdk.Duration_Seconds(jsii.Number(60)),
	EnablePerformanceInsights: jsii.Boolean(true),
	CloudwatchLogsExports: []*string{
		jsii.String("trace"),
		jsii.String("audit"),
		jsii.String("alert"),
		jsii.String("listener"),
	},
	CloudwatchLogsRetention: logs.RetentionDays_ONE_MONTH,
	AutoMinorVersionUpgrade: jsii.Boolean(true),
	 // required to be true if LOCATOR is used in the option group
	OptionGroup: OptionGroup,
	ParameterGroup: ParameterGroup,
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Allow connections on default port from any IPV4
instance.connections.AllowDefaultPortFromAnyIpv4()

// Rotate the master user password every 30 days
instance.addRotationSingleUser()

// Add alarm for high CPU
// Add alarm for high CPU
cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &AlarmProps{
	Metric: instance.metricCPUUtilization(),
	Threshold: jsii.Number(90),
	EvaluationPeriods: jsii.Number(1),
})

// Trigger Lambda function on instance availability events
fn := lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
})

availabilityRule := instance.OnEvent(jsii.String("Availability"), &OnEventOptions{
	Target: targets.NewLambdaFunction(fn),
})
availabilityRule.AddEventPattern(&EventPattern{
	Detail: map[string]interface{}{
		"EventCategories": []interface{}{
			jsii.String("availability"),
		},
	},
})
```

Add XMLDB and OEM with option group

```go
// Set open cursors with parameter group
parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	Parameters: map[string]*string{
		"open_cursors": jsii.String("2500"),
	},
})

optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &OptionGroupProps{
	Engine: rds.DatabaseInstanceEngine_*OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	Configurations: []optionConfiguration{
		&optionConfiguration{
			Name: jsii.String("LOCATOR"),
		},
		&optionConfiguration{
			Name: jsii.String("OEM"),
			Port: jsii.Number(1158),
			Vpc: *Vpc,
		},
	},
})

// Allow connections to OEM
optionGroup.OptionConnections.oEM.Connections.AllowDefaultPortFromAnyIpv4()

// Database instance with production values
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_*OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
	}),
	LicenseModel: rds.LicenseModel_BRING_YOUR_OWN_LICENSE,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_MEDIUM),
	MultiAz: jsii.Boolean(true),
	StorageType: rds.StorageType_IO1,
	Credentials: rds.Credentials_FromUsername(jsii.String("syscdk")),
	Vpc: Vpc,
	DatabaseName: jsii.String("ORCL"),
	StorageEncrypted: jsii.Boolean(true),
	BackupRetention: cdk.Duration_Days(jsii.Number(7)),
	MonitoringInterval: cdk.Duration_Seconds(jsii.Number(60)),
	EnablePerformanceInsights: jsii.Boolean(true),
	CloudwatchLogsExports: []*string{
		jsii.String("trace"),
		jsii.String("audit"),
		jsii.String("alert"),
		jsii.String("listener"),
	},
	CloudwatchLogsRetention: logs.RetentionDays_ONE_MONTH,
	AutoMinorVersionUpgrade: jsii.Boolean(true),
	 // required to be true if LOCATOR is used in the option group
	OptionGroup: OptionGroup,
	ParameterGroup: ParameterGroup,
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
})

// Allow connections on default port from any IPV4
instance.connections.AllowDefaultPortFromAnyIpv4()

// Rotate the master user password every 30 days
instance.addRotationSingleUser()

// Add alarm for high CPU
// Add alarm for high CPU
cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &AlarmProps{
	Metric: instance.metricCPUUtilization(),
	Threshold: jsii.Number(90),
	EvaluationPeriods: jsii.Number(1),
})

// Trigger Lambda function on instance availability events
fn := lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: lambda.Code_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
})

availabilityRule := instance.OnEvent(jsii.String("Availability"), &OnEventOptions{
	Target: targets.NewLambdaFunction(fn),
})
availabilityRule.AddEventPattern(&EventPattern{
	Detail: map[string]interface{}{
		"EventCategories": []interface{}{
			jsii.String("availability"),
		},
	},
})
```

Use the `storageType` property to specify the [type of storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)
to use for the instance:

```go
var vpc vpc


iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_30(),
	}),
	Vpc: Vpc,
	StorageType: rds.StorageType_IO1,
	Iops: jsii.Number(5000),
})

gp3Instance := rds.NewDatabaseInstance(this, jsii.String("Gp3Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_*Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_30(),
	}),
	Vpc: Vpc,
	AllocatedStorage: jsii.Number(500),
	StorageType: rds.StorageType_GP3,
	StorageThroughput: jsii.Number(500),
})
```

Use the `allocatedStorage` property to specify the amount of storage (in gigabytes) that is initially allocated for the instance
to use for the instance:

```go
var vpc vpc

// Setting allocatedStorage for DatabaseInstance replica
// Note: If allocatedStorage isn't set here, the replica instance will inherit the allocatedStorage of the source instance
var sourceInstance databaseInstance


// Setting allocatedStorage for DatabaseInstance
iopsInstance := rds.NewDatabaseInstance(this, jsii.String("IopsInstance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_30(),
	}),
	Vpc: Vpc,
	StorageType: rds.StorageType_IO1,
	Iops: jsii.Number(5000),
	AllocatedStorage: jsii.Number(500),
})
rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &DatabaseInstanceReadReplicaProps{
	SourceDatabaseInstance: sourceInstance,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_LARGE),
	Vpc: Vpc,
	AllocatedStorage: jsii.Number(500),
})
```

Use the `caCertificate` property to specify the [CA certificates](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html)
to use for the instance:

```go
var vpc vpc


rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_30(),
	}),
	Vpc: Vpc,
	CaCertificate: rds.CaCertificate_RDS_CA_RDS2048_G1(),
})
```

You can specify a custom CA certificate with:

```go
var vpc vpc


rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_30(),
	}),
	Vpc: Vpc,
	CaCertificate: rds.CaCertificate_Of(jsii.String("future-rds-ca")),
})
```

## Setting Public Accessibility

You can set public accessibility for the `DatabaseInstance` or the `ClusterInstance` using the `publiclyAccessible` property.
If you specify `true`, it creates an instance with a publicly resolvable DNS name, which resolves to a public IP address.
If you specify `false`, it creates an internal instance with a DNS name that resolves to a private IP address.

The default value will be `true` if `vpcSubnets` is `subnetType: SubnetType.PUBLIC`, `false` otherwise. In the case of a
cluster, the default value will be determined on the vpc placement of the `DatabaseCluster` otherwise it will be determined
based on the vpc placement of standalone `DatabaseInstance`.

```go
var vpc vpc

// Setting public accessibility for DB instance
// Setting public accessibility for DB instance
rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_19(),
	}),
	Vpc: Vpc,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	PubliclyAccessible: jsii.Boolean(true),
})

// Setting public accessibility for DB cluster instance
// Setting public accessibility for DB cluster instance
rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("Writer"), &ServerlessV2ClusterInstanceProps{
		PubliclyAccessible: jsii.Boolean(true),
	}),
	Vpc: Vpc,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
})
```

## Instance events

To define Amazon CloudWatch event rules for database instances, use the `onEvent`
method:

```go
var instance databaseInstance
var fn function

rule := instance.OnEvent(jsii.String("InstanceEvent"), &OnEventOptions{
	Target: targets.NewLambdaFunction(fn),
})
```

## Login credentials

By default, database instances and clusters (with the exception of `DatabaseInstanceFromSnapshot` and `ServerlessClusterFromSnapshot`) will have `admin` user with an auto-generated password.
An alternative username (and password) may be specified for the admin user instead of the default.

The following examples use a `DatabaseInstance`, but the same usage is applicable to `DatabaseCluster`.

```go
var vpc vpc

engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
	Version: rds.PostgresEngineVersion_VER_15_2(),
})
rds.NewDatabaseInstance(this, jsii.String("InstanceWithUsername"), &DatabaseInstanceProps{
	Engine: Engine,
	Vpc: Vpc,
	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres")),
})

rds.NewDatabaseInstance(this, jsii.String("InstanceWithUsernameAndPassword"), &DatabaseInstanceProps{
	Engine: Engine,
	Vpc: Vpc,
	Credentials: rds.Credentials_FromPassword(jsii.String("postgres"), awscdk.SecretValue_SsmSecure(jsii.String("/dbPassword"), jsii.String("1"))),
})

mySecret := secretsmanager.Secret_FromSecretName(this, jsii.String("DBSecret"), jsii.String("myDBLoginInfo"))
rds.NewDatabaseInstance(this, jsii.String("InstanceWithSecretLogin"), &DatabaseInstanceProps{
	Engine: Engine,
	Vpc: Vpc,
	Credentials: rds.Credentials_FromSecret(mySecret),
})
```

Secrets generated by `fromGeneratedSecret()` can be customized:

```go
var vpc vpc

engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
	Version: rds.PostgresEngineVersion_VER_15_2(),
})
myKey := kms.NewKey(this, jsii.String("MyKey"))

rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &DatabaseInstanceProps{
	Engine: Engine,
	Vpc: Vpc,
	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres"), &CredentialsBaseOptions{
		SecretName: jsii.String("my-cool-name"),
		EncryptionKey: myKey,
		ExcludeCharacters: jsii.String("!&*^#@()"),
		ReplicaRegions: []replicaRegion{
			&replicaRegion{
				Region: jsii.String("eu-west-1"),
			},
			&replicaRegion{
				Region: jsii.String("eu-west-2"),
			},
		},
	}),
})
```

### Snapshot credentials

As noted above, Databases created with `DatabaseInstanceFromSnapshot` or `ServerlessClusterFromSnapshot` will not create user and auto-generated password by default because it's not possible to change the master username for a snapshot. Instead, they will use the existing username and password from the snapshot. You can still generate a new password - to generate a secret similarly to the other constructs, pass in credentials with `fromGeneratedSecret()` or `fromGeneratedPassword()`.

```go
var vpc vpc

engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
	Version: rds.PostgresEngineVersion_VER_15_2(),
})
myKey := kms.NewKey(this, jsii.String("MyKey"))

rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("InstanceFromSnapshotWithCustomizedSecret"), &DatabaseInstanceFromSnapshotProps{
	Engine: Engine,
	Vpc: Vpc,
	SnapshotIdentifier: jsii.String("mySnapshot"),
	Credentials: rds.SnapshotCredentials_FromGeneratedSecret(jsii.String("username"), &SnapshotCredentialsFromGeneratedPasswordOptions{
		EncryptionKey: myKey,
		ExcludeCharacters: jsii.String("!&*^#@()"),
		ReplicaRegions: []replicaRegion{
			&replicaRegion{
				Region: jsii.String("eu-west-1"),
			},
			&replicaRegion{
				Region: jsii.String("eu-west-2"),
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

cluster.Connections.AllowFromAnyIpv4(ec2.Port_AllTraffic(), jsii.String("Open to the world"))
```

The endpoints to access your database cluster will be available as the `.clusterEndpoint` and `.readerEndpoint`
attributes:

```go
var cluster databaseCluster

writeAddress := cluster.ClusterEndpoint.SocketAddress
```

For an instance database:

```go
var instance databaseInstance

address := instance.InstanceEndpoint.SocketAddress
```

## Rotating credentials

When the master password is generated and stored in AWS Secrets Manager, it can be rotated automatically:

```go
var instance databaseInstance
var mySecurityGroup securityGroup


instance.addRotationSingleUser(&RotationSingleUserOptions{
	AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(7)),
	 // defaults to 30 days
	ExcludeCharacters: jsii.String("!@#$%^&*"),
	 // defaults to the set " %+~`#/// here*()|[]{}:;<>?!'/@\"\\"
	SecurityGroup: mySecurityGroup,
})
```

```go
cluster := rds.NewDatabaseCluster(stack, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA(),
	InstanceProps: &InstanceProps{
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
		Vpc: *Vpc,
	},
})

cluster.addRotationSingleUser()

clusterWithCustomRotationOptions := rds.NewDatabaseCluster(stack, jsii.String("CustomRotationOptions"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA(),
	InstanceProps: &InstanceProps{
		InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
		Vpc: *Vpc,
	},
})
clusterWithCustomRotationOptions.addRotationSingleUser(&RotationSingleUserOptions{
	AutomaticallyAfter: cdk.Duration_Days(jsii.Number(7)),
	ExcludeCharacters: jsii.String("!@#$%^&*"),
	SecurityGroup: SecurityGroup,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	Endpoint: endpoint,
})
```

The multi user rotation scheme is also available:

```go
var instance databaseInstance
var myImportedSecret databaseSecret

instance.addRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
	Secret: myImportedSecret,
})
```

It's also possible to create user credentials together with the instance/cluster and add rotation:

```go
var instance databaseInstance

myUserSecret := rds.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &DatabaseSecretProps{
	Username: jsii.String("myuser"),
	SecretName: jsii.String("my-user-secret"),
	 // optional, defaults to a CloudFormation-generated name
	Dbname: jsii.String("mydb"),
	 //optional, defaults to the main database of the RDS cluster this secret gets attached to
	MasterSecret: instance.Secret,
	ExcludeCharacters: jsii.String("{}[]()'\"/\\"),
})
myUserSecretAttached := myUserSecret.attach(instance) // Adds DB connections information in the secret

instance.addRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
	 // Add rotation using the multi user scheme
	Secret: myUserSecretAttached,
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


instance.addRotationSingleUser(&RotationSingleUserOptions{
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	 // Place rotation Lambda in private subnets
	Endpoint: myEndpoint,
})
```

See also [aws-cdk-lib/aws-secretsmanager](https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/aws-secretsmanager/README.md) for credentials rotation of existing clusters/instances.

By default, any stack updates will cause AWS Secrets Manager to rotate a secret immediately. To prevent this behavior and wait until the next scheduled rotation window specified via the `automaticallyAfter` property, set the `rotateImmediatelyOnUpdate` property to false:

```go
var instance databaseInstance
var mySecurityGroup securityGroup


instance.addRotationSingleUser(&RotationSingleUserOptions{
	AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(7)),
	 // defaults to 30 days
	ExcludeCharacters: jsii.String("!@#$%^&*"),
	 // defaults to the set " %+~`#/// here*()|[]{}:;<>?!'/@\"\\"
	SecurityGroup: mySecurityGroup,
	 // defaults to an auto-created security group
	RotateImmediatelyOnUpdate: jsii.Boolean(false),
})
```

## IAM Authentication

You can also authenticate to a database instance using AWS Identity and Access Management (IAM) database authentication;
See [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) for more information
and a list of supported versions and limitations.

The following example shows enabling IAM authentication for a database instance and granting connection access to an IAM role.

### Instance

```go
var vpc vpc

instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_19(),
	}),
	Vpc: Vpc,
	IamAuthentication: jsii.Boolean(true),
})
role := iam.NewRole(this, jsii.String("DBRole"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(this.Account),
})
instance.grantConnect(role)
```

### Proxy

The following example shows granting connection access for RDS Proxy to an IAM role.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
})

proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &DatabaseProxyProps{
	ProxyTarget: rds.ProxyTarget_FromCluster(cluster),
	Secrets: []iSecret{
		cluster.Secret,
	},
	Vpc: Vpc,
})

role := iam.NewRole(this, jsii.String("DBProxyRole"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(this.Account),
})
proxy.GrantConnect(role, jsii.String("admin"))
```

**Note**: In addition to the setup above, a database user will need to be created to support IAM auth.
See [https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.DBAccounts.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.DBAccounts.html) for setup instructions.

To specify the details of authentication used by a proxy to log in as a specific database
user use the `clientPasswordAuthType` property:

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
})

proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &DatabaseProxyProps{
	ProxyTarget: rds.ProxyTarget_FromCluster(cluster),
	Secrets: []iSecret{
		cluster.Secret,
	},
	Vpc: Vpc,
	ClientPasswordAuthType: rds.ClientPasswordAuthType_MYSQL_NATIVE_PASSWORD,
})
```

### Cluster

The following example shows granting connection access for an IAM role to an Aurora Cluster.

```go
var vpc vpc

cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
})
role := iam.NewRole(this, jsii.String("AppRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("someservice.amazonaws.com")),
})
cluster.GrantConnect(role, jsii.String("somedbuser"))
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

role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &RoleProps{
	AssumedBy: iam.NewCompositePrincipal(
	iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
	iam.NewServicePrincipal(jsii.String("directoryservice.rds.amazonaws.com"))),
	ManagedPolicies: []iManagedPolicy{
		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
	},
})
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_19(),
	}),
	Vpc: Vpc,
	Domain: jsii.String("d-????????"),
	 // The ID of the domain for the instance to join.
	DomainRole: role,
})
```

You can also use the Kerberos authentication for an Aurora database cluster.

```go
var vpc vpc

iamRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewCompositePrincipal(
	iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
	iam.NewServicePrincipal(jsii.String("directoryservice.rds.amazonaws.com"))),
	ManagedPolicies: []iManagedPolicy{
		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
	},
})

rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_05_1(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("Instance"), &ProvisionedClusterInstanceProps{
		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_MEDIUM),
	}),
	Vpc: Vpc,
	Domain: jsii.String("d-????????"),
	 // The ID of the domain for the cluster to join.
	DomainRole: iamRole,
})
```

**Note**: In addition to the setup above, you need to make sure that the database instance or cluster has network connectivity
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
readLatency := instance.metric(jsii.String("ReadLatency"), &MetricOptions{
	Statistic: jsii.String("Average"),
	Period: awscdk.Duration_Seconds(jsii.Number(60)),
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
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

importBucket := s3.NewBucket(this, jsii.String("importbucket"))
exportBucket := s3.NewBucket(this, jsii.String("exportbucket"))
rds.NewDatabaseCluster(this, jsii.String("dbcluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
	S3ImportBuckets: []iBucket{
		importBucket,
	},
	S3ExportBuckets: []*iBucket{
		exportBucket,
	},
})
```

## Creating a Database Proxy

Amazon RDS Proxy sits between your application and your relational database to efficiently manage
connections to the database and improve scalability of the application. Learn more about at [Amazon RDS Proxy](https://aws.amazon.com/rds/proxy/).

RDS Proxy is supported for MySQL, MariaDB, Postgres, and SQL Server.

The following code configures an RDS Proxy for a `DatabaseInstance`.

```go
var vpc vpc
var securityGroup securityGroup
var secrets []secret
var dbInstance databaseInstance


proxy := dbInstance.AddProxy(jsii.String("proxy"), &DatabaseProxyOptions{
	BorrowTimeout: awscdk.Duration_Seconds(jsii.Number(30)),
	MaxConnectionsPercent: jsii.Number(50),
	Secrets: Secrets,
	Vpc: Vpc,
})
```

## Exporting Logs

You can publish database logs to Amazon CloudWatch Logs. With CloudWatch Logs, you can perform real-time analysis of the log data,
store the data in highly durable storage, and manage the data with the CloudWatch Logs Agent. This is available for both database
instances and clusters; the types of logs available depend on the database type and engine being used.

```go
import "github.com/aws/aws-cdk-go/awscdk"
var myLogsPublishingRole role
var vpc vpc


// Exporting logs from a cluster
cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_Aurora(&AuroraClusterEngineProps{
		Version: rds.AuroraEngineVersion_VER_1_17_9(),
	}),
	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
	Vpc: Vpc,
	CloudwatchLogsExports: []*string{
		jsii.String("error"),
		jsii.String("general"),
		jsii.String("slowquery"),
		jsii.String("audit"),
	},
	 // Export all available MySQL-based logs
	CloudwatchLogsRetention: logs.RetentionDays_THREE_MONTHS,
	 // Optional - default is to never expire logs
	CloudwatchLogsRetentionRole: myLogsPublishingRole,
})

// When 'cloudwatchLogsExports' is set, each export value creates its own log group in DB cluster.
// Specify an export value to access its log group.
errorLogGroup := cluster.cloudwatchLogGroups["error"]
auditLogGroup := cluster.cloudwatchLogGroups.audit

// Exporting logs from an instance
instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
		Version: rds.PostgresEngineVersion_VER_15_2(),
	}),
	Vpc: Vpc,
	CloudwatchLogsExports: []*string{
		jsii.String("postgresql"),
	},
	 // Export the PostgreSQL logs
	CloudwatchLogsRetention: logs.RetentionDays_THREE_MONTHS,
})

// When 'cloudwatchLogsExports' is set, each export value creates its own log group in DB instance.
// Specify an export value to access its log group.
postgresqlLogGroup := instance.cloudwatchLogGroups["postgresql"]
```

## Option Groups

Some DB engines offer additional features that make it easier to manage data and databases, and to provide additional security for your database.
Amazon RDS uses option groups to enable and configure these features. An option group can specify features, called options,
that are available for a particular Amazon RDS DB instance.

```go
var vpc vpc
var securityGroup securityGroup


rds.NewOptionGroup(this, jsii.String("Options"), &OptionGroupProps{
	Engine: rds.DatabaseInstanceEngine_OracleSe2(&OracleSe2InstanceEngineProps{
		Version: rds.OracleEngineVersion_VER_19(),
	}),
	Configurations: []optionConfiguration{
		&optionConfiguration{
			Name: jsii.String("OEM"),
			Port: jsii.Number(5500),
			Vpc: *Vpc,
			SecurityGroups: []iSecurityGroup{
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


parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &ParameterGroupProps{
	Engine: rds.DatabaseInstanceEngine_SqlServerEe(&SqlServerEeInstanceEngineProps{
		Version: rds.SqlServerEngineVersion_VER_11(),
	}),
	Parameters: map[string]*string{
		"locks": jsii.String("100"),
	},
})

rds.NewDatabaseInstance(this, jsii.String("Database"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_SQL_SERVER_EE(),
	Vpc: Vpc,
	ParameterGroup: ParameterGroup,
})
```

Another way to specify parameters is to use the inline field `parameters` that creates an RDS parameter group for you.
You can use this if you do not want to reuse the parameter group instance for different instances:

```go
var vpc vpc


rds.NewDatabaseInstance(this, jsii.String("Database"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_SqlServerEe(&SqlServerEeInstanceEngineProps{
		Version: rds.SqlServerEngineVersion_VER_11(),
	}),
	Vpc: Vpc,
	Parameters: map[string]*string{
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


cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
	CopyTagsToSnapshot: jsii.Boolean(true),
	 // whether to save the cluster tags when creating the snapshot. Default is 'true'
	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql11")),
	Vpc: Vpc,
	Scaling: &ServerlessScalingOptions{
		AutoPause: awscdk.Duration_Minutes(jsii.Number(10)),
		 // default is to pause after 5 minutes of idle time
		MinCapacity: rds.AuroraCapacityUnit_ACU_8,
		 // default is 2 Aurora capacity units (ACUs)
		MaxCapacity: rds.AuroraCapacityUnit_ACU_32,
		 // default is 16 Aurora capacity units (ACUs)
		Timeout: awscdk.Duration_Seconds(jsii.Number(100)),
		 // default is 5 minutes
		TimeoutAction: rds.TimeoutAction_FORCE_APPLY_CAPACITY_CHANGE,
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

rds.NewServerlessClusterFromSnapshot(this, jsii.String("Cluster"), &ServerlessClusterFromSnapshotProps{
	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
	Vpc: Vpc,
	SnapshotIdentifier: jsii.String("mySnapshot"),
})
```

### Data API

You can access your Aurora DB cluster using the built-in Data API. The Data API doesn't require a persistent connection to the DB cluster. Instead, it provides a secure HTTP endpoint and integration with AWS SDKs.

The following example shows granting Data API access to a Lamba function.

```go
var vpc vpc
var fn function


// Create a serverless V1 cluster
serverlessV1Cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
	Vpc: Vpc,
	 // this parameter is optional for serverless Clusters
	EnableDataApi: jsii.Boolean(true),
})
serverlessV1Cluster.grantDataApiAccess(fn)

// Create an Aurora cluster
cluster := rds.NewDatabaseCluster(this, jsii.String("Cluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
	Vpc: Vpc,
	EnableDataApi: jsii.Boolean(true),
})
cluster.GrantDataApiAccess(fn)
// It is necessary to grant the function access to the secret associated with the cluster for `DatabaseCluster`.
cluster.Secret.GrantRead(fn)
```

**Note**: To invoke the Data API, the resource will need to read the secret associated with the cluster.

To learn more about using the Data API, see the [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html).

### Default VPC

The `vpc` parameter is optional.

If not provided, the cluster will be created in the default VPC of the account and region.
As this VPC is not deployed with AWS CDK, you can't configure the `vpcSubnets`, `subnetGroup` or `securityGroups` of the Aurora Serverless Cluster.
If you want to provide one of `vpcSubnets`, `subnetGroup` or `securityGroups` parameter, please provide a `vpc`.

### Preferred Maintenance Window

When creating an RDS cluster, it is possible to (optionally) specify a preferred maintenance window for the cluster as well as the instances under the cluster.
See [AWS docs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance) for more information regarding maintenance windows.

The following code snippet shows an example of setting the cluster's maintenance window to 22:15-22:45 (UTC) on Saturdays, but setting the instances' maintenance window
to 23:15-23:45 on Sundays

```go
var vpc vpc

rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA(),
	InstanceProps: &InstanceProps{
		Vpc: vpc,
		PreferredMaintenanceWindow: jsii.String("Sun:23:15-Sun:23:45"),
	},
	PreferredMaintenanceWindow: jsii.String("Sat:22:15-Sat:22:45"),
})
```
