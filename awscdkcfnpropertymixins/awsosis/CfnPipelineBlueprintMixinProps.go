package awsosis


// Properties for CfnPipelineBlueprintPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPipelineBlueprintMixinProps := &CfnPipelineBlueprintMixinProps{
//   	BlueprintName: jsii.String("blueprintName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineblueprint.html
//
type CfnPipelineBlueprintMixinProps struct {
	// The name of the blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineblueprint.html#cfn-osis-pipelineblueprint-blueprintname
	//
	BlueprintName *string `field:"optional" json:"blueprintName" yaml:"blueprintName"`
}

