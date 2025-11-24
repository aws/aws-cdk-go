package mixinsawscodebuild


// Information about the proxy configurations that apply network access control to your reserved capacity instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   proxyConfigurationProperty := &ProxyConfigurationProperty{
//   	DefaultBehavior: jsii.String("defaultBehavior"),
//   	OrderedProxyRules: []interface{}{
//   		&FleetProxyRuleProperty{
//   			Effect: jsii.String("effect"),
//   			Entities: []*string{
//   				jsii.String("entities"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-proxyconfiguration.html
//
type CfnFleetPropsMixin_ProxyConfigurationProperty struct {
	// The default behavior of outgoing traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-proxyconfiguration.html#cfn-codebuild-fleet-proxyconfiguration-defaultbehavior
	//
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// An array of `FleetProxyRule` objects that represent the specified destination domains or IPs to allow or deny network access control to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-proxyconfiguration.html#cfn-codebuild-fleet-proxyconfiguration-orderedproxyrules
	//
	OrderedProxyRules interface{} `field:"optional" json:"orderedProxyRules" yaml:"orderedProxyRules"`
}

