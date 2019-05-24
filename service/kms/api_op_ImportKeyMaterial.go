// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ImportKeyMaterialRequest
type ImportKeyMaterialInput struct {
	_ struct{} `type:"structure"`

	// The encrypted key material to import. It must be encrypted with the public
	// key that you received in the response to a previous GetParametersForImport
	// request, using the wrapping algorithm that you specified in that request.
	//
	// EncryptedKeyMaterial is automatically base64 encoded/decoded by the SDK.
	//
	// EncryptedKeyMaterial is a required field
	EncryptedKeyMaterial []byte `min:"1" type:"blob" required:"true"`

	// Specifies whether the key material expires. The default is KEY_MATERIAL_EXPIRES,
	// in which case you must include the ValidTo parameter. When this parameter
	// is set to KEY_MATERIAL_DOES_NOT_EXPIRE, you must omit the ValidTo parameter.
	ExpirationModel ExpirationModelType `type:"string" enum:"true"`

	// The import token that you received in the response to a previous GetParametersForImport
	// request. It must be from the same response that contained the public key
	// that you used to encrypt the key material.
	//
	// ImportToken is automatically base64 encoded/decoded by the SDK.
	//
	// ImportToken is a required field
	ImportToken []byte `min:"1" type:"blob" required:"true"`

	// The identifier of the CMK to import the key material into. The CMK's Origin
	// must be EXTERNAL.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// The time at which the imported key material expires. When the key material
	// expires, AWS KMS deletes the key material and the CMK becomes unusable. You
	// must omit this parameter when the ExpirationModel parameter is set to KEY_MATERIAL_DOES_NOT_EXPIRE.
	// Otherwise it is required.
	ValidTo *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s ImportKeyMaterialInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportKeyMaterialInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportKeyMaterialInput"}

	if s.EncryptedKeyMaterial == nil {
		invalidParams.Add(aws.NewErrParamRequired("EncryptedKeyMaterial"))
	}
	if s.EncryptedKeyMaterial != nil && len(s.EncryptedKeyMaterial) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EncryptedKeyMaterial", 1))
	}

	if s.ImportToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImportToken"))
	}
	if s.ImportToken != nil && len(s.ImportToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImportToken", 1))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ImportKeyMaterialResponse
type ImportKeyMaterialOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportKeyMaterialOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportKeyMaterial = "ImportKeyMaterial"

// ImportKeyMaterialRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Imports key material into an existing AWS KMS customer master key (CMK) that
// was created without key material. You cannot perform this operation on a
// CMK in a different AWS account. For more information about creating CMKs
// with no key material and then importing key material, see Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html)
// in the AWS Key Management Service Developer Guide.
//
// Before using this operation, call GetParametersForImport. Its response includes
// a public key and an import token. Use the public key to encrypt the key material.
// Then, submit the import token from the same GetParametersForImport response.
//
// When calling this operation, you must specify the following values:
//
//    * The key ID or key ARN of a CMK with no key material. Its Origin must
//    be EXTERNAL. To create a CMK with no key material, call CreateKey and
//    set the value of its Origin parameter to EXTERNAL. To get the Origin of
//    a CMK, call DescribeKey.)
//
//    * The encrypted key material. To get the public key to encrypt the key
//    material, call GetParametersForImport.
//
//    * The import token that GetParametersForImport returned. This token and
//    the public key used to encrypt the key material must have come from the
//    same response.
//
//    * Whether the key material expires and if so, when. If you set an expiration
//    date, you can change it only by reimporting the same key material and
//    specifying a new expiration date. If the key material expires, AWS KMS
//    deletes the key material and the CMK becomes unusable. To use the CMK
//    again, you must reimport the same key material.
//
// When this operation is successful, the key state of the CMK changes from
// PendingImport to Enabled, and you can use the CMK. After you successfully
// import key material into a CMK, you can reimport the same key material into
// that CMK, but you cannot import different key material.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ImportKeyMaterialRequest.
//    req := client.ImportKeyMaterialRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ImportKeyMaterial
func (c *Client) ImportKeyMaterialRequest(input *ImportKeyMaterialInput) ImportKeyMaterialRequest {
	op := &aws.Operation{
		Name:       opImportKeyMaterial,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportKeyMaterialInput{}
	}

	req := c.newRequest(op, input, &ImportKeyMaterialOutput{})
	return ImportKeyMaterialRequest{Request: req, Input: input, Copy: c.ImportKeyMaterialRequest}
}

// ImportKeyMaterialRequest is the request type for the
// ImportKeyMaterial API operation.
type ImportKeyMaterialRequest struct {
	*aws.Request
	Input *ImportKeyMaterialInput
	Copy  func(*ImportKeyMaterialInput) ImportKeyMaterialRequest
}

// Send marshals and sends the ImportKeyMaterial API request.
func (r ImportKeyMaterialRequest) Send(ctx context.Context) (*ImportKeyMaterialResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportKeyMaterialResponse{
		ImportKeyMaterialOutput: r.Request.Data.(*ImportKeyMaterialOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportKeyMaterialResponse is the response type for the
// ImportKeyMaterial API operation.
type ImportKeyMaterialResponse struct {
	*ImportKeyMaterialOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportKeyMaterial request.
func (r *ImportKeyMaterialResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
