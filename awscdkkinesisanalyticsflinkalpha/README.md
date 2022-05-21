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
import core "github.com/aws/aws-cdk-go/awscdk"
import flink "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

app := core.NewApp()
stack := core.NewStack(app, jsii.String("FlinkAppTest"))

flinkApp := flink.NewApplication(stack, jsii.String("App"), &applicationProps{
	code: flink.applicationCode.fromAsset(path.join(__dirname, jsii.String("code-asset"))),
	runtime: flink.runtime_FLINK_1_11(),
})

cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &alarmProps{
	metric: flinkApp.metricFullRestarts(),
	evaluationPeriods: jsii.Number(1),
	threshold: jsii.Number(3),
})

app.synth()
```

The `code` property can use `fromAsset` as shown above to reference a local jar
file in s3 or `fromBucket` to reference a file in s3.

```go
import path "github.com/aws-samples/dummy/path"
import assets "github.com/aws/aws-cdk-go/awscdk"
import core "github.com/aws/aws-cdk-go/awscdk"
import flink "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"

app := core.NewApp()
stack := core.NewStack(app, jsii.String("FlinkAppCodeFromBucketTest"))

asset := assets.NewAsset(stack, jsii.String("CodeAsset"), &assetProps{
	path: path.join(__dirname, jsii.String("code-asset")),
})
bucket := asset.bucket
fileKey := asset.s3ObjectKey

flink.NewApplication(stack, jsii.String("App"), &applicationProps{
	code: flink.applicationCode.fromBucket(bucket, fileKey),
	runtime: flink.runtime_FLINK_1_11(),
})

app.synth()
```

The `propertyGroups` property provides a way of passing arbitrary runtime
properties to your Flink application. You can use the
aws-kinesisanalytics-runtime library to [retrieve these
properties](https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-properties.html#how-properties-access).

```go
var bucket bucket

flinkApp := flink.NewApplication(this, jsii.String("Application"), &applicationProps{
	propertyGroups: &propertyGroups{
		flinkApplicationProperties: map[string]*string{
			"inputStreamName": jsii.String("my-input-kinesis-stream"),
			"outputStreamName": jsii.String("my-output-kinesis-stream"),
		},
	},
	// ...
	runtime: flink.runtime_FLINK_1_13(),
	code: flink.applicationCode.fromBucket(bucket, jsii.String("my-app.jar")),
})
```

Flink applications also have specific configuration for passing parameters
when the Flink job starts. These include parameters for checkpointing,
snapshotting, monitoring, and parallelism.

```go
var bucket bucket

flinkApp := flink.NewApplication(this, jsii.String("Application"), &applicationProps{
	code: flink.applicationCode.fromBucket(bucket, jsii.String("my-app.jar")),
	runtime: flink.runtime_FLINK_1_13(),
	checkpointingEnabled: jsii.Boolean(true),
	 // default is true
	checkpointInterval: awscdk.Duration.seconds(jsii.Number(30)),
	 // default is 1 minute
	minPauseBetweenCheckpoints: awscdk.Duration.seconds(jsii.Number(10)),
	 // default is 5 seconds
	logLevel: flink.logLevel_ERROR,
	 // default is INFO
	metricsLevel: flink.metricsLevel_PARALLELISM,
	 // default is APPLICATION
	autoScalingEnabled: jsii.Boolean(false),
	 // default is true
	parallelism: jsii.Number(32),
	 // default is 1
	parallelismPerKpu: jsii.Number(2),
	 // default is 1
	snapshotsEnabled: jsii.Boolean(false),
	 // default is true
	logGroup: logs.NewLogGroup(this, jsii.String("LogGroup")),
})
```
