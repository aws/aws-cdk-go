package awsappsync


// Properties for an AppSync http datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//
//   httpDataSourceProps := &httpDataSourceProps{
//   	api: graphqlApi,
//   	endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	authorizationConfig: &awsIamConfig{
//   		signingRegion: jsii.String("signingRegion"),
//   		signingServiceName: jsii.String("signingServiceName"),
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   }
//
type HttpDataSourceProps struct {
	// The API to attach this data source to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The http endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The authorization config in case the HTTP endpoint requires authorization.
	AuthorizationConfig *AwsIamConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

