package previewawsmediapackageevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for OriginEndpoint aws.mediapackage@MediaPackageHarvestJobNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaPackageHarvestJobNotificationProps := &MediaPackageHarvestJobNotificationProps{
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
//   	HarvestJob: &HarvestJob{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		CreatedAt: []*string{
//   			jsii.String("createdAt"),
//   		},
//   		EndTime: []*string{
//   			jsii.String("endTime"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		OriginEndpointId: []*string{
//   			jsii.String("originEndpointId"),
//   		},
//   		S3Destination: &S3Destination{
//   			BucketName: []*string{
//   				jsii.String("bucketName"),
//   			},
//   			ManifestKey: []*string{
//   				jsii.String("manifestKey"),
//   			},
//   			RoleArn: []*string{
//   				jsii.String("roleArn"),
//   			},
//   		},
//   		StartTime: []*string{
//   			jsii.String("startTime"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   	},
//   }
//
// Experimental.
type OriginEndpointEvents_MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// harvest_job property.
	//
	// Specify an array of string values to match this event if the actual value of harvest_job is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HarvestJob *OriginEndpointEvents_MediaPackageHarvestJobNotification_HarvestJob `field:"optional" json:"harvestJob" yaml:"harvestJob"`
}

