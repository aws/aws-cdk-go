package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationAzureBlob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationAzureBlobProps := &CfnLocationAzureBlobProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	AzureBlobAuthenticationType: jsii.String("azureBlobAuthenticationType"),
//
//   	// the properties below are optional
//   	AzureAccessTier: jsii.String("azureAccessTier"),
//   	AzureBlobContainerUrl: jsii.String("azureBlobContainerUrl"),
//   	AzureBlobSasConfiguration: &AzureBlobSasConfigurationProperty{
//   		AzureBlobSasToken: jsii.String("azureBlobSasToken"),
//   	},
//   	AzureBlobType: jsii.String("azureBlobType"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html
//
type CfnLocationAzureBlobProps struct {
	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container.
	//
	// You can specify more than one agent. For more information, see [Using multiple agents for your transfer](https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-agentarns
	//
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies the authentication method DataSync uses to access your Azure Blob Storage.
	//
	// DataSync can access blob storage using a shared access signature (SAS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobauthenticationtype
	//
	// Default: - "SAS".
	//
	AzureBlobAuthenticationType *string `field:"required" json:"azureBlobAuthenticationType" yaml:"azureBlobAuthenticationType"`
	// Specifies the access tier that you want your objects or files transferred into.
	//
	// This only applies when using the location as a transfer destination. For more information, see [Access tiers](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureaccesstier
	//
	// Default: - "HOT".
	//
	AzureAccessTier *string `field:"optional" json:"azureAccessTier" yaml:"azureAccessTier"`
	// Specifies the URL of the Azure Blob Storage container involved in your transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobcontainerurl
	//
	AzureBlobContainerUrl *string `field:"optional" json:"azureBlobContainerUrl" yaml:"azureBlobContainerUrl"`
	// Specifies the SAS configuration that allows DataSync to access your Azure Blob Storage.
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

