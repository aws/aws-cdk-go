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
type CfnFunctionConfiguration_AppSyncRuntimeProperty struct {
	// `CfnFunctionConfiguration.AppSyncRuntimeProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnFunctionConfiguration.AppSyncRuntimeProperty.RuntimeVersion`.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

