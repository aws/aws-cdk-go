package awsmacie


// Specifies the criteria for an allow list, which is a list that defines specific text or a text pattern to ignore when inspecting data sources for sensitive data.
//
// The criteria can be:
//
// - The location and name of an Amazon Simple Storage Service ( Amazon S3 ) object that lists specific predefined text to ignore ( `S3WordsList` ), or
// - A regular expression ( `Regex` ) that defines a text pattern to ignore.
//
// The criteria must specify either an S3 object or a regular expression. It can't specify both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   criteriaProperty := &CriteriaProperty{
//   	Regex: jsii.String("regex"),
//   	S3WordsList: &S3WordsListProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-allowlist-criteria.html
//
type CfnAllowList_CriteriaProperty struct {
	// The regular expression ( *regex* ) that defines the text pattern to ignore.
	//
	// The expression can contain 1-512 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-allowlist-criteria.html#cfn-macie-allowlist-criteria-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The location and name of an Amazon S3 object that lists specific text to ignore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-allowlist-criteria.html#cfn-macie-allowlist-criteria-s3wordslist
	//
	S3WordsList interface{} `field:"optional" json:"s3WordsList" yaml:"s3WordsList"`
}

