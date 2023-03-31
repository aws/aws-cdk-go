package awsecs


// Properties for defining a `CfnClusterCapacityProviderAssociations`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterCapacityProviderAssociationsProps := &cfnClusterCapacityProviderAssociationsProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	cluster: jsii.String("cluster"),
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyProperty{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnClusterCapacityProviderAssociationsProps struct {
	// The capacity providers to associate with the cluster.
	CapacityProviders *[]*string `field:"required" json:"capacityProviders" yaml:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	DefaultCapacityProviderStrategy interface{} `field:"required" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
}

