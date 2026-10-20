package awsworkspaces


// Describes the Global Accelerator for directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   globalAcceleratorForDirectoryProperty := &GlobalAcceleratorForDirectoryProperty{
//   	Mode: jsii.String("mode"),
//   	PreferredProtocol: jsii.String("preferredProtocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-globalacceleratorfordirectory.html
//
type CfnDirectoryPropsMixin_GlobalAcceleratorForDirectoryProperty struct {
	// Indicates if Global Accelerator for directory is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-globalacceleratorfordirectory.html#cfn-workspaces-directory-globalacceleratorfordirectory-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Indicates the preferred protocol for Global Accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-globalacceleratorfordirectory.html#cfn-workspaces-directory-globalacceleratorfordirectory-preferredprotocol
	//
	PreferredProtocol *string `field:"optional" json:"preferredProtocol" yaml:"preferredProtocol"`
}

