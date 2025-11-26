package previewawssecuritylakemixins


// Third-party custom log source that meets the requirements to be added to Amazon Security Lake .
//
// For more details, see [Custom log source](https://docs.aws.amazon.com//security-lake/latest/userguide/custom-sources.html#iam-roles-custom-sources) in the *Amazon Security Lake User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customLogSourceProperty := &CustomLogSourceProperty{
//   	SourceName: jsii.String("sourceName"),
//   	SourceVersion: jsii.String("sourceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-customlogsource.html
//
type CfnSubscriberPropsMixin_CustomLogSourceProperty struct {
	// The name of the custom log source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-customlogsource.html#cfn-securitylake-subscriber-customlogsource-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// The source version of the custom log source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-customlogsource.html#cfn-securitylake-subscriber-customlogsource-sourceversion
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

