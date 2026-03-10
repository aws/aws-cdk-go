package previewawsmediapackageevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.mediapackage@MediaPackageInputNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageInputNotificationProps := &MediaPackageInputNotificationProps{
//   	Event: []*string{
//   		jsii.String("event"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	ManifestUrls: []*string{
//   		jsii.String("manifestUrls"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	PackagingConfigurationId: []*string{
//   		jsii.String("packagingConfigurationId"),
//   	},
//   }
//
// Experimental.
type MediaPackageInputNotification_MediaPackageInputNotificationProps struct {
	// event property.
	//
	// Specify an array of string values to match this event if the actual value of event is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// manifest_urls property.
	//
	// Specify an array of string values to match this event if the actual value of manifest_urls is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ManifestUrls *[]*string `field:"optional" json:"manifestUrls" yaml:"manifestUrls"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// packaging_configuration_id property.
	//
	// Specify an array of string values to match this event if the actual value of packaging_configuration_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PackagingConfigurationId *[]*string `field:"optional" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
}

