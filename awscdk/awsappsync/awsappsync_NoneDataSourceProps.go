package awsappsync


// Properties for an AppSync dummy datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//
//   noneDataSourceProps := &noneDataSourceProps{
//   	api: graphqlApi,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   }
//
// Experimental.
type NoneDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

