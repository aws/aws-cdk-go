# CloudWatch Alarm Actions library

This library contains a set of classes which can be used as CloudWatch Alarm actions.

The currently implemented actions are: EC2 Actions, SNS Actions, SSM OpsCenter Actions, Autoscaling Actions and Application Autoscaling Actions

## EC2 Action Example

```go
// Alarm must be configured with an EC2 per-instance metric
var alarm Alarm

// Attach a reboot when alarm triggers
alarm.AddAlarmAction(
actions.NewEc2Action(actions.Ec2InstanceAction_REBOOT))
```

## SSM OpsCenter Action Example

```go
var alarm Alarm

// Create an OpsItem with specific severity and category when alarm triggers
alarm.AddAlarmAction(
actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
```

## SSM Incident Manager Action Example

```go
var alarm Alarm

// Create an Incident Manager incident based on a specific response plan
alarm.AddAlarmAction(
actions.NewSsmIncidentAction(jsii.String("ResponsePlanName")))
```

## Lambda Action Example

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var alarm Alarm
var fn Function
var alias Alias
var version Version


// Attach a Lambda Function when alarm triggers
alarm.AddAlarmAction(
actions.NewLambdaAction(fn))

// Attach a Lambda Function Alias when alarm triggers
alarm.AddAlarmAction(
actions.NewLambdaAction(alias))

// Attach a Lambda Function version when alarm triggers
alarm.AddAlarmAction(
actions.NewLambdaAction(version))
```

See `aws-cdk-lib/aws-cloudwatch` for more information.
