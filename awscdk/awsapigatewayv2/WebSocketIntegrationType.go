package awsapigatewayv2


// WebSocket Integration Types.
type WebSocketIntegrationType string

const (
	// AWS Proxy Integration Type.
	WebSocketIntegrationType_AWS_PROXY WebSocketIntegrationType = "AWS_PROXY"
	// Mock Integration Type.
	WebSocketIntegrationType_MOCK WebSocketIntegrationType = "MOCK"
	// AWS Integration Type.
	WebSocketIntegrationType_AWS WebSocketIntegrationType = "AWS"
)

