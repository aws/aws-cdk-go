package previewawsappflowmixins


// A class for modeling different type of tasks.
//
// Task implementation varies based on the `TaskType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskProperty := &TaskProperty{
//   	ConnectorOperator: &ConnectorOperatorProperty{
//   		Amplitude: jsii.String("amplitude"),
//   		CustomConnector: jsii.String("customConnector"),
//   		Datadog: jsii.String("datadog"),
//   		Dynatrace: jsii.String("dynatrace"),
//   		GoogleAnalytics: jsii.String("googleAnalytics"),
//   		InforNexus: jsii.String("inforNexus"),
//   		Marketo: jsii.String("marketo"),
//   		Pardot: jsii.String("pardot"),
//   		S3: jsii.String("s3"),
//   		Salesforce: jsii.String("salesforce"),
//   		SapoData: jsii.String("sapoData"),
//   		ServiceNow: jsii.String("serviceNow"),
//   		Singular: jsii.String("singular"),
//   		Slack: jsii.String("slack"),
//   		Trendmicro: jsii.String("trendmicro"),
//   		Veeva: jsii.String("veeva"),
//   		Zendesk: jsii.String("zendesk"),
//   	},
//   	DestinationField: jsii.String("destinationField"),
//   	SourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	TaskProperties: []interface{}{
//   		&TaskPropertiesObjectProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskType: jsii.String("taskType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html
//
type CfnFlowPropsMixin_TaskProperty struct {
	// The operation to be performed on the provided source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-connectoroperator
	//
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-destinationfield
	//
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// The source fields to which a particular task is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-sourcefields
	//
	SourceFields *[]*string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
	// A map used to store task-related information.
	//
	// The execution service looks for particular information based on the `TaskType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-taskproperties
	//
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	//
	// *Allowed values* : `Arithmetic` | `Filter` | `Map` | `Map_all` | `Mask` | `Merge` | `Truncate` | `Validate`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-task.html#cfn-appflow-flow-task-tasktype
	//
	TaskType *string `field:"optional" json:"taskType" yaml:"taskType"`
}

