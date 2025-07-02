package awsec2


// Specifies the market (purchasing) option for an instance.
//
// `InstanceMarketOptions` is a property of the [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMarketOptionsProperty := &InstanceMarketOptionsProperty{
//   	MarketType: jsii.String("marketType"),
//   	SpotOptions: &SpotOptionsProperty{
//   		BlockDurationMinutes: jsii.Number(123),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		MaxPrice: jsii.String("maxPrice"),
//   		SpotInstanceType: jsii.String("spotInstanceType"),
//   		ValidUntil: jsii.String("validUntil"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancemarketoptions.html
//
type CfnLaunchTemplate_InstanceMarketOptionsProperty struct {
	// The market type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancemarketoptions.html#cfn-ec2-launchtemplate-instancemarketoptions-markettype
	//
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// The options for Spot Instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancemarketoptions.html#cfn-ec2-launchtemplate-instancemarketoptions-spotoptions
	//
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

