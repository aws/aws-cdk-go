package awsmediapackage


// Sets a custom Amazon CloudWatch log group name for egress logs.
//
// If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packaginggroup-logconfiguration.html
//
type CfnPackagingGroup_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name for egress logs.
	//
	// If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packaginggroup-logconfiguration.html#cfn-mediapackage-packaginggroup-logconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

