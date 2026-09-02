package awsdevicefarm


// Properties for CfnUploadPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUploadMixinProps := &CfnUploadMixinProps{
//   	ContentType: jsii.String("contentType"),
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html
//
type CfnUploadMixinProps struct {
	// The upload's content type (for example, application/octet-stream).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The upload's file name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the project for the upload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-projectarn
	//
	ProjectArn *string `field:"optional" json:"projectArn" yaml:"projectArn"`
	// The upload's type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-upload.html#cfn-devicefarm-upload-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

