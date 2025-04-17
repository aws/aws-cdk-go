package awsappflow


// The properties that are applied when the custom connector is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorDestinationPropertiesProperty := &CustomConnectorDestinationPropertiesProperty{
//   	EntityName: jsii.String("entityName"),
//
//   	// the properties below are optional
//   	CustomProperties: map[string]*string{
//   		"customPropertiesKey": jsii.String("customProperties"),
//   	},
//   	ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		FailOnFirstError: jsii.Boolean(false),
//   	},
//   	IdFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	WriteOperationType: jsii.String("writeOperationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html
//
type CfnFlow_CustomConnectorDestinationPropertiesProperty struct {
	// The entity specified in the custom connector as a destination in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html#cfn-appflow-flow-customconnectordestinationproperties-entityname
	//
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// The custom properties that are specific to the connector when it's used as a destination in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html#cfn-appflow-flow-customconnectordestinationproperties-customproperties
	//
	CustomProperties interface{} `field:"optional" json:"customProperties" yaml:"customProperties"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html#cfn-appflow-flow-customconnectordestinationproperties-errorhandlingconfig
	//
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update, delete, or upsert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html#cfn-appflow-flow-customconnectordestinationproperties-idfieldnames
	//
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Specifies the type of write operation to be performed in the custom connector when it's used as destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-customconnectordestinationproperties.html#cfn-appflow-flow-customconnectordestinationproperties-writeoperationtype
	//
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

