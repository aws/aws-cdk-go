package awsconnect


// The type of notification recipient.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRecipientTypeProperty := &NotificationRecipientTypeProperty{
//   	UserArns: []*string{
//   		jsii.String("userArns"),
//   	},
//   	UserTags: map[string]*string{
//   		"userTagsKey": jsii.String("userTags"),
//   	},
//   }
//
type CfnRule_NotificationRecipientTypeProperty struct {
	// The Amazon Resource Name (ARN) of the user account.
	UserArns *[]*string `field:"optional" json:"userArns" yaml:"userArns"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }. Amazon Connect users with the specified tags will be notified.
	UserTags interface{} `field:"optional" json:"userTags" yaml:"userTags"`
}

