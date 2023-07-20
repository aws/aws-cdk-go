package awslex


// Provides information about the external source of the slot type's definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalSourceSettingProperty := &ExternalSourceSettingProperty{
//   	GrammarSlotTypeSetting: &GrammarSlotTypeSettingProperty{
//   		Source: &GrammarSlotTypeSourceProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-externalsourcesetting.html
//
type CfnBot_ExternalSourceSettingProperty struct {
	// Settings required for a slot type based on a grammar that you provide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-externalsourcesetting.html#cfn-lex-bot-externalsourcesetting-grammarslottypesetting
	//
	GrammarSlotTypeSetting interface{} `field:"optional" json:"grammarSlotTypeSetting" yaml:"grammarSlotTypeSetting"`
}

