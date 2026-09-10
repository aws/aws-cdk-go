package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   influxDBActionProperty := &InfluxDBActionProperty{
//   	BatchConfig: &InfluxDBBatchConfigProperty{
//   		BatchAcrossTopics: jsii.Boolean(false),
//   		MaxBatchOpenMs: jsii.Number(123),
//   		MaxBatchSize: jsii.Number(123),
//   		MaxBatchSizeBytes: jsii.Number(123),
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   	DestinationArn: jsii.String("destinationArn"),
//   	Organization: jsii.String("organization"),
//   	RoleArn: jsii.String("roleArn"),
//   	TableName: jsii.String("tableName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TimestampUnit: jsii.String("timestampUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html
//
type CfnTopicRulePropsMixin_InfluxDBActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-batchconfig
	//
	BatchConfig interface{} `field:"optional" json:"batchConfig" yaml:"batchConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-organization
	//
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-influxdbaction.html#cfn-iot-topicrule-influxdbaction-timestampunit
	//
	TimestampUnit *string `field:"optional" json:"timestampUnit" yaml:"timestampUnit"`
}

