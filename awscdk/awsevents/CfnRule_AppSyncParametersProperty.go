package awsevents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncParametersProperty := &AppSyncParametersProperty{
//   	GraphQlOperation: jsii.String("graphQlOperation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-appsyncparameters.html
//
type CfnRule_AppSyncParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-appsyncparameters.html#cfn-events-rule-appsyncparameters-graphqloperation
	//
	GraphQlOperation *string `field:"required" json:"graphQlOperation" yaml:"graphQlOperation"`
}

