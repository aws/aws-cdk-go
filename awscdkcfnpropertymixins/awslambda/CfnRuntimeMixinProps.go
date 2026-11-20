package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRuntimePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRuntimeMixinProps := &CfnRuntimeMixinProps{
//   	Description: jsii.String("description"),
//   	LicenseInfo: jsii.String("licenseInfo"),
//   	Name: jsii.String("name"),
//   	ReleaseLifecycle: &ReleaseLifecycleProperty{
//   		DeprecationDate: jsii.String("deprecationDate"),
//   		FunctionCreateBlockDate: jsii.String("functionCreateBlockDate"),
//   		FunctionUpdateBlockDate: jsii.String("functionUpdateBlockDate"),
//   		ReleaseState: jsii.String("releaseState"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html
//
type CfnRuntimeMixinProps struct {
	// A description of the runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html#cfn-lambda-runtime-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Licensing information for the runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html#cfn-lambda-runtime-licenseinfo
	//
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// The name of the runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html#cfn-lambda-runtime-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The publisher's lifecycle declaration for the runtime.
	//
	// Omitting it removes any stored declaration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html#cfn-lambda-runtime-releaselifecycle
	//
	ReleaseLifecycle interface{} `field:"optional" json:"releaseLifecycle" yaml:"releaseLifecycle"`
	// A list of tags to apply to the runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html#cfn-lambda-runtime-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

