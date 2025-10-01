package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   totalLocalStorageGBRequestProperty := &TotalLocalStorageGBRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html
//
type CfnCapacityProvider_TotalLocalStorageGBRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html#cfn-ecs-capacityprovider-totallocalstoragegbrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html#cfn-ecs-capacityprovider-totallocalstoragegbrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

