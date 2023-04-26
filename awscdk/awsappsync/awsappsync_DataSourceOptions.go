package awsappsync


// Optional configuration for data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceOptions := &DataSourceOptions{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type DataSourceOptions struct {
	// The description of the data source.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source, overrides the id given by cdk.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

