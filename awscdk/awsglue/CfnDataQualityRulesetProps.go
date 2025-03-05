package awsglue


// Properties for defining a `CfnDataQualityRuleset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDataQualityRulesetProps := &CfnDataQualityRulesetProps{
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Ruleset: jsii.String("ruleset"),
//   	Tags: tags,
//   	TargetTable: &DataQualityTargetTableProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html
//
type CfnDataQualityRulesetProps struct {
	// Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// A description of the data quality ruleset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data quality ruleset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Data Quality Definition Language (DQDL) ruleset.
	//
	// For more information see the AWS Glue Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-ruleset
	//
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// A list of tags applied to the data quality ruleset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An object representing an AWS Glue table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-dataqualityruleset.html#cfn-glue-dataqualityruleset-targettable
	//
	TargetTable interface{} `field:"optional" json:"targetTable" yaml:"targetTable"`
}

