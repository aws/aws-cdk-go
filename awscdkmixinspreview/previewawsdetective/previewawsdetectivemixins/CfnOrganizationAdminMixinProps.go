package previewawsdetectivemixins


// Properties for CfnOrganizationAdminPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationAdminMixinProps := &CfnOrganizationAdminMixinProps{
//   	AccountId: jsii.String("accountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-organizationadmin.html
//
type CfnOrganizationAdminMixinProps struct {
	// The AWS account identifier of the account to designate as the Detective administrator account for the organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-organizationadmin.html#cfn-detective-organizationadmin-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

