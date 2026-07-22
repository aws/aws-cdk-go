package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stepAliasMappingProperty := &StepAliasMappingProperty{
//   	StepAlias: jsii.String("stepAlias"),
//   	StepId: jsii.String("stepId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html
//
type CfnFlowPropsMixin_StepAliasMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html#cfn-quicksight-flow-stepaliasmapping-stepalias
	//
	StepAlias *string `field:"optional" json:"stepAlias" yaml:"stepAlias"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html#cfn-quicksight-flow-stepaliasmapping-stepid
	//
	StepId *string `field:"optional" json:"stepId" yaml:"stepId"`
}

