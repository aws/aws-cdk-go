package awsappfabric


// Contains information about how ingested data is processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   processingConfigurationProperty := &ProcessingConfigurationProperty{
//   	AuditLog: &AuditLogProcessingConfigurationProperty{
//   		Format: jsii.String("format"),
//   		Schema: jsii.String("schema"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-processingconfiguration.html
//
type CfnIngestionDestinationPropsMixin_ProcessingConfigurationProperty struct {
	// Contains information about an audit log processing configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-processingconfiguration.html#cfn-appfabric-ingestiondestination-processingconfiguration-auditlog
	//
	AuditLog interface{} `field:"optional" json:"auditLog" yaml:"auditLog"`
}

