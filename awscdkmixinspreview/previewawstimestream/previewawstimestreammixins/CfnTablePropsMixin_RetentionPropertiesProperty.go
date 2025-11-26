package previewawstimestreammixins


// Retention properties contain the duration for which your time-series data must be stored in the magnetic store and the memory store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionPropertiesProperty := &RetentionPropertiesProperty{
//   	MagneticStoreRetentionPeriodInDays: jsii.String("magneticStoreRetentionPeriodInDays"),
//   	MemoryStoreRetentionPeriodInHours: jsii.String("memoryStoreRetentionPeriodInHours"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-retentionproperties.html
//
type CfnTablePropsMixin_RetentionPropertiesProperty struct {
	// The duration for which data must be stored in the magnetic store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-retentionproperties.html#cfn-timestream-table-retentionproperties-magneticstoreretentionperiodindays
	//
	MagneticStoreRetentionPeriodInDays *string `field:"optional" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// The duration for which data must be stored in the memory store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-retentionproperties.html#cfn-timestream-table-retentionproperties-memorystoreretentionperiodinhours
	//
	MemoryStoreRetentionPeriodInHours *string `field:"optional" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

