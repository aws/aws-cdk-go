package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mSKSourceConfigurationProperty := &MSKSourceConfigurationProperty{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		Connectivity: jsii.String("connectivity"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	MskClusterArn: jsii.String("mskClusterArn"),
//   	TopicName: jsii.String("topicName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html
//
type CfnDeliveryStream_MSKSourceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-mskclusterarn
	//
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.html#cfn-kinesisfirehose-deliverystream-msksourceconfiguration-topicname
	//
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

