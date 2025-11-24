package mixinsawsopensearchserverless


// Describes IAM Identity Center options for an OpenSearch Serverless security configuration in the form of a key-value map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamIdentityCenterConfigOptionsProperty := &IamIdentityCenterConfigOptionsProperty{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	ApplicationName: jsii.String("applicationName"),
//   	GroupAttribute: jsii.String("groupAttribute"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	UserAttribute: jsii.String("userAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html
//
type CfnSecurityConfigPropsMixin_IamIdentityCenterConfigOptionsProperty struct {
	// The ARN of the IAM Identity Center application used to integrate with OpenSearch Serverless.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The description of the IAM Identity Center application used to integrate with OpenSearch Serverless.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-applicationdescription
	//
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// The name of the IAM Identity Center application used to integrate with OpenSearch Serverless.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The group attribute for this IAM Identity Center integration.
	//
	// Defaults to `GroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-groupattribute
	//
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// The ARN of the IAM Identity Center instance used to integrate with OpenSearch Serverless.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The user attribute for this IAM Identity Center integration.
	//
	// Defaults to `UserId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html#cfn-opensearchserverless-securityconfig-iamidentitycenterconfigoptions-userattribute
	//
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

