package previewawsquicksightmixins


// The image configuration of a table field URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableFieldImageConfigurationProperty := &TableFieldImageConfigurationProperty{
//   	SizingOptions: &TableCellImageSizingConfigurationProperty{
//   		TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldimageconfiguration.html
//
type CfnTemplatePropsMixin_TableFieldImageConfigurationProperty struct {
	// The sizing options for the table image configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldimageconfiguration.html#cfn-quicksight-template-tablefieldimageconfiguration-sizingoptions
	//
	SizingOptions interface{} `field:"optional" json:"sizingOptions" yaml:"sizingOptions"`
}

