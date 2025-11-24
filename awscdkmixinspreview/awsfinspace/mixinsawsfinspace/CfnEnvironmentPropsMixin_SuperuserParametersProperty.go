package mixinsawsfinspace


// Configuration information for the superuser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   superuserParametersProperty := &SuperuserParametersProperty{
//   	EmailAddress: jsii.String("emailAddress"),
//   	FirstName: jsii.String("firstName"),
//   	LastName: jsii.String("lastName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html
//
type CfnEnvironmentPropsMixin_SuperuserParametersProperty struct {
	// The email address of the superuser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-emailaddress
	//
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The first name of the superuser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name of the superuser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
}

