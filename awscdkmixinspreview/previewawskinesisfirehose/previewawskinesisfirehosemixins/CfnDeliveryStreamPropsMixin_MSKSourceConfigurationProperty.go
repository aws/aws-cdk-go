package previewawskinesisfirehosemixins


// The configuration for the Amazon MSK cluster to be used as the source for a delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mSKSourceConfigurationProperty := &MSKSourceConfigurationProperty{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		Connectivity: jsii.String("connectivity"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	MskClusterArn: jsii.String("mskClusterArn"),
//   	ReadFromTimestamp: jsii.String("readFromTimestamp"),
//   	TopicName: jsii.String("topicName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html
//
type CfnDeliveryStreamPropsMixin_MSKSourceConfigurationProperty struct {
	// The authentication configuration of the Amazon MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// The ARN of the Amazon MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-mskclusterarn
	//
	MskClusterArn *string `field:"optional" json:"mskClusterArn" yaml:"mskClusterArn"`
	// The start date and time in UTC for the offset position within your MSK topic from where Firehose begins to read.
	//
	// By default, this is set to timestamp when Firehose becomes Active.
	//
	// If you want to create a Firehose stream with Earliest start position from SDK or CLI, you need to set the `ReadFromTimestamp` parameter to Epoch (1970-01-01T00:00:00Z).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-readfromtimestamp
	//
	ReadFromTimestamp *string `field:"optional" json:"readFromTimestamp" yaml:"readFromTimestamp"`
	// The topic name within the Amazon MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-topicname
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

