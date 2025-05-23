package awsappsync


// Base properties for an AppSync datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var api iApi
//
//   appSyncBaseDataSourceProps := &AppSyncBaseDataSourceProps{
//   	Api: api,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
type AppSyncBaseDataSourceProps struct {
	// The API to attach this data source to.
	Api IApi `field:"required" json:"api" yaml:"api"`
	// The description of the data source.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	//
	// The only allowed pattern is: {[_A-Za-z][_0-9A-Za-z]*}.
	// Any invalid characters will be automatically removed.
	// Default: - id of data source.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

