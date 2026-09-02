package awsdevicefarm


// Properties for defining a `CfnUpload`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUploadProps := &CfnUploadProps{
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ContentType: jsii.String("contentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html
//
type CfnUploadProps struct {
	// The upload's file name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the project for the upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-projectarn
	//
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
	// The upload's type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The upload's content type (for example, application/octet-stream).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

