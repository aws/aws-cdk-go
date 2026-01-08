package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchConfigProperty := &BatchConfigProperty{
//   	MaxBatchOpenMs: jsii.Number(123),
//   	MaxBatchSize: jsii.Number(123),
//   	MaxBatchSizeBytes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-batchconfig.html
//
type CfnTopicRule_BatchConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-batchconfig.html#cfn-iot-topicrule-batchconfig-maxbatchopenms
	//
	MaxBatchOpenMs *float64 `field:"optional" json:"maxBatchOpenMs" yaml:"maxBatchOpenMs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-batchconfig.html#cfn-iot-topicrule-batchconfig-maxbatchsize
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-batchconfig.html#cfn-iot-topicrule-batchconfig-maxbatchsizebytes
	//
	MaxBatchSizeBytes *float64 `field:"optional" json:"maxBatchSizeBytes" yaml:"maxBatchSizeBytes"`
}

