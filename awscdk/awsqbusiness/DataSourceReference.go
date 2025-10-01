package awsqbusiness


// A reference to a DataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceReference := &DataSourceReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	DataSourceId: jsii.String("dataSourceId"),
//   	IndexId: jsii.String("indexId"),
//   }
//
type DataSourceReference struct {
	// The ApplicationId of the DataSource resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the DataSource resource.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The DataSourceId of the DataSource resource.
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// The IndexId of the DataSource resource.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
}

