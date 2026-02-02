package previewawsmwaaserverlessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-encryptionconfiguration.html
//
type CfnWorkflowPropsMixin_EncryptionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-encryptionconfiguration.html#cfn-mwaaserverless-workflow-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-encryptionconfiguration.html#cfn-mwaaserverless-workflow-encryptionconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

