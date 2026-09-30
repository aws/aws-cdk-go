package awscustomerprofiles


// Properties for CfnRecommenderSchemaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRecommenderSchemaMixinProps := &CfnRecommenderSchemaMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	Fields: map[string]interface{}{
//   		"fieldsKey": []interface{}{
//   			&RecommenderSchemaFieldProperty{
//   				"contentType": jsii.String("contentType"),
//   				"featureType": jsii.String("featureType"),
//   				"targetFieldName": jsii.String("targetFieldName"),
//   			},
//   		},
//   	},
//   	RecommenderSchemaName: jsii.String("recommenderSchemaName"),
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
type CfnRecommenderSchemaMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-recommenderschemaname
	//
	RecommenderSchemaName *string `field:"optional" json:"recommenderSchemaName" yaml:"recommenderSchemaName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-recommenderschema.html#cfn-customerprofiles-recommenderschema-tags
	//
	Tags *[]*CfnRecommenderSchemaPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

