package awsappsync


// Merge type used to associate the source API.
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
//   mergedApi := appsync.NewGraphqlApi(this, jsii.String("MergedAPI"), &GraphqlApiProps{
//   	Name: jsii.String("MergedAPI"),
//   	Definition: appsync.Definition_FromSourceApis(&SourceApiOptions{
//   		SourceApis: []sourceApi{
//   			&sourceApi{
//   				SourceApi: firstApi,
//   				MergeType: appsync.MergeType_MANUAL_MERGE,
//   			},
//   		},
//   	}),
//   })
//
type MergeType string

const (
	// Manual merge.
	//
	// The merge must be triggered manually when the source API has changed.
	MergeType_MANUAL_MERGE MergeType = "MANUAL_MERGE"
	// Auto merge.
	//
	// The merge is triggered automatically when the source API has changed.
	MergeType_AUTO_MERGE MergeType = "AUTO_MERGE"
)

