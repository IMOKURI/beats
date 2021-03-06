// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for DescribeStackEvents action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackEventsInput
type DescribeStackEventsInput struct {
	_ struct{} `type:"structure"`

	// A string that identifies the next page of events that you want to retrieve.
	NextToken *string `min:"1" type:"string"`

	// The name or the unique stack ID that is associated with the stack, which
	// are not always interchangeable:
	//
	//    * Running stacks: You can specify either the stack's name or its unique
	//    stack ID.
	//
	//    * Deleted stacks: You must specify the unique stack ID.
	//
	// Default: There is no default value.
	StackName *string `type:"string"`
}

// String returns the string representation
func (s DescribeStackEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackEventsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a DescribeStackEvents action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackEventsOutput
type DescribeStackEventsOutput struct {
	_ struct{} `type:"structure"`

	// If the output exceeds 1 MB in size, a string that identifies the next page
	// of events. If no additional page exists, this value is null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackEvents structures.
	StackEvents []StackEvent `type:"list"`
}

// String returns the string representation
func (s DescribeStackEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackEvents = "DescribeStackEvents"

// DescribeStackEventsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns all stack related events for a specified stack in reverse chronological
// order. For more information about a stack's event history, go to Stacks (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/concept-stack.html)
// in the AWS CloudFormation User Guide.
//
// You can list events for stacks that have failed to create or have been deleted
// by specifying the unique stack identifier (stack ID).
//
//    // Example sending a request using DescribeStackEventsRequest.
//    req := client.DescribeStackEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackEvents
func (c *Client) DescribeStackEventsRequest(input *DescribeStackEventsInput) DescribeStackEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeStackEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeStackEventsInput{}
	}

	req := c.newRequest(op, input, &DescribeStackEventsOutput{})
	return DescribeStackEventsRequest{Request: req, Input: input, Copy: c.DescribeStackEventsRequest}
}

// DescribeStackEventsRequest is the request type for the
// DescribeStackEvents API operation.
type DescribeStackEventsRequest struct {
	*aws.Request
	Input *DescribeStackEventsInput
	Copy  func(*DescribeStackEventsInput) DescribeStackEventsRequest
}

// Send marshals and sends the DescribeStackEvents API request.
func (r DescribeStackEventsRequest) Send(ctx context.Context) (*DescribeStackEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackEventsResponse{
		DescribeStackEventsOutput: r.Request.Data.(*DescribeStackEventsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeStackEventsRequestPaginator returns a paginator for DescribeStackEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeStackEventsRequest(input)
//   p := cloudformation.NewDescribeStackEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeStackEventsPaginator(req DescribeStackEventsRequest) DescribeStackEventsPaginator {
	return DescribeStackEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeStackEventsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeStackEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeStackEventsPaginator struct {
	aws.Pager
}

func (p *DescribeStackEventsPaginator) CurrentPage() *DescribeStackEventsOutput {
	return p.Pager.CurrentPage().(*DescribeStackEventsOutput)
}

// DescribeStackEventsResponse is the response type for the
// DescribeStackEvents API operation.
type DescribeStackEventsResponse struct {
	*DescribeStackEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackEvents request.
func (r *DescribeStackEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
