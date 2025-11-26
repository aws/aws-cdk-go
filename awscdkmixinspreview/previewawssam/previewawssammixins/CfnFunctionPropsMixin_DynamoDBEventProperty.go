package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamoDBEventProperty := &DynamoDBEventProperty{
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnFunctionError: jsii.Boolean(false),
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	ParallelizationFactor: jsii.Number(123),
//   	StartingPosition: jsii.String("startingPosition"),
//   	Stream: jsii.String("stream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html
//
type CfnFunctionPropsMixin_DynamoDBEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-bisectbatchonfunctionerror
	//
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-maximumrecordageinseconds
	//
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-parallelizationfactor
	//
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-startingposition
	//
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-dynamodbevent.html#cfn-serverless-function-dynamodbevent-stream
	//
	Stream *string `field:"optional" json:"stream" yaml:"stream"`
}

