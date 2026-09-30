package awshealthagent


// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
type CfnDomainProps struct {
	// The name of the Domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthagent-domain.html#cfn-healthagent-domain-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags associated with the Domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthagent-domain.html#cfn-healthagent-domain-tags
	//
	Tags *[]*CfnDomain_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

