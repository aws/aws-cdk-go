package awsec2


// Allows customer to specify ENA-SRD options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enaSrdSpecificationProperty := &EnaSrdSpecificationProperty{
//   	EnaSrdEnabled: jsii.Boolean(false),
//   	EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   		EnaSrdUdpEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdspecification.html
//
type CfnLaunchTemplate_EnaSrdSpecificationProperty struct {
	// Enables TCP ENA-SRD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdspecification.html#cfn-ec2-launchtemplate-enasrdspecification-enasrdenabled
	//
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// Allows customer to specify ENA-SRD (UDP) options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdspecification.html#cfn-ec2-launchtemplate-enasrdspecification-enasrdudpspecification
	//
	EnaSrdUdpSpecification interface{} `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

