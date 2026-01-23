package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationReviewNotificationRecipientValueProperty := &EvaluationReviewNotificationRecipientValueProperty{
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipientvalue.html
//
type CfnEvaluationFormPropsMixin_EvaluationReviewNotificationRecipientValueProperty struct {
	// The user identifier for the notification recipient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipientvalue.html#cfn-connect-evaluationform-evaluationreviewnotificationrecipientvalue-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

