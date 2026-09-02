package awsmsk


// Schema evolution configuration of the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   schemaEvolutionProperty := &SchemaEvolutionProperty{
//   	EnableSchemaEvolution: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-schemaevolution.html
//
type CfnChannelPropsMixin_SchemaEvolutionProperty struct {
	// Whether schema evolution is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-schemaevolution.html#cfn-msk-channel-schemaevolution-enableschemaevolution
	//
	// Default: - false.
	//
	EnableSchemaEvolution interface{} `field:"optional" json:"enableSchemaEvolution" yaml:"enableSchemaEvolution"`
}

