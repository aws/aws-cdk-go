package awsmanagedblockchain


// Configuration properties of a peer node within a membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeConfigurationProperty := &nodeConfigurationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	instanceType: jsii.String("instanceType"),
//   }
//
type CfnNode_NodeConfigurationProperty struct {
	// The Availability Zone in which the node exists.
	//
	// Required for Ethereum nodes.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The Amazon Managed Blockchain instance type for the node.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

