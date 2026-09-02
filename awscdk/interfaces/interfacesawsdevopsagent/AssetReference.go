package interfacesawsdevopsagent


// A reference to a Asset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetReference := &AssetReference{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	AssetArn: jsii.String("assetArn"),
//   	AssetId: jsii.String("assetId"),
//   }
//
type AssetReference struct {
	// The AgentSpaceId of the Asset resource.
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The ARN of the Asset resource.
	AssetArn *string `field:"required" json:"assetArn" yaml:"assetArn"`
	// The AssetId of the Asset resource.
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
}

