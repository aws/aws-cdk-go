package awscases

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnField`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFieldProps := &CfnFieldProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DomainId: jsii.String("domainId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html
//
type CfnFieldProps struct {
	// Name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html#cfn-cases-field-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html#cfn-cases-field-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Description of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html#cfn-cases-field-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html#cfn-cases-field-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-field.html#cfn-cases-field-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

