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
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html#aws-glue-api-regex-oneLine) .
	//
	// The key-name for the LF-tag.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// An array of UTF-8 strings, not less than 1 or more than 50 strings.
	//
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
	// Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html#aws-glue-api-regex-oneLine) .
	//
	// The identifier for the Data Catalog . By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

