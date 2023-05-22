// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an Asset Bundle import job. An Asset Bundle import job imports specified
// Amazon QuickSight assets into an Amazon QuickSight account. You can also choose
// to import a naming prefix and specified configuration overrides. The assets that
// are contained in the bundle file that you provide are used to create or update a
// new or existing asset in your Amazon QuickSight account. Each Amazon QuickSight
// account can run up to 10 import jobs concurrently. The API caller must have the
// necessary "create" , "describe" , and "update" permissions in their IAM role to
// access each resource type that is contained in the bundle file before the
// resources can be imported.
func (c *Client) StartAssetBundleImportJob(ctx context.Context, params *StartAssetBundleImportJobInput, optFns ...func(*Options)) (*StartAssetBundleImportJobOutput, error) {
	if params == nil {
		params = &StartAssetBundleImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAssetBundleImportJob", params, optFns, c.addOperationStartAssetBundleImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAssetBundleImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAssetBundleImportJobInput struct {

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	//
	// This member is required.
	AssetBundleImportJobId *string

	// The source of the asset bundle zip file that contains the data that you want to
	// import.
	//
	// This member is required.
	AssetBundleImportSource *types.AssetBundleImportSource

	// The ID of the Amazon Web Services account to import assets into.
	//
	// This member is required.
	AwsAccountId *string

	// The failure action for the import job. If you choose ROLLBACK , failed import
	// jobs will attempt to undo any asset changes caused by the failed job. If you
	// choose DO_NOTHING , failed import jobs will not attempt to roll back any asset
	// changes caused by the failed job, possibly leaving the Amazon QuickSight account
	// in an inconsistent state.
	FailureAction types.AssetBundleImportFailureAction

	// Optional overrides to be applied to the resource configuration before import.
	OverrideParameters *types.AssetBundleImportJobOverrideParameters

	noSmithyDocumentSerde
}

type StartAssetBundleImportJobOutput struct {

	// The Amazon Resource Name (ARN) for the import job.
	Arn *string

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	AssetBundleImportJobId *string

	// The Amazon Web Services response ID for this operation.
	RequestId *string

	// The HTTP status of the response.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAssetBundleImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartAssetBundleImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAssetBundleImportJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartAssetBundleImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "StartAssetBundleImportJob",
	}
}
