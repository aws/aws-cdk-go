package awsec2


// The minimum and maximum amount of total local storage, in GB.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-totallocalstoragegbrequest.html
//
type CfnSpotFleet_TotalLocalStorageGBRequestProperty struct {
	// The maximum amount of total local storage, in GB.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-totallocalstoragegbrequest.html#cfn-ec2-spotfleet-totallocalstoragegbrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of total local storage, in GB.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-totallocalstoragegbrequest.html#cfn-ec2-spotfleet-totallocalstoragegbrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

