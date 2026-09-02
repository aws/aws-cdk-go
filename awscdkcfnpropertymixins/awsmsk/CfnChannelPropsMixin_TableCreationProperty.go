package awsmsk


// Table creation configuration of the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tableCreationProperty := &TableCreationProperty{
//   	EnableTableCreation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-tablecreation.html
//
type CfnChannelPropsMixin_TableCreationProperty struct {
	// Whether table creation is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-tablecreation.html#cfn-msk-channel-tablecreation-enabletablecreation
	//
	// Default: - true.
	//
	EnableTableCreation interface{} `field:"optional" json:"enableTableCreation" yaml:"enableTableCreation"`
}

