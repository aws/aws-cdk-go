package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationReviewNotificationRecipientProperty := &EvaluationReviewNotificationRecipientProperty{
//   	Type: jsii.String("type"),
//   	Value: &EvaluationReviewNotificationRecipientValueProperty{
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipient.html
//
type CfnEvaluationForm_EvaluationReviewNotificationRecipientProperty struct {
	// The type of notification recipient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipient.html#cfn-connect-evaluationform-evaluationreviewnotificationrecipient-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipient.html#cfn-connect-evaluationform-evaluationreviewnotificationrecipient-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

