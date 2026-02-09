package previewawssesmixins


// An object that defines a MailManager archive that is used to preserve emails that you send using the configuration set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archivingOptionsProperty := &ArchivingOptionsProperty{
//   	ArchiveArn: jsii.String("archiveArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-archivingoptions.html
//
type CfnConfigurationSetPropsMixin_ArchivingOptionsProperty struct {
	// The ARN of the MailManager archive to associate with the configuration set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-archivingoptions.html#cfn-ses-configurationset-archivingoptions-archivearn
	//
	ArchiveArn *string `field:"optional" json:"archiveArn" yaml:"archiveArn"`
}

