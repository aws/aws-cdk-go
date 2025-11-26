package previewawsquicksightmixins


// The parameters for Presto.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   prestoParametersProperty := &PrestoParametersProperty{
//   	Catalog: jsii.String("catalog"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html
//
type CfnDataSourcePropsMixin_PrestoParametersProperty struct {
	// Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-catalog
	//
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

