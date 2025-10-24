package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for deploying with Stackset, which creates a StackSet constraint.
//
// Example:
//   var portfolio Portfolio
//   var product CloudFormationProduct
//
//
//   adminRole := iam.NewRole(this, jsii.String("AdminRole"), &RoleProps{
//   	AssumedBy: iam.NewAccountRootPrincipal(),
//   })
//
//   portfolio.deployWithStackSets(product, &StackSetsConstraintOptions{
//   	Accounts: []*string{
//   		jsii.String("012345678901"),
//   		jsii.String("012345678902"),
//   		jsii.String("012345678903"),
//   	},
//   	Regions: []*string{
//   		jsii.String("us-west-1"),
//   		jsii.String("us-east-1"),
//   		jsii.String("us-west-2"),
//   		jsii.String("us-east-1"),
//   	},
//   	AdminRole: adminRole,
//   	ExecutionRoleName: jsii.String("SCStackSetExecutionRole"),
//   	 // Name of role deployed in end users accounts.
//   	AllowStackSetInstanceOperations: jsii.Boolean(true),
//   })
//
type StackSetsConstraintOptions struct {
	// The description of the constraint.
	// Default: - No description provided.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Default: - English.
	//
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// List of accounts to deploy stacks to.
	Accounts *[]*string `field:"required" json:"accounts" yaml:"accounts"`
	// IAM role used to administer the StackSets configuration.
	AdminRole awsiam.IRoleRef `field:"required" json:"adminRole" yaml:"adminRole"`
	// IAM role used to provision the products in the Stacks.
	ExecutionRoleName *string `field:"required" json:"executionRoleName" yaml:"executionRoleName"`
	// List of regions to deploy stacks to.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Wether to allow end users to create, update, and delete stacks.
	// Default: false.
	//
	AllowStackSetInstanceOperations *bool `field:"optional" json:"allowStackSetInstanceOperations" yaml:"allowStackSetInstanceOperations"`
}

