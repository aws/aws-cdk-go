# Kinesis Analytics Flink

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This package provides constructs for creating Kinesis Analytics Flink
applications. To learn more about using using managed Flink applications, see
the [AWS developer
guide](https://docs.aws.amazon.com/kinesisanalytics/latest/java/).

## Creating Flink Applications

To create a new Flink application, use the `Application` construct:

```go
import path "github.com/aws-samples/dummy/path"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"

app := core.NewApp()
stack := core.NewStack(app, jsii.String("FlinkAppTest"))

flinkApp := flink.NewApplication(stack, jsii.String("App"), &ApplicationProps{
	Code: flink.ApplicationCode_FromAsset(path.join(__dirname, jsii.String("code-asset"))),
	Runtime: flink.Runtime_FLINK_1_11(),
})

cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &AlarmProps{
	Metric: flinkApp.metricFullRestarts(),
	EvaluationPeriods: jsii.Number(1),
	Threshold: jsii.Number(3),
})

app.Synth()
```

The `code` property can use `fromAsset` as shown above to reference a local jar
file in s3 or `fromBucket` to reference a file in s3.

```go
import path "github.com/aws-samples/dummy/path"
import assets "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"

app := core.NewApp()
stack := core.NewStack(app, jsii.String("FlinkAppCodeFromBucketTest"))

asset := assets.NewAsset(stack, jsii.String("CodeAsset"), &AssetProps{
	Path: path.join(__dirname, jsii.String("code-asset")),
})
bucket := asset.Bucket
fileKey := asset.S3ObjectKey

flink.NewApplication(stack, jsii.String("App"), &ApplicationProps{
	Code: flink.ApplicationCode_FromBucket(bucket, fileKey),
	Runtime: flink.Runtime_FLINK_1_11(),
})

app.Synth()
```

The `propertyGroups` property provides a way of passing arbitrary runtime
properties to your Flink application. You can use the
aws-kinesisanalytics-runtime library to [retrieve these
properties](https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-properties.html#how-properties-access).

```go
var bucket bucket

flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
	PropertyGroups: map[string]map[string]*string{
		"FlinkApplicationProperties": map[string]*string{
			"inputStreamName": jsii.String("my-input-kinesis-stream"),
			"outputStreamName": jsii.String("my-output-kinesis-stream"),
		},
	},
	// ...
	Runtime: flink.Runtime_FLINK_1_15(),
	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
})
```

Flink applications also have specific configuration for passing parameters
when the Flink job starts. These include parameters for checkpointing,
snapshotting, monitoring, and parallelism.

```go
var bucket bucket

flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
	Runtime: flink.Runtime_FLINK_1_15(),
	CheckpointingEnabled: jsii.Boolean(true),
	 // default is true
	CheckpointInterval: awscdk.Duration_Seconds(jsii.Number(30)),
	 // default is 1 minute
	MinPauseBetweenCheckpoints: awscdk.Duration_*Seconds(jsii.Number(10)),
	 // default is 5 seconds
	LogLevel: flink.LogLevel_ERROR,
	 // default is INFO
	MetricsLevel: flink.MetricsLevel_PARALLELISM,
	 // default is APPLICATION
	AutoScalingEnabled: jsii.Boolean(false),
	 // default is true
	Parallelism: jsii.Number(32),
	 // default is 1
	ParallelismPerKpu: jsii.Number(2),
	 // default is 1
	SnapshotsEnabled: jsii.Boolean(false),
	 // default is true
	LogGroup: logs.NewLogGroup(this, jsii.String("LogGroup")),
})
```

Flink applications can optionally be deployed in a VPC:

```go
var bucket bucket
var vpc vpc

flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
	Runtime: flink.Runtime_FLINK_1_15(),
	Vpc: Vpc,
})
```
