package previewawsec2mixins


// Describes the Traffic Mirror port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trafficMirrorPortRangeProperty := &TrafficMirrorPortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html
//
type CfnTrafficMirrorFilterRulePropsMixin_TrafficMirrorPortRangeProperty struct {
	// The start of the Traffic Mirror port range.
	//
	// This applies to the TCP and UDP protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrange-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The end of the Traffic Mirror port range.
	//
	// This applies to the TCP and UDP protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrange-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

