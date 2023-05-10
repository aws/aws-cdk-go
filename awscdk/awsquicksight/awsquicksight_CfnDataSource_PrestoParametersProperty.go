package awsquicksight


// The parameters for Presto.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prestoParametersProperty := &PrestoParametersProperty{
//   	Catalog: jsii.String("catalog"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
type CfnDataSource_PrestoParametersProperty struct {
	// Catalog.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

