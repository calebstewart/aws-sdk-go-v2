// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package servicecatalogiface provides an interface to enable mocking the AWS Service Catalog service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicecatalogiface

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

// ClientAPI provides an interface to enable mocking the
// servicecatalog.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog.
//    func myFunc(svc servicecatalogiface.ClientAPI) bool {
//        // Make svc.AcceptPortfolioShare request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := servicecatalog.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        servicecatalogiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptPortfolioShare(input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AcceptPortfolioShareRequest(*servicecatalog.AcceptPortfolioShareInput) servicecatalog.AcceptPortfolioShareRequest

	AssociateBudgetWithResourceRequest(*servicecatalog.AssociateBudgetWithResourceInput) servicecatalog.AssociateBudgetWithResourceRequest

	AssociatePrincipalWithPortfolioRequest(*servicecatalog.AssociatePrincipalWithPortfolioInput) servicecatalog.AssociatePrincipalWithPortfolioRequest

	AssociateProductWithPortfolioRequest(*servicecatalog.AssociateProductWithPortfolioInput) servicecatalog.AssociateProductWithPortfolioRequest

	AssociateServiceActionWithProvisioningArtifactRequest(*servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) servicecatalog.AssociateServiceActionWithProvisioningArtifactRequest

	AssociateTagOptionWithResourceRequest(*servicecatalog.AssociateTagOptionWithResourceInput) servicecatalog.AssociateTagOptionWithResourceRequest

	BatchAssociateServiceActionWithProvisioningArtifactRequest(*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactRequest

	BatchDisassociateServiceActionFromProvisioningArtifactRequest(*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactRequest

	CopyProductRequest(*servicecatalog.CopyProductInput) servicecatalog.CopyProductRequest

	CreateConstraintRequest(*servicecatalog.CreateConstraintInput) servicecatalog.CreateConstraintRequest

	CreatePortfolioRequest(*servicecatalog.CreatePortfolioInput) servicecatalog.CreatePortfolioRequest

	CreatePortfolioShareRequest(*servicecatalog.CreatePortfolioShareInput) servicecatalog.CreatePortfolioShareRequest

	CreateProductRequest(*servicecatalog.CreateProductInput) servicecatalog.CreateProductRequest

	CreateProvisionedProductPlanRequest(*servicecatalog.CreateProvisionedProductPlanInput) servicecatalog.CreateProvisionedProductPlanRequest

	CreateProvisioningArtifactRequest(*servicecatalog.CreateProvisioningArtifactInput) servicecatalog.CreateProvisioningArtifactRequest

	CreateServiceActionRequest(*servicecatalog.CreateServiceActionInput) servicecatalog.CreateServiceActionRequest

	CreateTagOptionRequest(*servicecatalog.CreateTagOptionInput) servicecatalog.CreateTagOptionRequest

	DeleteConstraintRequest(*servicecatalog.DeleteConstraintInput) servicecatalog.DeleteConstraintRequest

	DeletePortfolioRequest(*servicecatalog.DeletePortfolioInput) servicecatalog.DeletePortfolioRequest

	DeletePortfolioShareRequest(*servicecatalog.DeletePortfolioShareInput) servicecatalog.DeletePortfolioShareRequest

	DeleteProductRequest(*servicecatalog.DeleteProductInput) servicecatalog.DeleteProductRequest

	DeleteProvisionedProductPlanRequest(*servicecatalog.DeleteProvisionedProductPlanInput) servicecatalog.DeleteProvisionedProductPlanRequest

	DeleteProvisioningArtifactRequest(*servicecatalog.DeleteProvisioningArtifactInput) servicecatalog.DeleteProvisioningArtifactRequest

	DeleteServiceActionRequest(*servicecatalog.DeleteServiceActionInput) servicecatalog.DeleteServiceActionRequest

	DeleteTagOptionRequest(*servicecatalog.DeleteTagOptionInput) servicecatalog.DeleteTagOptionRequest

	DescribeConstraintRequest(*servicecatalog.DescribeConstraintInput) servicecatalog.DescribeConstraintRequest

	DescribeCopyProductStatusRequest(*servicecatalog.DescribeCopyProductStatusInput) servicecatalog.DescribeCopyProductStatusRequest

	DescribePortfolioRequest(*servicecatalog.DescribePortfolioInput) servicecatalog.DescribePortfolioRequest

	DescribePortfolioShareStatusRequest(*servicecatalog.DescribePortfolioShareStatusInput) servicecatalog.DescribePortfolioShareStatusRequest

	DescribeProductRequest(*servicecatalog.DescribeProductInput) servicecatalog.DescribeProductRequest

	DescribeProductAsAdminRequest(*servicecatalog.DescribeProductAsAdminInput) servicecatalog.DescribeProductAsAdminRequest

	DescribeProductViewRequest(*servicecatalog.DescribeProductViewInput) servicecatalog.DescribeProductViewRequest

	DescribeProvisionedProductRequest(*servicecatalog.DescribeProvisionedProductInput) servicecatalog.DescribeProvisionedProductRequest

	DescribeProvisionedProductPlanRequest(*servicecatalog.DescribeProvisionedProductPlanInput) servicecatalog.DescribeProvisionedProductPlanRequest

	DescribeProvisioningArtifactRequest(*servicecatalog.DescribeProvisioningArtifactInput) servicecatalog.DescribeProvisioningArtifactRequest

	DescribeProvisioningParametersRequest(*servicecatalog.DescribeProvisioningParametersInput) servicecatalog.DescribeProvisioningParametersRequest

	DescribeRecordRequest(*servicecatalog.DescribeRecordInput) servicecatalog.DescribeRecordRequest

	DescribeServiceActionRequest(*servicecatalog.DescribeServiceActionInput) servicecatalog.DescribeServiceActionRequest

	DescribeTagOptionRequest(*servicecatalog.DescribeTagOptionInput) servicecatalog.DescribeTagOptionRequest

	DisableAWSOrganizationsAccessRequest(*servicecatalog.DisableAWSOrganizationsAccessInput) servicecatalog.DisableAWSOrganizationsAccessRequest

	DisassociateBudgetFromResourceRequest(*servicecatalog.DisassociateBudgetFromResourceInput) servicecatalog.DisassociateBudgetFromResourceRequest

	DisassociatePrincipalFromPortfolioRequest(*servicecatalog.DisassociatePrincipalFromPortfolioInput) servicecatalog.DisassociatePrincipalFromPortfolioRequest

	DisassociateProductFromPortfolioRequest(*servicecatalog.DisassociateProductFromPortfolioInput) servicecatalog.DisassociateProductFromPortfolioRequest

	DisassociateServiceActionFromProvisioningArtifactRequest(*servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) servicecatalog.DisassociateServiceActionFromProvisioningArtifactRequest

	DisassociateTagOptionFromResourceRequest(*servicecatalog.DisassociateTagOptionFromResourceInput) servicecatalog.DisassociateTagOptionFromResourceRequest

	EnableAWSOrganizationsAccessRequest(*servicecatalog.EnableAWSOrganizationsAccessInput) servicecatalog.EnableAWSOrganizationsAccessRequest

	ExecuteProvisionedProductPlanRequest(*servicecatalog.ExecuteProvisionedProductPlanInput) servicecatalog.ExecuteProvisionedProductPlanRequest

	ExecuteProvisionedProductServiceActionRequest(*servicecatalog.ExecuteProvisionedProductServiceActionInput) servicecatalog.ExecuteProvisionedProductServiceActionRequest

	GetAWSOrganizationsAccessStatusRequest(*servicecatalog.GetAWSOrganizationsAccessStatusInput) servicecatalog.GetAWSOrganizationsAccessStatusRequest

	ListAcceptedPortfolioSharesRequest(*servicecatalog.ListAcceptedPortfolioSharesInput) servicecatalog.ListAcceptedPortfolioSharesRequest

	ListBudgetsForResourceRequest(*servicecatalog.ListBudgetsForResourceInput) servicecatalog.ListBudgetsForResourceRequest

	ListConstraintsForPortfolioRequest(*servicecatalog.ListConstraintsForPortfolioInput) servicecatalog.ListConstraintsForPortfolioRequest

	ListLaunchPathsRequest(*servicecatalog.ListLaunchPathsInput) servicecatalog.ListLaunchPathsRequest

	ListOrganizationPortfolioAccessRequest(*servicecatalog.ListOrganizationPortfolioAccessInput) servicecatalog.ListOrganizationPortfolioAccessRequest

	ListPortfolioAccessRequest(*servicecatalog.ListPortfolioAccessInput) servicecatalog.ListPortfolioAccessRequest

	ListPortfoliosRequest(*servicecatalog.ListPortfoliosInput) servicecatalog.ListPortfoliosRequest

	ListPortfoliosForProductRequest(*servicecatalog.ListPortfoliosForProductInput) servicecatalog.ListPortfoliosForProductRequest

	ListPrincipalsForPortfolioRequest(*servicecatalog.ListPrincipalsForPortfolioInput) servicecatalog.ListPrincipalsForPortfolioRequest

	ListProvisionedProductPlansRequest(*servicecatalog.ListProvisionedProductPlansInput) servicecatalog.ListProvisionedProductPlansRequest

	ListProvisioningArtifactsRequest(*servicecatalog.ListProvisioningArtifactsInput) servicecatalog.ListProvisioningArtifactsRequest

	ListProvisioningArtifactsForServiceActionRequest(*servicecatalog.ListProvisioningArtifactsForServiceActionInput) servicecatalog.ListProvisioningArtifactsForServiceActionRequest

	ListRecordHistoryRequest(*servicecatalog.ListRecordHistoryInput) servicecatalog.ListRecordHistoryRequest

	ListResourcesForTagOptionRequest(*servicecatalog.ListResourcesForTagOptionInput) servicecatalog.ListResourcesForTagOptionRequest

	ListServiceActionsRequest(*servicecatalog.ListServiceActionsInput) servicecatalog.ListServiceActionsRequest

	ListServiceActionsForProvisioningArtifactRequest(*servicecatalog.ListServiceActionsForProvisioningArtifactInput) servicecatalog.ListServiceActionsForProvisioningArtifactRequest

	ListTagOptionsRequest(*servicecatalog.ListTagOptionsInput) servicecatalog.ListTagOptionsRequest

	ProvisionProductRequest(*servicecatalog.ProvisionProductInput) servicecatalog.ProvisionProductRequest

	RejectPortfolioShareRequest(*servicecatalog.RejectPortfolioShareInput) servicecatalog.RejectPortfolioShareRequest

	ScanProvisionedProductsRequest(*servicecatalog.ScanProvisionedProductsInput) servicecatalog.ScanProvisionedProductsRequest

	SearchProductsRequest(*servicecatalog.SearchProductsInput) servicecatalog.SearchProductsRequest

	SearchProductsAsAdminRequest(*servicecatalog.SearchProductsAsAdminInput) servicecatalog.SearchProductsAsAdminRequest

	SearchProvisionedProductsRequest(*servicecatalog.SearchProvisionedProductsInput) servicecatalog.SearchProvisionedProductsRequest

	TerminateProvisionedProductRequest(*servicecatalog.TerminateProvisionedProductInput) servicecatalog.TerminateProvisionedProductRequest

	UpdateConstraintRequest(*servicecatalog.UpdateConstraintInput) servicecatalog.UpdateConstraintRequest

	UpdatePortfolioRequest(*servicecatalog.UpdatePortfolioInput) servicecatalog.UpdatePortfolioRequest

	UpdateProductRequest(*servicecatalog.UpdateProductInput) servicecatalog.UpdateProductRequest

	UpdateProvisionedProductRequest(*servicecatalog.UpdateProvisionedProductInput) servicecatalog.UpdateProvisionedProductRequest

	UpdateProvisionedProductPropertiesRequest(*servicecatalog.UpdateProvisionedProductPropertiesInput) servicecatalog.UpdateProvisionedProductPropertiesRequest

	UpdateProvisioningArtifactRequest(*servicecatalog.UpdateProvisioningArtifactInput) servicecatalog.UpdateProvisioningArtifactRequest

	UpdateServiceActionRequest(*servicecatalog.UpdateServiceActionInput) servicecatalog.UpdateServiceActionRequest

	UpdateTagOptionRequest(*servicecatalog.UpdateTagOptionInput) servicecatalog.UpdateTagOptionRequest
}

var _ ClientAPI = (*servicecatalog.Client)(nil)
