# CloudWatch Alarm Actions library

This library contains a set of classes which can be used as CloudWatch Alarm actions.

The currently implemented actions are: EC2 Actions, SNS Actions, SSM OpsCenter Actions, Autoscaling Actions and Application Autoscaling Actions

## EC2 Action Example

```go
// Alarm must be configured with an EC2 per-instance metric
var alarm alarm

// Attach a reboot when alarm triggers
alarm.AddAlarmAction(
actions.NewEc2Action(actions.Ec2InstanceAction_REBOOT))
```

## SSM OpsCenter Action Example

```go
var alarm alarm

// Create an OpsItem with specific severity and category when alarm triggers
alarm.AddAlarmAction(
actions.NewSsmAction(actions.OpsItemSeverity_CRITICAL, actions.OpsItemCategory_PERFORMANCE))
```

See `@aws-cdk/aws-cloudwatch` for more information.
