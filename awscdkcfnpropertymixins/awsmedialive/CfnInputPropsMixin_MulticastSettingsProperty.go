package awsmedialive


// Settings for a Multicast input.
//
// Contains a list of multicast Urls and optional source ip addresses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   multicastSettingsProperty := &MulticastSettingsProperty{
//   	Sources: []interface{}{
//   		&MulticastSourceProperty{
//   			SourceIp: jsii.String("sourceIp"),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettings.html
//
type CfnInputPropsMixin_MulticastSettingsProperty struct {
	// Placeholder documentation for __listOfMulticastSource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettings.html#cfn-medialive-input-multicastsettings-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

