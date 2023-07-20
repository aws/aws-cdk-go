package awsquicksight


// Parameters for Amazon Aurora.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraParametersProperty := &AuroraParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-auroraparameters.html
//
type CfnDataSource_AuroraParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-auroraparameters.html#cfn-quicksight-datasource-auroraparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-auroraparameters.html#cfn-quicksight-datasource-auroraparameters-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-auroraparameters.html#cfn-quicksight-datasource-auroraparameters-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

