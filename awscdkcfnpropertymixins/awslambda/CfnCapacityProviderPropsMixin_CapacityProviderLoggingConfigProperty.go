package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capacityProviderLoggingConfigProperty := &CapacityProviderLoggingConfigProperty{
//   	LogGroup: jsii.String("logGroup"),
//   	SystemLogLevel: jsii.String("systemLogLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html
//
type CfnCapacityProviderPropsMixin_CapacityProviderLoggingConfigProperty struct {
	// The log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html#cfn-lambda-capacityprovider-capacityproviderloggingconfig-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// System log granularity level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderloggingconfig.html#cfn-lambda-capacityprovider-capacityproviderloggingconfig-systemloglevel
	//
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

