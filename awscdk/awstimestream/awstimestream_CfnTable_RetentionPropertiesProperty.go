package awstimestream


// Retention properties contain the duration for which your time-series data must be stored in the magnetic store and the memory store.
//
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
	// The duration for which data must be stored in the magnetic store.
	MagneticStoreRetentionPeriodInDays *string `field:"optional" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// The duration for which data must be stored in the memory store.
	MemoryStoreRetentionPeriodInHours *string `field:"optional" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

