package mixinsawsnotificationscontacts


// Configures email contacts for AWS User Notifications .
//
// You must activate the email contact using the activation token that you will receive when the email contact is set up.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   emailContactProperty := &EmailContactProperty{
//   	Address: jsii.String("address"),
//   	Arn: jsii.String("arn"),
//   	CreationTime: jsii.String("creationTime"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   	UpdateTime: jsii.String("updateTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html
//
type CfnEmailContactPropsMixin_EmailContactProperty struct {
	// The email address of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The Amazon Resource Name (ARN) of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The creation time of the `EmailContact` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-creationtime
	//
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The name of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the contact.
	//
	// Only activated contacts receive emails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The time the `EmailContact` was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-updatetime
	//
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}

