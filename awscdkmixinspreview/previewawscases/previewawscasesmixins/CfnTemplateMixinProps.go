package previewawscasesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplateMixinProps := &CfnTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	DomainId: jsii.String("domainId"),
//   	LayoutConfiguration: &LayoutConfigurationProperty{
//   		DefaultLayout: jsii.String("defaultLayout"),
//   	},
//   	Name: jsii.String("name"),
//   	RequiredFields: []interface{}{
//   		&RequiredFieldProperty{
//   			FieldId: jsii.String("fieldId"),
//   		},
//   	},
//   	Rules: []interface{}{
//   		&TemplateRuleProperty{
//   			CaseRuleId: jsii.String("caseRuleId"),
//   			FieldId: jsii.String("fieldId"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html
//
type CfnTemplateMixinProps struct {
	// A brief description of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Object to store configuration of layouts associated to the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-layoutconfiguration
	//
	LayoutConfiguration interface{} `field:"optional" json:"layoutConfiguration" yaml:"layoutConfiguration"`
	// The template name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of fields that must contain a value for a case to be successfully created with this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-requiredfields
	//
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
	// A list of case rules (also known as [case field conditions](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) ) on a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The status of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-template.html#cfn-cases-template-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

