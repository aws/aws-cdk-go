package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRecipientTypeProperty := &notificationRecipientTypeProperty{
//   	userArns: []*string{
//   		jsii.String("userArns"),
//   	},
//   	userTags: map[string]*string{
//   		"userTagsKey": jsii.String("userTags"),
//   	},
//   }
//
type CfnRule_NotificationRecipientTypeProperty struct {
	// `CfnRule.NotificationRecipientTypeProperty.UserArns`.
	UserArns *[]*string `field:"optional" json:"userArns" yaml:"userArns"`
	// `CfnRule.NotificationRecipientTypeProperty.UserTags`.
	UserTags interface{} `field:"optional" json:"userTags" yaml:"userTags"`
}

