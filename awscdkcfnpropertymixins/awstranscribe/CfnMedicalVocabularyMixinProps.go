package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMedicalVocabularyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMedicalVocabularyMixinProps := &CfnMedicalVocabularyMixinProps{
//   	LanguageCode: jsii.String("languageCode"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html
//
type CfnMedicalVocabularyMixinProps struct {
	// The language code of the vocabulary entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Tags associated with the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-vocabularyfileuri
	//
	VocabularyFileUri *string `field:"optional" json:"vocabularyFileUri" yaml:"vocabularyFileUri"`
	// The name of the medical vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-medicalvocabulary.html#cfn-transcribe-medicalvocabulary-vocabularyname
	//
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

