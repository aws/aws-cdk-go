package awsapigatewayv2


// The attributes used to import existing HttpStage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi HttpApi
//
//   httpStageAttributes := &HttpStageAttributes{
//   	Api: httpApi,
//   	StageName: jsii.String("stageName"),
//   }
//
type HttpStageAttributes struct {
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The API to which this stage is associated.
	Api IHttpApi `field:"required" json:"api" yaml:"api"`
}

