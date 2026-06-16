package awselementalinference


// Properties for defining a `CfnDictionary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDictionaryProps := &CfnDictionaryProps{
//   	Language: jsii.String("language"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Entries: jsii.String("entries"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html
//
type CfnDictionaryProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-entries
	//
	Entries *string `field:"optional" json:"entries" yaml:"entries"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

