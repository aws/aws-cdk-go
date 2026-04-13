package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   containerDependencyProperty := &ContainerDependencyProperty{
//   	Condition: jsii.String("condition"),
//   	ContainerName: jsii.String("containerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-containerdependency.html
//
type CfnDaemonTaskDefinitionPropsMixin_ContainerDependencyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-containerdependency.html#cfn-ecs-daemontaskdefinition-containerdependency-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-containerdependency.html#cfn-ecs-daemontaskdefinition-containerdependency-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

