package awsappflow


// The properties that are applied when using SAPOData as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataSourcePropertiesProperty := &SAPODataSourcePropertiesProperty{
//   	ObjectPath: jsii.String("objectPath"),
//
//   	// the properties below are optional
//   	PaginationConfig: &SAPODataPaginationConfigProperty{
//   		MaxPageSize: jsii.Number(123),
//   	},
//   	ParallelismConfig: &SAPODataParallelismConfigProperty{
//   		MaxParallelism: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatasourceproperties.html
//
type CfnFlow_SAPODataSourcePropertiesProperty struct {
	// The object path specified in the SAPOData flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatasourceproperties.html#cfn-appflow-flow-sapodatasourceproperties-objectpath
	//
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// Sets the page size for each concurrent process that transfers OData records from your SAP instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatasourceproperties.html#cfn-appflow-flow-sapodatasourceproperties-paginationconfig
	//
	PaginationConfig interface{} `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// Sets the number of concurrent processes that transfers OData records from your SAP instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatasourceproperties.html#cfn-appflow-flow-sapodatasourceproperties-parallelismconfig
	//
	ParallelismConfig interface{} `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

