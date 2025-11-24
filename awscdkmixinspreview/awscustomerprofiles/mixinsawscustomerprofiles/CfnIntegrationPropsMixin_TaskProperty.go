package mixinsawscustomerprofiles


// The `Task` property type specifies the class for modeling different type of tasks.
//
// Task implementation varies based on the TaskType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskProperty := &TaskProperty{
//   	ConnectorOperator: &ConnectorOperatorProperty{
//   		Marketo: jsii.String("marketo"),
//   		S3: jsii.String("s3"),
//   		Salesforce: jsii.String("salesforce"),
//   		ServiceNow: jsii.String("serviceNow"),
//   		Zendesk: jsii.String("zendesk"),
//   	},
//   	DestinationField: jsii.String("destinationField"),
//   	SourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	TaskProperties: []interface{}{
//   		&TaskPropertiesMapProperty{
//   			OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   			Property: jsii.String("property"),
//   		},
//   	},
//   	TaskType: jsii.String("taskType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html
//
type CfnIntegrationPropsMixin_TaskProperty struct {
	// The operation to be performed on the provided source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html#cfn-customerprofiles-integration-task-connectoroperator
	//
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html#cfn-customerprofiles-integration-task-destinationfield
	//
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// The source fields to which a particular task is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html#cfn-customerprofiles-integration-task-sourcefields
	//
	SourceFields *[]*string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
	// A map used to store task-related information.
	//
	// The service looks for particular information based on the TaskType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html#cfn-customerprofiles-integration-task-taskproperties
	//
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-task.html#cfn-customerprofiles-integration-task-tasktype
	//
	TaskType *string `field:"optional" json:"taskType" yaml:"taskType"`
}

