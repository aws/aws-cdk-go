package awsappfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIngestionDestinationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnIngestionDestinationMixinProps := &CfnIngestionDestinationMixinProps{
//   	AppBundleArn: jsii.String("appBundleArn"),
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		AuditLog: &AuditLogDestinationConfigurationProperty{
//   			Destination: &DestinationProperty{
//   				FirehoseStream: &FirehoseStreamProperty{
//   					StreamName: jsii.String("streamName"),
//   				},
//   				S3Bucket: &S3BucketProperty{
//   					BucketName: jsii.String("bucketName"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   		},
//   	},
//   	IngestionArn: jsii.String("ingestionArn"),
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		AuditLog: &AuditLogProcessingConfigurationProperty{
//   			Format: jsii.String("format"),
//   			Schema: jsii.String("schema"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html
//
type CfnIngestionDestinationMixinProps struct {
	// The Amazon Resource Name (ARN) of the app bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html#cfn-appfabric-ingestiondestination-appbundlearn
	//
	AppBundleArn *string `field:"optional" json:"appBundleArn" yaml:"appBundleArn"`
	// Contains information about the destination of ingested data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html#cfn-appfabric-ingestiondestination-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// The Amazon Resource Name (ARN) of the ingestion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html#cfn-appfabric-ingestiondestination-ingestionarn
	//
	IngestionArn *string `field:"optional" json:"ingestionArn" yaml:"ingestionArn"`
	// Contains information about how ingested data is processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html#cfn-appfabric-ingestiondestination-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestiondestination.html#cfn-appfabric-ingestiondestination-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

