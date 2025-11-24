package mixinsawscodebuild


// Information about the proxy rule for your reserved capacity instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetProxyRuleProperty := &FleetProxyRuleProperty{
//   	Effect: jsii.String("effect"),
//   	Entities: []*string{
//   		jsii.String("entities"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-fleetproxyrule.html
//
type CfnFleetPropsMixin_FleetProxyRuleProperty struct {
	// The behavior of the proxy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-fleetproxyrule.html#cfn-codebuild-fleet-fleetproxyrule-effect
	//
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The destination of the proxy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-fleetproxyrule.html#cfn-codebuild-fleet-fleetproxyrule-entities
	//
	Entities *[]*string `field:"optional" json:"entities" yaml:"entities"`
	// The type of proxy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-fleetproxyrule.html#cfn-codebuild-fleet-fleetproxyrule-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

