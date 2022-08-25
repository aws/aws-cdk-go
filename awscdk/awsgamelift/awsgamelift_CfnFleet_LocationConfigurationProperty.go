package awsgamelift


// A remote location where a multi-location fleet can deploy EC2 instances for game hosting.
//
// *Related actions*
//
// [CreateFleet](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateFleet.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationConfigurationProperty := &locationConfigurationProperty{
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	locationCapacity: &locationCapacityProperty{
//   		desiredEc2Instances: jsii.Number(123),
//   		maxSize: jsii.Number(123),
//   		minSize: jsii.Number(123),
//   	},
//   }
//
type CfnFleet_LocationConfigurationProperty struct {
	// An AWS Region code, such as `us-west-2` .
	Location *string `field:"required" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// *Related actions*
	//
	// [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) | [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) | [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
	LocationCapacity interface{} `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
}

