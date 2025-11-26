package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisEventProperty := &KinesisEventProperty{
//   	BatchSize: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	FunctionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	Stream: jsii.String("stream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html
//
type CfnFunctionPropsMixin_KinesisEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html#cfn-serverless-function-kinesisevent-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html#cfn-serverless-function-kinesisevent-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html#cfn-serverless-function-kinesisevent-functionresponsetypes
	//
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html#cfn-serverless-function-kinesisevent-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-kinesisevent.html#cfn-serverless-function-kinesisevent-stream
	//
	Stream *string `field:"optional" json:"stream" yaml:"stream"`
}

