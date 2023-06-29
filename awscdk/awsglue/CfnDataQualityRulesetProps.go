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
type CfnDataQualityRulesetProps struct {
	// Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// A description of the data quality ruleset.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data quality ruleset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Data Quality Definition Language (DQDL) ruleset.
	//
	// For more information see the AWS Glue Developer Guide.
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// A list of tags applied to the data quality ruleset.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An object representing an AWS Glue table.
	TargetTable interface{} `field:"optional" json:"targetTable" yaml:"targetTable"`
}

