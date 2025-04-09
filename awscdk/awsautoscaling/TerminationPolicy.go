package awsautoscaling


// Specifies the termination criteria to apply before Amazon EC2 Auto Scaling chooses an instance for termination.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//   var arn string
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	TerminationPolicies: []terminationPolicy{
//   		autoscaling.*terminationPolicy_CUSTOM_LAMBDA_FUNCTION,
//   		autoscaling.*terminationPolicy_OLDEST_INSTANCE,
//   		autoscaling.*terminationPolicy_DEFAULT,
//   	},
//
//   	//terminationPolicyCustomLambdaFunctionArn property must be specified if the TerminationPolicy.CUSTOM_LAMBDA_FUNCTION is used
//   	TerminationPolicyCustomLambdaFunctionArn: arn,
//   })
//
type TerminationPolicy string

const (
	// Terminate instances in the Auto Scaling group to align the remaining instances to the allocation strategy for the type of instance that is terminating (either a Spot Instance or an On-Demand Instance).
	TerminationPolicy_ALLOCATION_STRATEGY TerminationPolicy = "ALLOCATION_STRATEGY"
	// Terminate instances that are closest to the next billing hour.
	TerminationPolicy_CLOSEST_TO_NEXT_INSTANCE_HOUR TerminationPolicy = "CLOSEST_TO_NEXT_INSTANCE_HOUR"
	// Terminate instances according to the default termination policy.
	TerminationPolicy_DEFAULT TerminationPolicy = "DEFAULT"
	// Terminate the newest instance in the group.
	TerminationPolicy_NEWEST_INSTANCE TerminationPolicy = "NEWEST_INSTANCE"
	// Terminate the oldest instance in the group.
	TerminationPolicy_OLDEST_INSTANCE TerminationPolicy = "OLDEST_INSTANCE"
	// Terminate instances that have the oldest launch configuration.
	TerminationPolicy_OLDEST_LAUNCH_CONFIGURATION TerminationPolicy = "OLDEST_LAUNCH_CONFIGURATION"
	// Terminate instances that have the oldest launch template.
	TerminationPolicy_OLDEST_LAUNCH_TEMPLATE TerminationPolicy = "OLDEST_LAUNCH_TEMPLATE"
	// Terminate instances using custom termination policy with lambda.
	//
	// If this is
	// specified, you must also supply a value of lambda arn in the terminationPolicyCustomLambdaFunctionArn property.
	//
	// If there are multiple termination policies specified, the custom termination policy with lambda
	// must be specified first in the order.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/lambda-custom-termination-policy.html#lambda-custom-termination-policy-limitations
	//
	TerminationPolicy_CUSTOM_LAMBDA_FUNCTION TerminationPolicy = "CUSTOM_LAMBDA_FUNCTION"
)

