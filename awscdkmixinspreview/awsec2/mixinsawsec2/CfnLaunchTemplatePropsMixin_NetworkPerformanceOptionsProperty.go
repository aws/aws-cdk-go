package mixinsawsec2


// Contains settings for the network performance options for the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkPerformanceOptionsProperty := &NetworkPerformanceOptionsProperty{
//   	BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkperformanceoptions.html
//
type CfnLaunchTemplatePropsMixin_NetworkPerformanceOptionsProperty struct {
	// Specify the bandwidth weighting option to boost the associated type of baseline bandwidth, as follows:.
	//
	// - **default** - This option uses the standard bandwidth configuration for your instance type.
	// - **vpc-1** - This option boosts your networking baseline bandwidth and reduces your EBS baseline bandwidth.
	// - **ebs-1** - This option boosts your EBS baseline bandwidth and reduces your networking baseline bandwidth.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkperformanceoptions.html#cfn-ec2-launchtemplate-networkperformanceoptions-bandwidthweighting
	//
	BandwidthWeighting *string `field:"optional" json:"bandwidthWeighting" yaml:"bandwidthWeighting"`
}

