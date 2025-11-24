package mixinsawsecs


// The minimum and maximum number of network interfaces for instance type selection.
//
// This is useful for workloads that require multiple network interfaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceCountRequestProperty := &NetworkInterfaceCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html
//
type CfnCapacityProviderPropsMixin_NetworkInterfaceCountRequestProperty struct {
	// The maximum number of network interfaces.
	//
	// Instance types that support more network interfaces are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html#cfn-ecs-capacityprovider-networkinterfacecountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of network interfaces.
	//
	// Instance types that support fewer network interfaces are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkinterfacecountrequest.html#cfn-ecs-capacityprovider-networkinterfacecountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

