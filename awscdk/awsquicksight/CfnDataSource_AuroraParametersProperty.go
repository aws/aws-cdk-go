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
type CfnDataSource_AuroraParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

