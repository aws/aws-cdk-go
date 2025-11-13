package interfacesawsappsync


// A reference to a DataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceReference := &DataSourceReference{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   }
//
type DataSourceReference struct {
	// The DataSourceArn of the DataSource resource.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
}

