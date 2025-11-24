package mixinsawsecs


// Properties for CfnClusterCapacityProviderAssociationsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterCapacityProviderAssociationsMixinProps := &CfnClusterCapacityProviderAssociationsMixinProps{
//   	CapacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	Cluster: jsii.String("cluster"),
//   	DefaultCapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html
//
type CfnClusterCapacityProviderAssociationsMixinProps struct {
	// The capacity providers to associate with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
	//
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// The cluster the capacity provider association is the target of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
	//
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// The default capacity provider strategy to associate with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
	//
	DefaultCapacityProviderStrategy interface{} `field:"optional" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
}

