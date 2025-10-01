package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorCountRequestProperty := &AcceleratorCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html
//
type CfnCapacityProvider_AcceleratorCountRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html#cfn-ecs-capacityprovider-acceleratorcountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-acceleratorcountrequest.html#cfn-ecs-capacityprovider-acceleratorcountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

