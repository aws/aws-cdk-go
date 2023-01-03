package awsquicksight


// The usage configuration to apply to child datasets that reference this dataset as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetUsageConfigurationProperty := &dataSetUsageConfigurationProperty{
//   	disableUseAsDirectQuerySource: jsii.Boolean(false),
//   	disableUseAsImportedSource: jsii.Boolean(false),
//   }
//
type CfnDataSet_DataSetUsageConfigurationProperty struct {
	// An option that controls whether a child dataset of a direct query can use this dataset as a source.
	DisableUseAsDirectQuerySource interface{} `field:"optional" json:"disableUseAsDirectQuerySource" yaml:"disableUseAsDirectQuerySource"`
	// An option that controls whether a child dataset that's stored in QuickSight can use this dataset as a source.
	DisableUseAsImportedSource interface{} `field:"optional" json:"disableUseAsImportedSource" yaml:"disableUseAsImportedSource"`
}

