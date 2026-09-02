package awscases

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCaseMixinProps := &CfnCaseMixinProps{
//   	CustomerId: jsii.String("customerId"),
//   	DomainId: jsii.String("domainId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateId: jsii.String("templateId"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html
//
type CfnCaseMixinProps struct {
	// The full customer profile ARN for the case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html#cfn-cases-case-customerid
	//
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html#cfn-cases-case-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// A list of tags for the case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html#cfn-cases-case-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A unique identifier of a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html#cfn-cases-case-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// The title of the case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-case.html#cfn-cases-case-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

