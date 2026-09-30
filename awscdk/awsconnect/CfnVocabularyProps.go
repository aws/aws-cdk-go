package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVocabulary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVocabularyProps := &CfnVocabularyProps{
//   	Content: jsii.String("content"),
//   	InstanceId: jsii.String("instanceId"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html
//
type CfnVocabularyProps struct {
	// The content of the custom vocabulary in plain-text format with a table of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-content
	//
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The language code of the vocabulary entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-languagecode
	//
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// A unique name of the custom vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-vocabularyname
	//
	VocabularyName *string `field:"required" json:"vocabularyName" yaml:"vocabularyName"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

