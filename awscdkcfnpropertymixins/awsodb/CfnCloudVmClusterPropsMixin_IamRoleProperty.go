package awsodb


// An AWS Identity and Access Management (IAM) service role associated with the VM cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   iamRoleProperty := &IamRoleProperty{
//   	AwsIntegration: jsii.String("awsIntegration"),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-iamrole.html
//
type CfnCloudVmClusterPropsMixin_IamRoleProperty struct {
	// The AWS integration configuration settings for the AWS Identity and Access Management (IAM) service role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-iamrole.html#cfn-odb-cloudvmcluster-iamrole-awsintegration
	//
	AwsIntegration *string `field:"optional" json:"awsIntegration" yaml:"awsIntegration"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-iamrole.html#cfn-odb-cloudvmcluster-iamrole-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The current status of the AWS Identity and Access Management (IAM) service role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-iamrole.html#cfn-odb-cloudvmcluster-iamrole-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

