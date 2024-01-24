package awsevents


// Contains the GraphQL operation to be parsed and executed, if the event target is an AWS AppSync API.
//
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
	// The GraphQL operation; that is, the query, mutation, or subscription to be parsed and executed by the GraphQL service.
	//
	// For more information, see [Operations](https://docs.aws.amazon.com/appsync/latest/devguide/graphql-architecture.html#graphql-operations) in the *AWS AppSync User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-appsyncparameters.html#cfn-events-rule-appsyncparameters-graphqloperation
	//
	GraphQlOperation *string `field:"required" json:"graphQlOperation" yaml:"graphQlOperation"`
}

