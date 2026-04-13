package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelCapabilitiesProperty := &KernelCapabilitiesProperty{
//   	Add: []*string{
//   		jsii.String("add"),
//   	},
//   	Drop: []*string{
//   		jsii.String("drop"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-kernelcapabilities.html
//
type CfnDaemonTaskDefinition_KernelCapabilitiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-kernelcapabilities.html#cfn-ecs-daemontaskdefinition-kernelcapabilities-add
	//
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-kernelcapabilities.html#cfn-ecs-daemontaskdefinition-kernelcapabilities-drop
	//
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

