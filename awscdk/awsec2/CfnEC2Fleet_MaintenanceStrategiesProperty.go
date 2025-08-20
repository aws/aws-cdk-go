package awsec2


// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceStrategiesProperty := &MaintenanceStrategiesProperty{
//   	CapacityRebalance: &CapacityRebalanceProperty{
//   		ReplacementStrategy: jsii.String("replacementStrategy"),
//   		TerminationDelay: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-maintenancestrategies.html
//
type CfnEC2Fleet_MaintenanceStrategiesProperty struct {
	// The strategy to use when Amazon EC2 emits a signal that your Spot Instance is at an elevated risk of being interrupted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-maintenancestrategies.html#cfn-ec2-ec2fleet-maintenancestrategies-capacityrebalance
	//
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

