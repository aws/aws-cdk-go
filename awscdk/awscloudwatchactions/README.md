# CloudWatch Alarm Actions library

This library contains a set of classes which can be used as CloudWatch Alarm actions.

The currently implemented actions are: EC2 Actions, SNS Actions, SSM OpsCenter Actions, Autoscaling Actions and Application Autoscaling Actions

## EC2 Action Example

```go
// Alarm must be configured with an EC2 per-instance metric
var alarm alarm

// Attach a reboot when alarm triggers
alarm.addAlarmAction(
actions.NewEc2Action(actions.ec2InstanceAction_REBOOT))
```

## SSM OpsCenter Action Example

```go
var alarm alarm

// Create an OpsItem with specific severity and category when alarm triggers
alarm.addAlarmAction(
actions.NewSsmAction(actions.opsItemSeverity_CRITICAL, actions.opsItemCategory_PERFORMANCE))
```

## SSM Incident Manager Action Example

```go
var alarm alarm

// Create an Incident Manager incident based on a specific response plan
alarm.addAlarmAction(
actions.NewSsmIncidentAction(jsii.String("ResponsePlanName")))
```

See `@aws-cdk/aws-cloudwatch` for more information.
