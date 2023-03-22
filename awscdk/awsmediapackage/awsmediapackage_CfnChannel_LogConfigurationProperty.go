package awsmediapackage


// The access log configuration parameters for your channel.
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
type CfnChannel_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

