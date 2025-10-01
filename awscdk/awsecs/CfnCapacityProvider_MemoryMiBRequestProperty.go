package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryMiBRequestProperty := &MemoryMiBRequestProperty{
//   	Min: jsii.Number(123),
//
//   	// the properties below are optional
//   	Max: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html
//
type CfnCapacityProvider_MemoryMiBRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html#cfn-ecs-capacityprovider-memorymibrequest-min
	//
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html#cfn-ecs-capacityprovider-memorymibrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

