package awsagentregistry


// Configuration that defines a typed metadata schema for a registry.
//
// A default schema and per-record-type schema overrides are both optional. Schema evolution is additive only: you can add fields and enum values, but you cannot remove fields, change types, or remove enum values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customMetadataSchemaConfigurationProperty := &CustomMetadataSchemaConfigurationProperty{
//   	DefaultSchema: jsii.String("defaultSchema"),
//   	RecordTypeSchemaOverrides: []interface{}{
//   		&RecordTypeSchemaOverrideProperty{
//   			RecordType: jsii.String("recordType"),
//   			Schema: jsii.String("schema"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-custommetadataschemaconfiguration.html
//
type CfnRegistry_CustomMetadataSchemaConfigurationProperty struct {
	// The default JSON Schema that applies to record types without a specific override.
	//
	// Supported property types are string, string with an enum constraint, string with a uri format, and boolean.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-custommetadataschemaconfiguration.html#cfn-agentregistry-registry-custommetadataschemaconfiguration-defaultschema
	//
	DefaultSchema *string `field:"optional" json:"defaultSchema" yaml:"defaultSchema"`
	// A list of per-record-type schema overrides.
	//
	// When a record's type matches an override, that override's schema is used instead of the default schema for validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-custommetadataschemaconfiguration.html#cfn-agentregistry-registry-custommetadataschemaconfiguration-recordtypeschemaoverrides
	//
	RecordTypeSchemaOverrides interface{} `field:"optional" json:"recordTypeSchemaOverrides" yaml:"recordTypeSchemaOverrides"`
}

