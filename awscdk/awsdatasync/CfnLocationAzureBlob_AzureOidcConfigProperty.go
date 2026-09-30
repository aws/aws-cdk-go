package awsdatasync


// Specifies the Microsoft Entra (Azure AD) identity that DataSync federates with to obtain an access token for your Azure Blob Storage container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureOidcConfigProperty := &AzureOidcConfigProperty{
//   	ClientId: jsii.String("clientId"),
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureoidcconfig.html
//
type CfnLocationAzureBlob_AzureOidcConfigProperty struct {
	// Specifies the client ID of the Microsoft Entra (Azure AD) identity that DataSync uses to obtain an access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureoidcconfig.html#cfn-datasync-locationazureblob-azureoidcconfig-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Specifies the Microsoft Entra (Azure AD) tenant ID that the identity belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureoidcconfig.html#cfn-datasync-locationazureblob-azureoidcconfig-tenantid
	//
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

