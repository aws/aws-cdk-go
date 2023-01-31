# Amazon CloudWatch Synthetics Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Developer Preview](https://img.shields.io/badge/cdk--constructs-developer--preview-informational.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are in **developer preview** before they
> become stable. We will only make breaking changes to address unforeseen API issues. Therefore,
> these APIs are not subject to [Semantic Versioning](https://semver.org/), and breaking changes
> will be announced in release notes. This means that while you may use them, you may need to
> update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

Amazon CloudWatch Synthetics allow you to monitor your application by generating **synthetic** traffic. The traffic is produced by a **canary**: a configurable script that runs on a schedule. You configure the canary script to follow the same routes and perform the same actions as a user, which allows you to continually verify your user experience even when you don't have any traffic on your applications.

## Canary

To illustrate how to use a canary, assume your application defines the following endpoint:

```console
% curl "https://api.example.com/user/books/topbook/"
The Hitchhikers Guide to the Galaxy

```

The below code defines a canary that will hit the `books/topbook` endpoint every 5 minutes:

```go
canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &canaryProps{
	schedule: synthetics.schedule.rate(awscdk.Duration.minutes(jsii.Number(5))),
	test: synthetics.test.custom(&customTestOptions{
		code: synthetics.code.fromAsset(path.join(__dirname, jsii.String("canary"))),
		handler: jsii.String("index.handler"),
	}),
	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
	environmentVariables: map[string]*string{
		"stage": jsii.String("prod"),
	},
})
```

The following is an example of an `index.js` file which exports the `handler` function:

```js
const synthetics = require('Synthetics');
const log = require('SyntheticsLogger');

const pageLoadBlueprint = async function () {
  // Configure the stage of the API using environment variables
  const url = `https://api.example.com/${process.env.stage}/user/books/topbook/`;

  const page = await synthetics.getPage();
  const response = await page.goto(url, { waitUntil: 'domcontentloaded', timeout: 30000 });
  // Wait for page to render. Increase or decrease wait time based on endpoint being monitored.
  await page.waitFor(15000);
  // This will take a screenshot that will be included in test output artifacts.
  await synthetics.takeScreenshot('loaded', 'loaded');
  const pageTitle = await page.title();
  log.info('Page title: ' + pageTitle);
  if (response.status() !== 200) {
    throw 'Failed to load page!';
  }
};

exports.handler = async () => {
  return await pageLoadBlueprint();
};
```

> **Note:** The function **must** be called `handler`.

The canary will automatically produce a CloudWatch Dashboard:

![UI Screenshot](images/ui-screenshot.png)

The Canary code will be executed in a lambda function created by Synthetics on your behalf. The Lambda function includes a custom [runtime](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) provided by Synthetics. The provided runtime includes a variety of handy tools such as [Puppeteer](https://www.npmjs.com/package/puppeteer-core) (for nodejs based one) and Chromium.

To learn more about Synthetics capabilities, check out the [docs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries.html).

### Canary Schedule

You can specify the schedule on which a canary runs by providing a
[`Schedule`](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-synthetics.Schedule.html)
object to the `schedule` property.

Configure a run rate of up to 60 minutes with `Schedule.rate`:

```go
schedule := synthetics.schedule.rate(awscdk.Duration.minutes(jsii.Number(5)))
```

You can also specify a [cron expression](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html) with `Schedule.cron`:

```go
schedule := synthetics.schedule.cron(&cronOptions{
	hour: jsii.String("0,8,16"),
})
```

If you want the canary to run just once upon deployment, you can use `Schedule.once()`.

### Canary DeleteLambdaResourcesOnCanaryDeletion

You can specify whether the AWS CloudFormation is to also delete the Lambda functions and layers used by this canary, when the canary is deleted.

This can be provisioned by setting the `DeleteLambdaResourcesOnCanaryDeletion` property to `true` when we define the canary.

```go
stack := awscdk.Newstack()

canary := synthetics.NewCanary(stack, jsii.String("Canary"), &canaryProps{
	test: synthetics.test.custom(&customTestOptions{
		handler: jsii.String("index.handler"),
		code: synthetics.code.fromInline(jsii.String("/* Synthetics handler code")),
	}),
	enableAutoDeleteLambdas: jsii.Boolean(true),
	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
})
```

Even when set to `true` there are resources such as S3 buckets/logs that do NOT get deleted and are to be deleted manually.

### Configuring the Canary Script

To configure the script the canary executes, use the `test` property. The `test` property accepts a `Test` instance that can be initialized by the `Test` class static methods. Currently, the only implemented method is `Test.custom()`, which allows you to bring your own code. In the future, other methods will be added. `Test.custom()` accepts `code` and `handler` properties -- both are required by Synthetics to create a lambda function on your behalf.

The `synthetics.Code` class exposes static methods to bundle your code artifacts:

* `code.fromInline(code)` - specify an inline script.
* `code.fromAsset(path)` - specify a .zip file or a directory in the local filesystem which will be zipped and uploaded to S3 on deployment. See the above Note for directory structure.
* `code.fromBucket(bucket, key[, objectVersion])` - specify an S3 object that contains the .zip file of your runtime code. See the above Note for directory structure.

Using the `Code` class static initializers:

```go
// To supply the code from a S3 bucket:
import s3 "github.com/aws/aws-cdk-go/awscdk"
// To supply the code inline:
// To supply the code inline:
synthetics.NewCanary(this, jsii.String("Inline Canary"), &canaryProps{
	test: synthetics.test.custom(&customTestOptions{
		code: synthetics.code.fromInline(jsii.String("/* Synthetics handler code */")),
		handler: jsii.String("index.handler"),
	}),
	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
})

// To supply the code from your local filesystem:
// To supply the code from your local filesystem:
synthetics.NewCanary(this, jsii.String("Asset Canary"), &canaryProps{
	test: synthetics.*test.custom(&customTestOptions{
		code: synthetics.*code.fromAsset(path.join(__dirname, jsii.String("canary"))),
		handler: jsii.String("index.handler"),
	}),
	runtime: synthetics.*runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
})
bucket := s3.NewBucket(this, jsii.String("Code Bucket"))
synthetics.NewCanary(this, jsii.String("Bucket Canary"), &canaryProps{
	test: synthetics.*test.custom(&customTestOptions{
		code: synthetics.*code.fromBucket(bucket, jsii.String("canary.zip")),
		handler: jsii.String("index.handler"),
	}),
	runtime: synthetics.*runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
})
```

> **Note:** Synthetics have a specified folder structure for canaries. For Node scripts supplied via `code.fromAsset()` or `code.fromBucket()`, the canary resource requires the following folder structure:
>
> ```plaintext
> canary/
> ├── nodejs/
>    ├── node_modules/
>         ├── <filename>.js
> ```
>
> For Python scripts supplied via `code.fromAsset()` or `code.fromBucket()`, the canary resource requires the following folder structure:
>
> ```plaintext
> canary/
> ├── python/
>     ├── <filename>.py
> ```
>
> See Synthetics [docs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html).

### Running a canary on a VPC

You can specify what [VPC a canary executes in](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).
This can allow for monitoring services that may be internal to a specific VPC. To place a canary within a VPC, you can specify the `vpc` property with the desired `VPC` to place then canary in.
This will automatically attach the appropriate IAM permissions to attach to the VPC. This will also create a Security Group and attach to the default subnets for the VPC unless specified via `vpcSubnets` and `securityGroups`.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var vpc iVpc

synthetics.NewCanary(this, jsii.String("Vpc Canary"), &canaryProps{
	test: synthetics.test.custom(&customTestOptions{
		code: synthetics.code.fromAsset(path.join(__dirname, jsii.String("canary"))),
		handler: jsii.String("index.handler"),
	}),
	runtime: synthetics.runtime_SYNTHETICS_NODEJS_PUPPETEER_3_8(),
	vpc: vpc,
})
```

> **Note:** By default, the Synthetics runtime needs access to the S3 and CloudWatch APIs, which will fail in a private subnet without internet access enabled (e.g. an isolated subnnet).
>
> Ensure that the Canary is placed in a VPC either with internet connectivity or with VPC Endpoints for S3 and CloudWatch enabled and configured.
>
> See [Synthetics VPC docs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).

### Alarms

You can configure a CloudWatch Alarm on a canary metric. Metrics are emitted by CloudWatch automatically and can be accessed by the following APIs:

* `canary.metricSuccessPercent()` - percentage of successful canary runs over a given time
* `canary.metricDuration()` - how much time each canary run takes, in seconds.
* `canary.metricFailed()` - number of failed canary runs over a given time

Create an alarm that tracks the canary metric:

```go
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

var canary canary

cloudwatch.NewAlarm(this, jsii.String("CanaryAlarm"), &alarmProps{
	metric: canary.metricSuccessPercent(),
	evaluationPeriods: jsii.Number(2),
	threshold: jsii.Number(90),
	comparisonOperator: cloudwatch.comparisonOperator_LESS_THAN_THRESHOLD,
})
```
