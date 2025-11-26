package previewawsbatchmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mountPointsProperty := &MountPointsProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	ReadOnly: jsii.Boolean(false),
//   	SourceVolume: jsii.String("sourceVolume"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html
//
type CfnJobDefinitionPropsMixin_MountPointsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-sourcevolume
	//
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

