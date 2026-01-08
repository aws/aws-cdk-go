package interfacesawsapigatewayv2


// A reference to a Stage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageReference := &StageReference{
//   	StageName: jsii.String("stageName"),
//   }
//
type StageReference struct {
	// The StageName of the Stage resource.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

