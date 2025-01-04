package awscloudformation


// The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
	// Limit deployment targets to individual accounts or include additional accounts with provided OUs.
	//
	// The following is a list of possible values for the `AccountFilterType` operation.
	//
	// - `INTERSECTION` : StackSet deploys to the accounts specified in the `Accounts` parameter.
	// - `DIFFERENCE` : StackSet deploys to the OU, excluding the accounts specified in the `Accounts` parameter.
	// - `UNION` StackSet deploys to the OU, and the accounts specified in the `Accounts` parameter. `UNION` is not supported for create operations when using StackSet as a resource.
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

