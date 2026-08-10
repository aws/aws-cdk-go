package awsosis


// Properties for defining a `CfnPipelineBlueprint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineBlueprintProps := &CfnPipelineBlueprintProps{
//   	BlueprintName: jsii.String("blueprintName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineblueprint.html
//
type CfnPipelineBlueprintProps struct {
	// The name of the blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineblueprint.html#cfn-osis-pipelineblueprint-blueprintname
	//
	BlueprintName *string `field:"optional" json:"blueprintName" yaml:"blueprintName"`
}

