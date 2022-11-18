// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// The attributes used to import existing Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   stageAttributes := &stageAttributes{
//   	stageName: jsii.String("stageName"),
//   }
//
// Experimental.
type StageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

