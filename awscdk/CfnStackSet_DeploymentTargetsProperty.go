package awscdk


// The AWS Organizations accounts or AWS accounts to deploy stacks to in the specified Regions.
//
// When deploying to AWS Organizations accounts with `SERVICE_MANAGED` permissions:
//
// - You must specify the `OrganizationalUnitIds` property.
// - If you specify organizational units (OUs) for `OrganizationalUnitIds` and use either the `Accounts` or `AccountsUrl` property, you must also specify the `AccountFilterType` property.
//
// When deploying to AWS accounts with `SELF_MANAGED` permissions:
//
// - You must specify either the `Accounts` or `AccountsUrl` property, but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentTargetsProperty := &DeploymentTargetsProperty{
//   	AccountFilterType: jsii.String("accountFilterType"),
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	AccountsUrl: jsii.String("accountsUrl"),
//   	OrganizationalUnitIds: []*string{
//   		jsii.String("organizationalUnitIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html
//
type CfnStackSet_DeploymentTargetsProperty struct {
	// Refines which accounts to deploy stacks to by specifying how to use the `Accounts` and `OrganizationalUnitIds` properties together.
	//
	// The following values determine how CloudFormation selects target accounts:
	//
	// - `INTERSECTION` : StackSet deploys to the accounts specified in the `Accounts` property.
	// - `DIFFERENCE` : StackSet deploys to the OU, excluding the accounts specified in the `Accounts` property.
	// - `UNION` : StackSet deploys to the OU, and the accounts specified in the `Accounts` property. `UNION` is not supported for create operations when using StackSet as a resource or the `CreateStackInstances` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accountfiltertype
	//
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// The account IDs of the AWS accounts .
	//
	// If you have many account numbers, you can provide those accounts using the `AccountsUrl` property instead.
	//
	// *Pattern* : `^[0-9]{12}$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accounts
	//
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The Amazon S3 URL path to a file that contains a list of AWS account IDs.
	//
	// The file format must be either `.csv` or `.txt` , and the data can be comma-separated or new-line-separated. There is currently a 10MB limit for the data (approximately 800,000 accounts).
	//
	// This property serves the same purpose as `Accounts` but allows you to specify a large number of accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accountsurl
	//
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// The organization root ID or organizational unit (OU) IDs.
	//
	// *Pattern* : `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-organizationalunitids
	//
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

