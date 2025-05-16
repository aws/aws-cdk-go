package awsopensearchservice


// Configuration settings for IAM Identity Center in an OpenSearch application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamIdentityCenterOptionsProperty := &IamIdentityCenterOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	IamIdentityCenterInstanceArn: jsii.String("iamIdentityCenterInstanceArn"),
//   	IamRoleForIdentityCenterApplicationArn: jsii.String("iamRoleForIdentityCenterApplicationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html
//
type CfnApplication_IamIdentityCenterOptionsProperty struct {
	// Indicates whether IAM Identity Center is enabled for the OpenSearch application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Amazon Resource Name (ARN) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamidentitycenterinstancearn
	//
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// The Amazon Resource Name (ARN) of the IAM role assigned to the IAM Identity Center application for the OpenSearch application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamroleforidentitycenterapplicationarn
	//
	IamRoleForIdentityCenterApplicationArn *string `field:"optional" json:"iamRoleForIdentityCenterApplicationArn" yaml:"iamRoleForIdentityCenterApplicationArn"`
}

