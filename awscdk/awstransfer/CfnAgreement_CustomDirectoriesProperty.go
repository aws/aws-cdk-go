package awstransfer


// Specifies a separate directory for each type of file to store for an AS2 message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDirectoriesProperty := &CustomDirectoriesProperty{
//   	FailedFilesDirectory: jsii.String("failedFilesDirectory"),
//   	MdnFilesDirectory: jsii.String("mdnFilesDirectory"),
//   	PayloadFilesDirectory: jsii.String("payloadFilesDirectory"),
//   	StatusFilesDirectory: jsii.String("statusFilesDirectory"),
//   	TemporaryFilesDirectory: jsii.String("temporaryFilesDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html
//
type CfnAgreement_CustomDirectoriesProperty struct {
	// Specifies a location to store the failed files for an AS2 message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html#cfn-transfer-agreement-customdirectories-failedfilesdirectory
	//
	FailedFilesDirectory *string `field:"required" json:"failedFilesDirectory" yaml:"failedFilesDirectory"`
	// Specifies a location to store the MDN file for an AS2 message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html#cfn-transfer-agreement-customdirectories-mdnfilesdirectory
	//
	MdnFilesDirectory *string `field:"required" json:"mdnFilesDirectory" yaml:"mdnFilesDirectory"`
	// Specifies a location to store the payload file for an AS2 message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html#cfn-transfer-agreement-customdirectories-payloadfilesdirectory
	//
	PayloadFilesDirectory *string `field:"required" json:"payloadFilesDirectory" yaml:"payloadFilesDirectory"`
	// Specifies a location to store the status file for an AS2 message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html#cfn-transfer-agreement-customdirectories-statusfilesdirectory
	//
	StatusFilesDirectory *string `field:"required" json:"statusFilesDirectory" yaml:"statusFilesDirectory"`
	// Specifies a location to store the temporary processing file for an AS2 message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-agreement-customdirectories.html#cfn-transfer-agreement-customdirectories-temporaryfilesdirectory
	//
	TemporaryFilesDirectory *string `field:"required" json:"temporaryFilesDirectory" yaml:"temporaryFilesDirectory"`
}

