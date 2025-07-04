package awseks


// Effect types of kubernetes node taint.
//
// Example:
//   var cluster cluster
//
//   cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
//   	InstanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	Taints: []taintSpec{
//   		&taintSpec{
//   			Effect: eks.TaintEffect_NO_SCHEDULE,
//   			Key: jsii.String("foo"),
//   			Value: jsii.String("bar"),
//   		},
//   	},
//   })
//
type TaintEffect string

const (
	// NoSchedule.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// NoExecute.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

