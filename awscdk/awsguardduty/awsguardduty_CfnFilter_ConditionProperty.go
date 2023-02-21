package awsguardduty


// Specifies the condition to apply to a single field when filtering through  findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &conditionProperty{
//   	eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	equalTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	greaterThan: jsii.Number(123),
//   	greaterThanOrEqual: jsii.Number(123),
//   	gt: jsii.Number(123),
//   	gte: jsii.Number(123),
//   	lessThan: jsii.Number(123),
//   	lessThanOrEqual: jsii.Number(123),
//   	lt: jsii.Number(123),
//   	lte: jsii.Number(123),
//   	neq: []*string{
//   		jsii.String("neq"),
//   	},
//   	notEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   }
//
type CfnFilter_ConditionProperty struct {
	// Represents the equal condition to apply to a single field when querying for findings.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Represents an *equal* ** condition to be applied to a single field when querying for findings.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Represents a *greater than* condition to be applied to a single field when querying for findings.
	GreaterThan *float64 `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// Represents a *greater than or equal* condition to be applied to a single field when querying for findings.
	GreaterThanOrEqual *float64 `field:"optional" json:"greaterThanOrEqual" yaml:"greaterThanOrEqual"`
	// Represents a *greater than* condition to be applied to a single field when querying for findings.
	Gt *float64 `field:"optional" json:"gt" yaml:"gt"`
	// Represents the greater than or equal condition to apply to a single field when querying for findings.
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// Represents a *less than* condition to be applied to a single field when querying for findings.
	LessThan *float64 `field:"optional" json:"lessThan" yaml:"lessThan"`
	// Represents a *less than or equal* condition to be applied to a single field when querying for findings.
	LessThanOrEqual *float64 `field:"optional" json:"lessThanOrEqual" yaml:"lessThanOrEqual"`
	// Represents the less than condition to apply to a single field when querying for findings.
	Lt *float64 `field:"optional" json:"lt" yaml:"lt"`
	// Represents the less than or equal condition to apply to a single field when querying for findings.
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
	// Represents the not equal condition to apply to a single field when querying for findings.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
	// Represents a *not equal* ** condition to be applied to a single field when querying for findings.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
}

