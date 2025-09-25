package awsapigatewayv2


// A reference to a Stage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageReference := &StageReference{
//   	StageId: jsii.String("stageId"),
//   }
//
type StageReference struct {
	// The Id of the Stage resource.
	StageId *string `field:"required" json:"stageId" yaml:"stageId"`
}

