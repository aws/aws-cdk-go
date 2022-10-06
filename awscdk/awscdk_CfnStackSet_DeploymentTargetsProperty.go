// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentTargetsProperty := &deploymentTargetsProperty{
//   	accountFilterType: jsii.String("accountFilterType"),
//   	accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	organizationalUnitIds: []*string{
//   		jsii.String("organizationalUnitIds"),
//   	},
//   }
//
type CfnStackSet_DeploymentTargetsProperty struct {
	// `CfnStackSet.DeploymentTargetsProperty.AccountFilterType`.
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// The names of one or more AWS accounts for which you want to deploy stack set updates.
	//
	// *Pattern* : `^[0-9]{12}$`.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	//
	// *Pattern* : `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

