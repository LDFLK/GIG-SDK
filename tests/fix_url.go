package tests

import (
	"GIG/sdk"
)

func (t *TestSDK) TestThatFixUrlWorks() {
	href := "/images/pdf/registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	baseUrl := "http://www.buildings.gov.lk/index.php?option=com_content&view=category&layout=blog&id=47&Itemid=128&lang=en"
	result := sdk.FixUrl(href, baseUrl)
	t.AssertEqual("http://www.buildings.gov.lk"+href,result)
}

func (t *TestSDK) TestThatFixUrlWorksForAbsoluteUrl() {
	href := "http://www.buildings.gov.lk/images/pdf/registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	baseUrl := "http://www.buildings.gov.lk/index.php?option=com_content&view=category&layout=blog&id=47&Itemid=128&lang=en"
	result := sdk.FixUrl(href, baseUrl)
	t.AssertEqual(result,href)
}