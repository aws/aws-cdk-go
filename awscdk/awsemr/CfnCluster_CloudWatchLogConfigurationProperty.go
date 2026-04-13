package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logTypes interface{}
//
//   cloudWatchLogConfigurationProperty := &CloudWatchLogConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   	LogTypes: logTypes,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html
//
type CfnCluster_CloudWatchLogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-logstreamnameprefix
	//
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-cloudwatchlogconfiguration.html#cfn-emr-cluster-cloudwatchlogconfiguration-logtypes
	//
	LogTypes interface{} `field:"optional" json:"logTypes" yaml:"logTypes"`
}

