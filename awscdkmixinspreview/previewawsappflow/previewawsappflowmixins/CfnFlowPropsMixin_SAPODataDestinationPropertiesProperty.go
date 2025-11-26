package previewawsappflowmixins


// The properties that are applied when using SAPOData as a flow destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sAPODataDestinationPropertiesProperty := &SAPODataDestinationPropertiesProperty{
//   	ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		FailOnFirstError: jsii.Boolean(false),
//   	},
//   	IdFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	ObjectPath: jsii.String("objectPath"),
//   	SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	WriteOperationType: jsii.String("writeOperationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html
//
type CfnFlowPropsMixin_SAPODataDestinationPropertiesProperty struct {
	// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html#cfn-appflow-flow-sapodatadestinationproperties-errorhandlingconfig
	//
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// A list of field names that can be used as an ID field when performing a write operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html#cfn-appflow-flow-sapodatadestinationproperties-idfieldnames
	//
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// The object path specified in the SAPOData flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html#cfn-appflow-flow-sapodatadestinationproperties-objectpath
	//
	ObjectPath *string `field:"optional" json:"objectPath" yaml:"objectPath"`
	// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data.
	//
	// For example, this setting would determine where to write the response from a destination connector upon a successful insert operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html#cfn-appflow-flow-sapodatadestinationproperties-successresponsehandlingconfig
	//
	SuccessResponseHandlingConfig interface{} `field:"optional" json:"successResponseHandlingConfig" yaml:"successResponseHandlingConfig"`
	// The possible write operations in the destination connector.
	//
	// When this value is not provided, this defaults to the `INSERT` operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatadestinationproperties.html#cfn-appflow-flow-sapodatadestinationproperties-writeoperationtype
	//
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

