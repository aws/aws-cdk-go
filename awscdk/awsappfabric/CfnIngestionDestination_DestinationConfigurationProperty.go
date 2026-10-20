package awsappfabric


// Contains information about the destination of ingested data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigurationProperty := &DestinationConfigurationProperty{
//   	AuditLog: &AuditLogDestinationConfigurationProperty{
//   		Destination: &DestinationProperty{
//   			FirehoseStream: &FirehoseStreamProperty{
//   				StreamName: jsii.String("streamName"),
//   			},
//   			S3Bucket: &S3BucketProperty{
//   				BucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-destinationconfiguration.html
//
type CfnIngestionDestination_DestinationConfigurationProperty struct {
	// Contains information about an audit log destination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-destinationconfiguration.html#cfn-appfabric-ingestiondestination-destinationconfiguration-auditlog
	//
	AuditLog interface{} `field:"required" json:"auditLog" yaml:"auditLog"`
}

