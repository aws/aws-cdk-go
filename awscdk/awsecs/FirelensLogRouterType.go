package awsecs


// Firelens log router type, fluentbit or fluentd.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html
type FirelensLogRouterType string

const (
	// fluentbit.
	FirelensLogRouterType_FLUENTBIT FirelensLogRouterType = "FLUENTBIT"
	// fluentd.
	FirelensLogRouterType_FLUENTD FirelensLogRouterType = "FLUENTD"
)

