package awsbedrock


// Contains details about the server-side encryption for the managed knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedKnowledgeBaseServerSideEncryptionConfigurationProperty := &ManagedKnowledgeBaseServerSideEncryptionConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseserversideencryptionconfiguration.html
//
type CfnKnowledgeBase_ManagedKnowledgeBaseServerSideEncryptionConfigurationProperty struct {
	// The ARN of the AWS KMS key used to encrypt the managed knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseserversideencryptionconfiguration.html#cfn-bedrock-knowledgebase-managedknowledgebaseserversideencryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

