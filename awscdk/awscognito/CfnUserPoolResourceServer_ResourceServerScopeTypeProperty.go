package awscognito


// One custom scope associated with a user pool resource server.
//
// This data type is a member of `ResourceServerScopeType` . For more information, see [Scopes, M2M, and API authorization with resource servers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceServerScopeTypeProperty := &ResourceServerScopeTypeProperty{
//   	ScopeDescription: jsii.String("scopeDescription"),
//   	ScopeName: jsii.String("scopeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.html
//
type CfnUserPoolResourceServer_ResourceServerScopeTypeProperty struct {
	// A friendly description of a custom scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.html#cfn-cognito-userpoolresourceserver-resourceserverscopetype-scopedescription
	//
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	//
	// Amazon Cognito renders custom scopes in the format `resourceServerIdentifier/ScopeName` . For example, if this parameter is `exampleScope` in the resource server with the identifier `exampleResourceServer` , you request and receive the scope `exampleResourceServer/exampleScope` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.html#cfn-cognito-userpoolresourceserver-resourceserverscopetype-scopename
	//
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

