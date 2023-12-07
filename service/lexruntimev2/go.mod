module github.com/aws/aws-sdk-go-v2/service/lexruntimev2

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.23.5
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.3
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.8
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.8
	github.com/aws/smithy-go v1.19.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
