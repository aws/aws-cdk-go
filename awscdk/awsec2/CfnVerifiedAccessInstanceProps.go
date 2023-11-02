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
//   	FipsEnabled: jsii.Boolean(false),
//   	LoggingConfigurations: &VerifiedAccessLogsProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		IncludeTrustContext: jsii.Boolean(false),
//   		KinesisDataFirehose: &KinesisDataFirehoseProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		LogVersion: jsii.String("logVersion"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html
//
type CfnVerifiedAccessInstanceProps struct {
	// A description for the AWS Verified Access instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether support for Federal Information Processing Standards (FIPS) is enabled on the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-fipsenabled
	//
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// The logging configuration for the Verified Access instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-loggingconfigurations
	//
	LoggingConfigurations interface{} `field:"optional" json:"loggingConfigurations" yaml:"loggingConfigurations"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The IDs of the AWS Verified Access trust providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustproviderids
	//
	VerifiedAccessTrustProviderIds *[]*string `field:"optional" json:"verifiedAccessTrustProviderIds" yaml:"verifiedAccessTrustProviderIds"`
	// The IDs of the AWS Verified Access trust providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustproviders
	//
	VerifiedAccessTrustProviders interface{} `field:"optional" json:"verifiedAccessTrustProviders" yaml:"verifiedAccessTrustProviders"`
}

