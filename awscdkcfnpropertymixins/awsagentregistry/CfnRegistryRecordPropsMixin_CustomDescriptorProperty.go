package awsagentregistry


// The custom descriptor, populated when the record type is CUSTOM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customDescriptorProperty := &CustomDescriptorProperty{
//   	Data: jsii.String("data"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-customdescriptor.html
//
type CfnRegistryRecordPropsMixin_CustomDescriptorProperty struct {
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-customdescriptor.html#cfn-agentregistry-registryrecord-customdescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
}

