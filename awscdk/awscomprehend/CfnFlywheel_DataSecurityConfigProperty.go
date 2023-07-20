package awscomprehend


// Data security configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSecurityConfigProperty := &DataSecurityConfigProperty{
//   	DataLakeKmsKeyId: jsii.String("dataLakeKmsKeyId"),
//   	ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html
//
type CfnFlywheel_DataSecurityConfigProperty struct {
	// ID for the AWS KMS key that Amazon Comprehend uses to encrypt the data in the data lake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-datalakekmskeyid
	//
	DataLakeKmsKeyId *string `field:"optional" json:"dataLakeKmsKeyId" yaml:"dataLakeKmsKeyId"`
	// ID for the AWS KMS key that Amazon Comprehend uses to encrypt trained custom models.
	//
	// The ModelKmsKeyId can be either of the following formats:
	//
	// - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
	// - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-modelkmskeyid
	//
	ModelKmsKeyId *string `field:"optional" json:"modelKmsKeyId" yaml:"modelKmsKeyId"`
	// ID for the AWS KMS key that Amazon Comprehend uses to encrypt the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for the job.
	//
	// For more information, see [Amazon VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

