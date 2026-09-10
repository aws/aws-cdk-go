package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBBatchConfigProperty := &InfluxDBBatchConfigProperty{
//   	BatchAcrossTopics: jsii.Boolean(false),
//   	MaxBatchOpenMs: jsii.Number(123),
//   	MaxBatchSize: jsii.Number(123),
//   	MaxBatchSizeBytes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbbatchconfig.html
//
type CfnTopicRule_InfluxDBBatchConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbbatchconfig.html#cfn-iot-topicrule-influxdbbatchconfig-batchacrosstopics
	//
	BatchAcrossTopics interface{} `field:"optional" json:"batchAcrossTopics" yaml:"batchAcrossTopics"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbbatchconfig.html#cfn-iot-topicrule-influxdbbatchconfig-maxbatchopenms
	//
	MaxBatchOpenMs *float64 `field:"optional" json:"maxBatchOpenMs" yaml:"maxBatchOpenMs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbbatchconfig.html#cfn-iot-topicrule-influxdbbatchconfig-maxbatchsize
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbbatchconfig.html#cfn-iot-topicrule-influxdbbatchconfig-maxbatchsizebytes
	//
	MaxBatchSizeBytes *float64 `field:"optional" json:"maxBatchSizeBytes" yaml:"maxBatchSizeBytes"`
}

