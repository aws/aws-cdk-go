package awshealthagent


// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	Name: jsii.String("name"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthagent-domain.html
//
type CfnDomainMixinProps struct {
	// The name of the Domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthagent-domain.html#cfn-healthagent-domain-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with the Domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthagent-domain.html#cfn-healthagent-domain-tags
	//
	Tags *[]*CfnDomainPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

