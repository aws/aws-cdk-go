package mixinsawsec2


// Specifies the CPU performance to consider when using an instance family as the baseline reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnLaunchTemplatePropsMixin_CpuProperty struct {
	// The instance family to use as the baseline reference for CPU performance.
	//
	// All instance types that match your specified attributes are compared against the CPU performance of the referenced instance family, regardless of CPU manufacturer or architecture differences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpu.html#cfn-ec2-launchtemplate-cpu-references
	//
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

