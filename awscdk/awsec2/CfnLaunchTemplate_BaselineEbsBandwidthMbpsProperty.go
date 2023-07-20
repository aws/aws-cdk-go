package awsec2


// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
//
// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineEbsBandwidthMbpsProperty := &BaselineEbsBandwidthMbpsProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html
//
type CfnLaunchTemplate_BaselineEbsBandwidthMbpsProperty struct {
	// The maximum baseline bandwidth, in Mbps.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html#cfn-ec2-launchtemplate-baselineebsbandwidthmbps-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum baseline bandwidth, in Mbps.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html#cfn-ec2-launchtemplate-baselineebsbandwidthmbps-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

