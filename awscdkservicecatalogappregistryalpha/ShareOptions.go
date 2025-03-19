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
//   application.shareApplication(jsii.String("MyShareId"), &ShareOptions{
//   	Name: jsii.String("MyShare"),
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
	// Name of the share.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of AWS accounts that the application will be shared with.
	// Default: - No accounts specified for share.
	//
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// A list of AWS Organization or Organizational Units (OUs) ARNs that the application will be shared with.
	// Default: - No AWS Organizations or OUs specified for share.
	//
	// Experimental.
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
	// A list of AWS IAM roles that the application will be shared with.
	// Default: - No IAM roles specified for share.
	//
	// Experimental.
	Roles *[]awsiam.IRole `field:"optional" json:"roles" yaml:"roles"`
	// An option to manage access to the application or attribute group.
	// Default: - Principals will be assigned read only permissions on the application or attribute group.
	//
	// Experimental.
	SharePermission interface{} `field:"optional" json:"sharePermission" yaml:"sharePermission"`
	// A list of AWS IAM users that the application will be shared with.
	// Default: - No IAM Users specified for share.
	//
	// Experimental.
	Users *[]awsiam.IUser `field:"optional" json:"users" yaml:"users"`
}

