package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/manishbhatt/gutils/configs"
	"github.com/manishbhatt/gutils/gtasks"

	"github.com/jessevdk/go-flags"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

var opts struct {
	Title string `short:"t" long:"taskname" description:"Title of the task" required:"true"`
	List  string `short:"l" long:"list" default:"Daily" description:"Task list name. Default is Daily"`
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Title %v : ", opts.Title)
	fmt.Printf("List %v : ", opts.List)
	fmt.Printf("All Args %s", strings.Join(args, " "))
}

func init() {
	ctx := context.Background()
	home, _ := os.UserHomeDir()
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/.got/go_creds.json", home))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, _ := configs.GetClient(config)

	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	gtasks.Service = srv
}
