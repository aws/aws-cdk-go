package awsecs


// An object that defines the status of Express service creation and information about the status of the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayServiceStatusProperty := &ExpressGatewayServiceStatusProperty{
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicestatus.html
//
type CfnExpressGatewayService_ExpressGatewayServiceStatusProperty struct {
	// The status of the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicestatus.html#cfn-ecs-expressgatewayservice-expressgatewayservicestatus-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

