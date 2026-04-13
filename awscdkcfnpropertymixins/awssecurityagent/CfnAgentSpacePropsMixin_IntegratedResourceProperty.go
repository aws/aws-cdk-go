package awssecurityagent


// Integrated Resource details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   integratedResourceProperty := &IntegratedResourceProperty{
//   	Integration: jsii.String("integration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-integratedresource.html
//
type CfnAgentSpacePropsMixin_IntegratedResourceProperty struct {
	// Unique identifier of the Provider Integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-integratedresource.html#cfn-securityagent-agentspace-integratedresource-integration
	//
	Integration *string `field:"optional" json:"integration" yaml:"integration"`
}

