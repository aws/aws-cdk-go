package awsmsk


// Details for IAM access control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamProperty := &iamProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_IamProperty struct {
	// Whether IAM access control is enabled.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

