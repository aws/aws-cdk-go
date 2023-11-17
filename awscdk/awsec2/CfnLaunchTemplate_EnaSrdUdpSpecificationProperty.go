package awsec2


// ENA Express is compatible with both TCP and UDP transport protocols.
//
// When it's enabled, TCP traffic automatically uses it. However, some UDP-based applications are designed to handle network packets that are out of order, without a need for retransmission, such as live video broadcasting or other near-real-time applications. For UDP traffic, you can specify whether to use ENA Express, based on your application environment needs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enaSrdUdpSpecificationProperty := &EnaSrdUdpSpecificationProperty{
//   	EnaSrdUdpEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdudpspecification.html
//
type CfnLaunchTemplate_EnaSrdUdpSpecificationProperty struct {
	// Indicates whether UDP traffic to and from the instance uses ENA Express.
	//
	// To specify this setting, you must first enable ENA Express.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdudpspecification.html#cfn-ec2-launchtemplate-enasrdudpspecification-enasrdudpenabled
	//
	EnaSrdUdpEnabled interface{} `field:"optional" json:"enaSrdUdpEnabled" yaml:"enaSrdUdpEnabled"`
}

