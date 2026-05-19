package awsbedrockagentcore


// Attributes for importing an existing Gateway Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gateway Gateway
//
//   gatewayTargetAttributes := &GatewayTargetAttributes{
//   	Gateway: gateway,
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   	TargetArn: jsii.String("targetArn"),
//   	TargetId: jsii.String("targetId"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Status: jsii.String("status"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
type GatewayTargetAttributes struct {
	// The gateway this target belongs to.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The name of the gateway target.
	GatewayTargetName *string `field:"required" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The ARN of the gateway target.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The ID of the gateway target.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Optional creation timestamp.
	// Default: - No creation timestamp.
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Optional status of the target.
	// Default: - No status.
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Optional last update timestamp.
	// Default: - No update timestamp.
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

