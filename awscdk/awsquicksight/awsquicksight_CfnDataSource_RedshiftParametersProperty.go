package awsquicksight


// The parameters for Amazon Redshift.
//
// The `ClusterId` field can be blank if `Host` and `Port` are both set. The `Host` and `Port` fields can be blank if the `ClusterId` field is set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftParametersProperty := &redshiftParametersProperty{
//   	database: jsii.String("database"),
//
//   	// the properties below are optional
//   	clusterId: jsii.String("clusterId"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_RedshiftParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cluster ID.
	//
	// This field can be blank if the `Host` and `Port` are provided.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Host.
	//
	// This field can be blank if `ClusterId` is provided.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Port.
	//
	// This field can be blank if the `ClusterId` is provided.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

