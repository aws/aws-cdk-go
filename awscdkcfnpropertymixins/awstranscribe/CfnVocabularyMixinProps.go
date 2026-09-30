package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVocabularyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnVocabularyMixinProps := &CfnVocabularyMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Phrases: []*string{
//   		jsii.String("phrases"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VocabularyFileUri: jsii.String("vocabularyFileUri"),
//   	VocabularyName: jsii.String("vocabularyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html
//
type CfnVocabularyMixinProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The language code that represents the language of the entries in your custom vocabulary.
	//
	// Each custom vocabulary must contain terms in only one language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Use this parameter if you want to create your custom vocabulary by including all desired terms, as comma-separated values, within your request.
	//
	// You cannot use this parameter together with VocabularyFileUri.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-phrases
	//
	Phrases *[]*string `field:"optional" json:"phrases" yaml:"phrases"`
	// Adds one or more custom tags, each in the form of a key:value pair, to the custom vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains your custom vocabulary.
	//
	// You cannot use this parameter together with Phrases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-vocabularyfileuri
	//
	VocabularyFileUri *string `field:"optional" json:"vocabularyFileUri" yaml:"vocabularyFileUri"`
	// A unique name, chosen by you, for your custom vocabulary.
	//
	// This name is case sensitive, cannot contain spaces, and must be unique within an AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabulary.html#cfn-transcribe-vocabulary-vocabularyname
	//
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

