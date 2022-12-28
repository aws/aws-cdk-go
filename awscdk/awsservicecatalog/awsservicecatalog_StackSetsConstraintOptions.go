package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for deploying with Stackset, which creates a StackSet constraint.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   adminRole := iam.NewRole(this, jsii.String("AdminRole"), &roleProps{
//   	assumedBy: iam.NewAccountRootPrincipal(),
//   })
//
//   portfolio.deployWithStackSets(product, &stackSetsConstraintOptions{
//   	accounts: []*string{
//   		jsii.String("012345678901"),
//   		jsii.String("012345678902"),
//   		jsii.String("012345678903"),
//   	},
//   	regions: []*string{
//   		jsii.String("us-west-1"),
//   		jsii.String("us-east-1"),
//   		jsii.String("us-west-2"),
//   		jsii.String("us-east-1"),
//   	},
//   	adminRole: adminRole,
//   	executionRoleName: jsii.String("SCStackSetExecutionRole"),
//   	 // Name of role deployed in end users accounts.
//   	allowStackSetInstanceOperations: jsii.Boolean(true),
//   })
//
// Experimental.
type StackSetsConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// List of accounts to deploy stacks to.
	// Experimental.
	Accounts *[]*string `field:"required" json:"accounts" yaml:"accounts"`
	// IAM role used to administer the StackSets configuration.
	// Experimental.
	AdminRole awsiam.IRole `field:"required" json:"adminRole" yaml:"adminRole"`
	// IAM role used to provision the products in the Stacks.
	// Experimental.
	ExecutionRoleName *string `field:"required" json:"executionRoleName" yaml:"executionRoleName"`
	// List of regions to deploy stacks to.
	// Experimental.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Wether to allow end users to create, update, and delete stacks.
	// Experimental.
	AllowStackSetInstanceOperations *bool `field:"optional" json:"allowStackSetInstanceOperations" yaml:"allowStackSetInstanceOperations"`
}

