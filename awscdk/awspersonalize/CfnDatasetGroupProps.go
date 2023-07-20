package awspersonalize


// Properties for defining a `CfnDatasetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetGroupProps := &CfnDatasetGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Domain: jsii.String("domain"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datasetgroup.html
//
type CfnDatasetGroupProps struct {
	// The name of the dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datasetgroup.html#cfn-personalize-datasetgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain of a Domain dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datasetgroup.html#cfn-personalize-datasetgroup-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key used to encrypt the datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datasetgroup.html#cfn-personalize-datasetgroup-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the IAM role that has permissions to create the dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datasetgroup.html#cfn-personalize-datasetgroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

