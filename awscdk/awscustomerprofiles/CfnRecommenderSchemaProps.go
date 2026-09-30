package awscustomerprofiles


// Properties for defining a `CfnRecommenderSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecommenderSchemaProps := &CfnRecommenderSchemaProps{
//   	DomainName: jsii.String("domainName"),
//   	Fields: map[string]interface{}{
//   		"fieldsKey": []interface{}{
//   			&RecommenderSchemaFieldProperty{
//   				"targetFieldName": jsii.String("targetFieldName"),
//
//   				// the properties below are optional
//   				"contentType": jsii.String("contentType"),
//   				"featureType": jsii.String("featureType"),
//   			},
//   		},
//   	},
//   	RecommenderSchemaName: jsii.String("recommenderSchemaName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html
//
type CfnRecommenderSchemaProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-recommenderschemaname
	//
	RecommenderSchemaName *string `field:"required" json:"recommenderSchemaName" yaml:"recommenderSchemaName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-tags
	//
	Tags *[]*CfnRecommenderSchema_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

