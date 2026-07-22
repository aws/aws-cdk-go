package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepAliasMappingProperty := &StepAliasMappingProperty{
//   	StepAlias: jsii.String("stepAlias"),
//   	StepId: jsii.String("stepId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html
//
type CfnFlow_StepAliasMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html#cfn-quicksight-flow-stepaliasmapping-stepalias
	//
	StepAlias *string `field:"required" json:"stepAlias" yaml:"stepAlias"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-flow-stepaliasmapping.html#cfn-quicksight-flow-stepaliasmapping-stepid
	//
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
}

