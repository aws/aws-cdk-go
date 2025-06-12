package awsapigatewayv2


// The attributes used to import existing Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageAttributes := &StageAttributes{
//   	StageName: jsii.String("stageName"),
//   }
//
type StageAttributes struct {
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

