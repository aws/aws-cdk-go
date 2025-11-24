package mixinsawsodb


// A contact to receive notification from Oracle about maintenance updates for a specific Exadata infrastructure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerContactProperty := &CustomerContactProperty{
//   	Email: jsii.String("email"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-customercontact.html
//
type CfnCloudExadataInfrastructurePropsMixin_CustomerContactProperty struct {
	// The email address of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-customercontact.html#cfn-odb-cloudexadatainfrastructure-customercontact-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
}

