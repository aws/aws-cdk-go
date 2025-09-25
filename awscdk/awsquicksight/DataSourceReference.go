package awsquicksight


// A reference to a DataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceReference := &DataSourceReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	DataSourceId: jsii.String("dataSourceId"),
//   }
//
type DataSourceReference struct {
	// The AwsAccountId of the DataSource resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the DataSource resource.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The DataSourceId of the DataSource resource.
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
}

