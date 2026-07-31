package awsmediaconnectalpha


// Attributes for importing an existing Router Output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerOutputAttributes := &RouterOutputAttributes{
//   	RouterOutputArn: jsii.String("routerOutputArn"),
//
//   	// the properties below are optional
//   	RouterOutputId: jsii.String("routerOutputId"),
//   	RouterOutputName: jsii.String("routerOutputName"),
//   }
//
// Experimental.
type RouterOutputAttributes struct {
	// The ARN of the router output.
	// Experimental.
	RouterOutputArn *string `field:"required" json:"routerOutputArn" yaml:"routerOutputArn"`
	// The unique identifier of the router output.
	// Default: - parsed from the ARN when available.
	//
	// Experimental.
	RouterOutputId *string `field:"optional" json:"routerOutputId" yaml:"routerOutputId"`
	// The name of the router output.
	//
	// The name is not encoded in the ARN (the ARN contains only the service-generated ID),
	// so provide it here if you need to access `routerOutputName` on the imported construct.
	// Default: - routerOutputName is not available on the imported construct.
	//
	// Experimental.
	RouterOutputName *string `field:"optional" json:"routerOutputName" yaml:"routerOutputName"`
}

