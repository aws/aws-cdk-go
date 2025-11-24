package mixinsawstransfer


// The basic building block of a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   workflowStepProperty := &WorkflowStepProperty{
//   	CopyStepDetails: copyStepDetails,
//   	CustomStepDetails: customStepDetails,
//   	DecryptStepDetails: &DecryptStepDetailsProperty{
//   		DestinationFileLocation: &InputFileLocationProperty{
//   			EfsFileLocation: &EfsInputFileLocationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				Path: jsii.String("path"),
//   			},
//   			S3FileLocation: &S3InputFileLocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		OverwriteExisting: jsii.String("overwriteExisting"),
//   		SourceFileLocation: jsii.String("sourceFileLocation"),
//   		Type: jsii.String("type"),
//   	},
//   	DeleteStepDetails: deleteStepDetails,
//   	TagStepDetails: tagStepDetails,
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html
//
type CfnWorkflowPropsMixin_WorkflowStepProperty struct {
	// Details for a step that performs a file copy.
	//
	// Consists of the following values:
	//
	// - A description
	// - An Amazon S3 location for the destination of the file copy.
	// - A flag that indicates whether to overwrite an existing file of the same name. The default is `FALSE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-copystepdetails
	//
	CopyStepDetails interface{} `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// Details for a step that invokes an AWS Lambda function.
	//
	// Consists of the Lambda function's name, target, and timeout (in seconds).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-customstepdetails
	//
	CustomStepDetails interface{} `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// Details for a step that decrypts an encrypted file.
	//
	// Consists of the following values:
	//
	// - A descriptive name
	// - An Amazon S3 or Amazon Elastic File System (Amazon EFS) location for the source file to decrypt.
	// - An S3 or Amazon EFS location for the destination of the file decryption.
	// - A flag that indicates whether to overwrite an existing file of the same name. The default is `FALSE` .
	// - The type of encryption that's used. Currently, only PGP encryption is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-decryptstepdetails
	//
	DecryptStepDetails interface{} `field:"optional" json:"decryptStepDetails" yaml:"decryptStepDetails"`
	// Details for a step that deletes the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-deletestepdetails
	//
	DeleteStepDetails interface{} `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// Details for a step that creates one or more tags.
	//
	// You specify one or more tags. Each tag contains a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-tagstepdetails
	//
	TagStepDetails interface{} `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
	// Currently, the following step types are supported.
	//
	// - *`COPY`* - Copy the file to another location.
	// - *`CUSTOM`* - Perform a custom step with an AWS Lambda function target.
	// - *`DECRYPT`* - Decrypt a file that was encrypted before it was uploaded.
	// - *`DELETE`* - Delete the file.
	// - *`TAG`* - Add a tag to the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-workflowstep.html#cfn-transfer-workflow-workflowstep-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

