package previewawsnotificationscontactsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEmailContactPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEmailContactMixinProps := &CfnEmailContactMixinProps{
//   	EmailAddress: jsii.String("emailAddress"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notificationscontacts-emailcontact.html
//
type CfnEmailContactMixinProps struct {
	// The email address of the contact.
	//
	// The activation and notification emails are sent here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notificationscontacts-emailcontact.html#cfn-notificationscontacts-emailcontact-emailaddress
	//
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The name of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notificationscontacts-emailcontact.html#cfn-notificationscontacts-emailcontact-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of tags to apply to the email contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notificationscontacts-emailcontact.html#cfn-notificationscontacts-emailcontact-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

