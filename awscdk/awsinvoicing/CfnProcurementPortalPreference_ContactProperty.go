package awsinvoicing


// Contact information for a person or role associated with the procurement portal preference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactProperty := &ContactProperty{
//   	Email: jsii.String("email"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-contact.html
//
type CfnProcurementPortalPreference_ContactProperty struct {
	// The email address of the contact person or role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-contact.html#cfn-invoicing-procurementportalpreference-contact-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the contact person or role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-contact.html#cfn-invoicing-procurementportalpreference-contact-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

