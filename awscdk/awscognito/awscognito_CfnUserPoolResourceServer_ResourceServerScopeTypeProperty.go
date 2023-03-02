package awscognito


// A resource server scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceServerScopeTypeProperty := &resourceServerScopeTypeProperty{
//   	scopeDescription: jsii.String("scopeDescription"),
//   	scopeName: jsii.String("scopeName"),
//   }
//
type CfnUserPoolResourceServer_ResourceServerScopeTypeProperty struct {
	// A description of the scope.
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// The name of the scope.
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

