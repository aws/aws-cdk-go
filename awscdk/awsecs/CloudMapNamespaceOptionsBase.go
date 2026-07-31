package awsecs


// Options shared by all ways of setting the default AWS Cloud Map namespace of a cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudMapNamespaceOptionsBase := &CloudMapNamespaceOptionsBase{
//   	UseForServiceConnect: jsii.Boolean(false),
//   }
//
type CloudMapNamespaceOptionsBase struct {
	// This property specifies whether to set the provided namespace as the service connect default in the cluster properties.
	// Default: false.
	//
	UseForServiceConnect *bool `field:"optional" json:"useForServiceConnect" yaml:"useForServiceConnect"`
}

