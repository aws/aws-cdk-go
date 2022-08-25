package awsquicksight


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
	// `CfnDataSet.DataSetUsageConfigurationProperty.DisableUseAsDirectQuerySource`.
	DisableUseAsDirectQuerySource interface{} `field:"optional" json:"disableUseAsDirectQuerySource" yaml:"disableUseAsDirectQuerySource"`
	// `CfnDataSet.DataSetUsageConfigurationProperty.DisableUseAsImportedSource`.
	DisableUseAsImportedSource interface{} `field:"optional" json:"disableUseAsImportedSource" yaml:"disableUseAsImportedSource"`
}

