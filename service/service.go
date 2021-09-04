package service

import (
	"bytes"
	"context"
	"home-24/model"
	"home-24/utils"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func crawl(node *html.Node, doc *model.HtmlDocument) {
	switch {
	case node.Type == html.TextNode && node.Parent.Type == html.ElementNode:
		switch strings.ToLower(node.Parent.Data) {
		case "title":
			doc.Title = node.Data
		case "h1":
			doc.H1_Count++
		case "h2":
			doc.H2_Count++
		case "h3":
			doc.H3_Count++
		case "h4":
			doc.H4_Count++
		case "h5":
			doc.H5_Count++
		case "h6":
			doc.H6_Count++
		case "a":
			for _, attr := range node.Parent.Attr {
				if strings.ToLower(attr.Key) == "href" {
					if u, err := url.Parse(attr.Val); err == nil {
						doc.AnchorsTags = append(doc.AnchorsTags, model.AnchorTag{
							Url:      attr.Val,
							External: u.Host != doc.Host,
						})
						break
					}
				}
			}
		}
	case node.Type == html.ElementNode && node.Data == "input":
		for _, attr := range node.Attr {
			if strings.ToLower(attr.Key) == "name" && attr.Val == "username" {
				doc.UserNameFieldExist = true
				break
			} else if strings.ToLower(attr.Key) == "name" && attr.Val == "password" {
				doc.PasswordFieldExist = true
				break
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		crawl(child, doc)
	}
}

func Analyze(ctx context.Context, client utils.HttpClient, urlToBeScanned string) (*model.HtmlDocument, error) {
	var stat model.HtmlDocument
	if url, err := url.Parse(urlToBeScanned); err == nil {
		stat.Host = url.Host
	} else {
		return nil, err
	}

	documentBody, err := client.ReadBytes(ctx, urlToBeScanned)
	if err != nil {
		return nil, err
	}

	doc, _ := html.Parse(bytes.NewBuffer(documentBody))
	crawl(doc, &stat)
	for i, aTag := range stat.AnchorsTags {
		stat.AnchorsTags[i].Accessible = client.IsAccessible(ctx, aTag.Url)
	}
	return &stat, nil
}
