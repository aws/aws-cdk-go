package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretProperty := &SecretProperty{
//   	Name: jsii.String("name"),
//   	ValueFrom: jsii.String("valueFrom"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html
//
type CfnExpressGatewayService_SecretProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html#cfn-ecs-expressgatewayservice-secret-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-secret.html#cfn-ecs-expressgatewayservice-secret-valuefrom
	//
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

