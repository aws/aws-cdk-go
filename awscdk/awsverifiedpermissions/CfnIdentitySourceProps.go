package awsverifiedpermissions


// Properties for defining a `CfnIdentitySource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentitySourceProps := &CfnIdentitySourceProps{
//   	Configuration: &IdentitySourceConfigurationProperty{
//   		CognitoUserPoolConfiguration: &CognitoUserPoolConfigurationProperty{
//   			UserPoolArn: jsii.String("userPoolArn"),
//
//   			// the properties below are optional
//   			ClientIds: []*string{
//   				jsii.String("clientIds"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   	PrincipalEntityType: jsii.String("principalEntityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html
//
type CfnIdentitySourceProps struct {
	// Contains configuration information used when creating or updating an identity source.
	//
	// > At this time, the only valid member of this structure is a Amazon Cognito user pool configuration.
	// >
	// > You must specify a `userPoolArn` , and optionally, a `ClientId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// Specifies the ID of the policy store in which you want to store this identity source.
	//
	// Only policies and requests made using this policy store can reference identities from the identity provider configured in the new identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-policystoreid
	//
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
	// Specifies the namespace and data type of the principals generated for identities authenticated by the new identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-principalentitytype
	//
	PrincipalEntityType *string `field:"optional" json:"principalEntityType" yaml:"principalEntityType"`
}

