package awslex


// Settings requried for a slot type based on a grammar that you provide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grammarSlotTypeSettingProperty := &GrammarSlotTypeSettingProperty{
//   	Source: &GrammarSlotTypeSourceProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesetting.html
//
type CfnBot_GrammarSlotTypeSettingProperty struct {
	// The source of the grammar used to create the slot type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesetting.html#cfn-lex-bot-grammarslottypesetting-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

