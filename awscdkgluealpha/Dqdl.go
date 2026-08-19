package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Data Quality Definition Language (DQDL) document for a `DataQualityRuleset`.
//
// DQDL is an authored string that Glue parses and validates at deploy time. Build
// one from a raw DQDL string with {@link Dqdl.fromString}.
//
// Example:
//   glue.NewDataQualityRuleset(this, jsii.String("MyRuleset"), &DataQualityRulesetProps{
//   	RulesetName: jsii.String("my_ruleset"),
//   	Dqdl: glue.Dqdl_FromString(jsii.String("Rules = [ RowCount > 100, IsComplete \"order_id\" ]")),
//   	TargetTable: glue.NewDataQualityTargetTable(jsii.String("my_database"), jsii.String("my_table")),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/dg/dqdl.html
//
// Experimental.
type Dqdl interface {
}

// The jsii proxy struct for Dqdl
type jsiiProxy_Dqdl struct {
	_ byte // padding
}

// Create a `Dqdl` from a raw DQDL string.
// Experimental.
func Dqdl_FromString(dqdl *string) Dqdl {
	_init_.Initialize()

	if err := validateDqdl_FromStringParameters(dqdl); err != nil {
		panic(err)
	}
	var returns Dqdl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Dqdl",
		"fromString",
		[]interface{}{dqdl},
		&returns,
	)

	return returns
}

