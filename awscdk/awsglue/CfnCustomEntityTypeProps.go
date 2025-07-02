package awsglue


// Properties for defining a `CfnCustomEntityType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCustomEntityTypeProps := &CfnCustomEntityTypeProps{
//   	ContextWords: []*string{
//   		jsii.String("contextWords"),
//   	},
//   	Name: jsii.String("name"),
//   	RegexString: jsii.String("regexString"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html
//
type CfnCustomEntityTypeProps struct {
	// A list of context words.
	//
	// If none of these context words are found within the vicinity of the regular expression the data will not be detected as sensitive data.
	//
	// If no context words are passed only a regular expression is checked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html#cfn-glue-customentitytype-contextwords
	//
	ContextWords *[]*string `field:"optional" json:"contextWords" yaml:"contextWords"`
	// A name for the custom pattern that allows it to be retrieved or deleted later.
	//
	// This name must be unique per AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html#cfn-glue-customentitytype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression string that is used for detecting sensitive data in a custom pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html#cfn-glue-customentitytype-regexstring
	//
	RegexString *string `field:"optional" json:"regexString" yaml:"regexString"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-customentitytype.html#cfn-glue-customentitytype-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

