package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMedicalVocabulary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMedicalVocabularyProps := &CfnMedicalVocabularyProps{
//   	LanguageCode: jsii.String("languageCode"),
//   	VocabularyName: jsii.String("vocabularyName"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VocabularyFileUri: jsii.String("vocabularyFileUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html
//
type CfnMedicalVocabularyProps struct {
	// The language code of the vocabulary entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-languagecode
	//
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// The name of the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-vocabularyname
	//
	VocabularyName *string `field:"required" json:"vocabularyName" yaml:"vocabularyName"`
	// Tags associated with the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-vocabularyfileuri
	//
	VocabularyFileUri *string `field:"optional" json:"vocabularyFileUri" yaml:"vocabularyFileUri"`
}

