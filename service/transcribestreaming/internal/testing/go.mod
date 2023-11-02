module github.com/aws/aws-sdk-go-v2/service/transcribestreaming/internal/testing

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.22.1
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.0
	github.com/aws/aws-sdk-go-v2/service/internal/eventstreamtesting v1.2.1
	github.com/aws/aws-sdk-go-v2/service/transcribestreaming v1.13.0
	github.com/aws/smithy-go v1.16.0
	github.com/google/go-cmp v0.5.8
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.1 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)

replace github.com/aws/aws-sdk-go-v2 => ../../../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/eventstreamtesting => ../../../../service/internal/eventstreamtesting/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/aws/aws-sdk-go-v2/service/ssooidc => ../../../../service/ssooidc/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../../service/sts/

replace github.com/aws/aws-sdk-go-v2/service/transcribestreaming => ../../../../service/transcribestreaming/
