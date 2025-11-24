package mixinsawsquicksight


// The usage configuration to apply to child datasets that reference this dataset as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetUsageConfigurationProperty := &DataSetUsageConfigurationProperty{
//   	DisableUseAsDirectQuerySource: jsii.Boolean(false),
//   	DisableUseAsImportedSource: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetusageconfiguration.html
//
type CfnDataSetPropsMixin_DataSetUsageConfigurationProperty struct {
	// An option that controls whether a child dataset of a direct query can use this dataset as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetusageconfiguration.html#cfn-quicksight-dataset-datasetusageconfiguration-disableuseasdirectquerysource
	//
	// Default: - false.
	//
	DisableUseAsDirectQuerySource interface{} `field:"optional" json:"disableUseAsDirectQuerySource" yaml:"disableUseAsDirectQuerySource"`
	// An option that controls whether a child dataset that's stored in Quick Sight can use this dataset as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetusageconfiguration.html#cfn-quicksight-dataset-datasetusageconfiguration-disableuseasimportedsource
	//
	// Default: - false.
	//
	DisableUseAsImportedSource interface{} `field:"optional" json:"disableUseAsImportedSource" yaml:"disableUseAsImportedSource"`
}

