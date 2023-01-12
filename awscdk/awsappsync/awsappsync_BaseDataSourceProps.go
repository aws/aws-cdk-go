package awsappsync


// Base properties for an AppSync datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//
//   baseDataSourceProps := &baseDataSourceProps{
//   	api: graphqlApi,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   }
//
type BaseDataSourceProps struct {
	// The API to attach this data source to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

