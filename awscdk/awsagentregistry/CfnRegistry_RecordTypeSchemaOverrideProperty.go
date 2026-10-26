package awsagentregistry


// A schema override for a specific record type within a custom metadata schema configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordTypeSchemaOverrideProperty := &RecordTypeSchemaOverrideProperty{
//   	RecordType: jsii.String("recordType"),
//   	Schema: jsii.String("schema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-recordtypeschemaoverride.html
//
type CfnRegistry_RecordTypeSchemaOverrideProperty struct {
	// The record type that this schema override applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-recordtypeschemaoverride.html#cfn-agentregistry-registry-recordtypeschemaoverride-recordtype
	//
	RecordType *string `field:"required" json:"recordType" yaml:"recordType"`
	// The JSON Schema for the specified record type.
	//
	// Must follow the same structural rules as the default schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-recordtypeschemaoverride.html#cfn-agentregistry-registry-recordtypeschemaoverride-schema
	//
	Schema *string `field:"required" json:"schema" yaml:"schema"`
}

