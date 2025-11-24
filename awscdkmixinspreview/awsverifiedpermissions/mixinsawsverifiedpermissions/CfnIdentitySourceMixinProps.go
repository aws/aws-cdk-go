package mixinsawsverifiedpermissions


// Properties for CfnIdentitySourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentitySourceMixinProps := &CfnIdentitySourceMixinProps{
//   	Configuration: &IdentitySourceConfigurationProperty{
//   		CognitoUserPoolConfiguration: &CognitoUserPoolConfigurationProperty{
//   			ClientIds: []*string{
//   				jsii.String("clientIds"),
//   			},
//   			GroupConfiguration: &CognitoGroupConfigurationProperty{
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
//   			UserPoolArn: jsii.String("userPoolArn"),
//   		},
//   		OpenIdConnectConfiguration: &OpenIdConnectConfigurationProperty{
//   			EntityIdPrefix: jsii.String("entityIdPrefix"),
//   			GroupConfiguration: &OpenIdConnectGroupConfigurationProperty{
//   				GroupClaim: jsii.String("groupClaim"),
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
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
//   		},
//   	},
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   	PrincipalEntityType: jsii.String("principalEntityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html
//
type CfnIdentitySourceMixinProps struct {
	// Contains configuration information used when creating a new identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html#cfn-verifiedpermissions-identitysource-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
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

