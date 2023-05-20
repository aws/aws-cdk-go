package awsmsk


// Broker access controls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessProperty := &PublicAccessProperty{
//   	Type: jsii.String("type"),
//   }
//
type CfnCluster_PublicAccessProperty struct {
	// DISABLED means that public access is turned off.
	//
	// SERVICE_PROVIDED_EIPS means that public access is turned on.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

