package awslambda


// The runtime configuration for the revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   runtimeConfigProperty := &RuntimeConfigProperty{
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-runtimeconfig.html
//
type CfnWebFunctionRevisionPropsMixin_RuntimeConfigProperty struct {
	// The runtime identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-runtimeconfig.html#cfn-lambda-webfunctionrevision-runtimeconfig-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

