package awsopensearchservice


// Container for IAM Identity Center Options settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityCenterOptionsProperty := &IdentityCenterOptionsProperty{
//   	EnabledApiAccess: jsii.Boolean(false),
//   	IdentityCenterApplicationArn: jsii.String("identityCenterApplicationArn"),
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	RolesKey: jsii.String("rolesKey"),
//   	SubjectKey: jsii.String("subjectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html
//
type CfnDomain_IdentityCenterOptionsProperty struct {
	// True to enable IAM Identity Center for API access in Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-enabledapiaccess
	//
	EnabledApiAccess interface{} `field:"optional" json:"enabledApiAccess" yaml:"enabledApiAccess"`
	// The ARN for IAM Identity Center Application which will integrate with Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitycenterapplicationarn
	//
	IdentityCenterApplicationArn *string `field:"optional" json:"identityCenterApplicationArn" yaml:"identityCenterApplicationArn"`
	// The ARN for IAM Identity Center Instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// The ID of IAM Identity Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitystoreid
	//
	IdentityStoreId *string `field:"optional" json:"identityStoreId" yaml:"identityStoreId"`
	// Specify the attribute that contains the backend role (groupName, groupID) of IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-roleskey
	//
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Specify the attribute that contains the subject (username, userID, email) of IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-subjectkey
	//
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

