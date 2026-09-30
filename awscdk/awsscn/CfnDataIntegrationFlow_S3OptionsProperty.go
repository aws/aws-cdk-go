package awsscn


// The Amazon S3 options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OptionsProperty := &S3OptionsProperty{
//   	FileType: jsii.String("fileType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3options.html
//
type CfnDataIntegrationFlow_S3OptionsProperty struct {
	// The file type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-s3options.html#cfn-scn-dataintegrationflow-s3options-filetype
	//
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

