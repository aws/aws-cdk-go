package awsmedialive


// Settings for a Multicast input.
//
// Contains a list of multicast Urls and optional source ip addresses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multicastSettingsProperty := &MulticastSettingsProperty{
//   	Sources: []interface{}{
//   		&MulticastSourceProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			SourceIp: jsii.String("sourceIp"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettings.html
//
type CfnInput_MulticastSettingsProperty struct {
	// Placeholder documentation for __listOfMulticastSource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsettings.html#cfn-medialive-input-multicastsettings-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

