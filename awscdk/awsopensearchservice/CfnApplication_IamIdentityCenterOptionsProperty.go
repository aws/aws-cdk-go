package awsopensearchservice


// Settings for IAM Identity Center for an OpenSearch Application.
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
	// IAM Identity Center is enabled for the OpenSearch Application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Amazon Resource Name (ARN) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamidentitycenterinstancearn
	//
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// Amazon Resource Name of the IAM Identity Center's Application created for the OpenSearch Application after enabling IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamroleforidentitycenterapplicationarn
	//
	IamRoleForIdentityCenterApplicationArn *string `field:"optional" json:"iamRoleForIdentityCenterApplicationArn" yaml:"iamRoleForIdentityCenterApplicationArn"`
}

