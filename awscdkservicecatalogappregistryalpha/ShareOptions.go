// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The options that are passed into a share of an Application or Attribute Group.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var application application
//   var myRole iRole
//   var myUser iUser
//
//   application.shareApplication(&ShareOptions{
//   	Accounts: []*string{
//   		jsii.String("123456789012"),
//   	},
//   	OrganizationArns: []*string{
//   		jsii.String("arn:aws:organizations::123456789012:organization/o-my-org-id"),
//   	},
//   	Roles: []*iRole{
//   		myRole,
//   	},
//   	Users: []*iUser{
//   		myUser,
//   	},
//   })
//
// Experimental.
type ShareOptions struct {
	// A list of AWS accounts that the application will be shared with.
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// A list of AWS Organization or Organizational Units (OUs) ARNs that the application will be shared with.
	// Experimental.
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
	// A list of AWS IAM roles that the application will be shared with.
	// Experimental.
	Roles *[]awsiam.IRole `field:"optional" json:"roles" yaml:"roles"`
	// An option to manage access to the application or attribute group.
	// Experimental.
	SharePermission interface{} `field:"optional" json:"sharePermission" yaml:"sharePermission"`
	// A list of AWS IAM users that the application will be shared with.
	// Experimental.
	Users *[]awsiam.IUser `field:"optional" json:"users" yaml:"users"`
}

