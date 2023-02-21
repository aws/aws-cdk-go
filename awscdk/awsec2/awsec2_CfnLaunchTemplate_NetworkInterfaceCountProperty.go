package awsec2


// The minimum and maximum number of network interfaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceCountProperty := &networkInterfaceCountProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnLaunchTemplate_NetworkInterfaceCountProperty struct {
	// The maximum number of network interfaces.
	//
	// To specify no maximum limit, omit this parameter.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of network interfaces.
	//
	// To specify no minimum limit, omit this parameter.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

