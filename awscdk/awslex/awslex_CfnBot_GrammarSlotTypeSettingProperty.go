package awslex


// Settings required for a slot type based on a grammar that you provide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grammarSlotTypeSettingProperty := &grammarSlotTypeSettingProperty{
//   	source: &grammarSlotTypeSourceProperty{
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBot_GrammarSlotTypeSettingProperty struct {
	// The source of the grammar used to create the slot type.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

