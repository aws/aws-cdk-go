package awsmanagedblockchain


// Configuration properties of a peer node within a membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeConfigurationProperty := &NodeConfigurationProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	InstanceType: jsii.String("instanceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html
//
type CfnNode_NodeConfigurationProperty struct {
	// The Availability Zone in which the node exists.
	//
	// Required for Ethereum nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html#cfn-managedblockchain-node-nodeconfiguration-availabilityzone
	//
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The Amazon Managed Blockchain instance type for the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html#cfn-managedblockchain-node-nodeconfiguration-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

