package previewawsquicksightmixins


// The sizing options for the table image configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableCellImageSizingConfigurationProperty := &TableCellImageSizingConfigurationProperty{
//   	TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablecellimagesizingconfiguration.html
//
type CfnTemplatePropsMixin_TableCellImageSizingConfigurationProperty struct {
	// The cell scaling configuration of the sizing options for the table image configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablecellimagesizingconfiguration.html#cfn-quicksight-template-tablecellimagesizingconfiguration-tablecellimagescalingconfiguration
	//
	TableCellImageScalingConfiguration *string `field:"optional" json:"tableCellImageScalingConfiguration" yaml:"tableCellImageScalingConfiguration"`
}

