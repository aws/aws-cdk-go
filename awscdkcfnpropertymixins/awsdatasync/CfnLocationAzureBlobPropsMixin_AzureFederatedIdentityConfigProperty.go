package awsdatasync


// Specifies the identity federation configuration that DataSync uses to access your Azure Blob Storage container using an OpenID Connect (OIDC) token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureFederatedIdentityConfigProperty := &AzureFederatedIdentityConfigProperty{
//   	AwsIamRole: jsii.String("awsIamRole"),
//   	AzureOidc: &AzureOidcConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azurefederatedidentityconfig.html
//
type CfnLocationAzureBlobPropsMixin_AzureFederatedIdentityConfigProperty struct {
	// Specifies the ARN of the AWS Identity and Access Management (IAM) role that DataSync assumes to mint the OIDC token used to authenticate with the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azurefederatedidentityconfig.html#cfn-datasync-locationazureblob-azurefederatedidentityconfig-awsiamrole
	//
	AwsIamRole *string `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// Specifies the Microsoft Entra (Azure AD) identity that DataSync federates with to obtain an access token for your Azure Blob Storage container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azurefederatedidentityconfig.html#cfn-datasync-locationazureblob-azurefederatedidentityconfig-azureoidc
	//
	AzureOidc interface{} `field:"optional" json:"azureOidc" yaml:"azureOidc"`
}

