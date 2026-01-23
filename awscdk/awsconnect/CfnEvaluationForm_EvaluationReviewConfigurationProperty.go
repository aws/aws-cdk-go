package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationReviewConfigurationProperty := &EvaluationReviewConfigurationProperty{
//   	ReviewNotificationRecipients: []interface{}{
//   		&EvaluationReviewNotificationRecipientProperty{
//   			Type: jsii.String("type"),
//   			Value: &EvaluationReviewNotificationRecipientValueProperty{
//   				UserId: jsii.String("userId"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	EligibilityDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewconfiguration.html
//
type CfnEvaluationForm_EvaluationReviewConfigurationProperty struct {
	// List of recipients who should be notified when a review is requested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewconfiguration.html#cfn-connect-evaluationform-evaluationreviewconfiguration-reviewnotificationrecipients
	//
	ReviewNotificationRecipients interface{} `field:"required" json:"reviewNotificationRecipients" yaml:"reviewNotificationRecipients"`
	// Number of days during which a request for review can be submitted for evaluations created from this form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationreviewconfiguration.html#cfn-connect-evaluationform-evaluationreviewconfiguration-eligibilitydays
	//
	EligibilityDays *float64 `field:"optional" json:"eligibilityDays" yaml:"eligibilityDays"`
}

