package awsquicksight


// The parameters that are required to connect to a Starburst data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   starburstParametersProperty := &StarburstParametersProperty{
//   	Catalog: jsii.String("catalog"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	ProductType: jsii.String("productType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html
//
type CfnDataSource_StarburstParametersProperty struct {
	// The catalog name for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-catalog
	//
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// The host name of the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The product type for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-producttype
	//
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
}

