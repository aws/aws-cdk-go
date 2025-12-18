package previewawsimagebuildermixins


// Contains deletion settings of underlying resources of an image when it is replaced or deleted, including its Amazon Machine Images (AMIs), snapshots, or containers.
//
// > If you specify the `Retain` option in the [DeletionPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-updatereplacepolicy.html) or [UpdateReplacePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-deletionpolicy.html) , the deletion of underlying resources will not be executed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deletionSettingsProperty := &DeletionSettingsProperty{
//   	ExecutionRole: jsii.String("executionRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-deletionsettings.html
//
type CfnImagePropsMixin_DeletionSettingsProperty struct {
	// The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to delete the image and its underlying resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-deletionsettings.html#cfn-imagebuilder-image-deletionsettings-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
}

