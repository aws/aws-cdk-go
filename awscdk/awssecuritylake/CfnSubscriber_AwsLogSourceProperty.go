package awssecuritylake


// Adds a natively supported AWS service as an Amazon Security Lake source.
//
// Enables source types for member accounts in required AWS Regions, based on the parameters you specify. You can choose any source type in any Region for either accounts that are part of a trusted organization or standalone accounts. Once you add an AWS service as a source, Security Lake starts collecting logs and events from it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsLogSourceProperty := &AwsLogSourceProperty{
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-awslogsource.html
//
type CfnSubscriber_AwsLogSourceProperty struct {
	// Source name of the natively supported AWS service that is supported as an Amazon Security Lake source.
	//
	// For the list of sources supported by Amazon Security Lake see [Collecting data from AWS services](https://docs.aws.amazon.com//security-lake/latest/userguide/internal-sources.html) in the Amazon Security Lake User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-awslogsource.html#cfn-securitylake-subscriber-awslogsource-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// Source version of the natively supported AWS service that is supported as an Amazon Security Lake source.
	//
	// For more details about source versions supported by Amazon Security Lake see [OCSF source identification](https://docs.aws.amazon.com//security-lake/latest/userguide/open-cybersecurity-schema-framework.html#ocsf-source-identification) in the Amazon Security Lake User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-awslogsource.html#cfn-securitylake-subscriber-awslogsource-sourceversion
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

