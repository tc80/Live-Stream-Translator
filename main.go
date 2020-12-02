package main

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

const (
	url              = "https://www.youtube.com/watch?v=5qap5aO4i9A"
	translationScope = "https://www.googleapis.com/auth/cloud-translation"
	projectID        = "live-stream-translator"
	location         = "global"
)

var (
	parent = fmt.Sprintf("projects/%s/locations/%s", projectID, location)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func translateText(client *translate.TranslationClient, contents []string, src, dst string) (*translatepb.TranslateTextResponse, error) {
	return client.TranslateText(context.Background(), &translatepb.TranslateTextRequest{
		Parent:             parent,
		Contents:           contents,
		SourceLanguageCode: src,
		TargetLanguageCode: dst,
	})
}

func main() {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("credentials.json"))
	check(err)
	defer client.Close()

	resp, err := translateText(client, []string{"hello", "world"}, "en-US", "bg")
	check(err)

	fmt.Println(resp.String())
}
