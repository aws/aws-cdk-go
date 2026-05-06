package awscustomerprofiles


// Represents a field in a DomainObjectType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   domainObjectTypeFieldProperty := &DomainObjectTypeFieldProperty{
//   	ContentType: jsii.String("contentType"),
//   	FeatureType: jsii.String("featureType"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domainobjecttype-domainobjecttypefield.html
//
type CfnDomainObjectTypePropsMixin_DomainObjectTypeFieldProperty struct {
	// The content type of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domainobjecttype-domainobjecttypefield.html#cfn-customerprofiles-domainobjecttype-domainobjecttypefield-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The feature type of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domainobjecttype-domainobjecttypefield.html#cfn-customerprofiles-domainobjecttype-domainobjecttypefield-featuretype
	//
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
	// The source field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domainobjecttype-domainobjecttypefield.html#cfn-customerprofiles-domainobjecttype-domainobjecttypefield-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The target field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domainobjecttype-domainobjecttypefield.html#cfn-customerprofiles-domainobjecttype-domainobjecttypefield-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

