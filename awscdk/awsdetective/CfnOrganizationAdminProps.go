package awsdetective


// Properties for defining a `CfnOrganizationAdmin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationAdminProps := &CfnOrganizationAdminProps{
//   	AccountId: jsii.String("accountId"),
//   }
//
type CfnOrganizationAdminProps struct {
	// The AWS account identifier of the account to designate as the Detective administrator account for the organization.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

