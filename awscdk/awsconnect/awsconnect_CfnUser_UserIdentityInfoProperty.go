package awsconnect


// Contains information about the identity of a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userIdentityInfoProperty := &userIdentityInfoProperty{
//   	email: jsii.String("email"),
//   	firstName: jsii.String("firstName"),
//   	lastName: jsii.String("lastName"),
//   	mobile: jsii.String("mobile"),
//   	secondaryEmail: jsii.String("secondaryEmail"),
//   }
//
type CfnUser_UserIdentityInfoProperty struct {
	// The email address.
	//
	// If you are using SAML for identity management and include this parameter, an error is returned.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The first name.
	//
	// This is required if you are using Amazon Connect or SAML for identity management.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name.
	//
	// This is required if you are using Amazon Connect or SAML for identity management.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// `CfnUser.UserIdentityInfoProperty.Mobile`.
	Mobile *string `field:"optional" json:"mobile" yaml:"mobile"`
	// `CfnUser.UserIdentityInfoProperty.SecondaryEmail`.
	SecondaryEmail *string `field:"optional" json:"secondaryEmail" yaml:"secondaryEmail"`
}

