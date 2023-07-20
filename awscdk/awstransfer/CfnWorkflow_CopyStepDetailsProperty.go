package awstransfer


// Details for a step that performs a file copy.
//
// Consists of the following values:
//
// - A description
// - An Amazon S3 location for the destination of the file copy.
// - A flag that indicates whether to overwrite an existing file of the same name. The default is `FALSE` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyStepDetailsProperty := &CopyStepDetailsProperty{
//   	DestinationFileLocation: &S3FileLocationProperty{
//   		S3FileLocation: &S3InputFileLocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OverwriteExisting: jsii.String("overwriteExisting"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-copystepdetails.html
//
type CfnWorkflow_CopyStepDetailsProperty struct {
	// Specifies the location for the file being copied.
	//
	// Use `${Transfer:UserName}` or `${Transfer:UploadDate}` in this field to parametrize the destination prefix by username or uploaded date.
	//
	// - Set the value of `DestinationFileLocation` to `${Transfer:UserName}` to copy uploaded files to an Amazon S3 bucket that is prefixed with the name of the Transfer Family user that uploaded the file.
	// - Set the value of `DestinationFileLocation` to `${Transfer:UploadDate}` to copy uploaded files to an Amazon S3 bucket that is prefixed with the date of the upload.
	//
	// > The system resolves `UploadDate` to a date format of *YYYY-MM-DD* , based on the date the file is uploaded in UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-copystepdetails.html#cfn-transfer-workflow-copystepdetails-destinationfilelocation
	//
	DestinationFileLocation interface{} `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// The name of the step, used as an identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-copystepdetails.html#cfn-transfer-workflow-copystepdetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A flag that indicates whether to overwrite an existing file of the same name. The default is `FALSE` .
	//
	// If the workflow is processing a file that has the same name as an existing file, the behavior is as follows:
	//
	// - If `OverwriteExisting` is `TRUE` , the existing file is replaced with the file being processed.
	// - If `OverwriteExisting` is `FALSE` , nothing happens, and the workflow processing stops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-copystepdetails.html#cfn-transfer-workflow-copystepdetails-overwriteexisting
	//
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow.
	//
	// - To use the previous file as the input, enter `${previous.file}` . In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value.
	// - To use the originally uploaded file location as input for this step, enter `${original.file}` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-copystepdetails.html#cfn-transfer-workflow-copystepdetails-sourcefilelocation
	//
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

