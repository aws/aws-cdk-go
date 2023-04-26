// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// The attributes used to import existing HttpStage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var httpApi httpApi
//
//   httpStageAttributes := &HttpStageAttributes{
//   	Api: httpApi,
//   	StageName: jsii.String("stageName"),
//   }
//
// Experimental.
type HttpStageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The API to which this stage is associated.
	// Experimental.
	Api IHttpApi `field:"required" json:"api" yaml:"api"`
}

