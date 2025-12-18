package awsworkspacesweb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageMetadataProperty := &ImageMetadataProperty{
//   	FileExtension: jsii.String("fileExtension"),
//   	LastUploadTimestamp: jsii.String("lastUploadTimestamp"),
//   	MimeType: jsii.String("mimeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html
//
type CfnUserSettings_ImageMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-fileextension
	//
	FileExtension *string `field:"required" json:"fileExtension" yaml:"fileExtension"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-lastuploadtimestamp
	//
	LastUploadTimestamp *string `field:"required" json:"lastUploadTimestamp" yaml:"lastUploadTimestamp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-imagemetadata.html#cfn-workspacesweb-usersettings-imagemetadata-mimetype
	//
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

