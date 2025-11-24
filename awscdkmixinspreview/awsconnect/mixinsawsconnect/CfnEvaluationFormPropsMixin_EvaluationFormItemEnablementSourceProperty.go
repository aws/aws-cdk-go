package mixinsawsconnect


// An enablement expression source item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormItemEnablementSourceProperty := &EvaluationFormItemEnablementSourceProperty{
//   	RefId: jsii.String("refId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormItemEnablementSourceProperty struct {
	// A referenceId of the source item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html#cfn-connect-evaluationform-evaluationformitemenablementsource-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// A type of source item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html#cfn-connect-evaluationform-evaluationformitemenablementsource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

