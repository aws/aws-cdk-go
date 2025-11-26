package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   expressGatewayServiceAwsLogsConfigurationProperty := &ExpressGatewayServiceAwsLogsConfigurationProperty{
//   	LogGroup: jsii.String("logGroup"),
//   	LogStreamPrefix: jsii.String("logStreamPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration.html
//
type CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceAwsLogsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration-logstreamprefix
	//
	// Default: - "ecs".
	//
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
}

