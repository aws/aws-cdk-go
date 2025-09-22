package awsbedrockalpha


// Properties for configuring a Prompt Router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var bedrockFoundationModel bedrockFoundationModel
//
//   promptRouterProps := &PromptRouterProps{
//   	PromptRouterId: jsii.String("promptRouterId"),
//   	RoutingModels: []*bedrockFoundationModel{
//   		bedrockFoundationModel,
//   	},
//   }
//
// Experimental.
type PromptRouterProps struct {
	// Prompt Router ID that identifies the routing configuration.
	// Experimental.
	PromptRouterId *string `field:"required" json:"promptRouterId" yaml:"promptRouterId"`
	// The foundation models this router will route to.
	//
	// The router will intelligently select between these models based on the request.
	// Experimental.
	RoutingModels *[]BedrockFoundationModel `field:"required" json:"routingModels" yaml:"routingModels"`
}

