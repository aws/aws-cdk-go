package awsdatabrew


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &InputProperty{
//   	DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   		TempDirectory: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	S3InputDefinition: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-input.html
//
type CfnRecipe_InputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-input.html#cfn-databrew-recipe-input-datacataloginputdefinition
	//
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-input.html#cfn-databrew-recipe-input-s3inputdefinition
	//
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

