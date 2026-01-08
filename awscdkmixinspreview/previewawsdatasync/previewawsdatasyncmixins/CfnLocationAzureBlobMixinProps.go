package previewawsdatasyncmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocationAzureBlobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationAzureBlobMixinProps := &CfnLocationAzureBlobMixinProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	AzureAccessTier: jsii.String("azureAccessTier"),
//   	AzureBlobAuthenticationType: jsii.String("azureBlobAuthenticationType"),
//   	AzureBlobContainerUrl: jsii.String("azureBlobContainerUrl"),
//   	AzureBlobSasConfiguration: &AzureBlobSasConfigurationProperty{
//   		AzureBlobSasToken: jsii.String("azureBlobSasToken"),
//   	},
//   	AzureBlobType: jsii.String("azureBlobType"),
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html
//
type CfnLocationAzureBlobMixinProps struct {
	// (Optional) Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container.
	//
	// If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter.
	//
	// You can specify more than one agent. For more information, see [Using multiple agents for your transfer](https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html) .
	//
	// > Make sure you configure this parameter correctly when you first create your storage location. You cannot add or remove agents from a storage location after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-agentarns
	//
	AgentArns *[]*string `field:"optional" json:"agentArns" yaml:"agentArns"`
	// Specifies the access tier that you want your objects or files transferred into.
	//
	// This only applies when using the location as a transfer destination. For more information, see [Access tiers](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureaccesstier
	//
	// Default: - "HOT".
	//
	AzureAccessTier *string `field:"optional" json:"azureAccessTier" yaml:"azureAccessTier"`
	// Specifies the authentication method DataSync uses to access your Azure Blob Storage.
	//
	// DataSync can access blob storage using a shared access signature (SAS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobauthenticationtype
	//
	// Default: - "SAS".
	//
	AzureBlobAuthenticationType *string `field:"optional" json:"azureBlobAuthenticationType" yaml:"azureBlobAuthenticationType"`
	// Specifies the URL of the Azure Blob Storage container involved in your transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobcontainerurl
	//
	AzureBlobContainerUrl *string `field:"optional" json:"azureBlobContainerUrl" yaml:"azureBlobContainerUrl"`
	// Specifies the SAS configuration that allows DataSync to access your Azure Blob Storage.
	//
	// > If you provide an authentication token using `SasConfiguration` , but do not provide secret configuration details using `CmkSecretConfig` or `CustomSecretConfig` , then DataSync stores the token using your AWS account's secrets manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobsasconfiguration
	//
	AzureBlobSasConfiguration interface{} `field:"optional" json:"azureBlobSasConfiguration" yaml:"azureBlobSasConfiguration"`
	// Specifies the type of blob that you want your objects or files to be when transferring them into Azure Blob Storage.
	//
	// Currently, DataSync only supports moving data into Azure Blob Storage as block blobs. For more information on blob types, see the [Azure Blob Storage documentation](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobtype
	//
	// Default: - "BLOCK".
	//
	AzureBlobType *string `field:"optional" json:"azureBlobType" yaml:"azureBlobType"`
	// Specifies configuration information for a DataSync-managed secret, such as an authentication token, secret key, password, or Kerberos keytab that DataSync uses to access a specific storage location, with a customer-managed AWS KMS key .
	//
	// > You can use either `CmkSecretConfig` or `CustomSecretConfig` to provide credentials for a `CreateLocation` request. Do not provide both parameters for the same request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-cmksecretconfig
	//
	CmkSecretConfig interface{} `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed Secrets Manager secret where a storage location credentials is stored in Secrets Manager as plain text (for authentication token, secret key, or password) or as binary (for Kerberos keytab).
	//
	// This configuration includes the secret ARN, and the ARN for an IAM role that provides access to the secret.
	//
	// > You can use either `CmkSecretConfig` or `CustomSecretConfig` to provide credentials for a `CreateLocation` request. Do not provide both parameters for the same request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-customsecretconfig
	//
	CustomSecretConfig interface{} `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// Specifies path segments if you want to limit your transfer to a virtual directory in your container (for example, `/my/images` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your transfer location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

