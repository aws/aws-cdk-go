package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecordingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordingConfigurationProps := &cfnRecordingConfigurationProps{
//   	destinationConfiguration: &destinationConfigurationProperty{
//   		s3: &s3DestinationConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	thumbnailConfiguration: &thumbnailConfigurationProperty{
//   		recordingMode: jsii.String("recordingMode"),
//
//   		// the properties below are optional
//   		targetIntervalSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnRecordingConfigurationProps struct {
	// A destination configuration contains information about where recorded video will be stored.
	//
	// See the [DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html) property type for more information.
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Recording-configuration name.
	//
	// The value does not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session.
	//
	// See the [ThumbnailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thunbnailconfiguration.html) property type for more information.
	ThumbnailConfiguration interface{} `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

