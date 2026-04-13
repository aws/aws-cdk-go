package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   systemControlProperty := &SystemControlProperty{
//   	Namespace: jsii.String("namespace"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-systemcontrol.html
//
type CfnDaemonTaskDefinitionPropsMixin_SystemControlProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-systemcontrol.html#cfn-ecs-daemontaskdefinition-systemcontrol-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-systemcontrol.html#cfn-ecs-daemontaskdefinition-systemcontrol-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

