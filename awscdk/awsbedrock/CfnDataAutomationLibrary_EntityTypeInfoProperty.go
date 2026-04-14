package awsbedrock


// Information about an entity type in the DataAutomationLibrary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityTypeInfoProperty := &EntityTypeInfoProperty{
//   	EntityType: jsii.String("entityType"),
//
//   	// the properties below are optional
//   	EntityMetadata: jsii.String("entityMetadata"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-entitytypeinfo.html
//
type CfnDataAutomationLibrary_EntityTypeInfoProperty struct {
	// Entity types supported in DataAutomationLibraries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-entitytypeinfo.html#cfn-bedrock-dataautomationlibrary-entitytypeinfo-entitytype
	//
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// JSON string representing relevant metadata for the entity type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-entitytypeinfo.html#cfn-bedrock-dataautomationlibrary-entitytypeinfo-entitymetadata
	//
	EntityMetadata *string `field:"optional" json:"entityMetadata" yaml:"entityMetadata"`
}

