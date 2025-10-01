package awscdk


// A reference to a HookTypeConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   hookTypeConfigReference := &HookTypeConfigReference{
//   	ConfigurationArn: jsii.String("configurationArn"),
//   }
//
type HookTypeConfigReference struct {
	// The ConfigurationArn of the HookTypeConfig resource.
	ConfigurationArn *string `field:"required" json:"configurationArn" yaml:"configurationArn"`
}

