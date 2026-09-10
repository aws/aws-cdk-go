package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBDestinationPropertiesProperty := &InfluxDBDestinationPropertiesProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	InfluxDbVersion: jsii.String("influxDbVersion"),
//   	SecretId: jsii.String("secretId"),
//
//   	// the properties below are optional
//   	SecretKey: jsii.String("secretKey"),
//   	SecretType: jsii.String("secretType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html
//
type CfnTopicRuleDestination_InfluxDBDestinationPropertiesProperty struct {
	// The endpoint URL of the InfluxDB database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html#cfn-iot-topicruledestination-influxdbdestinationproperties-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The version of the InfluxDB database (for example, V2 or V3).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html#cfn-iot-topicruledestination-influxdbdestinationproperties-influxdbversion
	//
	InfluxDbVersion *string `field:"required" json:"influxDbVersion" yaml:"influxDbVersion"`
	// The ARN or name of the Secrets Manager secret containing the InfluxDB API token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html#cfn-iot-topicruledestination-influxdbdestinationproperties-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The key name within the secret that contains the InfluxDB token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html#cfn-iot-topicruledestination-influxdbdestinationproperties-secretkey
	//
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// The type of the secret value (SecretString or SecretBinary).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicruledestination-influxdbdestinationproperties.html#cfn-iot-topicruledestination-influxdbdestinationproperties-secrettype
	//
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
}

