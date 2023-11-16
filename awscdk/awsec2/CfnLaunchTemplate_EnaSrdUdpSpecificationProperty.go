package awsec2


// Allows customer to specify ENA-SRD (UDP) options.
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
	// Enables UDP ENA-SRD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdudpspecification.html#cfn-ec2-launchtemplate-enasrdudpspecification-enasrdudpenabled
	//
	EnaSrdUdpEnabled interface{} `field:"optional" json:"enaSrdUdpEnabled" yaml:"enaSrdUdpEnabled"`
}

