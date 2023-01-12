package awsaps


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnWorkspace_LoggingConfigurationProperty struct {
	// `CfnWorkspace.LoggingConfigurationProperty.LogGroupArn`.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

