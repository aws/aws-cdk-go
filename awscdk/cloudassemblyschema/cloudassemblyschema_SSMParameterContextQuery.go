package cloudassemblyschema


// Query to SSM Parameter Context Provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSMParameterContextQuery := &sSMParameterContextQuery{
//   	account: jsii.String("account"),
//   	parameterName: jsii.String("parameterName"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
type SSMParameterContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Parameter name to query.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

