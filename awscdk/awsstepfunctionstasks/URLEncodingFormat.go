package awsstepfunctionstasks


// The style used when applying URL encoding to array values.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_Basic(jsii.String("username"), awscdk.SecretValue_UnsafePlainText(jsii.String("password"))),
//   })
//
//   tasks.NewHttpInvoke(this, jsii.String("Invoke HTTP API"), &HttpInvokeProps{
//   	ApiRoot: jsii.String("https://api.example.com"),
//   	ApiEndpoint: sfn.TaskInput_FromText(sfn.JsonPath_Format(jsii.String("resource/{}/details"), sfn.JsonPath_StringAt(jsii.String("$.resourceId")))),
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"foo": jsii.String("bar"),
//   	}),
//   	Connection: Connection,
//   	Headers: sfn.TaskInput_*FromObject(map[string]interface{}{
//   		"Content-Type": jsii.String("application/json"),
//   	}),
//   	Method: sfn.TaskInput_*FromText(jsii.String("POST")),
//   	QueryStringParameters: sfn.TaskInput_*FromObject(map[string]interface{}{
//   		"id": jsii.String("123"),
//   	}),
//   	UrlEncodingFormat: tasks.URLEncodingFormat_BRACKETS,
//   })
//
type URLEncodingFormat string

const (
	// Encode arrays using brackets.
	//
	// For example, {'array': ['a','b','c']} encodes to 'array[]=a&array[]=b&array[]=c'.
	URLEncodingFormat_BRACKETS URLEncodingFormat = "BRACKETS"
	// Encode arrays using commas.
	//
	// For example, {'array': ['a','b','c']} encodes to 'array=a,b,c,d'.
	URLEncodingFormat_COMMAS URLEncodingFormat = "COMMAS"
	// Apply the default URL encoding style (INDICES).
	URLEncodingFormat_DEFAULT URLEncodingFormat = "DEFAULT"
	// Encode arrays using the index value.
	//
	// For example, {'array': ['a','b','c']} encodes to 'array[0]=a&array[1]=b&array[2]=c'.
	URLEncodingFormat_INDICES URLEncodingFormat = "INDICES"
	// Do not apply URL encoding.
	URLEncodingFormat_NONE URLEncodingFormat = "NONE"
	// Repeat key for each item in the array.
	//
	// For example, {'array': ['a','b','c']} encodes to 'array[]=a&array[]=b&array[]=c'.
	URLEncodingFormat_REPEAT URLEncodingFormat = "REPEAT"
)

