package previewawsdeadlinemixins


// The allowable range of vCPU processing power for the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vCpuCountRangeProperty := &VCpuCountRangeProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-vcpucountrange.html
//
type CfnFleetPropsMixin_VCpuCountRangeProperty struct {
	// The maximum amount of vCPU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-vcpucountrange.html#cfn-deadline-fleet-vcpucountrange-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of vCPU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-vcpucountrange.html#cfn-deadline-fleet-vcpucountrange-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

