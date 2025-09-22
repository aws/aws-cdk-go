package awsdetective


// A reference to a OrganizationAdmin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationAdminReference := &OrganizationAdminReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type OrganizationAdminReference struct {
	// The AccountId of the OrganizationAdmin resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

