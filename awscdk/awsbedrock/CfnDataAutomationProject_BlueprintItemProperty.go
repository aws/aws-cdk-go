package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueprintItemProperty := &BlueprintItemProperty{
//   	BlueprintArn: jsii.String("blueprintArn"),
//
//   	// the properties below are optional
//   	BlueprintStage: jsii.String("blueprintStage"),
//   	BlueprintVersion: jsii.String("blueprintVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-blueprintitem.html
//
type CfnDataAutomationProject_BlueprintItemProperty struct {
	// ARN of a Blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-blueprintitem.html#cfn-bedrock-dataautomationproject-blueprintitem-blueprintarn
	//
	BlueprintArn *string `field:"required" json:"blueprintArn" yaml:"blueprintArn"`
	// Stage of the Blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-blueprintitem.html#cfn-bedrock-dataautomationproject-blueprintitem-blueprintstage
	//
	BlueprintStage *string `field:"optional" json:"blueprintStage" yaml:"blueprintStage"`
	// Blueprint Version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-blueprintitem.html#cfn-bedrock-dataautomationproject-blueprintitem-blueprintversion
	//
	BlueprintVersion *string `field:"optional" json:"blueprintVersion" yaml:"blueprintVersion"`
}

