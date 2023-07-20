package awstimestream


// Details on SNS that are required to send the notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsConfigurationProperty := &SnsConfigurationProperty{
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-snsconfiguration.html
//
type CfnScheduledQuery_SnsConfigurationProperty struct {
	// SNS topic ARN that the scheduled query status notifications will be sent to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-snsconfiguration.html#cfn-timestream-scheduledquery-snsconfiguration-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

