package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormItemEnablementSourceValueProperty := &EvaluationFormItemEnablementSourceValueProperty{
//   	RefId: jsii.String("refId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html
//
type CfnEvaluationForm_EvaluationFormItemEnablementSourceValueProperty struct {
	// The identifier to reference the item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html#cfn-connect-evaluationform-evaluationformitemenablementsourcevalue-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// Type of the source entity value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html#cfn-connect-evaluationform-evaluationformitemenablementsourcevalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

