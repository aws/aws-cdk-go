package awsappsync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncRuntimeProperty := &appSyncRuntimeProperty{
//   	name: jsii.String("name"),
//   	runtimeVersion: jsii.String("runtimeVersion"),
//   }
//
type CfnResolver_AppSyncRuntimeProperty struct {
	// `CfnResolver.AppSyncRuntimeProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnResolver.AppSyncRuntimeProperty.RuntimeVersion`.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

