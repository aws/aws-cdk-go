package awsquicksight


// The sizing options for the table image configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableCellImageSizingConfigurationProperty := &TableCellImageSizingConfigurationProperty{
//   	TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   }
//
type CfnTemplate_TableCellImageSizingConfigurationProperty struct {
	// The cell scaling configuration of the sizing options for the table image configuration.
	TableCellImageScalingConfiguration *string `field:"optional" json:"tableCellImageScalingConfiguration" yaml:"tableCellImageScalingConfiguration"`
}

