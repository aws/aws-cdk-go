package awseks


// Effect types of kubernetes node taint.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("custom-node-group"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("m5.large")),
//   	},
//   	taints: []taintSpec{
//   		&taintSpec{
//   			effect: eks.taintEffect_NO_SCHEDULE,
//   			key: jsii.String("foo"),
//   			value: jsii.String("bar"),
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

