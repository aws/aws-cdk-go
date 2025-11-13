package awsconnect


// An enablement expression source item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormItemEnablementSourceProperty := &EvaluationFormItemEnablementSourceProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	RefId: jsii.String("refId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html
//
type CfnEvaluationForm_EvaluationFormItemEnablementSourceProperty struct {
	// A type of source item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html#cfn-connect-evaluationform-evaluationformitemenablementsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A referenceId of the source item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementsource.html#cfn-connect-evaluationform-evaluationformitemenablementsource-refid
	//
	RefId *string `field:"optional" json:"refId" yaml:"refId"`
}

