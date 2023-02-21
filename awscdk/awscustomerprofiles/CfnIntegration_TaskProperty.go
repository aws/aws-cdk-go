package awscustomerprofiles


// The `Task` property type specifies the class for modeling different type of tasks.
//
// Task implementation varies based on the TaskType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskProperty := &TaskProperty{
//   	SourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	TaskType: jsii.String("taskType"),
//
//   	// the properties below are optional
//   	ConnectorOperator: &ConnectorOperatorProperty{
//   		Marketo: jsii.String("marketo"),
//   		S3: jsii.String("s3"),
//   		Salesforce: jsii.String("salesforce"),
//   		ServiceNow: jsii.String("serviceNow"),
//   		Zendesk: jsii.String("zendesk"),
//   	},
//   	DestinationField: jsii.String("destinationField"),
//   	TaskProperties: []interface{}{
//   		&TaskPropertiesMapProperty{
//   			OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   			Property: jsii.String("property"),
//   		},
//   	},
//   }
//
type CfnIntegration_TaskProperty struct {
	// The source fields to which a particular task is applied.
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The operation to be performed on the provided source fields.
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// A map used to store task-related information.
	//
	// The service looks for particular information based on the TaskType.
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

