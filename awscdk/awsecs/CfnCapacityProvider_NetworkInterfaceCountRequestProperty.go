package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceCountRequestProperty := &NetworkInterfaceCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html
//
type CfnCapacityProvider_NetworkInterfaceCountRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html#cfn-ecs-capacityprovider-networkinterfacecountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html#cfn-ecs-capacityprovider-networkinterfacecountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

