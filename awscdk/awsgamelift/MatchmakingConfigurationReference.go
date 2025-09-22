package awsgamelift


// A reference to a MatchmakingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchmakingConfigurationReference := &MatchmakingConfigurationReference{
//   	MatchmakingConfigurationArn: jsii.String("matchmakingConfigurationArn"),
//   	MatchmakingConfigurationName: jsii.String("matchmakingConfigurationName"),
//   }
//
type MatchmakingConfigurationReference struct {
	// The ARN of the MatchmakingConfiguration resource.
	MatchmakingConfigurationArn *string `field:"required" json:"matchmakingConfigurationArn" yaml:"matchmakingConfigurationArn"`
	// The Name of the MatchmakingConfiguration resource.
	MatchmakingConfigurationName *string `field:"required" json:"matchmakingConfigurationName" yaml:"matchmakingConfigurationName"`
}

