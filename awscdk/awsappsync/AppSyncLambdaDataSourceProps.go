package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for an AppSync Lambda datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var api IApi
//   var function_ Function
//   var role Role
//
//   appSyncLambdaDataSourceProps := &AppSyncLambdaDataSourceProps{
//   	Api: api,
//   	LambdaFunction: function_,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   }
//
type AppSyncLambdaDataSourceProps struct {
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
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Default: -  Create a new role.
	//
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The Lambda function to call to interact with this data source.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
}

