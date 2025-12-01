package previewawsecsmixins


// A key-value pair object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyValuePairProperty := &KeyValuePairProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-keyvaluepair.html
//
type CfnExpressGatewayServicePropsMixin_KeyValuePairProperty struct {
	// The name of the key-value pair.
	//
	// For environment variables, this is the name of the environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-keyvaluepair.html#cfn-ecs-expressgatewayservice-keyvaluepair-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair.
	//
	// For environment variables, this is the value of the environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-keyvaluepair.html#cfn-ecs-expressgatewayservice-keyvaluepair-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

