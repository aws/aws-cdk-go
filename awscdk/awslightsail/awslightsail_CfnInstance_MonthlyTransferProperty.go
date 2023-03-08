package awslightsail


// `MonthlyTransfer` is a property of the [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html) property. It describes the amount of allocated monthly data transfer (in GB) for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monthlyTransferProperty := &monthlyTransferProperty{
//   	gbPerMonthAllocated: jsii.String("gbPerMonthAllocated"),
//   }
//
type CfnInstance_MonthlyTransferProperty struct {
	// The amount of allocated monthly data transfer (in GB) for an instance.
	GbPerMonthAllocated *string `field:"optional" json:"gbPerMonthAllocated" yaml:"gbPerMonthAllocated"`
}

