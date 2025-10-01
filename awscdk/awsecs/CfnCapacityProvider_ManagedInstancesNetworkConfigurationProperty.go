package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesNetworkConfigurationProperty := &ManagedInstancesNetworkConfigurationProperty{
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html
//
type CfnCapacityProvider_ManagedInstancesNetworkConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html#cfn-ecs-capacityprovider-managedinstancesnetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html#cfn-ecs-capacityprovider-managedinstancesnetworkconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

