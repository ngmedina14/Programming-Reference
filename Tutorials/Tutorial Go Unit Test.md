# Tutorial Go Unit Test (Romnick's Standard)

**Step 1:** Initialize the example struct.

```go
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
```
  
Where the first group is the initialized fields and the second group are the test cases that you want to run.

**Step 2:** Initialize the GET and POST URL that you want to access.

```go
w := httptest.NewRecorder()
getURL := "http://localhost:8080/api/generate_excel_sales_report?from=%s&to=%s&status=%s"
postURL := "http://localhost:8080/api/generate_excel_sales_report"
```

Where getURL is the one that you access directly into the web address bar while the postURL is the one that you access through the POST AJAX in your code.

**Step 3:** Escape the values so that you can pass directly into the GET URL.

```go
escapeFrom := url.QueryEscape(e.From)
escapeTo := url.QueryEscape(e.To)
escapeStatus := url.QueryEscape(e.Status)
```

**Step 4:** Create the condition for the AJAX Method.

* First condition: **GET**

```go
r := httptest.NewRequest(e.Method, fmt.Sprintf(getURL, escapeFrom, escapeTo, escapeStatus), nil)

fmt.Printf(getURL+"\n", escapeFrom, escapeTo, escapeStatus)

GenerateExcelSalesReportHandler(w, r)

if w.Code != http.StatusOK {
	t.Errorf("Invalid code on requesting /. %d", w.Code)
}

Create a new request there and pass the values from the example struct. Afterwards, pass it into the handler function.

* Second condition: **POST**

```go
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
```

Create a new request there and set the values in the Post Form. Afterwards, pass it into the handler function.

* Third condition: **INVALID METHOD**

Just copy the code from the first condition.

**Step 5:** Check the results in your terminal.
