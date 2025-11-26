package previewawssecuritylakemixins


// Properties for CfnAwsLogSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAwsLogSourceMixinProps := &CfnAwsLogSourceMixinProps{
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	DataLakeArn: jsii.String("dataLakeArn"),
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html
//
type CfnAwsLogSourceMixinProps struct {
	// Specify the AWS account information where you want to enable Security Lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-accounts
	//
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The Amazon Resource Name (ARN) used to create the data lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-datalakearn
	//
	DataLakeArn *string `field:"optional" json:"dataLakeArn" yaml:"dataLakeArn"`
	// The name for a AWS source.
	//
	// This must be a Regionally unique value. For the list of sources supported by Amazon Security Lake see [Collecting data from AWS services](https://docs.aws.amazon.com//security-lake/latest/userguide/internal-sources.html) in the Amazon Security Lake User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// The version for a AWS source.
	//
	// For more details about source versions supported by Amazon Security Lake see [OCSF source identification](https://docs.aws.amazon.com//security-lake/latest/userguide/open-cybersecurity-schema-framework.html#ocsf-source-identification) in the Amazon Security Lake User Guide. This must be a Regionally unique value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#cfn-securitylake-awslogsource-sourceversion
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

