package ui

import (
	"goREST/internal/request"
	"strings"
)

func Details(request request.RequestObject) string {
	doc := strings.Builder{}

	doc.WriteString(HighlightText.Render(request.URL) + "\n\n")

	doc.WriteString("Method: " + renderMethod(request.Method) + "\n\n")

	if len(request.Headers) > 0 {
		doc.WriteString("Headers: \n")
		doc.WriteString(renderKeyValuePairs(request.Headers) + "\n\n")
	}

	if len(request.Params) > 0 {
		doc.WriteString("Parameters: \n")
		doc.WriteString(renderKeyValuePairs(request.Params) + "\n\n")
	}

	if len(request.Body) > 0 {
		doc.WriteString("Body: " + request.Body + "\n\n")
	}

	if request.Auth.Type == "Basic" {
		doc.WriteString("Basic Auth:\n")
		doc.WriteString("Username: " + request.Auth.Username + "\n")
		doc.WriteString("Password: ****************\n\n")
	}

	if request.Auth.Type == "Bearer" {
		doc.WriteString("Bearer Auth:\n")
		doc.WriteString("Token: " + request.Auth.Token + "\n")
	}

	doc.WriteString(BackdropText.Render("Last modified: " + request.LastMod.Format("02.01.2006 15:04:05")))

	return Detail.Render(doc.String())
}

func renderKeyValuePairs(pairs map[string]string) string {
	doc := strings.Builder{}
	dlDoc := strings.Builder{}
	i := 0

	for k, v := range pairs {

		dlDoc.WriteString(DefinitionTerm.Render(k))
		dlDoc.WriteString("\n")
		dlDoc.WriteString(DefinitionData.Render(v))

		if i < len(pairs)-1 {
			dlDoc.WriteString("\n")
		}

		i++
	}

	doc.WriteString(DefinitionList.Render(dlDoc.String()))
	return doc.String()
}

func renderMethod(m string) string {
	switch m {
	case "GET":
		return GET.Render(m)
	case "POST":
		return POST.Render(m)
	case "PUT":
		return PUT.Render(m)
	case "PATCH":
		return PATCH.Render(m)
	case "DELETE":
		return DELETE.Render(m)
	default:
		return Method.Render(m)
	}
}
