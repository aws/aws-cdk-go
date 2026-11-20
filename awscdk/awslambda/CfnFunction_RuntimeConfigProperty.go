package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeConfigProperty := &RuntimeConfigProperty{
//   	RuntimeAlias: jsii.String("runtimeAlias"),
//   	RuntimeArn: jsii.String("runtimeArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimeconfig.html
//
type CfnFunction_RuntimeConfigProperty struct {
	// The name of the runtime alias that controls the function's runtime update cadence (for example, stable or latest).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimeconfig.html#cfn-lambda-function-runtimeconfig-runtimealias
	//
	RuntimeAlias *string `field:"required" json:"runtimeAlias" yaml:"runtimeAlias"`
	// The Amazon Resource Name (ARN) of the runtime that the function uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimeconfig.html#cfn-lambda-function-runtimeconfig-runtimearn
	//
	RuntimeArn *string `field:"required" json:"runtimeArn" yaml:"runtimeArn"`
}

