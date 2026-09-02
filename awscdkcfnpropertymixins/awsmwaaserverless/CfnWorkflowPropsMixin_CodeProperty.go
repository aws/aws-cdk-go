package awsmwaaserverless


// The location of code artifacts in Amazon S3 for the workflow.
//
// Modeled as a single-member container so it stays extensible to future artifact types (e.g. OCI images).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeProperty := &CodeProperty{
//   	S3Location: &CodeS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		ObjectKey: jsii.String("objectKey"),
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-code.html
//
type CfnWorkflowPropsMixin_CodeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-code.html#cfn-mwaaserverless-workflow-code-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

