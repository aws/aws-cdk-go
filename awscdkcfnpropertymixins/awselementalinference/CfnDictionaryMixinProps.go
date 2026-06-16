package awselementalinference


// Properties for CfnDictionaryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDictionaryMixinProps := &CfnDictionaryMixinProps{
//   	Entries: jsii.String("entries"),
//   	Language: jsii.String("language"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html
//
type CfnDictionaryMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-entries
	//
	Entries *string `field:"optional" json:"entries" yaml:"entries"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-language
	//
	Language *string `field:"optional" json:"language" yaml:"language"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-dictionary.html#cfn-elementalinference-dictionary-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

