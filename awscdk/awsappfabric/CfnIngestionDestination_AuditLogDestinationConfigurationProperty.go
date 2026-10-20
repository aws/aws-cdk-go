package awsappfabric


// Contains information about an audit log destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditLogDestinationConfigurationProperty := &AuditLogDestinationConfigurationProperty{
//   	Destination: &DestinationProperty{
//   		FirehoseStream: &FirehoseStreamProperty{
//   			StreamName: jsii.String("streamName"),
//   		},
//   		S3Bucket: &S3BucketProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-auditlogdestinationconfiguration.html
//
type CfnIngestionDestination_AuditLogDestinationConfigurationProperty struct {
	// Contains information about an audit log destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-auditlogdestinationconfiguration.html#cfn-appfabric-ingestiondestination-auditlogdestinationconfiguration-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
}

