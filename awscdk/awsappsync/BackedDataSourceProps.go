package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// properties for an AppSync datasource backed by a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//   var role role
//
//   backedDataSourceProps := &BackedDataSourceProps{
//   	Api: graphqlApi,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   }
//
type BackedDataSourceProps struct {
	// The API to attach this data source to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Default: - id of data source.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Default: -  Create a new role.
	//
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
}

