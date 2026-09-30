package awsnetworkfirewall


// A key-value filter pair used in container association monitoring configurations to narrow which containers are tracked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   containerAttributeProperty := &ContainerAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containerattribute.html
//
type CfnContainerAssociationPropsMixin_ContainerAttributeProperty struct {
	// The attribute key to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containerattribute.html#cfn-networkfirewall-containerassociation-containerattribute-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The attribute value to match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-containerassociation-containerattribute.html#cfn-networkfirewall-containerassociation-containerattribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

