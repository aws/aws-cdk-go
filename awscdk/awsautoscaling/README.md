# Amazon EC2 Auto Scaling Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Auto Scaling Group

An `AutoScalingGroup` represents a number of instances on which you run your code. You
pick the size of the fleet, the instance type and the OS image:

```go
var vpc vpc


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_MICRO),
	machineImage: ec2.NewAmazonLinuxImage(),
})
```

NOTE: AutoScalingGroup has an property called `allowAllOutbound` (allowing the instances to contact the
internet) which is set to `true` by default. Be sure to set this to `false`  if you don't want
your instances to be able to start arbitrary connections. Alternatively, you can specify an existing security
group to attach to the instances that are launched, rather than have the group create a new one.

```go
var vpc vpc


mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &securityGroupProps{
	vpc: vpc,
})
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_MICRO),
	machineImage: ec2.NewAmazonLinuxImage(),
	securityGroup: mySecurityGroup,
})
```

Alternatively you can create an `AutoScalingGroup` from a `LaunchTemplate`:

```go
var vpc vpc
var launchTemplate launchTemplate


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	launchTemplate: launchTemplate,
})
```

To launch a mixture of Spot and on-demand instances, and/or with multiple instance types, you can create an `AutoScalingGroup` from a MixedInstancesPolicy:

```go
var vpc vpc
var launchTemplate1 launchTemplate
var launchTemplate2 launchTemplate


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	mixedInstancesPolicy: &mixedInstancesPolicy{
		instancesDistribution: &instancesDistribution{
			onDemandPercentageAboveBaseCapacity: jsii.Number(50),
		},
		launchTemplate: launchTemplate1,
		launchTemplateOverrides: []launchTemplateOverrides{
			&launchTemplateOverrides{
				instanceType: ec2.NewInstanceType(jsii.String("t3.micro")),
			},
			&launchTemplateOverrides{
				instanceType: ec2.NewInstanceType(jsii.String("t3a.micro")),
			},
			&launchTemplateOverrides{
				instanceType: ec2.NewInstanceType(jsii.String("t4g.micro")),
				launchTemplate: launchTemplate2,
			},
		},
	},
})
```

## Machine Images (AMIs)

AMIs control the OS that gets launched when you start your EC2 instance. The EC2
library contains constructs to select the AMI you want to use.

Depending on the type of AMI, you select it a different way.

The latest version of Amazon Linux and Microsoft Windows images are
selectable by instantiating one of these classes:

```go
// Pick a Windows edition to use
windows := ec2.NewWindowsImage(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)

// Pick the right Amazon Linux edition. All arguments shown are optional
// and will default to these values when omitted.
amznLinux := ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
	edition: ec2.amazonLinuxEdition_STANDARD,
	virtualization: ec2.amazonLinuxVirt_HVM,
	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
})

// For other custom (Linux) images, instantiate a `GenericLinuxImage` with
// a map giving the AMI to in for each region:

linux := ec2.NewGenericLinuxImage(map[string]*string{
	"us-east-1": jsii.String("ami-97785bed"),
	"eu-west-1": jsii.String("ami-12345678"),
})
```

> NOTE: The Amazon Linux images selected will be cached in your `cdk.json`, so that your
> AutoScalingGroups don't automatically change out from under you when you're making unrelated
> changes. To update to the latest version of Amazon Linux, remove the cache entry from the `context`
> section of your `cdk.json`.
>
> We will add command-line options to make this step easier in the future.

## AutoScaling Instance Counts

AutoScalingGroups make it possible to raise and lower the number of instances in the group,
in response to (or in advance of) changes in workload.

When you create your AutoScalingGroup, you specify a `minCapacity` and a
`maxCapacity`. AutoScaling policies that respond to metrics will never go higher
or lower than the indicated capacity (but scheduled scaling actions might, see
below).

There are three ways to scale your capacity:

* **In response to a metric** (also known as step scaling); for example, you
  might want to scale out if the CPU usage across your cluster starts to rise,
  and scale in when it drops again.
* **By trying to keep a certain metric around a given value** (also known as
  target tracking scaling); you might want to automatically scale out and in to
  keep your CPU usage around 50%.
* **On a schedule**; you might want to organize your scaling around traffic
  flows you expect, by scaling out in the morning and scaling in in the
  evening.

The general pattern of autoscaling will look like this:

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	minCapacity: jsii.Number(5),
	maxCapacity: jsii.Number(100),
})
```

### Step Scaling

This type of scaling scales in and out in deterministics steps that you
configure, in response to metric values. For example, your scaling strategy to
scale in response to a metric that represents your average worker pool usage
might look like this:

```plaintext
 Scaling        -1          (no change)          +1       +3
            │        │                       │        │        │
            ├────────┼───────────────────────┼────────┼────────┤
            │        │                       │        │        │
Worker use  0%      10%                     50%       70%     100%
```

(Note that this is not necessarily a recommended scaling strategy, but it's
a possible one. You will have to determine what thresholds are right for you).

Note that in order to set up this scaling strategy, you will have to emit a
metric representing your worker utilization from your instances. After that,
you would configure the scaling something like this:

```go
var autoScalingGroup autoScalingGroup


workerUtilizationMetric := cloudwatch.NewMetric(&metricProps{
	namespace: jsii.String("MyService"),
	metricName: jsii.String("WorkerUtilization"),
})

autoScalingGroup.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
	metric: workerUtilizationMetric,
	scalingSteps: []scalingInterval{
		&scalingInterval{
			upper: jsii.Number(10),
			change: -jsii.Number(1),
		},
		&scalingInterval{
			lower: jsii.Number(50),
			change: +jsii.Number(1),
		},
		&scalingInterval{
			lower: jsii.Number(70),
			change: +jsii.Number(3),
		},
	},

	// Change this to AdjustmentType.PERCENT_CHANGE_IN_CAPACITY to interpret the
	// 'change' numbers before as percentages instead of capacity counts.
	adjustmentType: autoscaling.adjustmentType_CHANGE_IN_CAPACITY,
})
```

The AutoScaling construct library will create the required CloudWatch alarms and
AutoScaling policies for you.

### Target Tracking Scaling

This type of scaling scales in and out in order to keep a metric around a value
you prefer. There are four types of predefined metrics you can track, or you can
choose to track a custom metric. If you do choose to track a custom metric,
be aware that the metric has to represent instance utilization in some way
(AutoScaling will scale out if the metric is higher than the target, and scale
in if the metric is lower than the target).

If you configure multiple target tracking policies, AutoScaling will use the
one that yields the highest capacity.

The following example scales to keep the CPU usage of your instances around
50% utilization:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnCpuUtilization(jsii.String("KeepSpareCPU"), &cpuUtilizationScalingProps{
	targetUtilizationPercent: jsii.Number(50),
})
```

To scale on average network traffic in and out of your instances:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnIncomingBytes(jsii.String("LimitIngressPerInstance"), &networkUtilizationScalingProps{
	targetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
})
autoScalingGroup.scaleOnOutgoingBytes(jsii.String("LimitEgressPerInstance"), &networkUtilizationScalingProps{
	targetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
})
```

To scale on the average request count per instance (only works for
AutoScalingGroups that have been attached to Application Load
Balancers):

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnRequestCount(jsii.String("LimitRPS"), &requestCountScalingProps{
	targetRequestsPerSecond: jsii.Number(1000),
})
```

### Scheduled Scaling

This type of scaling is used to change capacities based on time. It works by
changing `minCapacity`, `maxCapacity` and `desiredCapacity` of the
AutoScalingGroup, and so can be used for two purposes:

* Scale in and out on a schedule by setting the `minCapacity` high or
  the `maxCapacity` low.
* Still allow the regular scaling actions to do their job, but restrict
  the range they can scale over (by setting both `minCapacity` and
  `maxCapacity` but changing their range over time).

A schedule is expressed as a cron expression. The `Schedule` class has a `cron` method to help build cron expressions.

The following example scales the fleet out in the morning, going back to natural
scaling (all the way down to 1 instance if necessary) at night:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &basicScheduledActionProps{
	schedule: autoscaling.schedule.cron(&cronOptions{
		hour: jsii.String("8"),
		minute: jsii.String("0"),
	}),
	minCapacity: jsii.Number(20),
})

autoScalingGroup.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &basicScheduledActionProps{
	schedule: autoscaling.*schedule.cron(&cronOptions{
		hour: jsii.String("20"),
		minute: jsii.String("0"),
	}),
	minCapacity: jsii.Number(1),
})
```

### Block Devices

This type specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.

#### GP3 Volumes

You can only specify the `throughput` on GP3 volumes.

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,
	blockDevices: []blockDevice{
		&blockDevice{
			deviceName: jsii.String("gp3-volume"),
			volume: autoscaling.blockDeviceVolume.ebs(jsii.Number(15), &ebsDeviceOptions{
				volumeType: autoscaling.ebsDeviceVolumeType_GP3,
				throughput: jsii.Number(125),
			}),
		},
	},
})
```

## Configuring Instances using CloudFormation Init

It is possible to use the CloudFormation Init mechanism to configure the
instances in the AutoScalingGroup. You can write files to it, run commands,
start services, etc. See the documentation of
[AWS::CloudFormation::Init](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-init.html)
and the documentation of CDK's `aws-ec2` library for more information.

When you specify a CloudFormation Init configuration for an AutoScalingGroup:

* you *must* also specify `signals` to configure how long CloudFormation
  should wait for the instances to successfully configure themselves.
* you *should* also specify an `updatePolicy` to configure how instances
  should be updated when the AutoScalingGroup is updated (for example,
  when the AMI is updated). If you don't specify an update policy, a *rolling
  update* is chosen by default.

Here's an example of using CloudFormation Init to write a file to the
instance hosts on startup:

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	init: ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/my_instance"), jsii.String("This got written during instance startup"))),
	signals: autoscaling.signals.waitForAll(&signalsOptions{
		timeout: awscdk.Duration.minutes(jsii.Number(10)),
	}),
})
```

## Signals

In normal operation, CloudFormation will send a Create or Update command to
an AutoScalingGroup and proceed with the rest of the deployment without waiting
for the *instances in the AutoScalingGroup*.

Configure `signals` to tell CloudFormation to wait for a specific number of
instances in the AutoScalingGroup to have been started (or failed to start)
before moving on. An instance is supposed to execute the
[`cfn-signal`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-signal.html)
program as part of its startup to indicate whether it was started
successfully or not.

If you use CloudFormation Init support (described in the previous section),
the appropriate call to `cfn-signal` is automatically added to the
AutoScalingGroup's UserData. If you don't use the `signals` directly, you are
responsible for adding such a call yourself.

The following type of `Signals` are available:

* `Signals.waitForAll([options])`: wait for all of `desiredCapacity` amount of instances
  to have started (recommended).
* `Signals.waitForMinCapacity([options])`: wait for a `minCapacity` amount of instances
  to have started (use this if waiting for all instances takes too long and you are happy
  with a minimum count of healthy hosts).
* `Signals.waitForCount(count, [options])`: wait for a specific amount of instances to have
  started.

There are two `options` you can configure:

* `timeout`: maximum time a host startup is allowed to take. If a host does not report
  success within this time, it is considered a failure. Default is 5 minutes.
* `minSuccessPercentage`: percentage of hosts that needs to be healthy in order for the
  update to succeed. If you set this value lower than 100, some percentage of hosts may
  report failure, while still considering the deployment a success. Default is 100%.

## Update Policy

The *update policy* describes what should happen to running instances when the definition
of the AutoScalingGroup is changed. For example, if you add a command to the UserData
of an AutoScalingGroup, do the existing instances get replaced with new instances that
have executed the new UserData? Or do the "old" instances just keep on running?

It is recommended to always use an update policy, otherwise the current state of your
instances also depends the previous state of your instances, rather than just on your
source code. This degrades the reproducibility of your deployments.

The following update policies are available:

* `UpdatePolicy.none()`: leave existing instances alone (not recommended).
* `UpdatePolicy.rollingUpdate([options])`: progressively replace the existing
  instances with new instances, in small batches. At any point in time,
  roughly the same amount of total instances will be running. If the deployment
  needs to be rolled back, the fresh instances will be replaced with the "old"
  configuration again.
* `UpdatePolicy.replacingUpdate([options])`: build a completely fresh copy
  of the new AutoScalingGroup next to the old one. Once the AutoScalingGroup
  has been successfully created (and the instances started, if `signals` is
  configured on the AutoScalingGroup), the old AutoScalingGroup is deleted.
  If the deployment needs to be rolled back, the new AutoScalingGroup is
  deleted and the old one is left unchanged.

## Allowing Connections

See the documentation of the `@aws-cdk/aws-ec2` package for more information
about allowing connections between resources backed by instances.

## Max Instance Lifetime

To enable the max instance lifetime support, specify `maxInstanceLifetime` property
for the `AutoscalingGroup` resource. The value must be between 7 and 365 days(inclusive).
To clear a previously set value, leave this property undefined.

## Instance Monitoring

To disable detailed instance monitoring, specify `instanceMonitoring` property
for the `AutoscalingGroup` resource as `Monitoring.BASIC`. Otherwise detailed monitoring
will be enabled.

## Monitoring Group Metrics

Group metrics are used to monitor group level properties; they describe the group rather than any of its instances (e.g GroupMaxSize, the group maximum size). To enable group metrics monitoring, use the `groupMetrics` property.
All group metrics are reported in a granularity of 1 minute at no additional charge.

See [EC2 docs](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html#as-group-metrics) for a list of all available group metrics.

To enable group metrics monitoring using the `groupMetrics` property:

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


// Enable monitoring of all group metrics
// Enable monitoring of all group metrics
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	groupMetrics: []groupMetrics{
		autoscaling.*groupMetrics.all(),
	},
})

// Enable monitoring for a subset of group metrics
// Enable monitoring for a subset of group metrics
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	groupMetrics: []*groupMetrics{
		autoscaling.NewGroupMetrics(autoscaling.groupMetric_MIN_SIZE(), autoscaling.*groupMetric_MAX_SIZE()),
	},
})
```

## Termination policies

Auto Scaling uses [termination policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
to determine which instances it terminates first during scale-in events. You
can specify one or more termination policies with the `terminationPolicies`
property:

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	terminationPolicies: []terminationPolicy{
		autoscaling.*terminationPolicy_OLDEST_INSTANCE,
		autoscaling.*terminationPolicy_DEFAULT,
	},
})
```

## Protecting new instances from being terminated on scale-in

By default, Auto Scaling can terminate an instance at any time after launch when
scaling in an Auto Scaling Group, subject to the group's [termination
policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html).

However, you may wish to protect newly-launched instances from being scaled in
if they are going to run critical applications that should not be prematurely
terminated. EC2 Capacity Providers for Amazon ECS requires this attribute be
set to `true`.

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	newInstancesProtectedFromScaleIn: jsii.Boolean(true),
})
```

## Configuring Instance Metadata Service (IMDS)

### Toggling IMDSv1

You can configure [EC2 Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) options to either
allow both IMDSv1 and IMDSv2 or enforce IMDSv2 when interacting with the IMDS.

To do this for a single `AutoScalingGroup`, you can use set the `requireImdsv2` property.
The example below demonstrates IMDSv2 being required on a single `AutoScalingGroup`:

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	requireImdsv2: jsii.Boolean(true),
})
```

You can also use `AutoScalingGroupRequireImdsv2Aspect` to apply the operation to multiple AutoScalingGroups.
The example below demonstrates the `AutoScalingGroupRequireImdsv2Aspect` being used to require IMDSv2 for all AutoScalingGroups in a stack:

```go
aspect := autoscaling.NewAutoScalingGroupRequireImdsv2Aspect()

awscdk.Aspects.of(this).add(aspect)
```

## Warm Pool

Auto Scaling offers [a warm pool](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html) which gives an ability to decrease latency for applications that have exceptionally long boot times. You can create a warm pool with default parameters as below:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.addWarmPool()
```

You can also customize a warm pool by configuring parameters:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.addWarmPool(&warmPoolOptions{
	minSize: jsii.Number(1),
	reuseOnScaleIn: jsii.Boolean(true),
})
```

### Default Instance Warming

You can use the default instance warmup feature to improve the Amazon CloudWatch metrics used for dynamic scaling.
When default instance warmup is not enabled, each instance starts contributing usage data to the aggregated metrics
as soon as the instance reaches the InService state. However, if you enable default instance warmup, this lets
your instances finish warming up before they contribute the usage data.

To optimize the performance of scaling policies that scale continuously, such as target tracking and step scaling
policies, we strongly recommend that you enable the default instance warmup, even if its value is set to 0 seconds.

To set up Default Instance Warming for an autoscaling group, simply pass it in as a prop

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	defaultInstanceWarmup: awscdk.Duration.seconds(jsii.Number(5)),
})
```

## Future work

* [ ] CloudWatch Events (impossible to add currently as the AutoScalingGroup ARN is
  necessary to make this rule and this cannot be accessed from CloudFormation).
