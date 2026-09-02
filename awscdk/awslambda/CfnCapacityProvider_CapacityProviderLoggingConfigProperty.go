package awslambda


// The capacity provider's Amazon CloudWatch Logs configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderLoggingConfigProperty := &CapacityProviderLoggingConfigProperty{
//   	LogGroup: jsii.String("logGroup"),
//   	SystemLogLevel: jsii.String("systemLogLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html
//
type CfnCapacityProvider_CapacityProviderLoggingConfigProperty struct {
	// The name of the Amazon CloudWatch log group the capacity provider sends logs to.
	//
	// By default, Lambda capacity providers send logs to a default log group named ``/aws/lambda/capacity-provider/<capacity provider name>``. To use a different log group, enter an existing log group or enter a new log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html#cfn-lambda-capacityprovider-capacityproviderloggingconfig-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Set this property to filter the system logs for your capacity provider that Lambda sends to CloudWatch.
	//
	// Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html#cfn-lambda-capacityprovider-capacityproviderloggingconfig-systemloglevel
	//
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

