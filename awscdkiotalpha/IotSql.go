package awscdkiotalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines AWS IoT SQL.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	Actions: []IAction{
//   		actions.NewSnsTopicAction(topic, &SnsTopicActionProps{
//   			MessageFormat: actions.SnsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type IotSql interface {
	// Returns the IoT SQL configuration.
	// Experimental.
	Bind(scope constructs.Construct) *IotSqlConfig
}

// The jsii proxy struct for IotSql
type jsiiProxy_IotSql struct {
	_ byte // padding
}

// Experimental.
func NewIotSql_Override(i IotSql) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.IotSql",
		nil, // no parameters
		i,
	)
}

// Uses the original SQL version built on 2015-10-08.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVer20151008(sql *string) IotSql {
	_init_.Initialize()

	if err := validateIotSql_FromStringAsVer20151008Parameters(sql); err != nil {
		panic(err)
	}
	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20151008",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the SQL version built on 2016-03-23.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVer20160323(sql *string) IotSql {
	_init_.Initialize()

	if err := validateIotSql_FromStringAsVer20160323Parameters(sql); err != nil {
		panic(err)
	}
	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20160323",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the most recent beta SQL version.
//
// If you use this version, it might
// introduce breaking changes to your rules.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVerNewestUnstable(sql *string) IotSql {
	_init_.Initialize()

	if err := validateIotSql_FromStringAsVerNewestUnstableParameters(sql); err != nil {
		panic(err)
	}
	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVerNewestUnstable",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSql) Bind(scope constructs.Construct) *IotSqlConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *IotSqlConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

