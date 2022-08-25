package awsquicksight


// The parameters for Presto.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prestoParametersProperty := &prestoParametersProperty{
//   	catalog: jsii.String("catalog"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
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

