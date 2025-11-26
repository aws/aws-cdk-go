package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hooksProperty := &HooksProperty{
//   	PostTraffic: jsii.String("postTraffic"),
//   	PreTraffic: jsii.String("preTraffic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-hooks.html
//
type CfnFunctionPropsMixin_HooksProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-hooks.html#cfn-serverless-function-hooks-posttraffic
	//
	PostTraffic *string `field:"optional" json:"postTraffic" yaml:"postTraffic"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-hooks.html#cfn-serverless-function-hooks-pretraffic
	//
	PreTraffic *string `field:"optional" json:"preTraffic" yaml:"preTraffic"`
}

