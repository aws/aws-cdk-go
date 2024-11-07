package awsconnect


// Contains information about the identity of a user.
//
// > For Amazon Connect instances that are created with the `EXISTING_DIRECTORY` identity management type, `FirstName` , `LastName` , and `Email` cannot be updated from within Amazon Connect because they are managed by the directory. > The `FirstName` and `LastName` length constraints below apply only to instances using SAML for identity management. If you are using Amazon Connect for identity management, the length constraints are 1-255 for `FirstName` , and 1-256 for `LastName` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userIdentityInfoProperty := &UserIdentityInfoProperty{
//   	Email: jsii.String("email"),
//   	FirstName: jsii.String("firstName"),
//   	LastName: jsii.String("lastName"),
//   	Mobile: jsii.String("mobile"),
//   	SecondaryEmail: jsii.String("secondaryEmail"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html
//
type CfnUser_UserIdentityInfoProperty struct {
	// The email address.
	//
	// If you are using SAML for identity management and include this parameter, an error is returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html#cfn-connect-user-useridentityinfo-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The first name.
	//
	// This is required if you are using Amazon Connect or SAML for identity management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html#cfn-connect-user-useridentityinfo-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name.
	//
	// This is required if you are using Amazon Connect or SAML for identity management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html#cfn-connect-user-useridentityinfo-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The user's mobile number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html#cfn-connect-user-useridentityinfo-mobile
	//
	Mobile *string `field:"optional" json:"mobile" yaml:"mobile"`
	// The user's secondary email address.
	//
	// If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.
	//
	// *Pattern* : `(?=^.{0,265}$)[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,63}`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-useridentityinfo.html#cfn-connect-user-useridentityinfo-secondaryemail
	//
	SecondaryEmail *string `field:"optional" json:"secondaryEmail" yaml:"secondaryEmail"`
}

