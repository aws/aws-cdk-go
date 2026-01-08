package previewawsworkspaceswebmixins


// Metadata information about an uploaded image file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageMetadataProperty := &ImageMetadataProperty{
//   	FileExtension: jsii.String("fileExtension"),
//   	LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   	MimeType: jsii.String("mimeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html
//
type CfnUserSettingsPropsMixin_ImageMetadataProperty struct {
	// The file extension of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-fileextension
	//
	FileExtension *string `field:"optional" json:"fileExtension" yaml:"fileExtension"`
	// The timestamp when the image was last uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-lastuploadtimestamp
	//
	LastUploadTimestamp *string `field:"optional" json:"lastUploadTimestamp" yaml:"lastUploadTimestamp"`
	// The MIME type of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-mimetype
	//
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
}

