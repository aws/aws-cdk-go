package awsguardduty


// Specifies the condition to apply to a single field when filtering through GuardDuty findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &ConditionProperty{
//   	Eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	EqualTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	GreaterThan: jsii.Number(123),
//   	GreaterThanOrEqual: jsii.Number(123),
//   	Gt: jsii.Number(123),
//   	Gte: jsii.Number(123),
//   	LessThan: jsii.Number(123),
//   	LessThanOrEqual: jsii.Number(123),
//   	Lt: jsii.Number(123),
//   	Lte: jsii.Number(123),
//   	Neq: []*string{
//   		jsii.String("neq"),
//   	},
//   	NotEquals: []*string{
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

