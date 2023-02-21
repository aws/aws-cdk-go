package awslex


// Provides information about the external source of the slot type's definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalSourceSettingProperty := &externalSourceSettingProperty{
//   	grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   		source: &grammarSlotTypeSourceProperty{
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   			// the properties below are optional
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   }
//
type CfnBot_ExternalSourceSettingProperty struct {
	// Settings required for a slot type based on a grammar that you provide.
	GrammarSlotTypeSetting interface{} `field:"optional" json:"grammarSlotTypeSetting" yaml:"grammarSlotTypeSetting"`
}

