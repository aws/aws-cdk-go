package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceContainerConfigProperty := &InferenceContainerConfigProperty{
//   	ImageUri: jsii.String("imageUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-inferencecontainerconfig.html
//
type CfnConfiguredModelAlgorithm_InferenceContainerConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-inferencecontainerconfig.html#cfn-cleanroomsml-configuredmodelalgorithm-inferencecontainerconfig-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

