package awsdatazone


// Properties for defining a `CfnFormType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFormTypeProps := &CfnFormTypeProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Model: &ModelProperty{
//   		Smithy: jsii.String("smithy"),
//   	},
//   	Name: jsii.String("name"),
//   	OwningProjectIdentifier: jsii.String("owningProjectIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html
//
type CfnFormTypeProps struct {
	// The identifier of the Amazon DataZone domain in which the form type exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The model of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-model
	//
	Model interface{} `field:"required" json:"model" yaml:"model"`
	// The name of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the project that owns the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-owningprojectidentifier
	//
	OwningProjectIdentifier *string `field:"required" json:"owningProjectIdentifier" yaml:"owningProjectIdentifier"`
	// The description of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The status of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

