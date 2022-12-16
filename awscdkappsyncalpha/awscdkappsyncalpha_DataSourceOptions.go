// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Optional configuration for data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   dataSourceOptions := &dataSourceOptions{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
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

