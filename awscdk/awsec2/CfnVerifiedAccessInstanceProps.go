package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVerifiedAccessInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVerifiedAccessInstanceProps := &CfnVerifiedAccessInstanceProps{
//   	Description: jsii.String("description"),
//   	LoggingConfigurations: &VerifiedAccessLogsProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		KinesisDataFirehose: &KinesisDataFirehoseProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3Property{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VerifiedAccessTrustProviderIds: []*string{
//   		jsii.String("verifiedAccessTrustProviderIds"),
//   	},
//   	VerifiedAccessTrustProviders: []interface{}{
//   		&VerifiedAccessTrustProviderProperty{
//   			Description: jsii.String("description"),
//   			DeviceTrustProviderType: jsii.String("deviceTrustProviderType"),
//   			TrustProviderType: jsii.String("trustProviderType"),
//   			UserTrustProviderType: jsii.String("userTrustProviderType"),
//   			VerifiedAccessTrustProviderId: jsii.String("verifiedAccessTrustProviderId"),
//   		},
//   	},
//   }
//
type CfnVerifiedAccessInstanceProps struct {
	// A description for the AWS Verified Access instance.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The current logging configuration for the Verified Access instances.
	LoggingConfigurations interface{} `field:"optional" json:"loggingConfigurations" yaml:"loggingConfigurations"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The IDs of the AWS Verified Access trust providers.
	VerifiedAccessTrustProviderIds *[]*string `field:"optional" json:"verifiedAccessTrustProviderIds" yaml:"verifiedAccessTrustProviderIds"`
	// The IDs of the AWS Verified Access trust providers.
	VerifiedAccessTrustProviders interface{} `field:"optional" json:"verifiedAccessTrustProviders" yaml:"verifiedAccessTrustProviders"`
}

