package awspersonalize


// Properties for defining a `CfnDatasetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetGroupProps := &cfnDatasetGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDatasetGroupProps struct {
	// The name of the dataset group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain of a Domain dataset group.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key used to encrypt the datasets.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the IAM role that has permissions to create the dataset group.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

