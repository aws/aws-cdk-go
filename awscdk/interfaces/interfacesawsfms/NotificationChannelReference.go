package interfacesawsfms


// A reference to a NotificationChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationChannelReference := &NotificationChannelReference{
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type NotificationChannelReference struct {
	// The SnsTopicArn of the NotificationChannel resource.
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

