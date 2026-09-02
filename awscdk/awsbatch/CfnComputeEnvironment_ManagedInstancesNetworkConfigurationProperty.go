package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesNetworkConfigurationProperty := &ManagedInstancesNetworkConfigurationProperty{
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesnetworkconfiguration.html
//
type CfnComputeEnvironment_ManagedInstancesNetworkConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesnetworkconfiguration.html#cfn-batch-computeenvironment-managedinstancesnetworkconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesnetworkconfiguration.html#cfn-batch-computeenvironment-managedinstancesnetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

