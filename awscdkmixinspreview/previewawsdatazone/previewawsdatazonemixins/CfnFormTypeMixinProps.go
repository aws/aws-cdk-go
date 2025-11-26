package previewawsdatazonemixins


// Properties for CfnFormTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFormTypeMixinProps := &CfnFormTypeMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Model: &ModelProperty{
//   		Smithy: jsii.String("smithy"),
//   	},
//   	Name: jsii.String("name"),
//   	OwningProjectIdentifier: jsii.String("owningProjectIdentifier"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html
//
type CfnFormTypeMixinProps struct {
	// The description of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the Amazon DataZone domain in which the form type exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The model of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-model
	//
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// The name of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The identifier of the project that owns the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-owningprojectidentifier
	//
	OwningProjectIdentifier *string `field:"optional" json:"owningProjectIdentifier" yaml:"owningProjectIdentifier"`
	// The status of the form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html#cfn-datazone-formtype-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

