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
//   			GroupConfiguration: &CognitoGroupConfigurationProperty{
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
//   		},
//   		OpenIdConnectConfiguration: &OpenIdConnectConfigurationProperty{
//   			Issuer: jsii.String("issuer"),
//   			TokenSelection: &OpenIdConnectTokenSelectionProperty{
//   				AccessTokenOnly: &OpenIdConnectAccessTokenConfigurationProperty{
//   					Audiences: []*string{
//   						jsii.String("audiences"),
//   					},
//   					PrincipalIdClaim: jsii.String("principalIdClaim"),
//   				},
//   				IdentityTokenOnly: &OpenIdConnectIdentityTokenConfigurationProperty{
//   					ClientIds: []*string{
//   						jsii.String("clientIds"),
//   					},
//   					PrincipalIdClaim: jsii.String("principalIdClaim"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			EntityIdPrefix: jsii.String("entityIdPrefix"),
//   			GroupConfiguration: &OpenIdConnectGroupConfigurationProperty{
//   				GroupClaim: jsii.String("groupClaim"),
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
//   		},
//   	},
//   	PolicyStoreId: jsii.String("policyStoreId"),
//
//   	// the properties below are optional
//   	PrincipalEntityType: jsii.String("principalEntityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html
//
type CfnIdentitySourceProps struct {
	// Contains configuration information about an identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// Specifies the ID of the policy store in which you want to store this identity source.
	//
	// Only policies and requests made using this policy store can reference identities from the identity provider configured in the new identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-policystoreid
	//
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
	// Specifies the namespace and data type of the principals generated for identities authenticated by the new identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-principalentitytype
	//
	PrincipalEntityType *string `field:"optional" json:"principalEntityType" yaml:"principalEntityType"`
}

