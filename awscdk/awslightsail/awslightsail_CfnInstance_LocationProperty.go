package awslightsail


// `Location` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the location for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnInstance_LocationProperty struct {
	// The Availability Zone for the instance.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the AWS Region for the instance.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

