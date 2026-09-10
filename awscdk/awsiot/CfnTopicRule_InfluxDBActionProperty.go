package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBActionProperty := &InfluxDBActionProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	DestinationArn: jsii.String("destinationArn"),
//   	RoleArn: jsii.String("roleArn"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	BatchConfig: &InfluxDBBatchConfigProperty{
//   		BatchAcrossTopics: jsii.Boolean(false),
//   		MaxBatchOpenMs: jsii.Number(123),
//   		MaxBatchSize: jsii.Number(123),
//   		MaxBatchSizeBytes: jsii.Number(123),
//   	},
//   	Organization: jsii.String("organization"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TimestampUnit: jsii.String("timestampUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html
//
type CfnTopicRule_InfluxDBActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-destinationarn
	//
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-batchconfig
	//
	BatchConfig interface{} `field:"optional" json:"batchConfig" yaml:"batchConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-organization
	//
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-timestampunit
	//
	TimestampUnit *string `field:"optional" json:"timestampUnit" yaml:"timestampUnit"`
}

