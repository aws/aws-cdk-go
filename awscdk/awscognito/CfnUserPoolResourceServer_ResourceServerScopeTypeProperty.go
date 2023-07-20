package awscognito


// A resource server scope.
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
	// A description of the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.html#cfn-cognito-userpoolresourceserver-resourceserverscopetype-scopedescription
	//
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.html#cfn-cognito-userpoolresourceserver-resourceserverscopetype-scopename
	//
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

