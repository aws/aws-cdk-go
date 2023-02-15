package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSuppressionProperty := &availSuppressionProperty{
//   	mode: jsii.String("mode"),
//   	value: jsii.String("value"),
//   }
//
type CfnPlaybackConfiguration_AvailSuppressionProperty struct {
	// `CfnPlaybackConfiguration.AvailSuppressionProperty.Mode`.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// `CfnPlaybackConfiguration.AvailSuppressionProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

