package awslakeformation


// Properties for defining a `CfnTag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagProps := &cfnTagProps{
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//
//   	// the properties below are optional
//   	catalogId: jsii.String("catalogId"),
//   }
//
type CfnTagProps struct {
	// `AWS::LakeFormation::Tag.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `AWS::LakeFormation::Tag.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
	// `AWS::LakeFormation::Tag.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

