package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Source API configuration for creating a AppSync Merged API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // first source API
//   firstApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
//   	Name: jsii.String("FirstSourceAPI"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
//   })
//   firstApi.AddNoneDataSource(jsii.String("FirstSourceDS"), &DataSourceOptions{
//   	Name: cdk.Lazy_String(map[string]interface{}{
//   		(MethodDeclaration produce(): string { return 'FirstSourceDS'; }
//   				produce
//   				string
//   				{return jsii.String("FirstSourceDS")}),
//   	}),
//   })
//
//   // second source API
//   secondApi := appsync.NewGraphqlApi(this, jsii.String("SecondSourceAPI"), &GraphqlApiProps{
//   	Name: jsii.String("SecondSourceAPI"),
//   	Definition: appsync.Definition_*FromFile(path.join(__dirname, jsii.String("appsync.merged-api-2.graphql"))),
//   })
//   secondApi.AddNoneDataSource(jsii.String("SecondSourceDS"), &DataSourceOptions{
//   	Name: cdk.Lazy_*String(map[string]interface{}{
//   		(MethodDeclaration produce(): string { return 'SecondSourceDS'; }
//   				produce
//   				string
//   				{return jsii.String("SecondSourceDS")}),
//   	}),
//   })
//
//   // Merged API
//   // Merged API
//   appsync.NewGraphqlApi(this, jsii.String("MergedAPI"), &GraphqlApiProps{
//   	Name: jsii.String("MergedAPI"),
//   	Definition: appsync.Definition_FromSourceApis(&SourceApiOptions{
//   		SourceApis: []sourceApi{
//   			&sourceApi{
//   				SourceApi: firstApi,
//   				MergeType: appsync.MergeType_MANUAL_MERGE,
//   			},
//   			&sourceApi{
//   				SourceApi: secondApi,
//   				MergeType: appsync.MergeType_AUTO_MERGE,
//   			},
//   		},
//   	}),
//   })
//
type SourceApiOptions struct {
	// Definition of source APIs associated with this Merged API.
	SourceApis *[]*SourceApi `field:"required" json:"sourceApis" yaml:"sourceApis"`
	// IAM Role used to validate access to source APIs at runtime and to update the merged API endpoint with the source API changes.
	// Default: - An IAM Role with acccess to source schemas will be created.
	//
	MergedApiExecutionRole awsiam.Role `field:"optional" json:"mergedApiExecutionRole" yaml:"mergedApiExecutionRole"`
}

