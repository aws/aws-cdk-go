package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionConfigurationProperty := &SessionConfigurationProperty{
//   	SessionTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-sessionconfiguration.html
//
type CfnGateway_SessionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-sessionconfiguration.html#cfn-bedrockagentcore-gateway-sessionconfiguration-sessiontimeoutinseconds
	//
	SessionTimeoutInSeconds *float64 `field:"optional" json:"sessionTimeoutInSeconds" yaml:"sessionTimeoutInSeconds"`
}

