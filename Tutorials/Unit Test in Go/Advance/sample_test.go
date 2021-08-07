package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/* TestGenerateExcelSalesReportHandler is a unit testing function for GenerateExcelSalesReportHandler() function */
func TestGenerateExcelSalesReportHandler(t *testing.T) {
	examples := []struct {
		Method string
		From   string
		To     string
		Status string
	}{
		{Method: "GET", From: "", To: "", Status: ""},
		{Method: "POST", From: "", To: "", Status: ""},
		{Method: "NULL", From: "", To: "", Status: ""},
		{Method: "POST", From: "abcd", To: "efgh", Status: ""},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: ""},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: "G"},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: "3"},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: "4"},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: "3,G"},
		{Method: "POST", From: "2021-06-02", To: "2021-06-03", Status: "3,4"},
	}

	w := httptest.NewRecorder()
	getURL := "http://localhost:8080/api/generate_excel_sales_report?from=%s&to=%s&status=%s"
	postURL := "http://localhost:8080/api/generate_excel_sales_report"

	fmt.Println("\n-------------- GENERATE EXCEL SALES REPORT API HANDLER --------------")

	for _, e := range examples {
		escapeFrom := url.QueryEscape(e.From)
		escapeTo := url.QueryEscape(e.To)
		escapeStatus := url.QueryEscape(e.Status)

		fmt.Printf("\nMethod: %s\n", e.Method)

		if e.Method == "GET" {
			r := httptest.NewRequest(e.Method, fmt.Sprintf(getURL, escapeFrom, escapeTo, escapeStatus), nil)

			fmt.Printf(getURL+"\n", escapeFrom, escapeTo, escapeStatus)

			GenerateExcelSalesReportHandler(w, r)

			if w.Code != http.StatusOK {
				t.Errorf("Invalid code on requesting /. %d", w.Code)
			}

		} else if e.Method == "POST" {
			r := httptest.NewRequest(e.Method, fmt.Sprintf(postURL), nil)
			r.PostForm = url.Values{}
			r.PostForm.Set("from", e.From)
			r.PostForm.Set("to", e.To)
			r.PostForm.Set("status", e.Status)

			fmt.Printf(postURL+"\n", e.From, e.To, e.Status)

			GenerateExcelSalesReportHandler(w, r)

			if w.Code != http.StatusOK {
				t.Errorf("Invalid code on requesting /. %d", w.Code)
			}

		} else {
			r := httptest.NewRequest(e.Method, fmt.Sprintf(getURL, escapeFrom, escapeTo, escapeStatus), nil)

			fmt.Printf(getURL+"\n", escapeFrom, escapeTo, escapeStatus)

			GenerateExcelSalesReportHandler(w, r)

			if w.Code != http.StatusOK {
				t.Errorf("Invalid code on requesting /. %d", w.Code)
			}
		}
	}
}
