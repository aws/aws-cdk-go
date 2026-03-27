package awsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentStrategyOptionsProperty := &DeploymentStrategyOptionsProperty{
//   	DeploymentStrategy: jsii.String("deploymentStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-deploymentstrategyoptions.html
//
type CfnDomain_DeploymentStrategyOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-deploymentstrategyoptions.html#cfn-opensearchservice-domain-deploymentstrategyoptions-deploymentstrategy
	//
	DeploymentStrategy *string `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
}

