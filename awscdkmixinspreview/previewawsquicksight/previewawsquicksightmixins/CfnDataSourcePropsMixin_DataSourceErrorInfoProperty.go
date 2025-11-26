package previewawsquicksightmixins


// Error information for the data source creation or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceErrorInfoProperty := &DataSourceErrorInfoProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceerrorinfo.html
//
type CfnDataSourcePropsMixin_DataSourceErrorInfoProperty struct {
	// Error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceerrorinfo.html#cfn-quicksight-datasource-datasourceerrorinfo-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Error type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceerrorinfo.html#cfn-quicksight-datasource-datasourceerrorinfo-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

