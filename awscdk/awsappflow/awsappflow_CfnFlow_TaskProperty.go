package awsappflow


// A class for modeling different type of tasks.
//
// Task implementation varies based on the `TaskType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskProperty := &taskProperty{
//   	sourceFields: []*string{
//   		jsii.String("sourceFields"),
//   	},
//   	taskType: jsii.String("taskType"),
//
//   	// the properties below are optional
//   	connectorOperator: &connectorOperatorProperty{
//   		amplitude: jsii.String("amplitude"),
//   		customConnector: jsii.String("customConnector"),
//   		datadog: jsii.String("datadog"),
//   		dynatrace: jsii.String("dynatrace"),
//   		googleAnalytics: jsii.String("googleAnalytics"),
//   		inforNexus: jsii.String("inforNexus"),
//   		marketo: jsii.String("marketo"),
//   		s3: jsii.String("s3"),
//   		salesforce: jsii.String("salesforce"),
//   		sapoData: jsii.String("sapoData"),
//   		serviceNow: jsii.String("serviceNow"),
//   		singular: jsii.String("singular"),
//   		slack: jsii.String("slack"),
//   		trendmicro: jsii.String("trendmicro"),
//   		veeva: jsii.String("veeva"),
//   		zendesk: jsii.String("zendesk"),
//   	},
//   	destinationField: jsii.String("destinationField"),
//   	taskProperties: []interface{}{
//   		&taskPropertiesObjectProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFlow_TaskProperty struct {
	// The source fields to which a particular task is applied.
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Specifies the particular task implementation that Amazon AppFlow performs.
	//
	// *Allowed values* : `Arithmetic` | `Filter` | `Map` | `Map_all` | `Mask` | `Merge` | `Truncate` | `Validate`.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The operation to be performed on the provided source fields.
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// A field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// A map used to store task-related information.
	//
	// The execution service looks for particular information based on the `TaskType` .
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

