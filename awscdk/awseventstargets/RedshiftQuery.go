package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Schedule an Amazon Redshift Query to be run, using the Redshift Data API.
//
// If you would like Amazon Redshift to identify the Event Bridge rule, and present it in the Amazon Redshift console, append a `QS2-` prefix to both `statementName` and `ruleName`.
//
// Example:
//   import redshiftserverless "github.com/aws/aws-cdk-go/awscdk"
//
//   var workgroup CfnWorkgroup
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.AddTarget(targets.NewRedshiftQuery(workgroup.attrWorkgroupWorkgroupArn, &RedshiftQueryProps{
//   	Database: jsii.String("dev"),
//   	DeadLetterQueue: dlq,
//   	Sql: []*string{
//   		jsii.String("SELECT * FROM foo"),
//   		jsii.String("SELECT * FROM baz"),
//   	},
//   }))
//
type RedshiftQuery interface {
	awsevents.IRuleTarget
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for RedshiftQuery
type jsiiProxy_RedshiftQuery struct {
	internal.Type__awseventsIRuleTarget
}

func NewRedshiftQuery(clusterArn *string, props *RedshiftQueryProps) RedshiftQuery {
	_init_.Initialize()

	if err := validateNewRedshiftQueryParameters(clusterArn, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftQuery{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.RedshiftQuery",
		[]interface{}{clusterArn, props},
		&j,
	)

	return &j
}

func NewRedshiftQuery_Override(r RedshiftQuery, clusterArn *string, props *RedshiftQueryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.RedshiftQuery",
		[]interface{}{clusterArn, props},
		r,
	)
}

func (r *jsiiProxy_RedshiftQuery) Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig {
	if err := r.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

