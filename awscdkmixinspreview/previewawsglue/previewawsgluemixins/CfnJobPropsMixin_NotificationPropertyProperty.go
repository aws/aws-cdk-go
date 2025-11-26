package previewawsgluemixins


// Specifies configuration properties of a notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationPropertyProperty := &NotificationPropertyProperty{
//   	NotifyDelayAfter: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-notificationproperty.html
//
type CfnJobPropsMixin_NotificationPropertyProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-notificationproperty.html#cfn-glue-job-notificationproperty-notifydelayafter
	//
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

