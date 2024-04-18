package awssecuritylake


// Properties for defining a `CfnAwsLogSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAwsLogSourceProps := &CfnAwsLogSourceProps{
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//
//   	// the properties below are optional
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html
//
type CfnAwsLogSourceProps struct {
	// The ARN for the data lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-datalakearn
	//
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
	// The name for a AWS source.
	//
	// This must be a Regionally unique value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-sourcename
	//
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The version for a AWS source.
	//
	// This must be a Regionally unique value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-sourceversion
	//
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
	// AWS account where you want to collect logs from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-accounts
	//
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
}

