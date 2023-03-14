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
//   logConfigurationProperty := &logConfigurationProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnPackagingGroup_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name for egress logs.
	//
	// If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

