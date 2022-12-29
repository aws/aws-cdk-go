package awstimestream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPropertiesProperty := &retentionPropertiesProperty{
//   	magneticStoreRetentionPeriodInDays: jsii.String("magneticStoreRetentionPeriodInDays"),
//   	memoryStoreRetentionPeriodInHours: jsii.String("memoryStoreRetentionPeriodInHours"),
//   }
//
type CfnTable_RetentionPropertiesProperty struct {
	// `CfnTable.RetentionPropertiesProperty.MagneticStoreRetentionPeriodInDays`.
	MagneticStoreRetentionPeriodInDays *string `field:"optional" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// `CfnTable.RetentionPropertiesProperty.MemoryStoreRetentionPeriodInHours`.
	MemoryStoreRetentionPeriodInHours *string `field:"optional" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

