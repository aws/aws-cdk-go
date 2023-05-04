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

## Job

A `Job` encapsulates a script that connects to data sources, processes them, and then writes output to a data target.

There are 3 types of jobs supported by AWS Glue: Spark ETL, Spark Streaming, and Python Shell jobs.

The `glue.JobExecutable` allows you to specify the type of job, the language to use and the code assets required by the job.

`glue.Code` allows you to refer to the different code assets required by the job, either from an existing S3 location or from a local file path.

### Spark Jobs

These jobs run in an Apache Spark environment managed by AWS Glue.

#### ETL Jobs

An ETL job processes data in batches using Apache Spark.

```go
var bucket bucket

glue.NewJob(this, jsii.String("ScalaSparkEtlJob"), &JobProps{
	Executable: glue.JobExecutable_ScalaEtl(&ScalaJobExecutableProps{
		GlueVersion: glue.GlueVersion_V4_0(),
		Script: glue.Code_FromBucket(bucket, jsii.String("src/com/example/HelloWorld.scala")),
		ClassName: jsii.String("com.example.HelloWorld"),
		ExtraJars: []code{
			glue.*code_*FromBucket(bucket, jsii.String("jars/HelloWorld.jar")),
		},
	}),
	Description: jsii.String("an example Scala ETL job"),
})
```

#### Streaming Jobs

A Streaming job is similar to an ETL job, except that it performs ETL on data streams. It uses the Apache Spark Structured Streaming framework. Some Spark job features are not available to streaming ETL jobs.

```go
glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &JobProps{
	Executable: glue.JobExecutable_PythonStreaming(&PythonSparkJobExecutableProps{
		GlueVersion: glue.GlueVersion_V4_0(),
		PythonVersion: glue.PythonVersion_THREE,
		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
	}),
	Description: jsii.String("an example Python Streaming job"),
})
```

### Python Shell Jobs

A Python shell job runs Python scripts as a shell and supports a Python version that depends on the AWS Glue version you are using.
This can be used to schedule and run tasks that don't require an Apache Spark environment. Currently, three flavors are supported:

* PythonVersion.TWO (2.7; EOL)
* PythonVersion.THREE (3.6)
* PythonVersion.THREE_NINE (3.9)

```go
var bucket bucket

glue.NewJob(this, jsii.String("PythonShellJob"), &JobProps{
	Executable: glue.JobExecutable_PythonShell(&PythonShellExecutableProps{
		GlueVersion: glue.GlueVersion_V1_0(),
		PythonVersion: glue.PythonVersion_THREE,
		Script: glue.Code_FromBucket(bucket, jsii.String("script.py")),
	}),
	Description: jsii.String("an example Python Shell job"),
})
```

### Ray Jobs

These jobs run in a Ray environment managed by AWS Glue.

```go
glue.NewJob(this, jsii.String("RayJob"), &JobProps{
	Executable: glue.JobExecutable_PythonRay(&PythonRayExecutableProps{
		GlueVersion: glue.GlueVersion_V4_0(),
		PythonVersion: glue.PythonVersion_THREE_NINE,
		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
	}),
	WorkerType: glue.WorkerType_Z_2X(),
	WorkerCount: jsii.Number(2),
	Description: jsii.String("an example Ray job"),
})
```

See [documentation](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) for more information on adding jobs in Glue.

## Connection

A `Connection` allows Glue jobs, crawlers and development endpoints to access certain types of data stores. For example, to create a network connection to connect to a data source within a VPC:

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
glue.NewDatabase(this, jsii.String("MyDatabase"))
```

## Table

A Glue table describes a table of data in S3: its structure (column names and types), location of data (S3 objects with a common prefix in a S3 bucket), and format for the files (Json, Avro, Parquet, etc.):

```go
var myDatabase database

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

By default, an S3 bucket will be created to store the table's data and stored in the bucket root. You can also manually pass the `bucket` and `s3Prefix`:

### Partition Keys

To improve query performance, a table can specify `partitionKeys` on which data is stored and queried separately. For example, you might partition a table by `year` and `month` to optimize queries based on a time window:

```go
var myDatabase database

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

## [Encryption](https://docs.aws.amazon.com/athena/latest/ug/encryption.html)

You can enable encryption on a Table's data:

* [S3Managed](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) - (default) Server side encryption (`SSE-S3`) with an Amazon S3-managed key.

```go
var myDatabase database

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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
glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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
glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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
glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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
glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

*Note: you cannot provide a `Bucket` when creating the `Table` if you wish to use server-side encryption (`KMS`, `KMS_MANAGED` or `S3_MANAGED`)*.

## Types

A table's schema is a collection of columns, each of which have a `name` and a `type`. Types are recursive structures, consisting of primitive and complex types:

```go
var myDatabase database

glue.NewTable(this, jsii.String("MyTable"), &TableProps{
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

### Primitives

#### Numeric

| Name      	| Type     	| Comments                                                                                                          |
|-----------	|----------	|------------------------------------------------------------------------------------------------------------------	|
| FLOAT     	| Constant 	| A 32-bit single-precision floating point number                                                                   |
| INTEGER   	| Constant 	| A 32-bit signed value in two's complement format, with a minimum value of -2^31 and a maximum value of 2^31-1 	|
| DOUBLE    	| Constant 	| A 64-bit double-precision floating point number                                                                   |
| BIG_INT   	| Constant 	| A 64-bit signed INTEGER in two’s complement format, with a minimum value of -2^63 and a maximum value of 2^63 -1  |
| SMALL_INT 	| Constant 	| A 16-bit signed INTEGER in two’s complement format, with a minimum value of -2^15 and a maximum value of 2^15-1   |
| TINY_INT  	| Constant 	| A 8-bit signed INTEGER in two’s complement format, with a minimum value of -2^7 and a maximum value of 2^7-1      |

#### Date and time

| Name      	| Type     	| Comments                                                                                                                                                                	|
|-----------	|----------	|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| DATE      	| Constant 	| A date in UNIX format, such as YYYY-MM-DD.                                                                                                                              	|
| TIMESTAMP 	| Constant 	| Date and time instant in the UNiX format, such as yyyy-mm-dd hh:mm:ss[.f...]. For example, TIMESTAMP '2008-09-15 03:04:05.324'. This format uses the session time zone. 	|

#### String

| Name                                       	| Type     	| Comments                                                                                                                                                                                          	|
|--------------------------------------------	|----------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| STRING                                     	| Constant 	| A string literal enclosed in single or double quotes                                                                                                                                              	|
| decimal(precision: number, scale?: number) 	| Function 	| `precision` is the total number of digits. `scale` (optional) is the number of digits in fractional part with a default of 0. For example, use these type definitions: decimal(11,5), decimal(15) 	|
| char(length: number)                       	| Function 	| Fixed length character data, with a specified length between 1 and 255, such as char(10)                                                                                                          	|
| varchar(length: number)                    	| Function 	| Variable length character data, with a specified length between 1 and 65535, such as varchar(10)                                                                                                  	|

#### Miscellaneous

| Name    	| Type     	| Comments                      	|
|---------	|----------	|-------------------------------	|
| BOOLEAN 	| Constant 	| Values are `true` and `false` 	|
| BINARY  	| Constant 	| Value is in binary            	|

### Complex

| Name                                	| Type     	| Comments                                                          	|
|-------------------------------------	|----------	|-------------------------------------------------------------------	|
| array(itemType: Type)               	| Function 	| An array of some other type                                       	|
| map(keyType: Type, valueType: Type) 	| Function 	| A map of some primitive key type to any value type                	|
| struct(collumns: Column[])          	| Function 	| Nested structure containing individually named and typed collumns 	|
