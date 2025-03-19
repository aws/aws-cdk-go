package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkPerformanceOptionsProperty := &NetworkPerformanceOptionsProperty{
//   	BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkperformanceoptions.html
//
type CfnLaunchTemplate_NetworkPerformanceOptionsProperty struct {
	// Specifies the performance options of your instance or sets it to default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkperformanceoptions.html#cfn-ec2-launchtemplate-networkperformanceoptions-bandwidthweighting
	//
	BandwidthWeighting *string `field:"optional" json:"bandwidthWeighting" yaml:"bandwidthWeighting"`
}

