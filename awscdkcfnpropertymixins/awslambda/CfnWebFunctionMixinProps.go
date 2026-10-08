package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWebFunctionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWebFunctionMixinProps := &CfnWebFunctionMixinProps{
//   	FunctionName: jsii.String("functionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunction.html
//
type CfnWebFunctionMixinProps struct {
	// The name of the web function.
	//
	// The name can contain letters, numbers, hyphens (-), and underscores (_), and cannot begin or end with a hyphen or an underscore. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunction.html#cfn-lambda-webfunction-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// A list of tags to apply to the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunction.html#cfn-lambda-webfunction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

