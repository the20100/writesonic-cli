package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/vincentmaurin/writesonic-cli/internal/output"
)

var (
	blogTopic          string
	blogPrimaryKeyword string
)

var blogIdeasCmd = &cobra.Command{
	Use:   "blog-ideas",
	Short: "Generate blog post ideas for a topic",
	Example: `  writesonic blog-ideas --topic "sustainable fashion"
  writesonic blog-ideas --topic "AI tools" --copies 3 --engine premium`,
	RunE: runBlogIdeas,
}

func init() {
	blogIdeasCmd.Flags().StringVar(&blogTopic, "topic", "", "Topic to generate ideas for (required)")
	blogIdeasCmd.Flags().StringVar(&blogPrimaryKeyword, "keyword", "", "Primary keyword to focus on")
	blogIdeasCmd.MarkFlagRequired("topic")
	rootCmd.AddCommand(blogIdeasCmd)
}

func runBlogIdeas(cmd *cobra.Command, args []string) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	body := map[string]interface{}{
		"topic": blogTopic,
	}
	if blogPrimaryKeyword != "" {
		body["primary_keyword"] = blogPrimaryKeyword
	}

	results, err := client.PostResults("/blog-ideas", params, body)
	if err != nil {
		return err
	}

	if output.IsJSON(jsonFlag, prettyFlag) {
		return output.PrintJSON(results, prettyFlag)
	}

	texts := make([]string, len(results))
	for i, r := range results {
		texts[i] = r.Text
	}
	output.PrintText(texts)
	return nil
}
