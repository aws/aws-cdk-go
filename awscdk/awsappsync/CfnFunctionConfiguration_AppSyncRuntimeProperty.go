package awsappsync


// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function.
//
// Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncRuntimeProperty := &AppSyncRuntimeProperty{
//   	Name: jsii.String("name"),
//   	RuntimeVersion: jsii.String("runtimeVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-appsyncruntime.html
//
type CfnFunctionConfiguration_AppSyncRuntimeProperty struct {
	// The `name` of the runtime to use.
	//
	// Currently, the only allowed value is `APPSYNC_JS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-appsyncruntime.html#cfn-appsync-functionconfiguration-appsyncruntime-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The `version` of the runtime to use.
	//
	// Currently, the only allowed version is `1.0.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-appsyncruntime.html#cfn-appsync-functionconfiguration-appsyncruntime-runtimeversion
	//
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

