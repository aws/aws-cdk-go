package awslambda


// Document db event source config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentDBEventSourceConfigProperty := &DocumentDBEventSourceConfigProperty{
//   	CollectionName: jsii.String("collectionName"),
//   	DatabaseName: jsii.String("databaseName"),
//   	FullDocument: jsii.String("fullDocument"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-documentdbeventsourceconfig.html
//
type CfnEventSourceMapping_DocumentDBEventSourceConfigProperty struct {
	// The collection name to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-documentdbeventsourceconfig.html#cfn-lambda-eventsourcemapping-documentdbeventsourceconfig-collectionname
	//
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// The database name to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-documentdbeventsourceconfig.html#cfn-lambda-eventsourcemapping-documentdbeventsourceconfig-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Include full document in change stream response.
	//
	// The default option will only send the changes made to documents to Lambda. If you want the complete document sent to Lambda, set this to UpdateLookup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-documentdbeventsourceconfig.html#cfn-lambda-eventsourcemapping-documentdbeventsourceconfig-fulldocument
	//
	FullDocument *string `field:"optional" json:"fullDocument" yaml:"fullDocument"`
}

