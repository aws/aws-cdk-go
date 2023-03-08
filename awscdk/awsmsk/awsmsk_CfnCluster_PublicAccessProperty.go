package awsmsk


// Specifies whether the cluster's brokers are accessible from the internet.
//
// Public access is off by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessProperty := &publicAccessProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnCluster_PublicAccessProperty struct {
	// Set to `DISABLED` to turn off public access or to `SERVICE_PROVIDED_EIPS` to turn it on.
	//
	// Public access if off by default.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

