package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   secretProperty := &SecretProperty{
//   	Name: jsii.String("name"),
//   	ValueFrom: jsii.String("valueFrom"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html
//
type CfnExpressGatewayServicePropsMixin_SecretProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html#cfn-ecs-expressgatewayservice-secret-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html#cfn-ecs-expressgatewayservice-secret-valuefrom
	//
	ValueFrom *string `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

