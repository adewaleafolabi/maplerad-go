package maplerad

import "github.com/wirepay/maplerad-go/models"

type InstitutionService service

func (c *InstitutionService) GetInstitutions() (*models.GetInstitutionsResponse, error) {
	u := "/institutions"
	resp := &models.GetInstitutionsResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

func (c *InstitutionService) ResolveInstitution(account_number, bank_code string) (*models.ResolveInstitutionsResponse, error) {
	u := "/institutions/resolve"
	var body map[string]string
	body["account_number"] = account_number
	body["bank_code"] = bank_code
	resp := &models.ResolveInstitutionsResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}