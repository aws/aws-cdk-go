package awscloudtrail


// You can configure the `DataResource` in an `EventSelector` to log data events for the following three resource types:.
//
// - `AWS::DynamoDB::Table`
// - `AWS::Lambda::Function`
// - `AWS::S3::Object`
//
// To log data events for all other resource types including objects stored in [directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html) , you must use [AdvancedEventSelectors](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html) . You must also use `AdvancedEventSelectors` if you want to filter on the `eventName` field.
//
// Configure the `DataResource` to specify the resource type and resource ARNs for which you want to log data events.
//
// > The total number of allowed data resources is 250. This number can be distributed between 1 and 5 event selectors, but the total cannot exceed 250 across all selectors for the trail.
//
// The following example demonstrates how logging works when you configure logging of all data events for a general purpose bucket named `amzn-s3-demo-bucket1` . In this example, the CloudTrail user specified an empty prefix, and the option to log both `Read` and `Write` data events.
//
// - A user uploads an image file to `amzn-s3-demo-bucket1` .
// - The `PutObject` API operation is an Amazon S3 object-level API. It is recorded as a data event in CloudTrail. Because the CloudTrail user specified an S3 bucket with an empty prefix, events that occur on any object in that bucket are logged. The trail processes and logs the event.
// - A user uploads an object to an Amazon S3 bucket named `arn:aws:s3:::amzn-s3-demo-bucket1` .
// - The `PutObject` API operation occurred for an object in an S3 bucket that the CloudTrail user didn't specify for the trail. The trail doesn’t log the event.
//
// The following example demonstrates how logging works when you configure logging of AWS Lambda data events for a Lambda function named *MyLambdaFunction* , but not for all Lambda functions.
//
// - A user runs a script that includes a call to the *MyLambdaFunction* function and the *MyOtherLambdaFunction* function.
// - The `Invoke` API operation on *MyLambdaFunction* is an Lambda API. It is recorded as a data event in CloudTrail. Because the CloudTrail user specified logging data events for *MyLambdaFunction* , any invocations of that function are logged. The trail processes and logs the event.
// - The `Invoke` API operation on *MyOtherLambdaFunction* is an Lambda API. Because the CloudTrail user did not specify logging data events for all Lambda functions, the `Invoke` operation for *MyOtherLambdaFunction* does not match the function specified for the trail. The trail doesn’t log the event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataResourceProperty := &DataResourceProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html
//
type CfnTrail_DataResourceProperty struct {
	// The resource type in which you want to log data events.
	//
	// You can specify the following *basic* event selector resource types:
	//
	// - `AWS::DynamoDB::Table`
	// - `AWS::Lambda::Function`
	// - `AWS::S3::Object`
	//
	// Additional resource types are available through *advanced* event selectors. For more information about these additional resource types, see [AdvancedFieldSelector](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified resource type.
	//
	// - To log data events for all objects in all S3 buckets in your AWS account , specify the prefix as `arn:aws:s3` .
	//
	// > This also enables logging of data event activity performed by any user or role in your AWS account , even if that activity is performed on a bucket that belongs to another AWS account .
	// - To log data events for all objects in an S3 bucket, specify the bucket and an empty object prefix such as `arn:aws:s3:::amzn-s3-demo-bucket1/` . The trail logs data events for all objects in this S3 bucket.
	// - To log data events for specific objects, specify the S3 bucket and object prefix such as `arn:aws:s3:::amzn-s3-demo-bucket1/example-images` . The trail logs data events for objects in this S3 bucket that match the prefix.
	// - To log data events for all Lambda functions in your AWS account , specify the prefix as `arn:aws:lambda` .
	//
	// > This also enables logging of `Invoke` activity performed by any user or role in your AWS account , even if that activity is performed on a function that belongs to another AWS account .
	// - To log data events for a specific Lambda function, specify the function ARN.
	//
	// > Lambda function ARNs are exact. For example, if you specify a function ARN *arn:aws:lambda:us-west-2:111111111111:function:helloworld* , data events will only be logged for *arn:aws:lambda:us-west-2:111111111111:function:helloworld* . They will not be logged for *arn:aws:lambda:us-west-2:111111111111:function:helloworld2* .
	// - To log data events for all DynamoDB tables in your AWS account , specify the prefix as `arn:aws:dynamodb` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

