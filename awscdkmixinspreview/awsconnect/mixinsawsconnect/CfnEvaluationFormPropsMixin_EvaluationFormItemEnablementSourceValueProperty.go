package mixinsawsconnect


// An enablement expression source value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormItemEnablementSourceValueProperty := &EvaluationFormItemEnablementSourceValueProperty{
//   	RefId: jsii.String("refId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormItemEnablementSourceValueProperty struct {
	// A referenceId of the source value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html#cfn-connect-evaluationform-evaluationformitemenablementsourcevalue-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
	// A type of source item value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.html#cfn-connect-evaluationform-evaluationformitemenablementsourcevalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

