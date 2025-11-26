package previewawsgreengrassv2mixins


// Contains information about a device that Linux processes in a container can access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaDeviceMountProperty := &LambdaDeviceMountProperty{
//   	AddGroupOwner: jsii.Boolean(false),
//   	Path: jsii.String("path"),
//   	Permission: jsii.String("permission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdadevicemount.html
//
type CfnComponentVersionPropsMixin_LambdaDeviceMountProperty struct {
	// Whether or not to add the component's system user as an owner of the device.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdadevicemount.html#cfn-greengrassv2-componentversion-lambdadevicemount-addgroupowner
	//
	AddGroupOwner interface{} `field:"optional" json:"addGroupOwner" yaml:"addGroupOwner"`
	// The mount path for the device in the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdadevicemount.html#cfn-greengrassv2-componentversion-lambdadevicemount-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The permission to access the device: read/only ( `ro` ) or read/write ( `rw` ).
	//
	// Default: `ro`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdadevicemount.html#cfn-greengrassv2-componentversion-lambdadevicemount-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

