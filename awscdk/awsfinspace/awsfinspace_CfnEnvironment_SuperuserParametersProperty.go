package awsfinspace


// Configuration information for the superuser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   superuserParametersProperty := &superuserParametersProperty{
//   	emailAddress: jsii.String("emailAddress"),
//   	firstName: jsii.String("firstName"),
//   	lastName: jsii.String("lastName"),
//   }
//
type CfnEnvironment_SuperuserParametersProperty struct {
	// The email address of the superuser.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The first name of the superuser.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name of the superuser.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
}

