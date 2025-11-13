package interfacesawsbedrock


// A reference to a FlowAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowAliasReference := &FlowAliasReference{
//   	FlowAliasArn: jsii.String("flowAliasArn"),
//   	FlowArn: jsii.String("flowArn"),
//   }
//
type FlowAliasReference struct {
	// The Arn of the FlowAlias resource.
	FlowAliasArn *string `field:"required" json:"flowAliasArn" yaml:"flowAliasArn"`
	// The FlowArn of the FlowAlias resource.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
}

