package awsconnect

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
//   	Content: jsii.String("content"),
//   	InstanceId: jsii.String("instanceId"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VocabularyName: jsii.String("vocabularyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html
//
type CfnVocabularyMixinProps struct {
	// The content of the custom vocabulary in plain-text format with a table of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The language code of the vocabulary entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A unique name of the custom vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-vocabulary.html#cfn-connect-vocabulary-vocabularyname
	//
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

