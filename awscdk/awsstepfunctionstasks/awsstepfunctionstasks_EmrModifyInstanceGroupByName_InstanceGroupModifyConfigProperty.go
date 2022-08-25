package awsstepfunctionstasks


// Modify the size or configurations of an instance group.
//
// Example:
//   tasks.NewEmrModifyInstanceGroupByName(this, jsii.String("Task"), &emrModifyInstanceGroupByNameProps{
//   	clusterId: jsii.String("ClusterId"),
//   	instanceGroupName: sfn.jsonPath.stringAt(jsii.String("$.InstanceGroupName")),
//   	instanceGroup: &instanceGroupModifyConfigProperty{
//   		instanceCount: jsii.Number(1),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroupModifyConfig.html
//
type EmrModifyInstanceGroupByName_InstanceGroupModifyConfigProperty struct {
	// A list of new or modified configurations to apply for an instance group.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// The EC2 InstanceIds to terminate.
	//
	// After you terminate the instances, the instance group will not return to its original requested size.
	EC2InstanceIdsToTerminate *[]*string `field:"optional" json:"eC2InstanceIdsToTerminate" yaml:"eC2InstanceIdsToTerminate"`
	// Target size for the instance group.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Policy for customizing shrink operations.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ShrinkPolicy.html
	//
	ShrinkPolicy *EmrModifyInstanceGroupByName_ShrinkPolicyProperty `field:"optional" json:"shrinkPolicy" yaml:"shrinkPolicy"`
}

