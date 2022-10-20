package awscloudwatchactions


// Types of EC2 actions available.
//
// Example:
//   // Alarm must be configured with an EC2 per-instance metric
//   var alarm alarm
//
//   // Attach a reboot when alarm triggers
//   alarm.addAlarmAction(
//   actions.NewEc2Action(actions.ec2InstanceAction_REBOOT))
//
type Ec2InstanceAction string

const (
	// Stop the instance.
	Ec2InstanceAction_STOP Ec2InstanceAction = "STOP"
	// Terminatethe instance.
	Ec2InstanceAction_TERMINATE Ec2InstanceAction = "TERMINATE"
	// Recover the instance.
	Ec2InstanceAction_RECOVER Ec2InstanceAction = "RECOVER"
	// Reboot the instance.
	Ec2InstanceAction_REBOOT Ec2InstanceAction = "REBOOT"
)

