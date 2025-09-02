# AWS Glue Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## README

[AWS Glue](https://aws.amazon.com/glue/) is a serverless data integration
service that makes it easier to discover, prepare, move, and integrate data
from multiple sources for analytics, machine learning (ML), and application
development.

The Glue L2 construct has convenience methods working backwards from common
use cases and sets required parameters to defaults that align with recommended
best practices for each job type. It also provides customers with a balance
between flexibility via optional parameter overrides, and opinionated
interfaces that discouraging anti-patterns, resulting in reduced time to develop
and deploy new resources.

### References

* [Glue Launch Announcement](https://aws.amazon.com/blogs/aws/launch-aws-glue-now-generally-available/)
* [Glue Documentation](https://docs.aws.amazon.com/glue/index.html)
* [Glue L1 (CloudFormation) Constructs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_Glue.html)
* Prior version of the [@aws-cdk/aws-glue-alpha module](https://github.com/aws/aws-cdk/blob/v2.51.1/packages/%40aws-cdk/aws-glue/README.md)

## Create a Glue Job

A Job encapsulates a script that connects to data sources, processes
them, and then writes output to a data target. There are four types of Glue
Jobs: Spark (ETL and Streaming), Python Shell, Ray, and Flex Jobs. Most
of the required parameters for these jobs are common across all types,
but there are a few differences depending on the languages supported
and features provided by each type. For all job types, the L2 defaults
to AWS best practice recommendations, such as:

* Use of Secrets Manager for Connection JDBC strings
* Glue job autoscaling
* Default parameter values for Glue job creation

This iteration of the L2 construct introduces breaking changes to
the existing glue-alpha-module, but these changes streamline the developer
experience, introduce new constants for defaults, and replacing synth-time
validations with interface contracts for enforcement of the parameter combinations
that Glue supports. As an opinionated construct, the Glue L2 construct does
not allow developers to create resources that use non-current versions
of Glue or deprecated language dependencies (e.g. deprecated versions of Python).
As always, L1s allow you to specify a wider range of parameters if you need
or want to use alternative configurations.

Optional and required parameters for each job are enforced via interface
rather than validation; see [Glue's public documentation](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api.html)
for more granular details.

### Spark Jobs

#### ETL Jobs

ETL jobs support pySpark and Scala languages, for which there are separate but
similar constructors. ETL jobs default to the G2 worker type, but you can
override this default with other supported worker type values (G1, G2, G4
and G8). ETL jobs defaults to Glue version 4.0, which you can override to 3.0.
The following ETL features are enabled by default:
`—enable-metrics, —enable-spark-ui, —enable-continuous-cloudwatch-log.`
You can find more details about version, worker type and other features in
[Glue's public documentation](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html).

Reference the pyspark-etl-jobs.test.ts and scalaspark-etl-jobs.test.ts unit tests
for examples of required-only and optional job parameters when creating these
types of jobs.

For the sake of brevity, examples are shown using the pySpark job variety.

Example with only required parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkEtlJob(stack, jsii.String("PySparkETLJob"), &PySparkEtlJobProps{
	Role: Role,
	Script: Script,
	JobName: jsii.String("PySparkETLJob"),
})
```

Example with optional override parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkEtlJob(stack, jsii.String("PySparkETLJob"), &PySparkEtlJobProps{
	JobName: jsii.String("PySparkETLJobCustomName"),
	Description: jsii.String("This is a description"),
	Role: Role,
	Script: Script,
	GlueVersion: glue.GlueVersion_V3_0,
	ContinuousLogging: &ContinuousLoggingProps{
		Enabled: jsii.Boolean(false),
	},
	WorkerType: glue.WorkerType_G_2X,
	MaxConcurrentRuns: jsii.Number(100),
	Timeout: cdk.Duration_Hours(jsii.Number(2)),
	Connections: []iConnection{
		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
	},
	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
	Tags: map[string]*string{
		"FirstTagName": jsii.String("FirstTagValue"),
		"SecondTagName": jsii.String("SecondTagValue"),
		"XTagName": jsii.String("XTagValue"),
	},
	NumberOfWorkers: jsii.Number(2),
	MaxRetries: jsii.Number(2),
})
```

#### Streaming Jobs

Streaming jobs are similar to ETL jobs, except that they perform ETL on data
streams using the Apache Spark Structured Streaming framework. Some Spark
job features are not available to Streaming ETL jobs. They support Scala
and pySpark languages. PySpark streaming jobs default Python 3.9,
which you can override with any non-deprecated version of Python. It
defaults to the G2 worker type and Glue 4.0, both of which you can override.
The following best practice features are enabled by default:
`—enable-metrics, —enable-spark-ui, —enable-continuous-cloudwatch-log`.

Reference the pyspark-streaming-jobs.test.ts and scalaspark-streaming-jobs.test.ts
unit tests for examples of required-only and optional job parameters when creating
these types of jobs.

Example with only required parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkStreamingJob(stack, jsii.String("ImportedJob"), &PySparkStreamingJobProps{
	Role: Role,
	Script: Script,
})
```

Example with optional override parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkStreamingJob(stack, jsii.String("PySparkStreamingJob"), &PySparkStreamingJobProps{
	JobName: jsii.String("PySparkStreamingJobCustomName"),
	Description: jsii.String("This is a description"),
	Role: Role,
	Script: Script,
	GlueVersion: glue.GlueVersion_V3_0,
	ContinuousLogging: &ContinuousLoggingProps{
		Enabled: jsii.Boolean(false),
	},
	WorkerType: glue.WorkerType_G_2X,
	MaxConcurrentRuns: jsii.Number(100),
	Timeout: cdk.Duration_Hours(jsii.Number(2)),
	Connections: []iConnection{
		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
	},
	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
	Tags: map[string]*string{
		"FirstTagName": jsii.String("FirstTagValue"),
		"SecondTagName": jsii.String("SecondTagValue"),
		"XTagName": jsii.String("XTagValue"),
	},
	NumberOfWorkers: jsii.Number(2),
	MaxRetries: jsii.Number(2),
})
```

#### Flex Jobs

The flexible execution class is appropriate for non-urgent jobs such as
pre-production jobs, testing, and one-time data loads. Flexible jobs default
to Glue version 3.0 and worker type `G_2X`. The following best practice
features are enabled by default:
`—enable-metrics, —enable-spark-ui, —enable-continuous-cloudwatch-log`

Reference the pyspark-flex-etl-jobs.test.ts and scalaspark-flex-etl-jobs.test.ts
unit tests for examples of required-only and optional job parameters when creating
these types of jobs.

Example with only required parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkFlexEtlJob(stack, jsii.String("ImportedJob"), &PySparkFlexEtlJobProps{
	Role: Role,
	Script: Script,
})
```

Example with optional override parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkEtlJob(stack, jsii.String("pySparkEtlJob"), &PySparkEtlJobProps{
	JobName: jsii.String("pySparkEtlJob"),
	Description: jsii.String("This is a description"),
	Role: Role,
	Script: Script,
	GlueVersion: glue.GlueVersion_V3_0,
	ContinuousLogging: &ContinuousLoggingProps{
		Enabled: jsii.Boolean(false),
	},
	WorkerType: glue.WorkerType_G_2X,
	MaxConcurrentRuns: jsii.Number(100),
	Timeout: cdk.Duration_Hours(jsii.Number(2)),
	Connections: []iConnection{
		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
	},
	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
	Tags: map[string]*string{
		"FirstTagName": jsii.String("FirstTagValue"),
		"SecondTagName": jsii.String("SecondTagValue"),
		"XTagName": jsii.String("XTagValue"),
	},
	NumberOfWorkers: jsii.Number(2),
	MaxRetries: jsii.Number(2),
})
```

### Python Shell Jobs

Python shell jobs support a Python version that depends on the AWS Glue
version you use. These can be used to schedule and run tasks that don't
require an Apache Spark environment. Python shell jobs default to
Python 3.9 and a MaxCapacity of `0.0625`. Python 3.9 supports pre-loaded
analytics libraries using the `library-set=analytics` flag, which is
enabled by default.

Reference the pyspark-shell-job.test.ts unit tests for examples of
required-only and optional job parameters when creating these types of jobs.

Example with only required parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPythonShellJob(stack, jsii.String("ImportedJob"), &PythonShellJobProps{
	Role: Role,
	Script: Script,
})
```

Example with optional override parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPythonShellJob(stack, jsii.String("PythonShellJob"), &PythonShellJobProps{
	JobName: jsii.String("PythonShellJobCustomName"),
	Description: jsii.String("This is a description"),
	PythonVersion: glue.PythonVersion_TWO,
	MaxCapacity: glue.MaxCapacity_DPU_1,
	Role: Role,
	Script: Script,
	GlueVersion: glue.GlueVersion_V2_0,
	ContinuousLogging: &ContinuousLoggingProps{
		Enabled: jsii.Boolean(false),
	},
	WorkerType: glue.WorkerType_G_2X,
	MaxConcurrentRuns: jsii.Number(100),
	Timeout: cdk.Duration_Hours(jsii.Number(2)),
	Connections: []iConnection{
		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
	},
	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
	Tags: map[string]*string{
		"FirstTagName": jsii.String("FirstTagValue"),
		"SecondTagName": jsii.String("SecondTagValue"),
		"XTagName": jsii.String("XTagValue"),
	},
	NumberOfWorkers: jsii.Number(2),
	MaxRetries: jsii.Number(2),
})
```

### Ray Jobs

Glue Ray jobs use worker type Z.2X and Glue version 4.0. These are not
overrideable since these are the only configuration that Glue Ray jobs
currently support. The runtime defaults to Ray2.4 and min workers defaults to 3.

Reference the ray-job.test.ts unit tests for examples of required-only and
optional job parameters when creating these types of jobs.

Example with only required parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewRayJob(stack, jsii.String("ImportedJob"), &RayJobProps{
	Role: Role,
	Script: Script,
})
```

Example with optional override parameters:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewRayJob(stack, jsii.String("ImportedJob"), &RayJobProps{
	Role: Role,
	Script: Script,
	JobName: jsii.String("RayCustomJobName"),
	Description: jsii.String("This is a description"),
	WorkerType: glue.WorkerType_Z_2X,
	NumberOfWorkers: jsii.Number(5),
	Runtime: glue.Runtime_RAY_TWO_FOUR,
	MaxRetries: jsii.Number(3),
	MaxConcurrentRuns: jsii.Number(100),
	Timeout: cdk.Duration_Hours(jsii.Number(2)),
	Connections: []iConnection{
		glue.Connection_FromConnectionName(stack, jsii.String("Connection"), jsii.String("connectionName")),
	},
	SecurityConfiguration: glue.SecurityConfiguration_FromSecurityConfigurationName(stack, jsii.String("SecurityConfig"), jsii.String("securityConfigName")),
	Tags: map[string]*string{
		"FirstTagName": jsii.String("FirstTagValue"),
		"SecondTagName": jsii.String("SecondTagValue"),
		"XTagName": jsii.String("XTagValue"),
	},
})
```

### Metrics Control

By default, Glue jobs enable CloudWatch metrics (`--enable-metrics`) and observability metrics (`--enable-observability-metrics`) for monitoring and debugging. You can disable these metrics to reduce CloudWatch costs:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code


// Disable both metrics for cost optimization
// Disable both metrics for cost optimization
glue.NewPySparkEtlJob(stack, jsii.String("CostOptimizedJob"), &PySparkEtlJobProps{
	Role: Role,
	Script: Script,
	EnableMetrics: jsii.Boolean(false),
	EnableObservabilityMetrics: jsii.Boolean(false),
})

// Selective control - keep observability, disable profiling
// Selective control - keep observability, disable profiling
glue.NewPySparkEtlJob(stack, jsii.String("SelectiveJob"), &PySparkEtlJobProps{
	Role: Role,
	Script: Script,
	EnableMetrics: jsii.Boolean(false),
})
```

This feature is available for all Spark job types (ETL, Streaming, Flex) and Ray jobs.

### Enable Job Run Queuing

AWS Glue job queuing monitors your account level quotas and limits. If quotas or limits are insufficient to start a Glue job run, AWS Glue will automatically queue the job and wait for limits to free up. Once limits become available, AWS Glue will retry the job run. Glue jobs will queue for limits like max concurrent job runs per account, max concurrent Data Processing Units (DPU), and resource unavailable due to IP address exhaustion in Amazon Virtual Private Cloud (Amazon VPC).

Enable job run queuing by setting the `jobRunQueuingEnabled` property to `true`.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import iam "github.com/aws/aws-cdk-go/awscdk"
var stack stack
var role iRole
var script code

glue.NewPySparkEtlJob(stack, jsii.String("PySparkETLJob"), &PySparkEtlJobProps{
	Role: Role,
	Script: Script,
	JobName: jsii.String("PySparkETLJob"),
	JobRunQueuingEnabled: jsii.Boolean(true),
})
```

### Uploading scripts from the CDK app repository to S3

Similar to other L2 constructs, the Glue L2 automates uploading / updating
scripts to S3 via an optional fromAsset parameter pointing to a script
in the local file structure. You provide the existing S3 bucket and
path to which you'd like the script to be uploaded.

Reference the unit tests for examples of repo and S3 code target examples.

### Workflow Triggers

You can use Glue workflows to create and visualize complex
extract, transform, and load (ETL) activities involving multiple crawlers,
jobs, and triggers. Standalone triggers are an anti-pattern, so you must
create triggers from within a workflow using the L2 construct.

Within a workflow object, there are functions to create different
types of triggers with actions and predicates. You then add those triggers
to jobs.

StartOnCreation defaults to true for all trigger types, but you can
override it if you prefer for your trigger not to start on creation.

Reference the workflow-triggers.test.ts unit tests for examples of creating
workflows and triggers.

#### **1. On-Demand Triggers**

On-demand triggers can start glue jobs or crawlers. This construct provides
convenience functions to create on-demand crawler or job triggers. The constructor
takes an optional description parameter, but abstracts the requirement of an
actions list using the job or crawler objects using conditional types.

#### **2. Scheduled Triggers**

You can create scheduled triggers using cron expressions. This construct
provides daily, weekly, and monthly convenience functions,
as well as a custom function that allows you to create your own
custom timing using the [existing event Schedule class](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_events.Schedule.html)
without having to build your own cron expressions. The L2 extracts
the expression that Glue requires from the Schedule object. The constructor
takes an optional description and a list of jobs or crawlers as actions.

#### **3. Notify  Event Triggers**

There are two types of notify event triggers: batching and non-batching.
For batching triggers, you must specify `BatchSize`. For non-batching
triggers, `BatchSize` defaults to 1. For both triggers, `BatchWindow`
defaults to 900 seconds, but you can override the window to align with
your workload's requirements.

#### **4. Conditional Triggers**

Conditional triggers have a predicate and actions associated with them.
The trigger actions are executed when the predicateCondition is true.

### Connection Properties

A `Connection` allows Glue jobs, crawlers and development endpoints to access
certain types of data stores.

* **Secrets Management**
  You must specify JDBC connection credentials in Secrets Manager and
  provide the Secrets Manager Key name as a property to the job connection.
* **Networking - the CDK determines the best fit subnet for Glue connection
  configuration**
  The prior version of the glue-alpha-module requires the developer to
  specify the subnet of the Connection when it’s defined. Now, you can still
  specify the specific subnet you want to use, but are no longer required
  to. You are only required to provide a VPC and either a public or private
  subnet selection. Without a specific subnet provided, the L2 leverages the
  existing [EC2 Subnet Selection](https://docs.aws.amazon.com/cdk/api/v2/python/aws_cdk.aws_ec2/SubnetSelection.html)
  library to make the best choice selection for the subnet.

```go
var securityGroup securityGroup
var subnet subnet

glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
	Type: glue.ConnectionType_NETWORK(),
	// The security groups granting AWS Glue inbound access to the data source within the VPC
	SecurityGroups: []iSecurityGroup{
		securityGroup,
	},
	// The VPC subnet which contains the data source
	Subnet: Subnet,
})
```

For RDS `Connection` by JDBC, it is recommended to manage credentials using AWS Secrets Manager. To use Secret, specify `SECRET_ID` in `properties` like the following code. Note that in this case, the subnet must have a route to the AWS Secrets Manager VPC endpoint or to the AWS Secrets Manager endpoint through a NAT gateway.

```go
var securityGroup securityGroup
var subnet subnet
var db databaseCluster

glue.NewConnection(this, jsii.String("RdsConnection"), &ConnectionProps{
	Type: glue.ConnectionType_JDBC(),
	SecurityGroups: []iSecurityGroup{
		securityGroup,
	},
	Subnet: Subnet,
	Properties: map[string]*string{
		"JDBC_CONNECTION_URL": fmt.Sprintf("jdbc:mysql://%v/databasename", db.clusterEndpoint.socketAddress),
		"JDBC_ENFORCE_SSL": jsii.String("false"),
		"SECRET_ID": db.secret.secretName,
	},
})
```

If you need to use a connection type that doesn't exist as a static member on `ConnectionType`, you can instantiate a `ConnectionType` object, e.g: `new glue.ConnectionType('NEW_TYPE')`.

See [Adding a Connection to Your Data Store](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html) and [Connection Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-connections.html#aws-glue-api-catalog-connections-Connection) documentation for more information on the supported data stores and their configurations.

## SecurityConfiguration

A `SecurityConfiguration` is a set of security properties that can be used by AWS Glue to encrypt data at rest.

```go
glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
	CloudWatchEncryption: &CloudWatchEncryption{
		Mode: glue.CloudWatchEncryptionMode_KMS,
	},
	JobBookmarksEncryption: &JobBookmarksEncryption{
		Mode: glue.JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
	},
	S3Encryption: &S3Encryption{
		Mode: glue.S3EncryptionMode_KMS,
	},
})
```

By default, a shared KMS key is created for use with the encryption configurations that require one. You can also supply your own key for each encryption config, for example, for CloudWatch encryption:

```go
var key key

glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
	CloudWatchEncryption: &CloudWatchEncryption{
		Mode: glue.CloudWatchEncryptionMode_KMS,
		KmsKey: key,
	},
})
```

See [documentation](https://docs.aws.amazon.com/glue/latest/dg/encryption-security-configuration.html) for more info for Glue encrypting data written by Crawlers, Jobs, and Development Endpoints.

## Database

A `Database` is a logical grouping of `Tables` in the Glue Catalog.

```go
glue.NewDatabase(this, jsii.String("MyDatabase"), &DatabaseProps{
	DatabaseName: jsii.String("my_database"),
	Description: jsii.String("my_database_description"),
})
```

## Table

A Glue table describes a table of data in S3: its structure (column names and types), location of data (S3 objects with a common prefix in a S3 bucket), and format for the files (Json, Avro, Parquet, etc.):

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
		&column{
			Name: jsii.String("col2"),
			Type: glue.Schema_Array(glue.Schema_STRING()),
			Comment: jsii.String("col2 is an array of strings"),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

By default, a S3 bucket will be created to store the table's data but you can manually pass the `bucket` and `s3Prefix`:

```go
var myBucket bucket
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Bucket: myBucket,
	S3Prefix: jsii.String("my-table/"),
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

Glue tables can be configured to contain user-defined properties, to describe the physical storage of table data, through the `storageParameters` property:

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	StorageParameters: []storageParameter{
		glue.*storageParameter_SkipHeaderLineCount(jsii.Number(1)),
		glue.*storageParameter_CompressionType(glue.CompressionType_GZIP),
		glue.*storageParameter_Custom(jsii.String("separatorChar"), jsii.String(",")),
	},
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

Glue tables can also be configured to contain user-defined table properties through the [`parameters`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-tableinput.html#cfn-glue-table-tableinput-parameters) property:

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Parameters: map[string]*string{
		"key1": jsii.String("val1"),
		"key2": jsii.String("val2"),
	},
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

### Partition Keys

To improve query performance, a table can specify `partitionKeys` on which data is stored and queried separately. For example, you might partition a table by `year` and `month` to optimize queries based on a time window:

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	PartitionKeys: []*column{
		&column{
			Name: jsii.String("year"),
			Type: glue.Schema_SMALL_INT(),
		},
		&column{
			Name: jsii.String("month"),
			Type: glue.Schema_SMALL_INT(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

### Partition Indexes

Another way to improve query performance is to specify partition indexes. If no partition indexes are
present on the table, AWS Glue loads all partitions of the table and filters the loaded partitions using
the query expression. The query takes more time to run as the number of partitions increase. With an
index, the query will try to fetch a subset of the partitions instead of loading all partitions of the
table.

The keys of a partition index must be a subset of the partition keys of the table. You can have a
maximum of 3 partition indexes per table. To specify a partition index, you can use the `partitionIndexes`
property:

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	PartitionKeys: []*column{
		&column{
			Name: jsii.String("year"),
			Type: glue.Schema_SMALL_INT(),
		},
		&column{
			Name: jsii.String("month"),
			Type: glue.Schema_SMALL_INT(),
		},
	},
	PartitionIndexes: []partitionIndex{
		&partitionIndex{
			IndexName: jsii.String("my-index"),
			 // optional
			KeyNames: []*string{
				jsii.String("year"),
			},
		},
	},
	 // supply up to 3 indexes
	DataFormat: glue.DataFormat_JSON(),
})
```

Alternatively, you can call the `addPartitionIndex()` function on a table:

```go
var myTable table

myTable.AddPartitionIndex(&PartitionIndex{
	IndexName: jsii.String("my-index"),
	KeyNames: []*string{
		jsii.String("year"),
	},
})
```

### Partition Filtering

If you have a table with a large number of partitions that grows over time, consider using AWS Glue partition indexing and filtering.

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	PartitionKeys: []*column{
		&column{
			Name: jsii.String("year"),
			Type: glue.Schema_SMALL_INT(),
		},
		&column{
			Name: jsii.String("month"),
			Type: glue.Schema_SMALL_INT(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
	EnablePartitionFiltering: jsii.Boolean(true),
})
```

### Glue Connections

Glue connections allow external data connections to third party databases and data warehouses. However, these connections can also be assigned to Glue Tables, allowing you to query external data sources using the Glue Data Catalog.

Whereas `S3Table` will point to (and if needed, create) a bucket to store the tables' data, `ExternalTable` will point to an existing table in a data source. For example, to create a table in Glue that points to a table in Redshift:

```go
var myConnection connection
var myDatabase database

glue.NewExternalTable(this, jsii.String("MyTable"), &ExternalTableProps{
	Connection: myConnection,
	ExternalDataLocation: jsii.String("default_db_public_example"),
	 // A table in Redshift
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

## [Encryption](https://docs.aws.amazon.com/athena/latest/ug/encryption.html)

You can enable encryption on a Table's data:

* [S3Managed](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) - (default) Server side encryption (`SSE-S3`) with an Amazon S3-managed key.

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_S3_MANAGED,
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

* [Kms](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) - Server-side encryption (`SSE-KMS`) with an AWS KMS Key managed by the account owner.

```go
var myDatabase database

// KMS key is created automatically
// KMS key is created automatically
glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_KMS,
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})

// with an explicit KMS key
// with an explicit KMS key
glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_KMS,
	EncryptionKey: kms.NewKey(this, jsii.String("MyKey")),
	// ...
	Database: myDatabase,
	Columns: []*column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

* [KmsManaged](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) - Server-side encryption (`SSE-KMS`), like `Kms`, except with an AWS KMS Key managed by the AWS Key Management Service.

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_KMS_MANAGED,
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

* [ClientSideKms](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html#client-side-encryption-kms-managed-master-key-intro) - Client-side encryption (`CSE-KMS`) with an AWS KMS Key managed by the account owner.

```go
var myDatabase database

// KMS key is created automatically
// KMS key is created automatically
glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_CLIENT_SIDE_KMS,
	// ...
	Database: myDatabase,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})

// with an explicit KMS key
// with an explicit KMS key
glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Encryption: glue.TableEncryption_CLIENT_SIDE_KMS,
	EncryptionKey: kms.NewKey(this, jsii.String("MyKey")),
	// ...
	Database: myDatabase,
	Columns: []*column{
		&column{
			Name: jsii.String("col1"),
			Type: glue.Schema_STRING(),
		},
	},
	DataFormat: glue.DataFormat_JSON(),
})
```

*Note: you cannot provide a `Bucket` when creating the `S3Table` if you wish to use server-side encryption (`KMS`, `KMS_MANAGED` or `S3_MANAGED`)*.

## Types

A table's schema is a collection of columns, each of which have a `name` and a `type`. Types are recursive structures, consisting of primitive and complex types:

```go
var myDatabase database

glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
	Columns: []column{
		&column{
			Name: jsii.String("primitive_column"),
			Type: glue.Schema_STRING(),
		},
		&column{
			Name: jsii.String("array_column"),
			Type: glue.Schema_Array(glue.Schema_INTEGER()),
			Comment: jsii.String("array<integer>"),
		},
		&column{
			Name: jsii.String("map_column"),
			Type: glue.Schema_Map(glue.Schema_STRING(), glue.Schema_TIMESTAMP()),
			Comment: jsii.String("map<string,string>"),
		},
		&column{
			Name: jsii.String("struct_column"),
			Type: glue.Schema_Struct([]*column{
				&column{
					Name: jsii.String("nested_column"),
					Type: glue.Schema_DATE(),
					Comment: jsii.String("nested comment"),
				},
			}),
			Comment: jsii.String("struct<nested_column:date COMMENT 'nested comment'>"),
		},
	},
	// ...
	Database: myDatabase,
	DataFormat: glue.DataFormat_JSON(),
})
```

## Public FAQ

### What are we launching today?

We’re launching new features to an AWS CDK Glue L2 Construct to provide
best-practice defaults and convenience methods to create Glue Jobs, Connections,
Triggers, Workflows, and the underlying permissions and configuration.

### Why should I use this Construct?

Developers should use this Construct to reduce the amount of boilerplate
code and complexity each individual has to navigate, and make it easier to
create best-practice Glue resources.

### What’s not in scope?

Glue Crawlers and other resources that are now managed by the AWS LakeFormation
team are not in scope for this effort. Developers should use existing methods
to create these resources, and the new Glue L2 construct assumes they already
exist as inputs. While best practice is for application and infrastructure code
to be as close as possible for teams using fully-implemented DevOps mechanisms,
in practice these ETL scripts are likely managed by a data science team who
know Python or Scala and don’t necessarily own or manage their own
infrastructure deployments. We want to meet developers where they are, and not
assume that all of the code resides in the same repository, Developers can
automate this themselves via the CDK, however, if they do own both.

Validating Glue version and feature use per AWS region at synth time is also
not in scope. AWS’ intention is for all features to eventually be propagated to
all Global regions, so the complexity involved in creating and updating region-
specific configuration to match shifting feature sets does not out-weigh the
likelihood that a developer will use this construct to deploy resources to a
region without a particular new feature to a region that doesn’t yet support
it without researching or manually attempting to use that feature before
developing it via IaC. The developer will, of course, still get feedback from
the underlying Glue APIs as CloudFormation deploys the resources similar to the
current CDK L1 Glue experience.
