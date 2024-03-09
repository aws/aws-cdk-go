# Amazon EC2 Auto Scaling Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Auto Scaling Group

An `AutoScalingGroup` represents a number of instances on which you run your code. You
pick the size of the fleet, the instance type and the OS image:

```go
var vpc vpc


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),

	// The latest Amazon Linux 2 image
	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
})
```

Creating an `AutoScalingGroup` from a Launch Configuration has been deprecated. All new accounts created after December 31, 2023 will no longer be able to create Launch Configurations. With the `@aws-cdk/aws-autoscaling:generateLaunchTemplateInsteadOfLaunchConfig` feature flag set to true, `AutoScalingGroup` properties used to create a Launch Configuration will now be used to create a `LaunchTemplate` using a [Launch Configuration to `LaunchTemplate` mapping](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html#launch-configuration-mapping-reference). Specifically, the following `AutoScalingGroup` properties will be used to generate a `LaunchTemplate`:

* machineImage
* keyName
* instanceType
* instanceMonitoring
* securityGroup
* role
* userData
* associatePublicIpAddress
* spotPrice
* blockDevices

After the Launch Configuration is replaced with a `LaunchTemplate`, any new instances launched by the `AutoScalingGroup` will use the new `LaunchTemplate`. Existing instances are not affected. To update an existing instance, you can allow the `AutoScalingGroup` to gradually replace existing instances with new instances based on the `terminationPolicies` for the `AutoScalingGroup`. Alternatively, you can terminate them yourself and force the `AutoScalingGroup` to launch new instances to maintain the `desiredCapacity`.

Support for creating an `AutoScalingGroup` from a `LaunchTemplate` was added in CDK version 2.21.0. Users on a CDK version earlier than version 2.21.0 that need to create an `AutoScalingGroup` with an account created after December 31, 2023 must update their CDK version to 2.21.0 or later. Users on CDK versions 2.21.0 up to, but not including 2.86.0, must use a manually created `LaunchTemplate` to create an `AutoScalingGroup` for accounts created after December 31, 2023. CDK version 2.86.0 or later will automatically generate a `LaunchTemplate` using the `AutoScalingGroup` properties mentioned above.

For additional migration information, please see: [Migrating to a `LaunchTemplate` from a Launch Configuration](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-to-launch-templates.html)

NOTE: AutoScalingGroup has a property called `allowAllOutbound` (allowing the instances to contact the
internet) which is set to `true` by default. Be sure to set this to `false`  if you don't want
your instances to be able to start arbitrary connections. Alternatively, you can specify an existing security
group to attach to the instances that are launched, rather than have the group create a new one.

```go
var vpc vpc


mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
	Vpc: Vpc,
})
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
	SecurityGroup: mySecurityGroup,
})
```

Alternatively, to enable more advanced features, you can create an `AutoScalingGroup` from a supplied `LaunchTemplate`:

```go
var vpc vpc
var launchTemplate launchTemplate


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	LaunchTemplate: launchTemplate,
})
```

To launch a mixture of Spot and on-demand instances, and/or with multiple instance types, you can create an `AutoScalingGroup` from a `MixedInstancesPolicy`:

```go
var vpc vpc
var launchTemplate1 launchTemplate
var launchTemplate2 launchTemplate


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	MixedInstancesPolicy: &MixedInstancesPolicy{
		InstancesDistribution: &InstancesDistribution{
			OnDemandPercentageAboveBaseCapacity: jsii.Number(50),
		},
		LaunchTemplate: launchTemplate1,
		LaunchTemplateOverrides: []launchTemplateOverrides{
			&launchTemplateOverrides{
				InstanceType: ec2.NewInstanceType(jsii.String("t3.micro")),
			},
			&launchTemplateOverrides{
				InstanceType: ec2.NewInstanceType(jsii.String("t3a.micro")),
			},
			&launchTemplateOverrides{
				InstanceType: ec2.NewInstanceType(jsii.String("t4g.micro")),
				LaunchTemplate: launchTemplate2,
			},
		},
	},
})
```

You can specify instances requirements with the `instanceRequirements ` property:

```go
var vpc vpc
var launchTemplate1 launchTemplate


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	MixedInstancesPolicy: &MixedInstancesPolicy{
		LaunchTemplate: launchTemplate1,
		LaunchTemplateOverrides: []launchTemplateOverrides{
			&launchTemplateOverrides{
				InstanceRequirements: &InstanceRequirementsProperty{
					VCpuCount: &VCpuCountRequestProperty{
						Min: jsii.Number(4),
						Max: jsii.Number(8),
					},
					MemoryMiB: &MemoryMiBRequestProperty{
						Min: jsii.Number(16384),
					},
					CpuManufacturers: []*string{
						jsii.String("intel"),
					},
				},
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
windows := ec2.NewWindowsImage(ec2.WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)

// Pick the right Amazon Linux edition. All arguments shown are optional
// and will default to these values when omitted.
amznLinux := ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
	Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX,
	Edition: ec2.AmazonLinuxEdition_STANDARD,
	Virtualization: ec2.AmazonLinuxVirt_HVM,
	Storage: ec2.AmazonLinuxStorage_GENERAL_PURPOSE,
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


autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	MinCapacity: jsii.Number(5),
	MaxCapacity: jsii.Number(100),
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


workerUtilizationMetric := cloudwatch.NewMetric(&MetricProps{
	Namespace: jsii.String("MyService"),
	MetricName: jsii.String("WorkerUtilization"),
})

autoScalingGroup.scaleOnMetric(jsii.String("ScaleToCPU"), &BasicStepScalingPolicyProps{
	Metric: workerUtilizationMetric,
	ScalingSteps: []scalingInterval{
		&scalingInterval{
			Upper: jsii.Number(10),
			Change: -jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(50),
			Change: +jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(70),
			Change: +jsii.Number(3),
		},
	},
	EvaluationPeriods: jsii.Number(10),
	DatapointsToAlarm: jsii.Number(5),

	// Change this to AdjustmentType.PERCENT_CHANGE_IN_CAPACITY to interpret the
	// 'change' numbers before as percentages instead of capacity counts.
	AdjustmentType: autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
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


autoScalingGroup.scaleOnCpuUtilization(jsii.String("KeepSpareCPU"), &CpuUtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(50),
})
```

To scale on average network traffic in and out of your instances:

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnIncomingBytes(jsii.String("LimitIngressPerInstance"), &NetworkUtilizationScalingProps{
	TargetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
})
autoScalingGroup.scaleOnOutgoingBytes(jsii.String("LimitEgressPerInstance"), &NetworkUtilizationScalingProps{
	TargetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
})
```

To scale on the average request count per instance (only works for
AutoScalingGroups that have been attached to Application Load
Balancers):

```go
var autoScalingGroup autoScalingGroup


autoScalingGroup.scaleOnRequestCount(jsii.String("LimitRPS"), &RequestCountScalingProps{
	TargetRequestsPerSecond: jsii.Number(1000),
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


autoScalingGroup.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &BasicScheduledActionProps{
	Schedule: autoscaling.Schedule_Cron(&CronOptions{
		Hour: jsii.String("8"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(20),
})

autoScalingGroup.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &BasicScheduledActionProps{
	Schedule: autoscaling.Schedule_*Cron(&CronOptions{
		Hour: jsii.String("20"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(1),
})
```

### Instance Maintenance Policy

You can configure an instance maintenance policy for your Auto Scaling group to
meet specific capacity requirements during events that cause instances to be replaced,
such as an instance refresh or the health check process.

For example, suppose you have an Auto Scaling group that has a small number of instances.
You want to avoid the potential disruptions from terminating and then replacing an instance
when health checks indicate an impaired instance. With an instance maintenance policy, you
can make sure that Amazon EC2 Auto Scaling first launches a new instance and then waits for
it to be fully ready before terminating the unhealthy instance.

An instance maintenance policy also helps you minimize any potential disruptions in cases where
multiple instances are replaced at the same time. You set the `minHealthyPercentage`
and the `maxHealthyPercentage` for the policy, and your Auto Scaling group can only
increase and decrease capacity within that minimum-maximum range when replacing instances.
A larger range increases the number of instances that can be replaced at the same time.

```go
var vpc vpc


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
	MaxHealthyPercentage: jsii.Number(200),
	MinHealthyPercentage: jsii.Number(100),
})
```

> Visit [Instance maintenance policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html) for more details.

### Block Devices

This type specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.

#### GP3 Volumes

You can only specify the `throughput` on GP3 volumes.

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,
	BlockDevices: []blockDevice{
		&blockDevice{
			DeviceName: jsii.String("gp3-volume"),
			Volume: autoscaling.BlockDeviceVolume_Ebs(jsii.Number(15), &EbsDeviceOptions{
				VolumeType: autoscaling.EbsDeviceVolumeType_GP3,
				Throughput: jsii.Number(125),
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


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	Init: ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/my_instance"), jsii.String("This got written during instance startup"))),
	Signals: autoscaling.Signals_WaitForAll(&SignalsOptions{
		Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
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

See the documentation of the `aws-cdk-lib/aws-ec2` package for more information
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
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	GroupMetrics: []groupMetrics{
		autoscaling.*groupMetrics_All(),
	},
})

// Enable monitoring for a subset of group metrics
// Enable monitoring for a subset of group metrics
autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	GroupMetrics: []*groupMetrics{
		autoscaling.NewGroupMetrics(autoscaling.GroupMetric_MIN_SIZE(), autoscaling.GroupMetric_MAX_SIZE()),
	},
})
```

## Termination policies

Auto Scaling uses [termination policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
to determine which instances it terminates first during scale-in events. You
can specify one or more termination policies with the `terminationPolicies`
property:

[Custom termination policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lambda-custom-termination-policy.html) with lambda
can be used to determine which instances to terminate based on custom logic.
The custom termination policy can be specified using `TerminationPolicy.CUSTOM_LAMBDA_FUNCTION`. If this is
specified, you must also supply a value of lambda arn in the `terminationPolicyCustomLambdaFunctionArn` property and
attach necessary [permission](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lambda-custom-termination-policy.html#lambda-custom-termination-policy-create-function)
to invoke the lambda function.

If there are multiple termination policies specified,
custom termination policy with lambda `TerminationPolicy.CUSTOM_LAMBDA_FUNCTION`
must be specified first.

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage
var arn string


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	TerminationPolicies: []terminationPolicy{
		autoscaling.*terminationPolicy_CUSTOM_LAMBDA_FUNCTION,
		autoscaling.*terminationPolicy_OLDEST_INSTANCE,
		autoscaling.*terminationPolicy_DEFAULT,
	},

	//terminationPolicyCustomLambdaFunctionArn property must be specified if the TerminationPolicy.CUSTOM_LAMBDA_FUNCTION is used
	TerminationPolicyCustomLambdaFunctionArn: arn,
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


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	NewInstancesProtectedFromScaleIn: jsii.Boolean(true),
})
```

## Configuring Capacity Rebalancing

Indicates whether Capacity Rebalancing is enabled. Otherwise, Capacity Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of interruption. After launching a new instance, it then terminates an old instance. For more information, see [Use Capacity Rebalancing to handle Amazon EC2 Spot Interruptions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the in the Amazon EC2 Auto Scaling User Guide.

```go
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	CapacityRebalance: jsii.Boolean(true),
})
```

## Connecting to your instances using SSM Session Manager

SSM Session Manager makes it possible to connect to your instances from the
AWS Console, without preparing SSH keys.

To do so, you need to:

* Use an image with [SSM agent](https://docs.aws.amazon.com/systems-manager/latest/userguide/ssm-agent.html) installed
  and configured. [Many images come with SSM Agent
  preinstalled](https://docs.aws.amazon.com/systems-manager/latest/userguide/ami-preinstalled-agent.html), otherwise you
  may need to manually put instructions to [install SSM
  Agent](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-manual-agent-install.html) into your
  instance's UserData or use EC2 Init).
* Create the AutoScalingGroup with `ssmSessionPermissions: true`.

If these conditions are met, you can connect to the instance from the EC2 Console. Example:

```go
var vpc vpc


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_MICRO),

	// Amazon Linux 2 comes with SSM Agent by default
	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),

	// Turn on SSM
	SsmSessionPermissions: jsii.Boolean(true),
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


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	RequireImdsv2: jsii.Boolean(true),
})
```

You can also use `AutoScalingGroupRequireImdsv2Aspect` to apply the operation to multiple AutoScalingGroups.
The example below demonstrates the `AutoScalingGroupRequireImdsv2Aspect` being used to require IMDSv2 for all AutoScalingGroups in a stack:

```go
aspect := autoscaling.NewAutoScalingGroupRequireImdsv2Aspect()

awscdk.Aspects_Of(this).Add(aspect)
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


autoScalingGroup.addWarmPool(&WarmPoolOptions{
	MinSize: jsii.Number(1),
	ReuseOnScaleIn: jsii.Boolean(true),
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


autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: InstanceType,
	MachineImage: MachineImage,

	// ...

	DefaultInstanceWarmup: awscdk.Duration_Seconds(jsii.Number(5)),
})
```

## Future work

* [ ] CloudWatch Events (impossible to add currently as the AutoScalingGroup ARN is
  necessary to make this rule and this cannot be accessed from CloudFormation).
