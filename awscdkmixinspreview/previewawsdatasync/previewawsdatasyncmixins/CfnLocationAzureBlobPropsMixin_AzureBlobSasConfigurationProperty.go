package previewawsdatasyncmixins


// The shared access signature (SAS) configuration that allows AWS DataSync to access your Microsoft Azure Blob Storage.
//
// For more information, see [SAS tokens](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-sas-tokens) for accessing your Azure Blob Storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   azureBlobSasConfigurationProperty := &AzureBlobSasConfigurationProperty{
//   	AzureBlobSasToken: jsii.String("azureBlobSasToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureblobsasconfiguration.html
//
type CfnLocationAzureBlobPropsMixin_AzureBlobSasConfigurationProperty struct {
	// Specifies a SAS token that provides permissions to access your Azure Blob Storage.
	//
	// The token is part of the SAS URI string that comes after the storage resource URI and a question mark. A token looks something like this:
	//
	// `sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationazureblob-azureblobsasconfiguration.html#cfn-datasync-locationazureblob-azureblobsasconfiguration-azureblobsastoken
	//
	AzureBlobSasToken *string `field:"optional" json:"azureBlobSasToken" yaml:"azureBlobSasToken"`
}

