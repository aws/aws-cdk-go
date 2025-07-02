package awsbedrock


// Blueprints to apply to objects processed by the project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customOutputConfigurationProperty := &CustomOutputConfigurationProperty{
//   	Blueprints: []interface{}{
//   		&BlueprintItemProperty{
//   			BlueprintArn: jsii.String("blueprintArn"),
//
//   			// the properties below are optional
//   			BlueprintStage: jsii.String("blueprintStage"),
//   			BlueprintVersion: jsii.String("blueprintVersion"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-customoutputconfiguration.html
//
type CfnDataAutomationProject_CustomOutputConfigurationProperty struct {
	// A list of blueprints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-customoutputconfiguration.html#cfn-bedrock-dataautomationproject-customoutputconfiguration-blueprints
	//
	Blueprints interface{} `field:"optional" json:"blueprints" yaml:"blueprints"`
}

