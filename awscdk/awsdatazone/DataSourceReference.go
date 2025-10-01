package awsdatazone


// A reference to a DataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceReference := &DataSourceReference{
//   	DataSourceId: jsii.String("dataSourceId"),
//   	DomainId: jsii.String("domainId"),
//   }
//
type DataSourceReference struct {
	// The Id of the DataSource resource.
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// The DomainId of the DataSource resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
}

