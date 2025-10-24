package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to write data to an Amazon OpenSearch Service domain.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//   var domain Domain
//
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   })
//
//   topicRule.AddAction(actions.NewOpenSearchAction(domain, &OpenSearchActionProps{
//   	Id: jsii.String("my-id"),
//   	Index: jsii.String("my-index"),
//   	Type: jsii.String("my-type"),
//   }))
//
// Experimental.
type OpenSearchAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for OpenSearchAction
type jsiiProxy_OpenSearchAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewOpenSearchAction(domain awsopensearchservice.Domain, props *OpenSearchActionProps) OpenSearchAction {
	_init_.Initialize()

	if err := validateNewOpenSearchActionParameters(domain, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.OpenSearchAction",
		[]interface{}{domain, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchAction_Override(o OpenSearchAction, domain awsopensearchservice.Domain, props *OpenSearchActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.OpenSearchAction",
		[]interface{}{domain, props},
		o,
	)
}

