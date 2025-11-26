package previewawslightsailmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	DomainEntries: []interface{}{
//   		&DomainEntryProperty{
//   			Id: jsii.String("id"),
//   			IsAlias: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Target: jsii.String("target"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	DomainName: jsii.String("domainName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-domain.html
//
type CfnDomainMixinProps struct {
	// An array of key-value pairs containing information about the domain entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-domain.html#cfn-lightsail-domain-domainentries
	//
	DomainEntries interface{} `field:"optional" json:"domainEntries" yaml:"domainEntries"`
	// The fully qualified domain name in the certificate request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-domain.html#cfn-lightsail-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The tag keys and optional values for the resource.
	//
	// For more information about tags in Lightsail, see the [Amazon Lightsail Developer Guide](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-tags) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-domain.html#cfn-lightsail-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

