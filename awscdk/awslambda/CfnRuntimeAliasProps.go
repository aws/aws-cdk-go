package awslambda


// Properties for defining a `CfnRuntimeAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuntimeAliasProps := &CfnRuntimeAliasProps{
//   	Name: jsii.String("name"),
//   	RuntimeName: jsii.String("runtimeName"),
//   	RuntimeVersion: jsii.String("runtimeVersion"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ReleaseLifecycle: &ReleaseLifecycleProperty{
//   		DeprecationDate: jsii.String("deprecationDate"),
//   		FunctionCreateBlockDate: jsii.String("functionCreateBlockDate"),
//   		FunctionUpdateBlockDate: jsii.String("functionUpdateBlockDate"),
//   		ReleaseState: jsii.String("releaseState"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html
//
type CfnRuntimeAliasProps struct {
	// The alias name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html#cfn-lambda-runtimealias-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the parent runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html#cfn-lambda-runtimealias-runtimename
	//
	RuntimeName *string `field:"required" json:"runtimeName" yaml:"runtimeName"`
	// The runtime version this alias points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html#cfn-lambda-runtimealias-runtimeversion
	//
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// A description of the runtime alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html#cfn-lambda-runtimealias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The publisher's lifecycle declaration for the alias.
	//
	// Omitting it removes any stored declaration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimealias.html#cfn-lambda-runtimealias-releaselifecycle
	//
	ReleaseLifecycle interface{} `field:"optional" json:"releaseLifecycle" yaml:"releaseLifecycle"`
}

