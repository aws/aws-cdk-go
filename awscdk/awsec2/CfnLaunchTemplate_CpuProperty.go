package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cpuProperty := &CpuProperty{
//   	References: []interface{}{
//   		&ReferenceProperty{
//   			InstanceFamily: jsii.String("instanceFamily"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpu.html
//
type CfnLaunchTemplate_CpuProperty struct {
	// A list of references to be used as baseline for the CPU performance.
	//
	// Currently, you can only specify a single reference across different instance type variations such as CPU manufacturers, architectures etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpu.html#cfn-ec2-launchtemplate-cpu-references
	//
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

