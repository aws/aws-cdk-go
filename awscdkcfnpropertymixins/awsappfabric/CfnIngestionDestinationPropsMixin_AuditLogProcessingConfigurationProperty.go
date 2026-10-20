package awsappfabric


// Contains information about an audit log processing configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   auditLogProcessingConfigurationProperty := &AuditLogProcessingConfigurationProperty{
//   	Format: jsii.String("format"),
//   	Schema: jsii.String("schema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-auditlogprocessingconfiguration.html
//
type CfnIngestionDestinationPropsMixin_AuditLogProcessingConfigurationProperty struct {
	// The format in which the audit logs need to be formatted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-auditlogprocessingconfiguration.html#cfn-appfabric-ingestiondestination-auditlogprocessingconfiguration-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The event schema in which the audit logs need to be formatted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-auditlogprocessingconfiguration.html#cfn-appfabric-ingestiondestination-auditlogprocessingconfiguration-schema
	//
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

