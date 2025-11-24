package mixinsawsopensearchservice


// Settings container for integrating IAM Identity Center with OpenSearch UI applications, which enables enabling secure user authentication and access control across multiple data sources.
//
// This setup supports single sign-on (SSO) through IAM Identity Center, allowing centralized user management.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnDomainPropsMixin_IdentityCenterOptionsProperty struct {
	// Indicates whether IAM Identity Center is enabled for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-enabledapiaccess
	//
	EnabledApiAccess interface{} `field:"optional" json:"enabledApiAccess" yaml:"enabledApiAccess"`
	// The ARN of the IAM Identity Center application that integrates with Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitycenterapplicationarn
	//
	IdentityCenterApplicationArn *string `field:"optional" json:"identityCenterApplicationArn" yaml:"identityCenterApplicationArn"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// The identifier of the IAM Identity Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-identitystoreid
	//
	IdentityStoreId *string `field:"optional" json:"identityStoreId" yaml:"identityStoreId"`
	// Specifies the attribute that contains the backend role identifier (such as group name or group ID) in IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-roleskey
	//
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Specifies the attribute that contains the subject identifier (such as username, user ID, or email) in IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-identitycenteroptions.html#cfn-opensearchservice-domain-identitycenteroptions-subjectkey
	//
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

