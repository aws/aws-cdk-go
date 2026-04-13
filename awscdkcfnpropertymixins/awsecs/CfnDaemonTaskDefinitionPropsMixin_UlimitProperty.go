package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ulimitProperty := &UlimitProperty{
//   	HardLimit: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SoftLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-ulimit.html
//
type CfnDaemonTaskDefinitionPropsMixin_UlimitProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-ulimit.html#cfn-ecs-daemontaskdefinition-ulimit-hardlimit
	//
	HardLimit *float64 `field:"optional" json:"hardLimit" yaml:"hardLimit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-ulimit.html#cfn-ecs-daemontaskdefinition-ulimit-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-ulimit.html#cfn-ecs-daemontaskdefinition-ulimit-softlimit
	//
	SoftLimit *float64 `field:"optional" json:"softLimit" yaml:"softLimit"`
}

