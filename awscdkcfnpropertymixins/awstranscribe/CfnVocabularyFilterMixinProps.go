package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVocabularyFilterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnVocabularyFilterMixinProps := &CfnVocabularyFilterMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VocabularyFilterFileUri: jsii.String("vocabularyFilterFileUri"),
//   	VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   	Words: []*string{
//   		jsii.String("words"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html
//
type CfnVocabularyFilterMixinProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The language code that represents the language of the entries in your vocabulary filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Tags associated with the vocabulary filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains your custom vocabulary filter terms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-vocabularyfilterfileuri
	//
	VocabularyFilterFileUri *string `field:"optional" json:"vocabularyFilterFileUri" yaml:"vocabularyFilterFileUri"`
	// A unique name, chosen by you, for your custom vocabulary filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-vocabularyfiltername
	//
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html#cfn-transcribe-vocabularyfilter-words
	//
	Words *[]*string `field:"optional" json:"words" yaml:"words"`
}

