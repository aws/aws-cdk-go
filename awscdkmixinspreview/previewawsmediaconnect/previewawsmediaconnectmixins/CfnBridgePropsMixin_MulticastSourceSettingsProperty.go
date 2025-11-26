package previewawsmediaconnectmixins


// The settings related to the multicast source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multicastSourceSettingsProperty := &MulticastSourceSettingsProperty{
//   	MulticastSourceIp: jsii.String("multicastSourceIp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-multicastsourcesettings.html
//
type CfnBridgePropsMixin_MulticastSourceSettingsProperty struct {
	// The IP address of the source for source-specific multicast (SSM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-multicastsourcesettings.html#cfn-mediaconnect-bridge-multicastsourcesettings-multicastsourceip
	//
	MulticastSourceIp *string `field:"optional" json:"multicastSourceIp" yaml:"multicastSourceIp"`
}

