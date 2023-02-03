package awsappsync


// Optional configuration for data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceOptions := &dataSourceOptions{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   }
//
type DataSourceOptions struct {
	// The description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source, overrides the id given by cdk.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

