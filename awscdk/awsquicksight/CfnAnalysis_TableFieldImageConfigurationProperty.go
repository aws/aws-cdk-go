package awsquicksight


// The image configuration of a table field URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldImageConfigurationProperty := &TableFieldImageConfigurationProperty{
//   	SizingOptions: &TableCellImageSizingConfigurationProperty{
//   		TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   	},
//   }
//
type CfnAnalysis_TableFieldImageConfigurationProperty struct {
	// The sizing options for the table image configuration.
	SizingOptions interface{} `field:"optional" json:"sizingOptions" yaml:"sizingOptions"`
}

