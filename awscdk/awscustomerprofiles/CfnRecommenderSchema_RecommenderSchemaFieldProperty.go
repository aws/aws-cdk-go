package awscustomerprofiles


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recommenderSchemaFieldProperty := &RecommenderSchemaFieldProperty{
//   	TargetFieldName: jsii.String("targetFieldName"),
//
//   	// the properties below are optional
//   	ContentType: jsii.String("contentType"),
//   	FeatureType: jsii.String("featureType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommenderschema-recommenderschemafield.html
//
type CfnRecommenderSchema_RecommenderSchemaFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommenderschema-recommenderschemafield.html#cfn-customerprofiles-recommenderschema-recommenderschemafield-targetfieldname
	//
	TargetFieldName *string `field:"required" json:"targetFieldName" yaml:"targetFieldName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommenderschema-recommenderschemafield.html#cfn-customerprofiles-recommenderschema-recommenderschemafield-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-recommenderschema-recommenderschemafield.html#cfn-customerprofiles-recommenderschema-recommenderschemafield-featuretype
	//
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

