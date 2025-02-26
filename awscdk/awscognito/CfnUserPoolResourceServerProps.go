package awscognito


// Properties for defining a `CfnUserPoolResourceServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolResourceServerProps := &CfnUserPoolResourceServerProps{
//   	Identifier: jsii.String("identifier"),
//   	Name: jsii.String("name"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	Scopes: []interface{}{
//   		&ResourceServerScopeTypeProperty{
//   			ScopeDescription: jsii.String("scopeDescription"),
//   			ScopeName: jsii.String("scopeName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html
//
type CfnUserPoolResourceServerProps struct {
	// A unique resource server identifier for the resource server.
	//
	// The identifier can be an API friendly name like `solar-system-data` . You can also set an API URL like `https://solar-system-data-api.example.com` as your identifier.
	//
	// Amazon Cognito represents scopes in the access token in the format `$resource-server-identifier/$scope` . Longer scope-identifier strings increase the size of your access tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html#cfn-cognito-userpoolresourceserver-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// A friendly name for the resource server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html#cfn-cognito-userpoolresourceserver-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the user pool where you want to create a resource server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html#cfn-cognito-userpoolresourceserver-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A list of scopes.
	//
	// Each scope is a map with keys `ScopeName` and `ScopeDescription` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolresourceserver.html#cfn-cognito-userpoolresourceserver-scopes
	//
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
}

